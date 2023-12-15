package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// WorkoutRouteData holds the schema definition for the WorkoutRouteData entity.
type WorkoutRouteData struct {
	ent.Schema
}

// Mixin of the WorkoutRouteData.
func (WorkoutRouteData) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the WorkoutRouteData.
func (WorkoutRouteData) Fields() []ent.Field {
	return nil
}

// Edges of the WorkoutRouteData.
func (WorkoutRouteData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("workout", Workout.Type).Unique().Required(),
		edge.From("workout_data", WorkoutData.Type).
			Ref("workout_route_data").Unique(),
	}
}
