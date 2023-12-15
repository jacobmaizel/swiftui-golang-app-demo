package shared

import (
	"time"

	"github.com/google/uuid"
)

type PerWorkoutGoalDataEntry struct {
	Workout_id                 uuid.UUID `json:"workout_id"`
	Status                     int       `json:"status"`
	Healthkit_workout_end_date time.Time `json:"healthkit_workout_end_date"`
}

type NotificationPreferenceSettings struct {
	Squad_invites           bool `json:"squad_invites"`
	Squad_competition_state bool `json:"squad_competition_state"`
	Workout_invites         bool `json:"workout_invites"`
	Post_workout_summary    bool `json:"post_workout_summary"`
	Competition_invites     bool `json:"competition_invites"`
	Competition_state       bool `json:"competition_state"`
}

type NotificationData struct {
	Notification_id uuid.UUID  `json:"notification_id"`
	Sender_id       *uuid.UUID `json:"sender_id"`
	Sender_image    string     `json:"sender_image"`
	Receiver_id     *uuid.UUID `json:"receiver_id"`

	Workout_id     *uuid.UUID `json:"workout_id"`
	Squad_id       *uuid.UUID `json:"squad_id"`
	Competition_id *uuid.UUID `json:"competition_id"`
	Invite_id      *uuid.UUID `json:"invite_id"`
}

type PartialProfile struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Picture   string    `json:"picture"`
	Public    *bool     `json:"public"`
}

type PartialProfileWithInviteStatus struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Picture      string    `json:"picture"`
	InviteStatus string    `json:"invite_status"`
	Public       bool      `json:"public"`
}

type AppConfig struct {
	Auto_import_workouts *bool      `json:"auto_import_workouts"`
	Denied_location_date *time.Time `json:"denied_location_date"`
}

type WorkoutDataWeather struct {
	Temp_with_unit string `json:"temp_with_unit"`
	Symbol         string `json:"symbol"`
	Location_city  string `json:"location_city"`
}
