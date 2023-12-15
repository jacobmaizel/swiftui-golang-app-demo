package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	u "github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func AddActionRoutes(rg *gin.RouterGroup) {

	actionRoutes := rg.Group("/actions")

	actionRoutes.POST("/", TrackAction)

}

type Action struct {
	Title string `json:"title" binding:"required"`
}

func TrackAction(c *gin.Context) {

	prof := u.ProfileFromCtx(c)

	var action Action

	if bindErrs := u.BindAndValidate(c, &action); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return
	}

	_, err := s.Db.Action.Create().SetSender(prof).SetTitle(action.Title).Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Action tracked!"})

}
