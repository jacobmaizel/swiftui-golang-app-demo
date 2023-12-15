package notifications

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

type RequestNotification struct {
}

type ResponseNotification struct {
	Id        uuid.UUID                `json:"id"`
	Title     string                   `json:"title"`
	Body      string                   `json:"body"`
	Sent      time.Time                `json:"sent"`
	Opened    *time.Time               `json:"opened"`
	CreatedAt time.Time                `json:"created_at"`
	Data      *shared.NotificationData `json:"data"`
}

type RequestUpdateNotification struct {
	Id     uuid.UUID `json:"id" binding:"required"`
	Opened time.Time `json:"opened" binding:"required"`
}

func GenerateNotificationResponse(c *gin.Context, noti *ent.Notification) ResponseNotification {

	res := ResponseNotification{
		Id:        noti.ID,
		Title:     noti.Title,
		Body:      noti.Body,
		Sent:      noti.Sent,
		Opened:    noti.Opened,
		CreatedAt: noti.CreatedAt,
		Data:      noti.Data,
	}

	return res
}
