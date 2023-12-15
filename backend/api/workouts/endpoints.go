package workouts

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/invite"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"

	"strings"

	u "github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

// func handlePagination(c *gin.Context, q *ent.WorkoutQuery) (*ent.WorkoutQuery, error) {

// 	// current_page := u.Sfq(c, "page")
// 	// current_limit := u.Sfq(c, "limit")

// 	return q, nil
// }

func handleWorkoutFilters(c *gin.Context, q *ent.WorkoutQuery) (*ent.WorkoutQuery, error) {
	// pull out query params here, and add them to the query

	// if s := u.Sfq(c, "status"); s != "" {
	// 	q.Where(workout.Status(c.Query("status")))
	// }

	if c, ok := u.UuidFromQuery(c, "competition"); ok == nil {

		q.Where(workout.HasCompetitionWith(competition.ID(c)))

	}

	if for_profile, ok := u.UuidFromQuery(c, "for_profile"); ok == nil {

		q.Where(
			workout.HasWorkoutDataWith(
				workoutdata.HasProfileWith(
					profile.ID(for_profile),
				),
			),
		).
			WithWorkoutData(
				func(wdq *ent.WorkoutDataQuery) {
					wdq.
						WithCompetition(
							func(cq *ent.CompetitionQuery) {
								cq.WithOwner()
							},
						).
						WithProfile(
							func(pq *ent.ProfileQuery) {
								pq.Select("id", "first_name", "last_name", "picture", "public")
							},
						)
				},
			)

	}

	if lead_by, err := u.UuidFromQuery(c, "lead_by"); err == nil {
		s.L.Println("Error getting lead_by query param: ", err)
		q.Where(workout.HasLeaderWith(
			profile.ID(lead_by),
		))
	}

	// incomplete workouts -> current prof has no workout data for it yet

	if incomplete := u.Sfq(c, "incomplete"); incomplete == "true" {
		// completed workout is when the profile has submitted workout data for the workout

		s.L.Println("handling incomplete workouts")

		q.Where(
			workout.Not(
				workout.HasWorkoutDataWith(
					workoutdata.HasProfileWith(
						profile.ID(u.ProfileFromCtx(c).ID),
					),
				),
			),
		)

	}

	if accepted_invites := u.Sfq(c, "accepted_invites"); accepted_invites == "true" {
		// pending workout is when the profile has accepted an invite to the workout, but has not submitted workout data

		s.L.Println("handling accepted_invites workouts")

		q.Where(
			workout.And(
				workout.HasInviteWith(
					invite.And(
						invite.HasReceiverWith(
							profile.ID(u.ProfileFromCtx(c).ID),
						),
						invite.Status("accepted"),
					),
				),
				workout.Not(
					workout.HasWorkoutDataWith(
						workoutdata.HasProfileWith(
							profile.ID(u.ProfileFromCtx(c).ID),
						),
					),
				),
			),
		)

	}

	// * Get applicable workouts for a competition
	/*
		1. have a activity type that is included in the competitions activity types
		2. must be < 30 days old
		3. must NOT be already applied to another one of the users competitions (active OR ended) because we dont want a user finishing a comp, then adding that same workout towards another comp. single application only
	*/

	// applicable for comp id
	if comp_id, err := u.UuidFromQuery(c, "applicable_for_comp"); err == nil {
		s.L.Println("handling applicable_for_comp workouts")

		// get the competition
		comp, err := s.Db.Competition.Get(c, comp_id)

		if err != nil {
			s.L.Println("FAILED to get comp for id ", comp_id, " err: ", err)
			return q, err
		}

		s.L.Println("activity types", comp.WorkoutTypes)

		q.Where(
			workout.And(

				workout.HealthkitWorkoutActivityTypeIn(
					comp.WorkoutTypes...,
				),

				workout.HasWorkoutDataWith(
					workoutdata.And(

						workoutdata.HealthkitWorkoutEndDateGTE(
							time.Now().AddDate(0, 0, -30),
						),

						workoutdata.HasProfileWith(
							profile.ID(u.ProfileFromCtx(c).ID),
						),

						workoutdata.Not(
							workoutdata.HasCompetition(),
						),
					),
				),
			),
		).
			WithWorkoutData(
				func(wdq *ent.WorkoutDataQuery) {
					wdq.
						WithProfile(
							func(pq *ent.ProfileQuery) {
								pq.Select("id", "first_name", "last_name", "picture", "public")
							},
						)
				},
			)

	}

	return q, nil

}

