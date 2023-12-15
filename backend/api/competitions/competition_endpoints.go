package competitions

import (
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/invite"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	u "github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func AddCompetitionRoutes(rg *gin.RouterGroup) {
	compRoutes := rg.Group("/competitions")

	compRoutes.GET("/", ListCompetition)
	compRoutes.POST("/", CreateCompetition)
	compRoutes.PATCH("/", UpdateCompetition)
	compRoutes.GET("/:competition_id", GetCompetition)

	// compRoutes.GET("/:comp_id/participants", ListParticipantsByCompID)

	compRoutes.POST("/:competition_id/accept_invite/:invite_id", AcceptCompetitionInvite)
	compRoutes.POST("/:competition_id/decline_invite/:invite_id", DeclineCompetitionInvite)
	compRoutes.POST("/:competition_id/invite/:profile_id_to_invite", CreateCompetitionInvite)

	compRoutes.POST("/:competition_id/join", JoinComp)
	compRoutes.POST("/:competition_id/leave", LeaveComp)
	compRoutes.POST("/:competition_id/apply_workout_data", ApplyWorkoutDataToComp)
	compRoutes.GET("/:competition_id/participants", GetCompetitionParticipants)

	compRoutes.GET("/:competition_id/leaderboard", GetLeaderboard)
	compRoutes.GET("/:competition_id/graph", GetGraphData)
	compRoutes.GET("/:competition_id/stats", GetStats)

	compRoutes.POST("/:competition_id/invite", CreateCompetitionInvite)

}

func handleCompetitionFilters(c *gin.Context, cq *ent.CompetitionQuery) *ent.CompetitionQuery {

	prof := u.ProfileFromCtx(c)

	if owned := u.BoolFromQuery(c, "owned"); owned {
		cq.Where(
			competition.HasOwnerWith(profile.ID(prof.ID)),
		)
	}

	if joined := u.BoolFromQuery(c, "joined"); joined {
		cq.Where(
			competition.HasParticipantsWith(profile.ID(prof.ID)),
		)
	}

	if public := u.BoolFromQuery(c, "public"); public {
		cq.Where(
			competition.Public(public),
		)
	}

	if scheduled := u.BoolFromQuery(c, "scheduled"); scheduled {
		cq.Where(
			competition.Scheduled(scheduled),
		)
	}

	if for_participant, err := u.UuidFromQuery(c, "for_participant"); err == nil {
		cq.Where(
			competition.HasParticipantsWith(profile.ID(for_participant)),
		)
	}

	return cq

}

func GetCompetition(c *gin.Context) {

	comp_id, err := u.UuidFromPath(c, "competition_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comp, err := s.Db.Competition.Query().
		Where(
			competition.ID(comp_id),
		).
		WithOwner().
		// WithParticipants().
		// WithWorkoutData(func(wdq *ent.WorkoutDataQuery) {
		// 	wdq.WithWorkout()
		// }).
		First(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resComp := GenerateCompetitionResponse(c, comp)

	c.JSON(http.StatusOK, resComp)

}

func CreateCompetitionInvite(c *gin.Context) {

	sender_profile := u.ProfileFromCtx(c)

	request_competition_id, err := u.UuidFromPath(c, "competition_id")

	prof_ids_to_invite := u.Sfq(c, "send_to")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var uuids_to_inv []uuid.UUID

	for _, s := range strings.Split(prof_ids_to_invite, ",") {

		u, err := uuid.Parse(s)

		if err == nil {
			uuids_to_inv = append(uuids_to_inv, u)
		}

	}

	for _, to_recieve_id := range uuids_to_inv {
		invite_already_exists, err := s.Db.Invite.Query().
			Where(
				invite.And(
					invite.HasSenderWith(profile.ID(sender_profile.ID)),
					invite.HasReceiverWith(profile.ID(to_recieve_id)),
					invite.HasCompetitionWith(competition.ID(request_competition_id)),
					invite.Status("pending"),
				)).
			Exist(c)

		if invite_already_exists {
			s.L.Printf("Invite already exists, skipping for id: %s ", to_recieve_id.String())
			continue
		}

		if err != nil {
			s.L.Println("Error checking if invite exists", err)
			continue
		}

		// cant invite a user to a competition they are already in
		already_in_competition, err := s.Db.Competition.Query().
			Where(
				competition.And(
					competition.ID(request_competition_id),
					competition.HasParticipantsWith(profile.ID(to_recieve_id)),
				)).
			Exist(c)

		if err != nil {
			s.L.Println("Error checking if user is already in competition", err)
			continue
		}

		if already_in_competition {
			s.L.Printf("User already in competition, skipping for id: %s ", to_recieve_id.String())
			continue

		}

		created_invite, err := s.Db.Invite.Create().
			SetReceiverID(to_recieve_id).
			SetSenderID(sender_profile.ID).
			SetCompetitionID(request_competition_id).
			SetStatus("pending").
			Save(c)

		if err != nil {
			s.L.Println("INVITETO comp:::failed to create invite", err)
		} else {
			// send notification

			title := "Invited to a competition by " + sender_profile.FirstName + " " + sender_profile.LastName + "!"
			body := "Tap to view"
			data := shared.NotificationData{
				Competition_id: &request_competition_id,
				Invite_id:      &created_invite.ID,
				Sender_id:      &sender_profile.ID,
				Sender_image:   sender_profile.Picture,
				Receiver_id:    &to_recieve_id,
			}

			profile_to_send_noti_to, err := s.Db.Profile.Get(c, to_recieve_id)

			if err != nil {
				s.L.Println("invite to competition:::failed to get profile", err)
			}

			u.SendPushNotification(c, nil, profile_to_send_noti_to, title, body, data)
		}

		if err != nil {
			s.L.Printf("Error creating invite for id: %s %s ", to_recieve_id.String(), err)
			continue
		}

	}

	c.JSON(http.StatusOK, gin.H{"success": "Invite Created"})

}

func AcceptCompetitionInvite(c *gin.Context) {

	prof_accepting := u.ProfileFromCtx(c)

	request_competition_id, err := u.UuidFromPath(c, "competition_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error getting competition id from path", err)
		return
	}

	request_invite_id, err := u.UuidFromPath(c, "invite_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error getting invite id from path", err)
		return
	}

	invite_exists, err := s.Db.Invite.Query().
		Where(
			invite.And(
				invite.ID(request_invite_id),
				invite.HasReceiverWith(profile.ID(prof_accepting.ID)),
				invite.HasCompetitionWith(competition.ID(request_competition_id)),
				invite.Status("pending"))).
		Exist(c)

	if !invite_exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invite does not exist"})
		s.L.Println("Invite does not exist")
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error checking if invite exists", err)
		return
	}

	updated_invite, err := s.Db.Invite.UpdateOneID(request_invite_id).SetStatus("accepted").Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error updating invite status", err)
		return
	}

	joined_comp, err := s.Db.Competition.UpdateOneID(request_competition_id).AddParticipants(prof_accepting).Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error adding participant to competition", err)
		return
	}

	title := prof_accepting.FirstName + " " + prof_accepting.LastName + " accepted your invite to " + joined_comp.Title + "!"
	// body := "Tap to view"

	creator, err := updated_invite.QuerySender().Only(c)

	if err != nil {
		s.L.Println("Error getting invite sender", err)
	} else {
		u.SendPushNotification(c, nil, creator, title, "", shared.NotificationData{
			Competition_id: &request_competition_id,
			Invite_id:      &updated_invite.ID,
			Sender_id:      &prof_accepting.ID,
			Sender_image:   prof_accepting.Picture,
			Receiver_id:    &creator.ID,
		})
	}

	c.JSON(http.StatusOK, gin.H{"success": "Invite accepted"})
}

