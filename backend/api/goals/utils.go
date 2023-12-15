package goals

// import (
// 	"log"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
// 	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
// 	s "github.com/jacobmaizel/swiftui-golang-app-demo/server"
// 	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
// )

// /*
// PER WORKOUT
// - when per workout has expired,

// TOTAL
// - when total goal has expired, do a final check for completion.
// there is a chance that it was completed, but the last time we re ran the calculation
// the goal was NOT completed yet.

// so we should not conclude that it is expired unless we confirm.
// */
// const (
// 	FAILED = iota
// 	COMPLETED
// )

// func GoalStatus(g *ent.Goal, c *gin.Context) (string, string, []shared.PerWorkoutGoalDataEntry) {

// 	total := g.ValueAggregationType == "" && g.TimeInterval == ""
// 	per_workout := g.ValueAggregationType == "per_workout" && g.TimeInterval == ""

// 	if g.Status == "active" {

// 		if total {
// 			// is it expired (failed)?
// 			if g.End.Before(time.Now()) {
// 				_, err := g.Update().SetStatus("failed").Save(c)
// 				if err != nil {
// 					log.Printf("ERROR when updating goal from active to expired/failed:: %s", err.Error())
// 					return "active", g.CurrentTotalValue, nil
// 				}
// 				log.Println("Goal was updated from active to failed")
// 				return "failed", g.CurrentTotalValue, nil
// 			}

// 			completed, calculatedNewValue, _ := calcGoalCompletionStatusAndCurrentValue(c, g)

// 			log.Println("Goal was not expired ", completed, " with value ", calculatedNewValue)
// 			// is it completed?
// 			if completed {
// 				_, err := g.Update().SetStatus("completed").SetCurrentTotalValue(calculatedNewValue).Save(c)

// 				if err != nil {
// 					log.Printf("ERROR when updating goal from active to completed:: %s", err.Error())

// 					return "active", g.CurrentTotalValue, nil
// 				}

// 				log.Println("Goal was updated from active to COMPLETED!! ")

// 				return "completed", calculatedNewValue, nil

// 			} else {
// 				// even if its not completed we still need to update the current value
// 				// if calculated value == current value, do nothing

// 				if calculatedNewValue == g.CurrentTotalValue {
// 					log.Println("\n\nGoal was not completed, but current value is the same, returning status ", g.Status, " and current value ", g.CurrentTotalValue)
// 					return g.Status, g.CurrentTotalValue, nil
// 				}

// 				log.Println("Goal was not completed, updating current value to ", calculatedNewValue)

// 				_, err := g.Update().SetCurrentTotalValue(calculatedNewValue).Save(c)

// 				if err != nil {
// 					log.Printf("ERROR when updating goal from active to completed:: %s", err.Error())
// 					return "active", g.CurrentTotalValue, nil
// 				}

// 				return g.Status, calculatedNewValue, nil

// 			}
// 		} else if per_workout {
// 			// HANDLE PER WORKOUT status calcs
// 			// log.Println("\n\nHANDLING PER WORKOUT STATUS CALCULATION")

// 			_, _, newlyCalculatedPerWorkoutData := calcGoalCompletionStatusAndCurrentValue(c, g)
// 			// we are only checking completion after it is expired.
// 			if g.End.Before(time.Now()) {
// 				// go through all of the returned data. if every entry is COMPLETED (0), set as complete.

// 				if len(newlyCalculatedPerWorkoutData) <= 0 {
// 					// an expired entry with no data is considered failed.

// 					// _, err := g.Update().SetStatus("failed").SetPerWorkoutData(newlyCalculatedPerWorkoutData).Save(c)

// 					// if err != nil {
// 					// 	log.Printf("ERROR when updating goal from active to failed, len of data was 0:: %s", err.Error())
// 					// 	return "active", "", g.PerWorkoutData
// 					// }

// 					log.Println("\n\nPer workout Goal was updated from active to FAILED!!!! because an entry was not completed")
// 					return "failed", "", newlyCalculatedPerWorkoutData
// 				}

// 				for _, item := range newlyCalculatedPerWorkoutData {
// 					if item.Status == FAILED {
// 						// we found an entry that did not satisfy the target value.
// 						// update to failed and return
// 						// _, err := g.Update().SetStatus("failed").SetPerWorkoutData(newlyCalculatedPerWorkoutData).Save(c)

// 						// if err != nil {
// 						// 	log.Printf("ERROR when updating goal from active to failed:: %s", err.Error())
// 						// 	return "active", "", g.PerWorkoutData
// 						// }

// 						log.Println("\n\nPer workout Goal was updated from active to FAILED!!!! because an entry was not completed")
// 						return "failed", "", newlyCalculatedPerWorkoutData
// 					}
// 				}

// 				// FOR LOOP IS ENDED, GOAL IS COMPLETE -> end has passed, and all workouts satisfy the target value.

// 				_, err := g.Update().SetStatus("completed").SetPerWorkoutData(newlyCalculatedPerWorkoutData).Save(c)

// 				if err != nil {
// 					log.Printf("ERROR when updating goal from active to completed:: %s", err.Error())
// 					return "active", "", g.PerWorkoutData
// 				}

// 				log.Println("\n\nPer workout Goal was updated from active to COMPLETED!!!! because all entries were completed", newlyCalculatedPerWorkoutData)
// 				return "completed", "", newlyCalculatedPerWorkoutData

// 			} else {
// 				// it has not expired, but we still want to update the per workout data.

// 				_, err := g.Update().SetPerWorkoutData(newlyCalculatedPerWorkoutData).Save(c)

// 				if err != nil {
// 					log.Printf("ERROR when updating goal from active to completed:: %s", err.Error())
// 					return "active", "", g.PerWorkoutData
// 				}

