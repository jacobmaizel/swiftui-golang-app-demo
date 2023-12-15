package goals

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
	u "github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

type GoalResponse struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`

	Type                            string `json:"type"`
	Healthkit_workout_activity_type string `json:"healthkit_workout_activity_type"`
	Action                          string `json:"action"`
	Value                           string `json:"value"`
	Unit                            string `json:"unit"`
	Value_aggregation_type          string `json:"value_aggregation_type"`
	Time_interval                   string `json:"time_interval"`

	Status string `json:"status"`

	Current_total_value string `json:"current_total_value"`
	// Current_average_value string `json:"current_average_value"`
	Per_workout_data []shared.PerWorkoutGoalDataEntry `json:"per_workout_data"`

	Profile     uuid.UUID `json:"profile"`
	Competition *string   `json:"competition"`
	Squad       *string   `json:"squad"`
}

type GoalInput struct {
	Start time.Time `json:"start" binding:"required"`
	End   time.Time `json:"end" binding:"required"`

	Type                            string `json:"type"`
	Healthkit_workout_activity_type string `json:"healthkit_workout_activity_type" binding:"required"`
	Action                          string `json:"action" binding:"required"`
	Value                           string `json:"value" binding:"required"`
	Unit                            string `json:"unit" binding:"required"`
	Value_aggregation_type          string `json:"value_aggregation_type"`
	Time_interval                   string `json:"time_interval"`

	Competition *uuid.UUID `json:"competition"`
	Squad       *uuid.UUID `json:"squad"`
}

func GenerateGoalResponse(c *gin.Context, dbGoal *ent.Goal) GoalResponse {

	// span := sentry.StartSpan(c, "Test generate goal response")
	// doWork(span.Context(), item) // doWork may create additional spans

	resprof := dbGoal.QueryProfile().OnlyIDX(c)
	rescomp, err2 := dbGoal.QueryCompetition().OnlyID(c)
	ressquad, err3 := dbGoal.QuerySquad().OnlyID(c)

	rc := u.UuidFromDbToStrPtr(rescomp, err2)
	rs := u.UuidFromDbToStrPtr(ressquad, err3)

	var status, current_total_value string

	status = dbGoal.Status
	current_total_value = dbGoal.CurrentTotalValue
	// perWorkoutGoalData := dbGoal.PerWorkoutData

	// span := sentry.StartSpan(c, "Calculate goal status")

	if status == "active" {
		// if its active we want to calculate the state

		// status, current_total_value, perWorkoutGoalData = GoalStatus(dbGoal, c)

		log.Println("\nSERIALIZER::calculating status for goal: ", dbGoal.ID, " status: ", status, " current_total_value(before adjustment): ", current_total_value, " goal value: ", dbGoal.Value, " unit ", dbGoal.Unit)

	}

	// span.Finish()

	return GoalResponse{
		Id:                              dbGoal.ID,
		CreatedAt:                       dbGoal.CreatedAt,
		UpdatedAt:                       dbGoal.UpdatedAt,
		Start:                           dbGoal.Start,
		End:                             dbGoal.End,
		Type:                            dbGoal.Type,
		Healthkit_workout_activity_type: dbGoal.HealthkitWorkoutActivityType,
		Action:                          dbGoal.Action,
		Value:                           dbGoal.Value,
		Unit:                            dbGoal.Unit,
		Value_aggregation_type:          dbGoal.ValueAggregationType,
		Time_interval:                   dbGoal.TimeInterval,
		Status:                          status,
		Current_total_value:             current_total_value,
		// Per_workout_data:                perWorkoutGoalData,
		Profile:     resprof,
		Competition: rc,
		Squad:       rs,
	}
}
