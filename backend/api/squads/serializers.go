package squads

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competitionresult"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

type ResponseSquad struct {
	ID           uuid.UUID               `json:"id"`
	Title        string                  `json:"title"`
	CreatedAt    time.Time               `json:"created_at"`
	UpdatedAt    time.Time               `json:"updated_at"`
	Owner        shared.PartialProfile   `json:"owner"`
	Members      []shared.PartialProfile `json:"members"`
	Joined       bool                    `json:"joined"`
	Public       bool                    `json:"public"`
	Stats        *SquadStatsResponse     `json:"stats"`
	Member_count int                     `json:"member_count"`
}
type SquadStatsResponse struct {
	MemberCount     int     `json:"member_count"`
	CompetitionWins int     `json:"competition_wins"`
	TotalDistanceFt float64 `json:"total_distance_ft"`
}

type CreateSquadStruct struct {
	Title string `json:"title" binding:"required"`
}

func generateSquadResponse(c *gin.Context, dbSquad *ent.Squad) ResponseSquad {

	prof := utils.ProfileFromCtx(c)
	prof_squads, _ := s.Db.Profile.Query().Where(profile.ID(prof.ID)).QuerySquad().All(c)

	profJoinedLookup := make(map[uuid.UUID]bool)

	for _, ps := range prof_squads {
		profJoinedLookup[ps.ID] = true
	}

	var res ResponseSquad

	var stats *SquadStatsResponse = nil

	stats = generateSquadStatResponse(c, dbSquad)

	// make sure owner and members edges are loaded

	var member_list []shared.PartialProfile = []shared.PartialProfile{}

	// ensure that members and owner are loaded
	// dbSquad.Edges.Owner = dbSquad.QueryOwner().OnlyX(c)
	// dbSquad.Edges.Members = dbSquad.QueryMembers().AllX(c)
	var err error

	if dbSquad.Edges.Owner == nil {
		dbSquad.Edges.Owner, err = dbSquad.QueryOwner().Only(c)
		if err != nil {
			dbSquad.Edges.Owner = &ent.Profile{
				ID:        uuid.New(),
				FirstName: "Missing",
				LastName:  "Owner Info",
				Picture:   "https://lh3.googleusercontent.com/a/AAcHTtdX9WVfVj5dWY5vVLyJozVxIdjY76d7qLJjgI7u2a4TK1xU=s20-c",
				Public:    true,
			}

		}

	}

	if dbSquad.Edges.Members == nil {
		dbSquad.Edges.Members, _ = dbSquad.QueryMembers().All(c)
	}

	for _, member := range dbSquad.Edges.Members {
		member_list = append(member_list, shared.PartialProfile{
			ID:        member.ID,
			FirstName: member.FirstName,
			LastName:  member.LastName,
			Picture:   member.Picture,
			Public:    &member.Public,
		})
	}

	res = ResponseSquad{
		ID:           dbSquad.ID,
		Title:        dbSquad.Title,
		Public:       dbSquad.Public,
		Joined:       profJoinedLookup[dbSquad.ID],
		Members:      member_list,
		Owner:        utils.GeneratePartialProfileResponse(c, dbSquad.Edges.Owner),
		CreatedAt:    dbSquad.CreatedAt,
		UpdatedAt:    dbSquad.UpdatedAt,
		Stats:        stats,
		Member_count: dbSquad.QueryMembers().CountX(c),
	}

	return res

}

func generateSquadListResponse(c *gin.Context, squadsInput []*ent.Squad) []ResponseSquad {

	prof := utils.ProfileFromCtx(c)
	prof_squads, _ := s.Db.Profile.Query().Where(profile.ID(prof.ID)).QuerySquad().All(c)

	profJoinedLookup := make(map[uuid.UUID]bool)

	for _, ps := range prof_squads {
		profJoinedLookup[ps.ID] = true
	}

	var res []ResponseSquad = []ResponseSquad{}

	// make sure owner and members edges are loaded

	for _, dbSquad := range squadsInput {

		var member_list []shared.PartialProfile = []shared.PartialProfile{}
		var err error

		if dbSquad.Edges.Owner == nil {
			dbSquad.Edges.Owner, err = dbSquad.QueryOwner().Only(c)
			if err != nil {
				dbSquad.Edges.Owner = &ent.Profile{
					ID:        uuid.New(),
					FirstName: "Missing",
					LastName:  "Owner Info",
					Picture:   "https://lh3.googleusercontent.com/a/AAcHTtdX9WVfVj5dWY5vVLyJozVxIdjY76d7qLJjgI7u2a4TK1xU=s20-c",
					Public:    true,
				}

			}

		}

		if dbSquad.Edges.Members == nil {
			dbSquad.Edges.Members, _ = dbSquad.QueryMembers().All(c)
		}
		for _, member := range dbSquad.Edges.Members {
			member_list = append(member_list, utils.GeneratePartialProfileResponse(c, member))
		}

		var stats *SquadStatsResponse = nil
		stats = generateSquadStatResponse(c, dbSquad)

		res = append(res, ResponseSquad{

			ID:      dbSquad.ID,
			Title:   dbSquad.Title,
			Public:  dbSquad.Public,
			Joined:  profJoinedLookup[dbSquad.ID],
			Members: member_list,

			Owner:     utils.GeneratePartialProfileResponse(c, dbSquad.Edges.Owner),
			CreatedAt: dbSquad.CreatedAt,
			UpdatedAt: dbSquad.UpdatedAt,
			Stats:     stats,

			Member_count: dbSquad.QueryMembers().CountX(c),
		})

	}

	return res
}

func generateSquadStatResponse(c *gin.Context, dbSquad *ent.Squad) *SquadStatsResponse {

	//! Member count
	member_count := len(dbSquad.Edges.Members)

	//! Competition wins
	comp_wins, err := dbSquad.QueryCompetitionResults().
		Where(
			competitionresult.Place("1"),
		).
		Count(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		s.L.Println("ERROR GETTING COMP WINS")
		return nil
	}

	//! Total Distance -> for each member of the squad, aggre. all distance for their workoutDatas
	total_distance_ft, err := dbSquad.QueryMembers().
		QueryWorkoutData().
		Aggregate(
			ent.Sum(workoutdata.FieldDistance),
		).
		Float64(c)

	if err != nil {
		// c.AbortWithStatusJSON(http.StatusOK, gin.H{"errors": err})

		total_distance_ft = 0.0
		s.L.Println("ERROR GETTING SQUAD total distance ft, but continuing", err.Error())
	}

	squad_stats_response := SquadStatsResponse{
		MemberCount:     member_count,
		CompetitionWins: comp_wins,
		TotalDistanceFt: total_distance_ft,
	}

	return &squad_stats_response

}
