package workouts

import "github.com/gin-gonic/gin"

func AddWorkoutRoutes(rg *gin.RouterGroup) {
	workoutRoutes := rg.Group("/workouts")

	workoutRoutes.POST("/:workout_id/accept_invite/:invite_id", AcceptWorkoutInvite)
	workoutRoutes.POST("/:workout_id/decline_invite/:invite_id", DeclineWorkoutInvite)
	workoutRoutes.POST("/:workout_id/remove_invite", RemoveAcceptedInvite)

	workoutRoutes.DELETE("/:workout_id", DeleteWorkout)

	workoutRoutes.POST("/:workout_id/invite", InviteToWorkout)
	workoutRoutes.POST("/:workout_id/submit_data", SubmitWorkoutData)
	workoutRoutes.POST("/import", ImportWorkoutData)
	workoutRoutes.GET("/:workout_id/invite_statuses", GetMemberInviteStatusesByWorkout)
	workoutRoutes.GET("/:workout_id", GetWorkoutById)
	// workoutRoutes.GET("/applicable", GetApplicableWorkoutsForComp)
	workoutRoutes.GET("/:workout_id/data", GetWorkoutData)

	workoutRoutes.POST("/", CreateWorkout)
	workoutRoutes.GET("/", ListWorkout)

	// workoutRoutes.GET("/:workout_id", GetWorkout)
}
