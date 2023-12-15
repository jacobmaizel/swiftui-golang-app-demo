package competitions

import (
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

type RequestCompetition struct {
	Title             string    `json:"title" binding:"required"`
	Description       string    `json:"description" binding:"required"`
	Start             time.Time `json:"start" binding:"required"`
	End               time.Time `json:"end" binding:"required"`
	Scheduled         bool      `json:"scheduled" `
	Participant_types []string  `json:"participant_types" binding:"required,min=1"`
	Workout_types     []string  `json:"workout_types" binding:"required,min=1"`
	Type              string    `json:"type" binding:"required"`
	Status            string    `json:"status" binding:"required"`
	Public            bool      `json:"public"`
}

type UpdateCompetitionRequest struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

type ResponseCompetition struct {
	ID                   uuid.UUID               `json:"id"`
	Title                string                  `json:"title"`
	Description          string                  `json:"description"`
	Start                time.Time               `json:"start"`
	End                  time.Time               `json:"end"`
	Scheduled            *bool                   `json:"scheduled"`
	Participant_types    []string                `json:"participant_types"`
	Workout_types        []string                `json:"workout_types"`
	Type                 string                  `json:"type"`
	Status               string                  `json:"status"`
	Public               *bool                   `json:"public"`
	Joined               bool                    `json:"joined"`
	Profile_participants []shared.PartialProfile `json:"profile_participants"`
	Place                string                  `json:"place"`
	Distance_total       float64                 `json:"distance_total"`
	Participant_count    int                     `json:"participant_count"`

	Owner     shared.PartialProfile `json:"owner"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

func generateCompetitionListResponse(c *gin.Context, comps []*ent.Competition) []ResponseCompetition {

	var resCompList []ResponseCompetition = []ResponseCompetition{}

	prof := utils.ProfileFromCtx(c)

	allProfCompList, _ := s.Db.Profile.Query().Where(profile.ID(prof.ID)).QueryCompetitions().All(c)

	profCompLookup := make(map[uuid.UUID]bool)

	for _, joinedComp := range allProfCompList {
		profCompLookup[joinedComp.ID] = true
	}

	for _, comp := range comps {

		var profile_participants []shared.PartialProfile = []shared.PartialProfile{}

		for _, prof := range comp.Edges.Participants {

			profile_participants = append(profile_participants, utils.GeneratePartialProfileResponse(c, prof))

		}

		place, distance, err := CalcLeaderboardPlace(c, prof.ID.String(), comp.ID)

		var realPlace string = ""

		if err != nil {
			log.Println("GEN COMP RESPONSE::: FAILED TO CALC LEADERBOARD PLACE", err.Error())
			realPlace = ""
		} else {
			realPlace = place
		}

		resCompList = append(resCompList, ResponseCompetition{
			ID:                   comp.ID,
			Title:                comp.Title,
			Description:          comp.Description,
			Scheduled:            &comp.Scheduled,
			Participant_types:    comp.ParticipantTypes,
			Type:                 comp.Type,
			Status:               comp.Status,
			Start:                comp.Start,
			End:                  comp.End,
			Public:               &comp.Public,
			Workout_types:        comp.WorkoutTypes,
			Profile_participants: profile_participants,
			Owner:                utils.GeneratePartialProfileResponse(c, comp.Edges.Owner),
			Place:                realPlace,
			Distance_total:       distance,
			Participant_count:    comp.QueryParticipants().CountX(c),

			Joined:    profCompLookup[comp.ID],
			CreatedAt: comp.CreatedAt,
			UpdatedAt: comp.UpdatedAt,
		})

	}

	return resCompList

}

func GenerateCompetitionResponse(c *gin.Context, comp *ent.Competition) ResponseCompetition {

	if comp.Edges.Owner == nil {
		log.Println("GEN COPM RESPONSE:::OWNER WAS NIL")
		comp.Edges.Owner = comp.QueryOwner().FirstX(c)
	}

	prof := utils.ProfileFromCtx(c)

	allProfCompList, _ := s.Db.Profile.Query().Where(profile.ID(prof.ID)).QueryCompetitions().All(c)

	profCompLookup := make(map[uuid.UUID]bool)

	for _, joinedComp := range allProfCompList {
		profCompLookup[joinedComp.ID] = true
	}

	place, distance, err := CalcLeaderboardPlace(c, prof.ID.String(), comp.ID)

	var realPlace string = ""

	if err != nil {
		// log.Println("GEN COMP RESPONSE::: FAILED TO CALC LEADERBOARD PLACE", err.Error())
		realPlace = ""
	} else {
		realPlace = place
	}

	var resComp ResponseCompetition = ResponseCompetition{
		ID:                comp.ID,
		Title:             comp.Title,
		Description:       comp.Description,
		Scheduled:         &comp.Scheduled,
		Participant_types: comp.ParticipantTypes,
		Type:              comp.Type,
		Status:            comp.Status,
		Start:             comp.Start,
		End:               comp.End,
		Public:            &comp.Public,
		Workout_types:     comp.WorkoutTypes,
		Owner:             utils.GeneratePartialProfileResponse(c, comp.Edges.Owner),
		Place:             realPlace,
		Distance_total:    distance,
		Joined:            profCompLookup[comp.ID],
		CreatedAt:         comp.CreatedAt,
		UpdatedAt:         comp.UpdatedAt,
		Participant_count: comp.QueryParticipants().CountX(c),
	}

	// log.Println("GEN COMP RESPONSE::: res", resComp)

	return resComp

}

func CalcLeaderboardPlace(c *gin.Context, prof string, comp_id uuid.UUID) (string, float64, error) {
	var res []LeaderBoardEntry = []LeaderBoardEntry{}

	err := s.Db.Profile.Query().
		GroupBy(
			profile.FieldID, profile.FieldFirstName, profile.FieldLastName, profile.FieldPicture,
		).
		Aggregate(
			func(s *sql.Selector) string {

				workoutDataTable := sql.Table(workoutdata.Table)

				s.Join(workoutDataTable).On(s.C(profile.FieldID), workoutDataTable.C(workoutdata.ProfileColumn))

				s.Where(
					sql.EQ(workoutdata.CompetitionColumn, comp_id.String()),
				)

				s.OrderBy(sql.Desc(sql.Sum(workoutDataTable.C(workoutdata.FieldDistance))))

				return sql.As(sql.Sum(workoutDataTable.C(workoutdata.FieldDistance)), "distance_total")

			},
		).Scan(c, &res)

	if err != nil {
		// log.Println("CALC LEADERBOARD PLACE::: FAILED TO QUERY LEADERBOARD ENTRIES", err.Error())
		return "", 0, err

	}

	for i, entry := range res {
		// log.Println("Leaderboard entry", fmt.Sprintf("%d", i+1), entry.First_name, entry.Last_name, entry.Distance_total)

		if entry.ID == prof {
			// log.Println("CALC LEADERBOARD PLACE::: found prof in leaderboard entries", i+1, res)
			return fmt.Sprintf("%d", i+1), entry.Distance_total, nil
			// return i + 1, nil
		}
	}

	// return 0, nil
	return "", 0, nil

}
