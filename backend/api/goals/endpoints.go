package goals

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/goal"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	u "github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func ListGoals(c *gin.Context) {

	// if hub := sentrygin.GetHubFromContext(c); hub != nil {
	// 	// // hub.WithScope(func(scope *sentry.Scope) {

	// 	// // })
	// 	// log.Println("hub here")

	// 	// // hub.

	// 	// // log.Println(" got da hubbb\n\n", hub.Client().)
	// }

	// sentry

	// sentrygin

	q := s.Db.Goal.Query().Where(goal.HasProfileWith(profile.ID(u.ProfileFromCtx(c).ID)))
	q = handleGoalFilters(c, q)

	goals, err := q.All(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	responseGoals := make([]GoalResponse, 0)

	for _, goal := range goals {

		responseGoals = append(responseGoals, GenerateGoalResponse(c, goal))

		// log.Println("goal completed: ", goalCompleted)

	}

	c.JSON(200, responseGoals)

}

func CreateGoal(c *gin.Context) {

	var goalInput GoalInput

	if bindErrs := u.BindAndValidate(c, &goalInput); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return
	}

	log.Println(goalInput)

	prof := u.ProfileFromCtx(c)

	var goalType string
	if u.SafeDeref(goalInput.Competition) != uuid.Nil {
		goalType = "competition"
	} else if u.SafeDeref(goalInput.Squad) != uuid.Nil {
		goalType = "squad"
	} else {
		goalType = "personal"
	}

	newGoal, errr := s.Db.Goal.Create().
		SetStart(goalInput.Start).
		SetEnd(goalInput.End).
		SetType(goalType).
		SetHealthkitWorkoutActivityType(goalInput.Healthkit_workout_activity_type).
		SetAction(goalInput.Action).
		SetValue(goalInput.Value).
		SetUnit(goalInput.Unit).
		SetValueAggregationType(goalInput.Value_aggregation_type).
		SetTimeInterval(goalInput.Time_interval).
		SetProfile(prof).
		SetNillableCompetitionID(goalInput.Competition).
		SetNillableSquadID(goalInput.Squad).
		Save(c)

	if errr != nil {
		c.JSON(500, gin.H{"error": errr.Error()})
		return
	}

	res := GenerateGoalResponse(c, newGoal)

	c.JSON(200, res)

}

func ClearMyGoals(c *gin.Context) {

	prof := u.ProfileFromCtx(c)
	s.Db.Goal.Delete().Where(goal.HasProfileWith(profile.ID(prof.ID))).ExecX(c)

	c.JSON(200, gin.H{"message": "cleared all goals"})
}
