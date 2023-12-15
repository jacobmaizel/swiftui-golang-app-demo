package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/appconfig"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func AddConfigRoutes(rg *gin.RouterGroup) {
	configRoutes := rg.Group("/config")

	configRoutes.GET("/", GetConfig)
	configRoutes.PATCH("/", UpdateConfig)

}

func GetConfig(c *gin.Context) {

	prof := utils.ProfileFromCtx(c)

	config, err := s.Db.AppConfig.Query().Where(
		appconfig.
			HasProfileWith(
				profile.ID(prof.ID),
			),
	).First(c)

	switch err.(type) {
	case *ent.NotFoundError:
		config, err = s.Db.AppConfig.Create().SetProfile(prof).Save(c)
		if err != nil {
			s.L.Println(err)
		}
	case nil:
		break
	default:
		s.L.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get config"})
		return

	}

	res := GenerateConfigResponse(config)

	c.JSON(http.StatusOK, res)

}

func UpdateConfig(c *gin.Context) {

	var req ConfigUpdate

	if errs := utils.BindAndValidate(c, &req); errs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errs)
		return

	}

	prof := utils.ProfileFromCtx(c)

	config, err := s.Db.AppConfig.Query().Where(
		appconfig.
			HasProfileWith(
				profile.ID(prof.ID),
			),
	).First(c)

	if err != nil {
		s.L.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get config"})
		return
	}

	base := config.Update()

	if req.Auto_sync_workouts != nil {
		base.SetAutoSyncWorkouts(*req.Auto_sync_workouts)
	}

	config, err = base.Save(c)

	if err != nil {
		s.L.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to update config"})
		return
	}

	confRes := GenerateConfigResponse(config)

	c.JSON(http.StatusOK, confRes)
}
