package config

import (
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
)

type ConfigResponse struct {
	Id                 uuid.UUID `json:"id"`
	Auto_sync_workouts *bool     `json:"auto_sync_workouts"`
}

type ConfigUpdate struct {
	Auto_sync_workouts *bool `json:"auto_sync_workouts"`
}

func GenerateConfigResponse(config *ent.AppConfig) *ConfigResponse {
	res := &ConfigResponse{}

	res.Auto_sync_workouts = &config.AutoSyncWorkouts
	res.Id = config.ID
	return res
}
