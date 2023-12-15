package workouts

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/competitions"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

type WorkoutCreateResponse struct {
	Id uuid.UUID `json:"id"`
	// Location_type                   string    `json:"location_type"`
	Healthkit_workout_activity_type string    `json:"healthkit_workout_activity_type"`
	CreatedAt                       time.Time `json:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at"`
}

type WorkoutCreateRequest struct {
	Healthkit_workout_activity_type string `json:"healthkit_workout_activity_type" binding:"required"`
}

type WorkoutDataSubmission struct {
	// Workout_id                   uuid.UUID `json:"workout_id" binding:"required"`
	Healthkit_workout_id         string                     `json:"healthkit_workout_id" binding:"required"`
	Healthkit_workout_start_date time.Time                  `json:"healthkit_workout_start_date" binding:"required"`
	Healthkit_workout_end_date   time.Time                  `json:"healthkit_workout_end_date" binding:"required"`
	Average_heart_rate           string                     `json:"average_heart_rate"`
	Distance                     float64                    `json:"distance"`
	Duration                     string                     `json:"duration"`
	Energy_burned                string                     `json:"energy_burned"`
	Source                       string                     `json:"source"`
	Location_type                string                     `json:"location_type"`
	Weather                      *shared.WorkoutDataWeather `json:"weather"`
}

type WorkoutDataImportRequest struct {
	// Workout_id                   uuid.UUID `json:"workout_id" binding:"required"`
	Healthkit_workout_activity_type string    `json:"healthkit_workout_activity_type" binding:"required"`
	Healthkit_workout_id            string    `json:"healthkit_workout_id" binding:"required"`
	Healthkit_workout_start_date    time.Time `json:"healthkit_workout_start_date" binding:"required"`
	Healthkit_workout_end_date      time.Time `json:"healthkit_workout_end_date" binding:"required"`
	Average_heart_rate              string    `json:"average_heart_rate"`
	Distance                        float64   `json:"distance"`
	Duration                        string    `json:"duration"`
	Energy_burned                   string    `json:"energy_burned"`
	Source                          string    `json:"source"`
	Location_type                   string    `json:"location_type"`
}

type WorkoutDataResponse struct {
	Id                           uuid.UUID                         `json:"id"`
	CreatedAt                    time.Time                         `json:"created_at"`
	UpdatedAt                    time.Time                         `json:"updated_at"`
	Healthkit_workout_id         string                            `json:"healthkit_workout_id"`
	Healthkit_workout_start_date time.Time                         `json:"healthkit_workout_start_date"`
	Healthkit_workout_end_date   time.Time                         `json:"healthkit_workout_end_date"`
	Distance                     float64                           `json:"distance"`
	Duration                     string                            `json:"duration"`
	Energy_burned                string                            `json:"energy_burned"`
	Average_heart_rate           string                            `json:"average_heart_rate"`
	Source                       string                            `json:"source"`
	Location_type                string                            `json:"location_type"`
	Competition                  *competitions.ResponseCompetition `json:"competition"`
	Profile                      shared.PartialProfile             `json:"profile"`
	Weather                      *shared.WorkoutDataWeather        `json:"weather"`
}

type CompleteWorkoutResponse struct {
	Id                              uuid.UUID `json:"id"`
	CreatedAt                       time.Time `json:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at"`
	Healthkit_workout_activity_type string    `json:"healthkit_workout_activity_type"`
	// Location_type                   string                            `json:"location_type"`

	Leader      shared.PartialProfile `json:"leader"`
	WorkoutData *WorkoutDataResponse  `json:"workout_data"`
}

func GenerateWorkoutCreateResponse(c *gin.Context, dbWorkout *ent.Workout) WorkoutCreateResponse {

	res := WorkoutCreateResponse{
		Id:                              dbWorkout.ID,
		Healthkit_workout_activity_type: dbWorkout.HealthkitWorkoutActivityType,
		CreatedAt:                       dbWorkout.CreatedAt,
		UpdatedAt:                       dbWorkout.UpdatedAt,
	}

	return res

}

