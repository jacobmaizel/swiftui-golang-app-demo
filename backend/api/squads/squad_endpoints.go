package squads

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competitionresult"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/invite"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/squad"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	u "github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func AddSquadRoutes(rg *gin.RouterGroup) {
	SquadRoutes := rg.Group("/squads")

	SquadRoutes.POST("/", CreateSquad)
	SquadRoutes.GET("/", ListSquads)
	SquadRoutes.PATCH("/:squad_id", PatchSquad)
	SquadRoutes.GET("/:squad_id", GetSquad)

	SquadRoutes.POST("/:squad_id/join", JoinSquad)
	SquadRoutes.POST("/:squad_id/leave", LeaveSquad)

	SquadRoutes.POST("/:squad_id/accept_invite/:invite_id", AcceptSquadInvite)
	SquadRoutes.POST("/:squad_id/decline_invite/:invite_id", DeclineSquadInvite)
	SquadRoutes.POST("/:squad_id/invite", CreateSquadInvite)

	SquadRoutes.GET("/:squad_id/stats", GetSquadStats)

	SquadRoutes.DELETE("/:squad_id", DeleteSquad)

}

func handleSquadFilters(c *gin.Context, sq *ent.SquadQuery) *ent.SquadQuery {

	// log.Println("HANDLE SQUAD FILTERS")

	prof := u.ProfileFromCtx(c).ID

	if joined := u.BoolFromQuery(c, "joined"); joined {
		sq.Where(
			squad.HasMembersWith(profile.ID(prof)),
		)
	}

	if public := u.BoolFromQuery(c, "public"); public {
		sq.Where(
			squad.Public(public),
		)
	}

	if owned := u.BoolFromQuery(c, "owned"); owned {
		sq.Where(
			squad.HasOwnerWith(profile.ID(prof)),
		)
	}

	if with_member, err := u.UuidFromQuery(c, "with_member"); with_member != uuid.Nil && err == nil {
		sq.Where(
			squad.HasMembersWith(profile.ID(with_member)),
		)
	}

	return sq
}

func CreateSquadInvite(c *gin.Context) {

	sender_profile := u.ProfileFromCtx(c)

	list_of_profiles_to_invite_path := u.Sfq(c, "send_to")

	request_squad_id, err := u.UuidFromPath(c, "squad_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var profUUIDList []uuid.UUID

	for _, strProfId := range strings.Split(list_of_profiles_to_invite_path, ",") {

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

	for _, receiving_profile_id := range profUUIDList {
		invite_already_exists, err := s.Db.Invite.Query().
			Where(
				invite.And(
					invite.HasSenderWith(profile.ID(sender_profile.ID)),
					invite.HasReceiverWith(profile.ID(receiving_profile_id)),
					invite.HasSquadWith(squad.ID(request_squad_id)),
					invite.Status("pending"),
				)).
			Exist(c)

		if err != nil {
			s.L.Println("error checking for existing invite: ", err)
			continue
		}

		if invite_already_exists {
			s.L.Println("invite already exists")
			continue
		}

		invitee_already_member, err := s.Db.Squad.Query().
			Where(
				squad.And(
					squad.ID(request_squad_id),
					squad.HasMembersWith(profile.ID(receiving_profile_id)),
				)).
			Exist(c)

		if err != nil {
			s.L.Println("error checking for existing invite: ", err)
			continue
		}

		if invitee_already_member {
			s.L.Println("invitee already member")
			continue
		}

		created_invite, err := s.Db.Invite.Create().
			SetReceiverID(receiving_profile_id).
			SetSenderID(sender_profile.ID).
			SetSquadID(request_squad_id).
			SetStatus("pending").
			Save(c)

		if err != nil {
			s.L.Println("error creating invite: ", err)
			continue
		}

		invitee, err2 := s.Db.Profile.Get(c, receiving_profile_id)
		squad_invited_to, err := s.Db.Squad.Get(c, request_squad_id)

		if err != nil || err2 != nil {
			s.L.Println("error getting invitee: ", err)
		} else {
			title := fmt.Sprintf("%s %s invited you to join %s", sender_profile.FirstName, sender_profile.LastName, squad_invited_to.Title)
			u.SendPushNotification(c, nil, invitee, title, "Squad invite", shared.NotificationData{
				Squad_id:     &request_squad_id,
				Receiver_id:  &invitee.ID,
				Sender_id:    &sender_profile.ID,
				Sender_image: sender_profile.Picture,
				Invite_id:    &created_invite.ID,
			})
		}

	}

	c.JSON(http.StatusOK, gin.H{"success": "Invites Sent!"})

}

func AcceptSquadInvite(c *gin.Context) {

	prof := u.ProfileFromCtx(c)

	request_squad_id, err := u.UuidFromPath(c, "squad_id")

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
				invite.HasSquadWith(squad.ID(request_squad_id)),
				invite.Status("pending"))).
		Exist(c)

	if !invite_exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invite does not exist"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if invitee is already a member of the squad
	alreadyMember := s.Db.Squad.Query().
		Where(
			squad.And(
				squad.ID(request_squad_id),
				squad.HasMembersWith(profile.ID(prof.ID)),
			),
		).ExistX(c)

	if alreadyMember {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "ALready a member of this squad"})
		return
	}

	_, err = s.Db.Invite.UpdateOneID(request_invite_id).SetStatus("accepted").Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = s.Db.Squad.UpdateOneID(request_squad_id).AddMembers(prof).Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Invite accepted"})
}

