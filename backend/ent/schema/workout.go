package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// Workout holds the schema definition for the Workout entity.
type Workout struct {
	ent.Schema
}

func (Workout) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the Workout.
func (Workout) Fields() []ent.Field {
	return []ent.Field{

		// * healthkit info
		field.String("healthkit_workout_activity_type").NotEmpty(),
	}
}

// Edges of the Workout.
func (Workout) Edges() []ent.Edge {
	return []ent.Edge{
		//reverse relationship for when an invite has a workout on it
		edge.From("invite", Invite.Type).
			Ref("workout"),

		edge.To("leader", Profile.Type).Unique(),
		edge.To("competition", Competition.Type).Unique(),

		edge.From("workout_data", WorkoutData.Type).
			Ref("workout"),

		edge.From("workout_route_data", WorkoutRouteData.Type).
			Ref("workout"),
	}

}
