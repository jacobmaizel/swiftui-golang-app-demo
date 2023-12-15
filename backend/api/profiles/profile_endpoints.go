package profiles

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competitionresult"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/fcmtoken"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

type UpdateProfileStruct struct {
	First_name          string     `json:"first_name"`
	Last_name           string     `json:"last_name"`
	Birthday            *time.Time `json:"birthday"`
	OnboardingCompleted *bool      `json:"onboarding_completed"`
	Public              *bool      `json:"public"`
}

type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type TokenResponse struct {
	UpdatedAt time.Time `json:"updated_at"`
}
type ResponseProfile struct {
	ID                  uuid.UUID  `json:"id"`
	First_name          string     `json:"first_name"`
	Last_name           string     `json:"last_name"`
	Birthday            *time.Time `json:"birthday"`
	OnboardingCompleted *bool      `json:"onboarding_completed"`
	Picture             string     `json:"picture"`
	Created_at          time.Time  `json:"created_at"`
	Updated_at          time.Time  `json:"updated_at"`
	Public              *bool      `json:"public"`
}

// TODO: Create struct for my personal profile that includes notitication perms, etc. full data
// adjust responseprofile to be for other profiles and not include sensitive data, just enought to display and link

func AddProfileRoutes(rg *gin.RouterGroup) {
	profileRoutes := rg.Group("/profiles")

	profileRoutes.GET("/:profile_id", GetProfileByID)

	profileRoutes.GET("/", ListProfiles)
	profileRoutes.GET("/me", GetMyProfile)
	profileRoutes.PATCH("/", UpdateProfile)
	profileRoutes.POST("/fcm_token", UpdateFcmToken)
	profileRoutes.GET("/:profile_id/stats", GetProfileStats)
}