func DeclineSquadInvite(c *gin.Context) {
	prof := u.ProfileFromCtx(c)

	request_squad_id, err := u.UuidFromPath(c, "squad_id")

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
				invite.HasSquadWith(squad.ID(request_squad_id)),
				invite.Status("pending"))).
		Exist(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !invite_exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invite does not exist"})
		return
	}

	_, err = s.Db.Invite.UpdateOneID(request_invite_id).SetStatus("declined").Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Invite declined"})
}

func CreateSquad(c *gin.Context) {

	var newSquad CreateSquadStruct

	if bindErrs := u.BindAndValidate(c, &newSquad); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return
	}

	alreadyExists := s.Db.Squad.Query().Where(squad.Title(newSquad.Title)).ExistX(c)

	if alreadyExists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Squad with that title already exists"})
		s.L.Print("Squad with that title already exists")
		return
	}
	prof := u.ProfileFromCtx(c)

	created_squad, err := s.Db.Squad.Create().SetTitle(newSquad.Title).SetOwnerID(prof.ID).AddMemberIDs(prof.ID).Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	responseSquad := generateSquadResponse(c, created_squad)

	c.JSON(http.StatusOK, responseSquad)

}

func ListSquads(c *gin.Context) {
	base_query := s.Db.Squad.Query()
	// prof := u.ProfileFromCtx(c)

	base := base_query.
		WithMembers(
			func(query *ent.ProfileQuery) {
				query.Select(profile.FieldID, profile.FieldFirstName, profile.FieldLastName, profile.FieldPicture)
			}).
		WithOwner().
		Order(ent.Desc(squad.FieldCreatedAt))

	// log.Println("LIST SQUADS")

	squad_list, err := handleSquadFilters(c, base).All(c)

	if err != nil {
		log.Println("ERROR LISTING SQUADS", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	squad_response_list := generateSquadListResponse(c, squad_list)

	c.JSON(http.StatusOK, squad_response_list)
}

func JoinSquad(c *gin.Context) {

	// If user is already a member of the squad, return error
	prof := u.ProfileFromCtx(c)

	squad_to_join_id, err := u.UuidFromPath(c, "squad_id")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid Squad ID"})
		log.Println("ERROR PARSING SQUAD ID")
		return
	}

	alreadyMember, err := s.Db.Squad.Query().
		Where(
			squad.And(
				squad.ID(squad_to_join_id),
				squad.HasMembersWith(profile.ID(prof.ID)))).
		Exist(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		log.Println("ERROR CHECKING IF ALREADY A MEMBER")
		return
	}

	if alreadyMember {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Already a member of this squad"})
		log.Println("ALREADY A MEMBER OF THIS SQUAD")
		return
	}

	// public, not joined, matching id
	squad_to_join, err := s.Db.Squad.Query().
		Where(
			squad.And(
				squad.ID(squad_to_join_id),
				squad.Public(true),
				squad.Not(squad.HasMembersWith(profile.ID(prof.ID))))).
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": "Squad not found"})
		log.Println("ERROR CHECKING IF SQUAD EXISTS")
		return
	}

	// If user is not already a member of the squad, add them to the squad

	_, err = squad_to_join.Update().AddMembers(prof).Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		log.Println("ERROR ADDING MEMBER TO SQUAD")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully joined squad"})

}

func LeaveSquad(c *gin.Context) {

	// if user is not a member of the squad, return error

	squad_to_leave_id, err := uuid.Parse(c.Param("squad_id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid Squad ID"})
		return
	}

	// check if squad exists
	_, err = s.Db.Squad.Query().Where(squad.ID(squad_to_leave_id)).Only(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": "Squad not found"})
		return
	}

	prof := u.ProfileFromCtx(c)

	alreadyMember := s.Db.Squad.Query().Where(squad.ID(squad_to_leave_id)).Where(squad.HasMembersWith(profile.ID(prof.ID))).ExistX(c)

	if !alreadyMember {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Not a member of this squad"})
		return
	}

	// if user is a member of the squad, remove them from the squad

	_, err = s.Db.Squad.UpdateOneID(squad_to_leave_id).RemoveMembers(prof).Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully left squad"})

}

