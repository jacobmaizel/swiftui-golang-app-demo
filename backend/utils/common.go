package utils

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"

	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
)

func PrintContextInternals(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				PrintContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}

func DerefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func SafeDeref[T any](p *T) T {
	if p == nil {
		var v T
		return v
	}
	return *p
}

func UserFromCtx(c *gin.Context) *ent.User {
	user, _ := c.Get("user")
	return user.(*ent.User)
}

func ProfileFromCtx(c *gin.Context) *ent.Profile {
	prof, _ := c.Get("profile")
	return prof.(*ent.Profile)
}

func UuidFromPath(c *gin.Context, key string) (uuid.UUID, error) {
	value, err := uuid.Parse(c.Param(key))
	return value, err
}

func UuidFromQuery(c *gin.Context, key string) (uuid.UUID, error) {
	value, err := uuid.Parse(c.Query(key))
	return value, err
}

func BoolFromQuery(c *gin.Context, key string) bool {
	value := c.Query(key)
	val_bool, err := strconv.ParseBool(value)

	if value == "" || err != nil {
		return false
	}

	return val_bool
}

func UnwrapBool(b *bool) bool {
	if b == nil || !*b {
		return false
	}
	return true
}

func Sfq(c *gin.Context, key string) string {
	// Returns string from query parameter
	value := c.Query(key)
	return value
}

func StrPtr(s string) *string {
	return &s
}

func UuidFromDbToStrPtr(u uuid.UUID, err error) *string {
	if err != nil {
		return nil
	}

	if u != uuid.Nil {
		return StrPtr(u.String())
	} else {
		return nil
	}
}

func StringToUUID(s string) *uuid.UUID {
	if s == "" {
		return nil
	}

	u, err := uuid.Parse(s)

	if err != nil {
		return nil
	}

	return &u
}

// func PaginateResponse[T any](items []T, page, perPage int) []T {
// 	if page < 1 {
// 		page = 1
// 	}

// 	if perPage < 1 {
// 		perPage = 1
// 	}

// 	start := (page - 1) * perPage
// 	end := start + perPage

// 	if start > len(items) {
// 		start = len(items)
// 	}

// 	if end > len(items) {
// 		end = len(items)
// 	}

// 	return items[start:end]
// }

// func InsertPaginationJsonIntoResponse[T any](items []T, page, perPage int, response *map[string]interface{}) {
// 	(*response)["items"] = PaginateResponse(items, page, perPage)
// 	(*response)["page"] = page
// 	(*response)["perPage"] = perPage
// 	(*response)["total"] = len(items)
// }

//lint:ignore U1000 Ignore unused function temporarily for debugging
func fixWorkoutActivityTypeLetterCase(c *gin.Context) {
	// update all workouts healthkit activity types to be the lowercase version of themselves
	all_workouts := s.Db.Workout.Query().AllX(c)

	for _, w := range all_workouts {
		_, err := s.Db.Workout.UpdateOneID(w.ID).SetHealthkitWorkoutActivityType(strings.ToLower(w.HealthkitWorkoutActivityType)).Save(c)

		if err != nil {
			s.L.Println("Error updating workout: ", err)
		}
	}
}

func ClearUsersWorkouts(prof *ent.Profile) {

	c := context.Background()
	wdres := s.Db.WorkoutData.Delete().Where(workoutdata.HasProfileWith(profile.ID(prof.ID))).ExecX(c)

	wres := s.Db.Workout.Delete().Where(workout.HasLeaderWith(profile.ID(prof.ID))).ExecX(c)

	log.Println(wdres, wres)

}