func DeclineCompetitionInvite(c *gin.Context) {
	prof := u.ProfileFromCtx(c)

	request_competition_id, err := u.UuidFromPath(c, "competition_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error getting competition id from path", err)
		return
	}

	request_invite_id, err := u.UuidFromPath(c, "invite_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error getting invite id from path", err)
		return
	}

	invite_exists, err := s.Db.Invite.Query().
		Where(
			invite.And(
				invite.ID(request_invite_id),
				invite.HasReceiverWith(profile.ID(prof.ID)),
				invite.HasCompetitionWith(competition.ID(request_competition_id)),
				invite.Status("pending"))).
		Exist(c)

	if err != nil || !invite_exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invite does not exist"})
		s.L.Println("Invite does not exist")
		return
	}

	_, err = s.Db.Invite.UpdateOneID(request_invite_id).SetStatus("declined").Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error updating invite status", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Invite declined"})
}

// Join competition
func JoinComp(c *gin.Context) {

	s.L.Println("JOINCOMP:::Joining competition")

	comp_id, err := uuid.Parse(c.Param("competition_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	prof := u.ProfileFromCtx(c)

	comp, err := s.Db.Competition.Query().WithOwner().Where(competition.ID(comp_id), competition.Not(competition.HasParticipantsWith(profile.ID(prof.ID)))).First(c)

	// s.Scream.Println("JOINCOMP:::Found competition: ")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to join competition"})
		return
	}

	if comp.End.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Competition has ended"})
		return
	}

	if !comp.Public {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to join competition"})
		return
	}

	updated, err := comp.Update().AddParticipants(prof).Save(c)

	// s.Scream.Println("JOINCOMP:::Updated competition with participant")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to join competition"})
		return
	}

	// s.Scream.Println("JOINCOMP:::Creating response")
	resComp := GenerateCompetitionResponse(c, updated)

	s.L.Println("Updated competition: ", resComp)

	c.JSON(http.StatusOK, resComp)

}

