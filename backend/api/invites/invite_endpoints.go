package invites

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/invite"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	u "github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func AddInviteRoutes(rg *gin.RouterGroup) {
	InviteRoutes := rg.Group("/invites")

	InviteRoutes.GET("/", ListMyInvites)
	InviteRoutes.GET("/:invite_id", GetInvite)

}

type ListInviteStruct struct {
	Id             uuid.UUID  `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	Status         string     `json:"status"`
	Workout_id     *uuid.UUID `json:"workout_id"`
	Squad_id       *uuid.UUID `json:"squad_id"`
	Competition_id *uuid.UUID `json:"competition_id"`

	Sender   shared.PartialProfile `json:"sender"`
	Receiver shared.PartialProfile `json:"receiver"`
}

func GetInvite(c *gin.Context) {

	invite_id, err := u.UuidFromPath(c, "invite_id")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		s.L.Println("err", err)
		return
	}

	invite, err := s.Db.Invite.Query().Where(
		invite.ID(invite_id),
	).
		WithSender().
		WithReceiver().
		WithWorkout().
		WithSquad().
		WithCompetition().
		First(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		s.L.Println("err", err)
		return
	}

	var workout_id *uuid.UUID
	var squad_id *uuid.UUID
	var competition_id *uuid.UUID

	if invite.Edges.Workout != nil {
		workout_id = &invite.Edges.Workout.ID
	}

	if invite.Edges.Squad != nil {
		squad_id = &invite.Edges.Squad.ID
	}

	if invite.Edges.Competition != nil {
		competition_id = &invite.Edges.Competition.ID
	}

	res := ListInviteStruct{
		Id:        invite.ID,
		CreatedAt: invite.CreatedAt,
		UpdatedAt: invite.UpdatedAt,
		Status:    invite.Status,
		Sender:    u.GeneratePartialProfileResponse(c, invite.Edges.Sender),
		Receiver:  u.GeneratePartialProfileResponse(c, invite.Edges.Receiver),

		Workout_id:     workout_id,
		Squad_id:       squad_id,
		Competition_id: competition_id,
	}

	c.JSON(http.StatusOK, res)

}

func ListMyInvites(c *gin.Context) {
	s.L.Println("ListMyInvites")
	prof := u.ProfileFromCtx(c)

	base := s.Db.Invite.Query().Where(invite.HasReceiverWith(profile.ID(prof.ID)))

	pending_filter_bool := u.BoolFromQuery(c, "pending")

	s.L.Println("pending_filter_bool", pending_filter_bool)

	invite_type := u.Sfq(c, "type")

	s.L.Println("invite_type", invite_type)

	if invite_type == "squad" {

		s.L.Println("invite_type == squad")

		base.Where(invite.HasSquad())

	} else if invite_type == "competition" {

		s.L.Println("invite_type == competition")

		base.Where(invite.HasCompetition())

	} else if invite_type == "workout" {

		s.L.Println("invite_type == workout")

		base.Where(invite.HasWorkout())

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invite type"})

		s.L.Println("Invalid invite type")
		return
	}

	if pending_filter_bool {

		s.L.Println("pending_filter_bool == true")
		base.Where(invite.Status("pending"))
	}

	invites, err := base.WithSender().
		WithCompetition().
		WithSquad().
		WithWorkout().
		WithReceiver().
		All(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})

		s.L.Println("err", err)
		return
	}

	res_invites := []ListInviteStruct{}

	// if the workout, squad or comp edge is not there, set to null

	for _, invite := range invites {
		var workout_id *uuid.UUID
		var squad_id *uuid.UUID
		var competition_id *uuid.UUID

		if invite.Edges.Workout != nil {
			workout_id = &invite.Edges.Workout.ID
		}

		if invite.Edges.Squad != nil {
			squad_id = &invite.Edges.Squad.ID
		}

		if invite.Edges.Competition != nil {
			competition_id = &invite.Edges.Competition.ID
		}

		res_invites = append(res_invites, ListInviteStruct{
			Id:        invite.ID,
			CreatedAt: invite.CreatedAt,
			UpdatedAt: invite.UpdatedAt,
			Status:    invite.Status,
			Sender:    u.GeneratePartialProfileResponse(c, invite.Edges.Sender),
			Receiver:  u.GeneratePartialProfileResponse(c, invite.Edges.Receiver),

			Workout_id:     workout_id,
			Squad_id:       squad_id,
			Competition_id: competition_id,
		},
		)
	}

	c.JSON(http.StatusOK, res_invites)

}