func GetWorkoutById(c *gin.Context) {
	// prof := u.ProfileFromCtx(c)

	workout_id, err := u.UuidFromPath(c, "workout_id")

	log.Println("GETWORKOUTBY ID:::workout id: ", workout_id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		log.Println("GETWORKOUTBY ID:::Error getting workout id from path: ", err)
		return
	}

	workout, err := s.Db.Workout.Query().
		Where(
			workout.And(
				workout.ID(workout_id),
				// workout.HasLeaderWith(profile.ID(prof.ID)),
			),
		).
		WithLeader(
			func(q *ent.ProfileQuery) {
				q.Select("id", "first_name", "last_name", "picture", "public")
			},
		).
		WithCompetition(
			func(q *ent.CompetitionQuery) {
				q.WithOwner()
			},
		).
		Only(c)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		log.Println("GETWORKOUTBY ID:::Error getting workout: ", err)
		return
	}

	response, err := GenerateFullWorkoutResForProf(c, workout)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error generating full workout response: ", err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InviteToWorkout(c *gin.Context) {
	/*
		 1.  split query param "send_to" by ,
		 2. for each of those ID's,
		 	1. Do they already have an invite to this workout that is pending?
			if so do not create another
			2. if they dont, create the invite, and send the notification
	*/

	prof := u.ProfileFromCtx(c)

	workout_id, err := u.UuidFromPath(c, "workout_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("workout id: ", workout_id)

	profile_id_strings := u.Sfq(c, "send_to")

	log.Println("profile id strings: ", profile_id_strings)

	var profUUIDList []uuid.UUID

	for _, strProfId := range strings.Split(profile_id_strings, ",") {

		parsed, err := uuid.Parse(strProfId)

		if err != nil {
			log.Println("adding to profUUIDList failed: ", err)
		} else {
			profUUIDList = append(profUUIDList, parsed)
		}

	}

	log.Println("done parsing profile id strings: ")

	if len(profUUIDList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no send_to query param"})
		log.Println("no send_to query param")
		return
	}

	for _, idToInvite := range profUUIDList {

		log.Println("\n\nstarting main loop for id: ", idToInvite)

		inviteAcceptedOrPending, err := s.Db.Invite.Query().Where(
			invite.And(
				invite.Or(
					invite.Status("pending"),
					invite.Status("accepted"),
				),
				invite.HasReceiverWith(
					profile.ID(idToInvite),
				),
				invite.HasWorkoutWith(workout.ID(workout_id)),
			),
		).
			Exist(c)

		log.Println("\n\nID::: ", idToInvite, "inviteAcceptedOrPending::: ", inviteAcceptedOrPending, "err::: ", err)

		if err != nil {
			log.Println("INVITETOWORKOUT:::failed to check  if they had an invite existing", err)
		}

		if err == nil && !inviteAcceptedOrPending {
			// no existing invite OR accepted invite, good to send
			log.Println("INVITETOWORKOUT:::no existing invite OR accepted invite, good to send for id: ", idToInvite)

			created_invite, err := s.Db.Invite.Create().
				SetSender(prof).
				SetReceiverID(idToInvite).
				SetWorkoutID(workout_id).
				SetStatus("pending").
				Save(c)

			if err != nil {
				log.Println("INVITETOWORKOUT:::failed to create invite", err)
			} else {
				// send notification

				title := "You've been invited to a workout by " + prof.FirstName + " " + prof.LastName + "!"
				body := "Tap to view"
				data := shared.NotificationData{
					Workout_id:   &workout_id,
					Invite_id:    &created_invite.ID,
					Sender_id:    &prof.ID,
					Sender_image: prof.Picture,
					Receiver_id:  &idToInvite,
				}

				profile_to_send_noti_to, err := s.Db.Profile.Get(c, idToInvite)

				if err != nil {
					log.Println("INVITETOWORKOUT:::failed to get profile", err)
				}

				log.Println("About to call Send push notification!!!!!!!!!!!!!!!")
				u.SendPushNotification(c, nil, profile_to_send_noti_to, title, body, data)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": "Invites sent"})
}

func CreateWorkout(c *gin.Context) {

	prof := u.ProfileFromCtx(c)

	var req WorkoutCreateRequest

	if bindErrs := u.BindAndValidate(c, &req); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return

	}

	w, err := s.Db.Workout.Create().
		SetHealthkitWorkoutActivityType(strings.ToLower(req.Healthkit_workout_activity_type)).
		SetLeader(prof).
		Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	response, err := GenerateFullWorkoutResForProf(c, w)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response)
}

func SubmitWorkoutData(c *gin.Context) {
	/*
	 User submitting workout data to a workout after they are done working out

	 * Constraints: Who can submit workout data
	 1. leader of the workout
	 2. Users who have accepted an invite to that workout
	*/

	prof := u.ProfileFromCtx(c)
	workout_id_submitting_to, err := u.UuidFromPath(c, "workout_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error getting workout id from path: ", err)
		return
	}

	var req WorkoutDataSubmission

	if bindErrs := u.BindAndValidate(c, &req); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		log.Println("Error binding workout data submission: ", bindErrs)
		return
	}

	// does user already have workout data for this workout?
	// if so, 400

	already_submitted, err := s.Db.WorkoutData.Query().
		Where(
			workoutdata.And(
				workoutdata.HasProfileWith(profile.ID(prof.ID)),
				workoutdata.HasWorkoutWith(workout.ID(workout_id_submitting_to)),
			),
		).Exist(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error checking if user has already submitted workout data: ", err)
		return
	}

	if already_submitted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User has already submitted workout data"})
		log.Println("User has already submitted workout data")
		return
	}

	isLeader, err := s.Db.Workout.Query().
		Where(
			workout.And(
				workout.HasLeaderWith(profile.ID(prof.ID)),
				workout.ID(workout_id_submitting_to),
			),
		).Exist(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error checking if user is leader: ", err)
		return
	}

	hasAcceptedInvite, err := s.Db.Invite.Query().
		Where(
			invite.And(
				invite.HasReceiverWith(
					profile.ID(prof.ID),
				),
				invite.Status(
					"accepted",
				),
				invite.HasWorkoutWith(
					workout.ID(workout_id_submitting_to),
				),
			),
		).Exist(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error checking if user has accepted invite: ", err)
		return
	}

	if isLeader || hasAcceptedInvite {
		// log.Println("User is leader or has accepted invite, is leader: ", isLeader, " has accepted invite: ", hasAcceptedInvite)
		_, err := s.Db.WorkoutData.Create().
			SetHealthkitWorkoutID(req.Healthkit_workout_id).
			SetHealthkitWorkoutStartDate(req.Healthkit_workout_start_date).
			SetHealthkitWorkoutEndDate(req.Healthkit_workout_end_date).
			SetDistance(req.Distance).
			SetDuration(req.Duration).
			SetEnergyBurned(req.Energy_burned).
			SetAverageHeartRate(req.Average_heart_rate).
			SetSource(req.Source).
			SetLocationType(req.Location_type).
			SetWorkoutID(workout_id_submitting_to).
			SetProfileID(prof.ID).
			SetWeather(req.Weather).
			Save(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println("Error creating workout data: ", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "Workout data submitted"})
		return
	}

}

