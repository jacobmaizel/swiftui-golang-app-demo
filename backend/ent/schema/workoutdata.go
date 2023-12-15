package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// WorkoutData holds the schema definition for the WorkoutData entity.
type WorkoutData struct {
	ent.Schema
}

// Mixin of the WorkoutData.
func (WorkoutData) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the WorkoutData.
func (WorkoutData) Fields() []ent.Field {
	return []ent.Field{

		// * healthkit info
		field.String("healthkit_workout_id").Optional(),

		// * start end dates
		field.Time("healthkit_workout_start_date").Default(time.Now).Optional(),
		field.Time("healthkit_workout_end_date").Default(time.Now).Optional(),

		// * stats
		field.Float("distance").Optional(),            // in feet
		field.String("duration").Optional(),           // in seconds
		field.String("energy_burned").Optional(),      // in calories
		field.String("average_heart_rate").Optional(), // in bpm

		// * used to determine if a workout came from swiftui-golang-app-demo or import/healthkit
		//? swiftui-golang-app-demo OR import
		field.String("source").NotEmpty().Default("swiftui-golang-app-demo"),

		// * location -> indoor or outdoor
		field.String("location_type").NotEmpty().Default("indoor"),

		// * weather
		field.JSON("weather", &shared.WorkoutDataWeather{}).Optional(),
	}
}

// Edges of the WorkoutData.
func (WorkoutData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("workout", Workout.Type).Unique(),
		edge.To("profile", Profile.Type).Unique().Required(),
		edge.To("workout_route_data", WorkoutRouteData.Type).Unique(),
		edge.To("competition", Competition.Type).Unique(),
	}
}
