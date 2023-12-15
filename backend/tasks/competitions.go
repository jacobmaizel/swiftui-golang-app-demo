package tasks

import (
	"context"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notification"
	"github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func CompAboutToStartNotiTask(c context.Context, done chan bool) {
	log.Print("STARTUP:::Starting Competition Start Time Checker\n\n")
	for {
		select {
		case <-done:
			log.Println("Got done signal, shutting down gracefully")
			return
		default:
			time.Sleep(30 * time.Minute)
			log.Println("CompAboutToStartTask:::Starting Check")
			NotificationsForStartingCompetitions(c)
			log.Println("CompAboutToStartTask:::Ending Check")
		}
	}
}

func CompAboutToEndNotiTask(c context.Context, done chan bool) {
	log.Print("STARTUP:::Starting Competition End Time Checker\n\n")

	for {
		select {
		case <-done:
			log.Println("Got done signal, shutting down gracefully")
			return
		default:
			time.Sleep(30 * time.Minute)
			log.Println("CompAboutToENDTask:::Starting Check")
			NotificationsForEndingCompetitions(c)
			log.Println("CompAboutToENDTask:::Ending Check")
		}
	}
}

func NotificationsForEndingCompetitions(c context.Context) {
	// Looks at all competitions that are ending in the next hour
	// notifies all users in the competition that it is ending
	ending_comps, err := server.Db.Competition.Query().
		Where(
			competition.EndLTE(time.Now().Add(30*time.Minute)),
			competition.EndGTE(time.Now()),
		).
		All(c)

	if err != nil {
		log.Println("error getting ending competitions:::", err)
		return
	}

	for _, comp := range ending_comps {
		log.Println("ending comp:::", comp.ID)
		// get all users in the competition
		users, err := comp.QueryParticipants().All(c)
		if err != nil {
			log.Println("error getting users for ending comp:::", err)
			continue
		}

		for _, prof := range users {

			data := shared.NotificationData{
				Competition_id: &comp.ID,
			}

			title := fmt.Sprintf("%s is ending soon", comp.Title)
			timeToEnd := time.Until(comp.End).Minutes()
			body := fmt.Sprintf("You have %.0f minutes until the end!", timeToEnd)

			previouslySent, err := server.Db.Notification.Query().
				Where(func(s *sql.Selector) {
					s.Where(
						sql.And(
							sqljson.ValueEQ(notification.FieldData, comp.ID.String(), sqljson.Path("competition_id")),

							sql.EQ(notification.ProfileColumn, prof.ID),
						),
					)

				},
				).
				Where(
					notification.SentGTE(time.Now().Add(-30 * time.Minute)),
				).
				Exist(c)

			if err == nil && !previouslySent {
				utils.SendPushNotification(c, nil, prof, title, body, data)
			} else {
				log.Println("NotificationsForEndingCompetitions:::previously sent:::", previouslySent, err)
			}

		}
	}

	if len(ending_comps) > 0 {
		log.Println("CheckForEndingCompetitions:::", len(ending_comps), "competitions ending soon")
	}

}

func NotificationsForStartingCompetitions(c context.Context) {
	// Looks at all competitions that are ending in the next hour
	// notifies all users in the competition that it is ending
	starting_comps, err := server.Db.Competition.Query().
		Where(
			competition.StartGTE(time.Now()),
			competition.StartLTE(time.Now().Add(30*time.Minute)),
		).
		All(c)

	if err != nil {
		log.Println("error getting starting competitions:::", err)
		return
	}

	for _, comp := range starting_comps {
		log.Println("starting comp:::", comp.ID)
		// get all users in the competition
		users, err := comp.QueryParticipants().All(c)
		if err != nil {
			log.Println("error getting users for starting comp:::", err)
			continue
		}

		for _, prof := range users {

			data := shared.NotificationData{
				Competition_id: &comp.ID,
			}

			title := fmt.Sprintf("%s is starting soon", comp.Title)
			timeToStart := time.Until(comp.Start).Minutes()
			body := fmt.Sprintf("You have %.0f minutes until the start!", timeToStart)

			previouslySent, err := server.Db.Notification.Query().
				Where(func(s *sql.Selector) {
					s.Where(
						sql.And(
							sqljson.ValueEQ(notification.FieldData, comp.ID.String(), sqljson.Path("competition_id")),

							sql.EQ(notification.ProfileColumn, prof.ID),
						),
					)

				},
				).
				Where(
					notification.SentGTE(time.Now().Add(-30 * time.Minute)),
				).
				Exist(c)

			if err == nil && !previouslySent {
				utils.SendPushNotification(c, nil, prof, title, body, data)
			} else {
				log.Println("NotificationsForEndingCompetitions:::previously sent:::", previouslySent, err)
			}

		}
	}

	if len(starting_comps) > 0 {
		log.Println("CheckForStartingCompetitions:::", len(starting_comps), "competitions starting soon")
	}
}