func LeaveComp(c *gin.Context) {

	comp_id, err := uuid.Parse(c.Param("competition_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		s.L.Println("error getting competition id", err)
		return
	}

	prof := u.ProfileFromCtx(c)

	compToLeave, err := s.Db.Competition.Query().WithOwner().Where(competition.ID(comp_id), competition.HasParticipantsWith(profile.ID(prof.ID))).Only(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to leave competiton, not joined"})
		s.L.Println("error leaving competition", err)
		return
	}

	if compToLeave.End.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Competition has ended"})
		return
	}

	updated, errr := compToLeave.Update().RemoveParticipants(prof).Save(c)

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to leave competition"})
		s.L.Println("error leaving competition", errr)
		return
	}

	s.L.Println("left competition: ", updated.ID.String())

	c.JSON(http.StatusOK, gin.H{"success": "Left Competition"})

}

//list my competitions

func ListCompetition(c *gin.Context) {
	var comps []*ent.Competition

	// dbspan := sentry.StartSpan(span.Context(), "query competitions")

	// prof := u.ProfileFromCtx(c)

	base := s.Db.Competition.Query().WithOwner().WithParticipants()

	base = handleCompetitionFilters(c, base)

	comps, err := base.Order(ent.Desc(competition.FieldCreatedAt)).All(c)

	if err != nil {
		s.L.Println("error getting competitions", err)
		c.JSON(http.StatusBadRequest, gin.H{"fail": err})
		return
	}

	var resComps []ResponseCompetition = generateCompetitionListResponse(c, comps)

	c.JSON(http.StatusOK, resComps)
}

func GetCompetitionParticipants(c *gin.Context) {

	comp_id, err := u.UuidFromPath(c, "competition_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		s.L.Println("error getting competition id", err)
		return
	}

	comp, err := s.Db.Competition.Query().Where(competition.ID(comp_id)).WithParticipants().Only(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		s.L.Println("error getting competition", err)
		return
	}

	var resComps []shared.PartialProfile = []shared.PartialProfile{}

	for _, prof := range comp.Edges.Participants {
		resComps = append(resComps, u.GeneratePartialProfileResponse(c, prof))

	}

	c.JSON(http.StatusOK, resComps)
}

func UpdateCompetition(c *gin.Context) {

	var comp UpdateCompetitionRequest

	if bindErrs := u.BindAndValidate(c, &comp); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return
	}

	updated_comp, err := s.Db.Competition.
		UpdateOneID(comp.Id).
		SetTitle(comp.Title).
		SetDescription(comp.Description).
		Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"fail": err})
		return
	}

	resComp := GenerateCompetitionResponse(c, updated_comp)

	c.JSON(http.StatusOK, resComp)

}

func CreateCompetition(c *gin.Context) {

	var comp RequestCompetition

	prof := u.ProfileFromCtx(c)

	if bindErrs := u.BindAndValidate(c, &comp); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return
	}

	s.L.Println("Creating competition: ", comp)

	created_comp, err := s.Db.Competition.Create().
		SetOwnerID(prof.ID).
		SetTitle(comp.Title).
		SetDescription(comp.Description).
		SetScheduled(comp.Scheduled).
		SetParticipantTypes(comp.Participant_types).
		SetWorkoutTypes(comp.Workout_types).
		SetType(comp.Type).
		SetStatus(comp.Status).
		SetPublic(comp.Public).
		SetStart(comp.Start).
		SetEnd(comp.End).
		AddParticipantIDs(prof.ID).
		Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"fail": err})
		return
	}

	db_comp, _ := s.Db.Competition.Query().Where(competition.ID(created_comp.ID)).WithOwner().Only(c)
	resComp := GenerateCompetitionResponse(c, db_comp)

	c.JSON(http.StatusOK, resComp)

}