func ListWorkout(c *gin.Context) {
	// prof := u.ProfileFromCtx(c)

	base_query := s.Db.Workout.Query()

	base_query, err := handleWorkoutFilters(c, base_query)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		s.L.Println("Error handling workout filters: ", err)
		return
	}

	// s.Scream.Println("base query after: ", base_query)

	workouts, err := base_query.
		WithLeader(
			func(q *ent.ProfileQuery) {
				q.Select("id", "first_name", "last_name", "picture", "public")
			},
		).All(c)
	// All(c)

	s.L.Println("Workouts for listing: ", len(workouts))

	// log.Println("Workouts for listing: ", workouts)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		log.Println("Error getting workouts: ", err)
		return
	}

	var response []CompleteWorkoutResponse = []CompleteWorkoutResponse{}

	for _, w := range workouts {

		// log.Println("Workout: ", w.ID, " leader: ", w.Edges.Leader.ID)

		if w.Edges.Leader == nil {
			s.L.Println("Failed to get leader on workout : ", w.ID)
			continue
		}

		to_append, err := GenerateFullWorkoutResForProf(c, w)

		if err != nil {
			log.Println("Error generating full workout response: ", err)
			continue
		}

		response = append(response, *to_append)

	}

	c.JSON(http.StatusOK, response)
}