func GetSquadStats(c *gin.Context) {
	squad_id_for_lookup, err := u.UuidFromPath(c, "squad_id")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})

		s.L.Println("ERROR PARSING SQUAD ID")
		return
	}

	dbSquad, err := s.Db.Squad.Query().
		Where(squad.ID(squad_id_for_lookup)).
		WithOwner().
		WithMembers(
			func(query *ent.ProfileQuery) {
				query.Select(profile.FieldID, profile.FieldFirstName, profile.FieldLastName, profile.FieldPicture)
			}).
		Only(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		s.L.Println("ERROR CHECKING IF SQUAD EXISTS")
		return
	}

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
		return
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

	c.JSON(http.StatusOK, squad_stats_response)

}

func GetSquad(c *gin.Context) {

	squad_id_for_lookup, err := u.UuidFromPath(c, "squad_id")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})

		s.L.Println("ERROR PARSING SQUAD ID")
		return
	}

	dbSquad, err := s.Db.Squad.Query().
		Where(squad.ID(squad_id_for_lookup)).
		WithOwner().
		WithMembers(
			func(query *ent.ProfileQuery) {
				query.Select(profile.FieldID, profile.FieldFirstName, profile.FieldLastName, profile.FieldPicture)
			}).
		Only(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		s.L.Println("ERROR CHECKING IF SQUAD EXISTS")
		return
	}

	res := generateSquadResponse(c, dbSquad)

	c.JSON(http.StatusOK, res)

}

func DeleteSquad(c *gin.Context) {
	prof := u.ProfileFromCtx(c)

	squad_id_to_delete, err := u.UuidFromPath(c, "squad_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	squad_exists, err := s.Db.Squad.Query().
		Where(
			squad.And(
				squad.HasOwnerWith(
					profile.ID(prof.ID),
				),
				squad.ID(squad_id_to_delete),
			),
		).Exist(c)

	if err != nil || !squad_exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete"})
		s.L.Println("Squad with prof as owner does not exist")
		return
	}

	members_to_notify, _ := s.Db.Profile.Query().
		// Select(
		// 	profile.FieldID,
		// ).
		Where(
			profile.HasSquadWith(
				squad.ID(squad_id_to_delete),
			),
		).All(c)

	squad_to_be_deleted, _ := s.Db.Squad.Get(c, squad_id_to_delete)

	err = s.Db.Squad.DeleteOneID(squad_id_to_delete).Exec(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete"})
		s.L.Println("db query failed to delete squad err: ", err)
		return
	}

	if err != nil {
		s.L.Println("error getting members to notify: ", err)
	} else {

		s.L.Println("Sending push notifications to members of squad")

		title := fmt.Sprintf("%s %s deleted the %s squad", prof.FirstName, prof.LastName, squad_to_be_deleted.Title)
		for _, member := range members_to_notify {
			s.L.Println("Sending push notification to member: ", member.ID)
			if member.ID == prof.ID {
				continue
			}
			u.SendPushNotification(c, nil, member, title, "Squad deleted", shared.NotificationData{
				Sender_id:    &prof.ID,
				Sender_image: prof.Picture,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": "Squad deleted"})

}

type PatchSquadStruct struct {
	Title  string `json:"title"`
	Public bool   `json:"public"`
}

func PatchSquad(c *gin.Context) {

	prof := u.ProfileFromCtx(c)

	squad_id_to_patch, err := u.UuidFromPath(c, "squad_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("ERROR PARSING SQUAD ID")
		return
	}

	squad_exists, err := s.Db.Squad.Query().
		Where(
			squad.HasOwnerWith(
				profile.ID(prof.ID),
			),
		).Exist(c)

	if err != nil || !squad_exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to patch"})
		s.L.Println("Squad with prof as owner does not exist")
		return
	}

	var patchSquad PatchSquadStruct

	if bindErrs := u.BindAndValidate(c, &patchSquad); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return
	}

	updated_squad, err := s.Db.Squad.UpdateOneID(squad_id_to_patch).SetTitle(patchSquad.Title).SetPublic(patchSquad.Public).Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		s.L.Println("ERROR UPDATING SQUAD")
		return
	}

	responseSquad := generateSquadResponse(c, updated_squad)

	c.JSON(http.StatusOK, responseSquad)

}