func GetProfileByID(c *gin.Context) {

	profile_id, err := utils.UuidFromPath(c, "profile_id")

	if err != nil {
		log.Println("ERROR GETTING PROFILE ID FROM PATH", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get Profile"})
		return
	}

	profile, err := s.Db.Profile.Query().
		Where(profile.ID(profile_id)).
		Only(c)

	if err != nil {
		log.Println("ERROR GETTING PROFILE", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get Single profile"})
		return
	}

	responseProf := utils.GeneratePartialProfileResponse(c, profile)

	c.JSON(http.StatusOK, responseProf)
}

func ListProfiles(c *gin.Context) {

	profiles, err := s.Db.Profile.Query().All(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	var responseProfiles []ResponseProfile = make([]ResponseProfile, len(profiles))

	for i, prof := range profiles {
		responseProfiles[i] = ResponseProfile{
			ID:                  prof.ID,
			First_name:          prof.FirstName,
			Last_name:           prof.LastName,
			Birthday:            prof.Birthday,
			OnboardingCompleted: &prof.OnboardingCompleted,
			Picture:             prof.Picture,
			Public:              &prof.Public,
			Created_at:          prof.CreatedAt,
			Updated_at:          prof.UpdatedAt,
		}

	}

	c.JSON(http.StatusOK, responseProfiles)
}

func UpdateProfile(c *gin.Context) {

	var uProf UpdateProfileStruct

	if bindErrs := utils.BindAndValidate(c, &uProf); bindErrs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bindErrs})
		return
	}

	contextProfile := utils.ProfileFromCtx(c)

	base_update_query := s.Db.Profile.
		UpdateOneID(contextProfile.ID)
	// SetNillableBirthday(uProf.Birthday).
	// 	SetFirstName(uProf.First_name).
	// 	SetLastName(uProf.Last_name).
	// 	SetNillableOnboardingCompleted(uProf.OnboardingCompleted).
	//  Save(c)

	if uProf.Birthday != nil {
		base_update_query.SetNillableBirthday(uProf.Birthday)
	}

	if uProf.First_name != "" {
		base_update_query.SetFirstName(uProf.First_name)
	}

	if uProf.Last_name != "" {
		base_update_query.SetLastName(uProf.Last_name)
	}

	if uProf.OnboardingCompleted != nil {
		base_update_query.SetNillableOnboardingCompleted(uProf.OnboardingCompleted)
	}

	if uProf.Public != nil {
		base_update_query.SetNillablePublic(uProf.Public)
	}

	updated_prof, err := base_update_query.Save(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	responseProf := ResponseProfile{
		ID:                  updated_prof.ID,
		First_name:          updated_prof.FirstName,
		Last_name:           updated_prof.LastName,
		Birthday:            updated_prof.Birthday,
		OnboardingCompleted: &updated_prof.OnboardingCompleted,
		Public:              &updated_prof.Public,
		Picture:             updated_prof.Picture,
		Created_at:          updated_prof.CreatedAt,
		Updated_at:          updated_prof.UpdatedAt,
	}

	c.JSON(http.StatusOK, responseProf)
}

func GetMyProfile(c *gin.Context) {

	prof := utils.ProfileFromCtx(c)

	// utils.ClearUsersWorkouts(prof)

	responseProf := ResponseProfile{
		ID:                  prof.ID,
		First_name:          prof.FirstName,
		Last_name:           prof.LastName,
		Birthday:            prof.Birthday,
		OnboardingCompleted: &prof.OnboardingCompleted,
		Public:              &prof.Public,
		Picture:             prof.Picture,
		Created_at:          prof.CreatedAt,
		Updated_at:          prof.UpdatedAt,
	}

	c.JSON(http.StatusOK, responseProf)

}

func UpdateFcmToken(c *gin.Context) {

	prof := utils.ProfileFromCtx(c)

	var token_payload TokenRequest

	if err := utils.BindAndValidate(c, &token_payload); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	existing_token, err := s.Db.FcmToken.Query().
		Where(fcmtoken.HasProfileWith(
			profile.ID(prof.ID),
		)).
		FirstID(c)

	if err != nil {
		// create new
		newToken, err := s.Db.FcmToken.Create().
			SetProfileID(prof.ID).
			SetToken(token_payload.Token).
			Save(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdNewResponse := TokenResponse{
			UpdatedAt: newToken.UpdatedAt,
		}

		c.JSON(http.StatusCreated, createdNewResponse)
		return
	}

	// update existing
	updatedToken, err := s.Db.FcmToken.UpdateOneID(existing_token).SetToken(token_payload.Token).Save(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTokenResponse := TokenResponse{
		UpdatedAt: updatedToken.UpdatedAt,
	}

	c.JSON(http.StatusOK, updatedTokenResponse)

}

type ProfileStatsResponse struct {
	Workout_count     int     `json:"workout_count"`
	Competition_wins  int     `json:"competition_wins"`
	Total_distance_ft float64 `json:"total_distance_ft"`
}

func GetProfileStats(c *gin.Context) {
	// prof := utils.ProfileFromCtx(c).ID
	prof, err := utils.UuidFromPath(c, "profile_id")

	if err != nil {
		log.Println("ERROR GETTING PROFILE ID FROM PATH", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var prof_stat_response ProfileStatsResponse = ProfileStatsResponse{}

	//!!! QUERY 1: Count of workout datas with this profile
	workout_count, err := s.Db.WorkoutData.Query().
		Where(
			workoutdata.HasProfileWith(
				profile.ID(prof),
			),
		).
		Count(c)

	if err != nil {

		log.Println("ERROR GETTING WORKOUT COUNT", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//!!! Count of CompetitionResults with this profile where the place was 1
	comp_win_count, err := s.Db.CompetitionResult.Query().
		Where(competitionresult.HasProfileWith(profile.ID(prof)), competitionresult.Place("1")).Count(c)

	if err != nil {

		log.Println("ERROR GETTING COMP WIN COUNT", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//!!! Sum of distance of workouts with this profile
	total_distance_ft, err := s.Db.WorkoutData.Query().
		Where(
			workoutdata.HasProfileWith(
				profile.ID(prof),
			),
		).Aggregate(
		ent.Sum(workoutdata.FieldDistance),
	).
		Float64(c)

	if err != nil {
		log.Println("ERROR GETTING TOTAL DISTANCE FT", err)
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
		prof_stat_response.Total_distance_ft = 0
	} else {
		prof_stat_response.Total_distance_ft = total_distance_ft
	}

	prof_stat_response.Workout_count = workout_count
	prof_stat_response.Competition_wins = comp_win_count

	c.JSON(http.StatusOK, prof_stat_response)
}