func AcceptWorkoutInvite(c *gin.Context) {
	prof := u.ProfileFromCtx(c)

	s.L.Println("Accepting workout invite called")

	request_workout_id, err := u.UuidFromPath(c, "workout_id")

	s.L.Println("request_workout_id: ", request_workout_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		s.L.Println("Error getting workout id from path: ", err)
		return
	}
	request_invite_id, err := u.UuidFromPath(c, "invite_id")

	s.L.Println("request_invite_id: ", request_invite_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		s.L.Println("Error getting invite id from path: ", err)

		return
	}

	invite_exists, err := s.Db.Invite.Query().
		Where(
			invite.And(
				invite.ID(request_invite_id),
				invite.HasReceiverWith(profile.ID(prof.ID)),
				invite.HasWorkoutWith(workout.ID(request_workout_id)),
				invite.Status("pending"))).
		Exist(c)

	if !invite_exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invite does not exist"})

		s.L.Println("Invite does not exist")
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		s.L.Println("Error checking if invite exists: ", err)

		return
	}

	_, err = s.Db.Invite.UpdateOneID(request_invite_id).SetStatus("accepted").Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		s.L.Println("Error updating invite status: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Invite accepted"})
}

func DeclineWorkoutInvite(c *gin.Context) {
	prof := u.ProfileFromCtx(c)

	request_workout_id, err := u.UuidFromPath(c, "workout_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request_invite_id, err := u.UuidFromPath(c, "invite_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invite_exists, err := s.Db.Invite.Query().
		Where(
			invite.And(
				invite.ID(request_invite_id),
				invite.HasReceiverWith(profile.ID(prof.ID)),
				invite.HasWorkoutWith(workout.ID(request_workout_id)),
				invite.StatusIn("pending", "accepted"))).
		Exist(c)

	if !invite_exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invite does not exist"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = s.Db.Invite.UpdateOneID(request_invite_id).SetStatus("declined").Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Invite declined"})
}

func GetMemberInviteStatusesByWorkout(c *gin.Context) {
	workout_id, err := u.UuidFromPath(c, "workout_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		s.L.Println("Error getting for_workout query param: ", err)
		return
	}

	// get all invites for this workout

	invites, err := s.Db.Invite.Query().
		Where(
			invite.HasWorkoutWith(workout.ID(workout_id)),
		).
		WithReceiver(
			func(q *ent.ProfileQuery) {
				q.Select("id", "first_name", "last_name", "picture", profile.FieldPublic)
			},
		).
		All(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("Error getting invites for workout: ", err)
		return
	}

	var response []shared.PartialProfileWithInviteStatus = []shared.PartialProfileWithInviteStatus{}

	for _, invite := range invites {

		response = append(response, shared.PartialProfileWithInviteStatus{
			ID:           invite.Edges.Receiver.ID,
			FirstName:    invite.Edges.Receiver.FirstName,
			LastName:     invite.Edges.Receiver.LastName,
			Picture:      invite.Edges.Receiver.Picture,
			InviteStatus: invite.Status,
			Public:       invite.Edges.Receiver.Public,
		})

	}

	c.JSON(http.StatusOK, response)

}

