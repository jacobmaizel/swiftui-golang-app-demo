package utils

import (
	"context"
	"log"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

func SendPushNotification(c context.Context, sender, receiver *ent.Profile, title, body string, data shared.NotificationData) error {
	// creates a DB notification, pulls the profiles fcm token and

	noti, err := s.Db.Notification.Create().
		SetProfileID(receiver.ID).
		SetTitle(title).
		SetBody(body).
		SetSent(time.Now()).
		Save(c)

	if err != nil {
		log.Println("\n\nSendPushNotification:::FAILED TO CREATE NOTI" + err.Error())
		return err
	}

	msgClient, err := s.Firebase.Messaging(c)
	if err != nil {
		return err
	}

	var msgData map[string]string = map[string]string{
		"notification_id": noti.ID.String(),
		"sender_id":       sender.ID.String(),
		"sender_image":    sender.Picture,
		"receiver_id":     receiver.ID.String(),
	}

	data.Notification_id = noti.ID
	data.Receiver_id = &receiver.ID
	data.Sender_id = &sender.ID
	data.Sender_image = sender.Picture

	if data.Workout_id != nil {
		msgData["workout_id"] = data.Workout_id.String()
	}

	if data.Squad_id != nil {
		msgData["squad_id"] = data.Squad_id.String()
	}

	if data.Competition_id != nil {
		msgData["competition_id"] = data.Competition_id.String()
	}

	if data.Invite_id != nil {
		msgData["invite_id"] = data.Invite_id.String()
	}

	_, err = noti.Update().SetData(&data).Save(c)
	if err != nil {
		log.Println("\n\nSendPushNotification:::FAILED TO UPDATE NOTI with data" + err.Error())
		return err
	}
	/*
			type NotificationData struct {
			Notification_id uuid.UUID `json:"notification_id"`
			Sender_id       uuid.UUID `json:"sender_id"`
			Sender_image    string    `json:"sender_image"`
			Receiver_id     uuid.UUID `json:"receiver_id"`

			Workout_id     *uuid.UUID `json:"workout_id"`
			Squad_id       *uuid.UUID `json:"squad_id"`
			Competition_id *uuid.UUID `json:"competition_id"`
			Invite_id      *uuid.UUID `json:"invite_id"`
		}
	*/

	// log.Println("\n\nSendPushNotification:::CREATED MSG DATA")

	var msgNoti = messaging.Notification{
		Title:    title,
		Body:     body,
		ImageURL: data.Sender_image,
	}

	dbTok, err := receiver.QueryFcmTokens().First(c)

	// log.Println("\n\nSendPushNotification:::CREATED MSG NOTI")

	if err != nil {
		log.Println("\n\nSendPushNotification::: receiver did not have an fcm token, failed" + receiver.ID.String())
		return err
	}

	var msgToSend = messaging.Message{
		Data:         msgData,
		Notification: &msgNoti,
		Token:        dbTok.Token,
	}

	_, err = msgClient.Send(c, &msgToSend)

	if err != nil {
		log.Println("\n\nSendPushNotification::: failed to send message" + err.Error())
		return err
	}

	log.Println("\n\nSUCCESFULLY SENT NOTIFICATION:: data::" + msgToSend.Data["notification_id"] + " to receiver::" + receiver.ID.String() + ":::name::: " + receiver.FirstName + " " + receiver.LastName)

	return nil

}