// 				// log.Println("\n\nper workout goal was not expired, but we still updated the per workout data", newlyCalculatedPerWorkoutData)
// 				return g.Status, "", newlyCalculatedPerWorkoutData
// 			}
// 		}
// 	}
// 	// if we got here, it is either still active, OR completed/failed already.
// 	log.Println("SHORT CIRCUITING GOAL CALS::WAS ALREADY FAILED OR COMPLETED", g.Status, " and current value ", g.CurrentTotalValue, " and per workout data ", g.PerWorkoutData)

// 	return g.Status, g.CurrentTotalValue, g.PerWorkoutData

// }

// func calcGoalCompletionStatusAndCurrentValue(c *gin.Context, g *ent.Goal) (bool, string, []shared.PerWorkoutGoalDataEntry) {

// 	total := g.ValueAggregationType == "" && g.TimeInterval == ""
// 	per_workout := g.ValueAggregationType == "per_workout" && g.TimeInterval == ""

// 	// prof := u.ProfileFromCtx(c).ID

// 	base_query := s.Db.Workout.Query().
// 		Where(
// 			workout.HealthkitWorkoutActivityType(g.HealthkitWorkoutActivityType),
// 		)

// 	// average := g.ValueAggregationType == "average"
// 	// per_workout := g.ValueAggregationType == "per_workout" && g.TimeInterval == ""

// 	// check and  update goal states if needed

// 	if total {

// 		// 		log.Println("Calculating total goal")

// 		// 		// Total
// 		// 		// ie. run 10 miles -> between start and end date, sum distance (filtering workouts by activity type of goal) and compare to target

// 		// 		log.Println("\n\n\nCalculating total for goal", g.ID.String(), g.HealthkitWorkoutActivityType, g.Start, g.End, g.Value, g.Unit, g.ValueAggregationType, g.TimeInterval, g.Status, g.CurrentTotalValue)

// 		// 		newlyAggregatedWorkoutDistanceTotalFeet, err := base_query.Aggregate(ent.As(ent.Sum(workout.FieldDistance), "workout_total")).Float64(c)

// 		// 		if err != nil {

// 		// 			log.Println("\n\nfailed to calculate total for goal, returning false", err.Error(), " goal info", g.ID.String(), g.HealthkitWorkoutActivityType, g.Start, g.End, g.Value, g.Unit, g.ValueAggregationType, g.TimeInterval, g.Status, g.CurrentTotalValue)

// 		// 			return false, "", nil
// 		// 		}

// 		// 		targetValueOnGoal, _ := strconv.ParseFloat(g.Value, 64)

// 		// 		adjustedNewTotal := adjustForUnit(newlyAggregatedWorkoutDistanceTotalFeet, g.Unit)

// 		// 		log.Println("\n\ncomparison after adjusting for unit", adjustedNewTotal, ">=", targetValueOnGoal, " unit ", g.Unit)

// 		// 		return adjustedNewTotal >= targetValueOnGoal, strconv.FormatFloat(newlyAggregatedWorkoutDistanceTotalFeet, 'f', 6, 64), nil

// 	} else if per_workout {

// 		// HANDLE PER WORKOUT CALCS

// 		// log.Println("\n\nCalculating per workout goal")

// 		goalDataRes := calculatePerWorkoutGoalData(c, g, base_query)

// 		if len(goalDataRes) == 0 {
// 			// log.Println("There were no goal data results for per workout calculation for goal", g.ID)
// 			return false, "", nil
// 		} else {
// 			// we have data

// 			return false, "", goalDataRes

// 		}

// 	}

// 	return false, "", nil

// }

// func calculatePerWorkoutGoalData(c *gin.Context, g *ent.Goal, wq *ent.WorkoutQuery) []shared.PerWorkoutGoalDataEntry {

// 	// get all workouts with base query. my workouts, done with the right activity type, between start -> end dates.
// 	// for each of those workouts
// 	// 		is the adjusted DISTANCE >= the TARGET "value" of the goal?

// 	// for each of those workouts, push a piece of data into a list.
// 	var ans = make([]shared.PerWorkoutGoalDataEntry, 0)

// 	// dbValues, err := wq.All(c)

// 	// if err != nil {
// 	// 	log.Println("ERROR in getting DB values while calculating per workout goal data::: ", err.Error())
// 	// 	return nil
// 	// }

// 	// for _, dbWorkout := range dbValues {
// 	// goalVal, err := strconv.ParseFloat(g.Value, 64)

// 	// if err != nil {
// 	// 	log.Println("ERROR in calculating per workout goal data::: ", err.Error())
// 	// 	return nil
// 	// }

// 	// if adjustForUnit(dbWorkout.Distance, g.Unit) >= goalVal {
// 	// 	// this workout has higher distance than target
// 	// 	ans = append(ans, shared.PerWorkoutGoalDataEntry{Workout_id: dbWorkout.ID, Status: COMPLETED, Healthkit_workout_end_date: dbWorkout.HealthkitWorkoutEndDate})
// 	// } else {
// 	// 	ans = append(ans, shared.PerWorkoutGoalDataEntry{Workout_id: dbWorkout.ID, Status: FAILED, Healthkit_workout_end_date: dbWorkout.HealthkitWorkoutEndDate})
// 	// }

// 	// }

// 	// log.Println("\n\nCALCULATED INFO FROM PER WORKOUT GOAL DATA\n\n", ans)

// 	return ans
// }

// // func adjustForUnit(value float64, unit string) float64 {
// // 	// // base values are in feet, should we convert to miles?

// // 	// if unit == "miles" {
// // 	// 	return value / 5280
// // 	// }

// // 	// return value
// // }
