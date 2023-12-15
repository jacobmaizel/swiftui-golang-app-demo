package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

func GeneratePartialProfileResponse(c *gin.Context, prof *ent.Profile) shared.PartialProfile {
	// log.Println("generating partial profile response for first name: ", prof.FirstName, " last name: ", prof.LastName, " id: ", prof.ID, " picture: ", prof.Picture)
	return shared.PartialProfile{
		ID:        prof.ID,
		FirstName: prof.FirstName,
		LastName:  prof.LastName,
		Picture:   prof.Picture,
		Public:    &prof.Public,
	}
}
