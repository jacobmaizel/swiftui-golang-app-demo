package notifications

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notification"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func AddNotificationRoutes(rg *gin.RouterGroup) {
	notificationRoutes := rg.Group("/notifications")
	notificationRoutes.GET("/", ListMyNotifications)
	notificationRoutes.PATCH("/", OpenNotification)

	if os.Getenv("ENVIRONMENT") == "DEV" {
		notificationRoutes.POST("/load_test", TestNotification)
		notificationRoutes.POST("/send_test", SendTestNotiToDevice)
	}

}

func OpenNotification(c *gin.Context) {

	var update_body RequestUpdateNotification

	if err := utils.BindAndValidate(c, &update_body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	updated_noti, err := s.Db.Notification.UpdateOneID(update_body.Id).
		SetOpened(update_body.Opened).Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GenerateNotificationResponse(c, updated_noti))

}

func ListMyNotifications(c *gin.Context) {

	prof := utils.ProfileFromCtx(c)

	notifications, err := s.Db.Notification.Query().
		Where(
			notification.HasProfileWith(profile.ID(prof.ID)),
		).
		Order(ent.Desc(notification.FieldSent)).
		All(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var noti_res []ResponseNotification = []ResponseNotification{}

	for _, noti := range notifications {
		noti_res = append(noti_res, GenerateNotificationResponse(c, noti))
	}

	c.JSON(http.StatusOK, noti_res)

}

func TestNotification(c *gin.Context) {

	for _, prof := range s.Db.Profile.Query().AllX(c) {

		// create 10 test notifications for each profile
		for i := 0; i < 10; i++ {

			noti, err := s.Db.Notification.Create().
				SetProfileID(prof.ID).
				SetTitle("Test Notification" + fmt.Sprint(i)).
				SetBody("This is a test notification").
				SetSent(time.Now()).
				Save(c)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			noti_data := shared.NotificationData{
				Notification_id: noti.ID,
			}

			noti.Update().SetData(&noti_data).Save(c)

		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "created test notifications"})
}

func SendTestNotiToDevice(c *gin.Context) {

	prof := utils.ProfileFromCtx(c)

	// delete all notis
	// s.Db.Notification.Delete().ExecX(c)

	err := utils.SendPushNotification(c, nil, prof, utils.Sfq(c, "title"), utils.Sfq(c, "body"), shared.NotificationData{
		Sender_image: utils.Sfq(c, "image"),
		Receiver_id:  &prof.ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		s.L.Println("error sending test notification:::", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sent test notification"})

}
