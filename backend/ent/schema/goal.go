package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// Goal holds the schema definition for the Goal entity.
type Goal struct {
	ent.Schema
}

func (Goal) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the Goal.
func (Goal) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"), // personal, squad

		field.Time("start"), // start date
		field.Time("end"),   // end date

		field.String("healthkit_workout_activity_type").Optional(), // ie. running, walking, cycling
		field.String("action"),                 // ie. run, burn, compete
		field.String("value"),                  // ie. 1
		field.String("unit"),                   // miles, feet
		field.String("value_aggregation_type"), // average, total
		field.String("time_interval"),          // daily, weekly, per workout, monthly

		// states
		field.String("status").Default("active"), // active, completed, failed

		// calculated values
		field.String("current_total_value").Optional(), // current total value
		field.JSON("per_workout_data", []shared.PerWorkoutGoalDataEntry{}).Optional(),
	}
}

// Edges of the Goal.
func (Goal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Unique().Required(),
		edge.To("competition", Competition.Type).Unique(),
		edge.To("squad", Squad.Type).Unique(),
	}
}

func (Goal) Hooks() []ent.Hook {
	return nil
}

func (Goal) Indexes() []ent.Index {
	return nil
}