type LeaderBoardEntry struct {
	ID         string `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Picture    string `json:"picture"`

	Distance_total float64 `json:"distance_total"`
}

func GetLeaderboard(c *gin.Context) {
	comp_id, err := u.UuidFromPath(c, "competition_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		s.L.Println("error getting competition id", err)
		return
	}

	var res []LeaderBoardEntry = []LeaderBoardEntry{}

	err = s.Db.Profile.Query().
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
		s.L.Println("error calculating leaderboard", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get competition"})
		return
	}

	// s.Scream.Println("Leaderboard response: ", res)

	c.JSON(http.StatusOK, res)

}

type CompetitionStatChartDataPoint struct {
	Id                      string  `json:"id"`
	HealthkitWorkoutEndDate string  `json:"healthkit_workout_end_date"`
	Distance                float64 `json:"distance"`
}

func GetGraphData(c *gin.Context) {
	// get competition id from path
	comp_id, err := u.UuidFromPath(c, "competition_id")
	prof := u.ProfileFromCtx(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	var res []CompetitionStatChartDataPoint = []CompetitionStatChartDataPoint{}

	err = s.Db.WorkoutData.Query().
		Where(
			workoutdata.HasCompetitionWith(
				competition.ID(comp_id),
			),
			workoutdata.HasProfileWith(profile.ID(prof.ID)),
		).
		Select(
			workoutdata.FieldHealthkitWorkoutEndDate, workoutdata.FieldDistance, workoutdata.FieldID,
		).Scan(c, &res)

	if err != nil {
		s.L.Println("error calculating stats", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get competition"})
		return
	}

	// s.Scream.Println("Stats response: ", len(res))

	c.JSON(http.StatusOK, res)

}

type MyCompStats struct {
	Place          string  `json:"place"`
	Distance_total float64 `json:"distance_total"`
}

func GetStats(c *gin.Context) {
	comp_id, err := u.UuidFromPath(c, "competition_id")
	prof := u.ProfileFromCtx(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param"})
		return
	}

	place, distance, err := CalcLeaderboardPlace(c, prof.ID.String(), comp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param"})
		s.L.Println("error calculating stats", err)
		return
	}

	res := MyCompStats{
		Place:          place,
		Distance_total: distance,
	}

	c.JSON(http.StatusOK, res)

}

type ApplyWorkoutRequest struct {
	Workout_data_ids []uuid.UUID `json:"workout_data_ids" binding:"required,min=1"`
}

func ApplyWorkoutDataToComp(c *gin.Context) {

	var req ApplyWorkoutRequest

	if bindErrs := u.BindAndValidate(c, &req); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return
	}

	s.L.Println("list of workout data ids: ", req.Workout_data_ids)

	prof := u.ProfileFromCtx(c)

	comp_id_path, err := u.UuidFromPath(c, "competition_id")

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		s.L.Println("error getting competition id", err)
		return

	}

	comp_to_add, err := s.Db.Competition.Get(c, comp_id_path)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		s.L.Println("error getting competition", err)
		return

	}

	// check if prof is member of comp
	prof_in_comp, err := comp_to_add.QueryParticipants().Where(profile.ID(prof.ID)).Exist(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		s.L.Println("error getting competition", err)
		return
	}

	if !prof_in_comp {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not in competition"})
		s.L.Println("error getting competition", err)
		return
	}

	// for all ids in req, if that workout data has prof matching our prof, and no comp,
	// add comp to workout data

	for _, workout_data_id := range req.Workout_data_ids {

		workout_exists, err := s.Db.WorkoutData.Query().Where(
			workoutdata.And(
				workoutdata.ID(
					workout_data_id,
				),
				workoutdata.HasProfileWith(profile.ID(prof.ID)),
				// workoutdata.Not(
				// 	workoutdata.HasCompetition(),
				// ),
			),
		).Exist(c)

		if !workout_exists {
			s.L.Println("workout did not exist, continuing")
			continue
		}

		if err != nil {
			s.L.Println("error getting workout data", err)
			continue
		}

		_, err = s.Db.WorkoutData.UpdateOneID(workout_data_id).
			SetCompetitionID(comp_id_path).
			Save(c)

		if err != nil {
			s.L.Println("error updating workout data", err)
			continue
		}

	}

	c.JSON(http.StatusOK, gin.H{"success": "Applied workout data to competition"})

}