func ImportWorkoutData(c *gin.Context) {

	prof := u.ProfileFromCtx(c)

	var req WorkoutDataImportRequest

	if bindErrs := u.BindAndValidate(c, &req); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		s.L.Println("Error binding workout data import: ", bindErrs)
		return
	}

	// 1. Create a workout for this data
	// 2. apply the data from the res to the workout
	// 3. generate complete workout response

	created_workout, err := s.Db.Workout.Create().
		SetHealthkitWorkoutActivityType(strings.ToLower(req.Healthkit_workout_activity_type)).
		SetLeader(prof).
		Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		s.L.Println("Error creating workout: ", err)
		return
	}

	_, err = s.Db.WorkoutData.Create().
		SetHealthkitWorkoutID(req.Healthkit_workout_id).
		SetHealthkitWorkoutStartDate(req.Healthkit_workout_start_date).
		SetHealthkitWorkoutEndDate(req.Healthkit_workout_end_date).
		SetDistance(req.Distance).
		SetDuration(req.Duration).
		SetEnergyBurned(req.Energy_burned).
		SetAverageHeartRate(req.Average_heart_rate).
		SetSource(req.Source).
		SetLocationType(req.Location_type).
		SetWorkoutID(created_workout.ID).
		SetProfileID(prof.ID).
		Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error creating workout data: ", err)
		return
	}

	response, err := GenerateFullWorkoutResForProf(c, created_workout)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error generating full workout response: ", err)
		return
	}

	c.JSON(http.StatusOK, response)

}

func GetWorkoutData(c *gin.Context) {
	workout_id, err := u.UuidFromPath(c, "workout_id")
	prof := u.ProfileFromCtx(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		s.L.Println("Error getting workout id from path: ", err)
		return
	}

	workout_data, err := s.Db.WorkoutData.Query().
		Where(
			workoutdata.And(
				workoutdata.HasWorkoutWith(workout.ID(workout_id)),
				workoutdata.Not(
					workoutdata.HasProfileWith(
						profile.ID(prof.ID),
					),
				),
			),
		).
		WithProfile(
			func(q *ent.ProfileQuery) {
				q.Select("id", "first_name", "last_name", "picture", "public")
			},
		).
		All(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get workout data"})
		s.L.Println("Error getting workout data: ", err)
		return
	}

	res := GenerateWorkoutDataListRes(c, workout_data)

	c.JSON(http.StatusOK, res)

}

func DeleteWorkout(c *gin.Context) {
	workout_id, err := u.UuidFromPath(c, "workout_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		s.L.Println("Error getting workout id from path: ", err)
		return
	}

	_, err = s.Db.Workout.Delete().Where(workout.ID(workout_id)).Exec(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete workout"})
		s.L.Println("Error deleting workout: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Workout deleted"})

}

func RemoveAcceptedInvite(c *gin.Context) {
	workout_id, err := u.UuidFromPath(c, "workout_id")
	prof := u.ProfileFromCtx(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		s.L.Println("Error getting workout id from path: ", err)
		return
	}

	invite_to_remove, err := s.Db.Invite.Query().
		Where(
			invite.And(
				invite.HasReceiverWith(profile.ID(prof.ID)),
				invite.HasWorkoutWith(workout.ID(workout_id)),
				invite.Status("accepted"))).
		Select(invite.FieldID).
		First(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found"})
		return
	}

	_, err = s.Db.Invite.UpdateOneID(invite_to_remove.ID).SetStatus("declined").Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Invite removed"})

}
