package goals

import "github.com/gin-gonic/gin"

func AddGoalRoutes(rg *gin.RouterGroup) {
	goalRoutes := rg.Group("/goals")
	goalRoutes.GET("/", ListGoals)
	goalRoutes.POST("/", CreateGoal)
	goalRoutes.POST("/clear", ClearMyGoals)

}