func GenerateWorkoutDataResponse(c *gin.Context, dbWorkoutData *ent.WorkoutData) *WorkoutDataResponse {

	// log.Println("\n\nGENERATE WORKOUT DATA RESPONSE\n\n startedd ")

	var competition_response *competitions.ResponseCompetition = nil

	// if dbWorkoutData.Edges.Competition == nil {

	// 	// var err error
	// 	// log.Println("competition was nil WHILE GETTING WORKOUT DATA RESPONSE, querying competition")
	// 	// var err error
	// 	dbWorkoutData.Edges.Competition, _ = dbWorkoutData.QueryCompetition().WithOwner().Only(c)

	// 	// 		if err != nil {
	// 	// 			server.Scream.Println("error querying competition: ", err)
	// 	// 		}
	// }

	// server.Scream.Println("Comp is:::::::::::::::::, ", dbWorkoutData.Edges.Competition)

	if dbWorkoutData.Edges.Competition != nil {
		competition_response = &competitions.ResponseCompetition{
			ID:                dbWorkoutData.Edges.Competition.ID,
			Title:             dbWorkoutData.Edges.Competition.Title,
			Start:             dbWorkoutData.Edges.Competition.Start,
			End:               dbWorkoutData.Edges.Competition.End,
			Scheduled:         &dbWorkoutData.Edges.Competition.Scheduled,
			Participant_types: dbWorkoutData.Edges.Competition.ParticipantTypes,
			Workout_types:     dbWorkoutData.Edges.Competition.WorkoutTypes,
			Type:              dbWorkoutData.Edges.Competition.Type,
			Status:            dbWorkoutData.Edges.Competition.Status,
			Public:            &dbWorkoutData.Edges.Competition.Public,
			CreatedAt:         dbWorkoutData.Edges.Competition.CreatedAt,
			UpdatedAt:         dbWorkoutData.Edges.Competition.UpdatedAt,
			Owner:             utils.GeneratePartialProfileResponse(c, dbWorkoutData.Edges.Competition.Edges.Owner),
		}
	}

	// server.Scream.Println("Comp response is:::::::::::::::::, ", competition_response)

	res := &WorkoutDataResponse{
		Id:                           dbWorkoutData.ID,
		Healthkit_workout_id:         dbWorkoutData.HealthkitWorkoutID,
		Healthkit_workout_start_date: dbWorkoutData.HealthkitWorkoutStartDate,
		Healthkit_workout_end_date:   dbWorkoutData.HealthkitWorkoutEndDate,
		Average_heart_rate:           dbWorkoutData.AverageHeartRate,
		Distance:                     dbWorkoutData.Distance,
		Duration:                     dbWorkoutData.Duration,
		Energy_burned:                dbWorkoutData.EnergyBurned,
		CreatedAt:                    dbWorkoutData.CreatedAt,
		UpdatedAt:                    dbWorkoutData.UpdatedAt,
		Source:                       dbWorkoutData.Source,
		Location_type:                dbWorkoutData.LocationType,
		Competition:                  competition_response,
		Profile:                      utils.GeneratePartialProfileResponse(c, dbWorkoutData.Edges.Profile),
		Weather:                      dbWorkoutData.Weather,
	}

	return res

}

func GenerateWorkoutDataListRes(c *gin.Context, dbWorkoutData []*ent.WorkoutData) []WorkoutDataResponse {

	var res []WorkoutDataResponse = []WorkoutDataResponse{}

	for _, wd := range dbWorkoutData {

		item := GenerateWorkoutDataResponse(c, wd)

		if item != nil {
			res = append(res, *item)
		}

	}

	return res

}

func GenerateFullWorkoutResForProf(c *gin.Context, dbWorkout *ent.Workout) (*CompleteWorkoutResponse, error) {
	// log.Println("GENERATE COMPLETE WORKOUT RESPONSE startedd for workout: ", dbWorkout.ID)

	var res CompleteWorkoutResponse

	// prof := utils.ProfileFromCtx(c)
	var err error

	if dbWorkout.Edges.Leader == nil {

		log.Println("leader was nil, querying leader")
		dbWorkout.Edges.Leader, err = dbWorkout.QueryLeader().
			Select(
				profile.FieldID,
				profile.FieldFirstName,
				profile.FieldLastName,
				profile.FieldPicture,
				profile.FieldPublic,
			).Only(c)

		if err != nil {
			log.Println("GENERATE COMPLETE WORKOUT RESPONSE ERROR: ", err)
		}

	}

	if err != nil {
		server.L.Println("GENERATE COMPLETE WORKOUT RESPONSE ERROR:::: ", err)
		return nil, err
	}

	var workout_data_response *WorkoutDataResponse = nil

	if dbWorkout.Edges.WorkoutData != nil && len(dbWorkout.Edges.WorkoutData) > 0 {
		// server.Scream.Println("workout data was not nil!!!")
		workout_data_response = GenerateWorkoutDataResponse(c, dbWorkout.Edges.WorkoutData[0])
	}

	// server.Scream.Println("after generating workout data response, about to generate competition response")

	// log.Println("\n\n\nabout to generate workout data response BODY", dbWorkout.Edges.WorkoutData)
	res = CompleteWorkoutResponse{
		Id:                              dbWorkout.ID,
		CreatedAt:                       dbWorkout.CreatedAt,
		UpdatedAt:                       dbWorkout.UpdatedAt,
		Healthkit_workout_activity_type: dbWorkout.HealthkitWorkoutActivityType,
		// Location_type:                   dbWorkout.LocationType,
		// Competition: competition_response,
		Leader:      utils.GeneratePartialProfileResponse(c, dbWorkout.Edges.Leader),
		WorkoutData: workout_data_response,
	}

	// log.Println(res)

	return &res, nil

}
