package schema

import (
	"context"
	"log"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
	"github.com/jacobmaizel/swiftui-golang-app-demo/redis"
)

// Competition holds the schema definition for the Competition entity.
type Competition struct {
	ent.Schema
}

func (Competition) Mixin() []ent.Mixin {
	return append(mixins.BaseMixin(), mixins.PrivacyMixin{})
}

func (Competition) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {

			// log.Println("hooking competition mutation")

			return ent.MutateFunc(

				func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if m.Op() == ent.OpDeleteOne || m.Op() == ent.OpDelete {
						return next.Mutate(ctx, m)
					}
					log.Println("Comp Hook::: ", m.Op(), m.Type())

					// log.Println("hooking competition mutation 2")
					// Pre mutation action.
					// competition_mut := m.(*gen.CompetitionMutation)

					// log.Println("comp mutation:", competition_mut)

					v, err := next.Mutate(ctx, m)

					updated_competition := v.(*gen.Competition)

					// log.Println("updated competition:", updated_competition)

					// toIndex :=

					// Post mutation action.
					// fmt.Println("new value:", v)

					redis.IndexCompetitionDoc(ctx, updated_competition)
					return v, err

				},
			)
		},
	}
}

// Fields of the Competition.
func (Competition) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description"),

		field.Time("start"),
		field.Time("end"),
		field.Bool("scheduled").Default(false),

		// pending, active, ended
		field.String("status").Default("pending"),

		// characteristics
		// solo, squad, all
		field.Strings("participant_types").Optional(),
		// list of HKWorkoutActivity Types (lowercase!! )
		field.Strings("workout_types").Optional(),
		// official / community created
		field.String("type").Default("official").NotEmpty(),
	}
}

// Edges of the Competition.
func (Competition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", Profile.Type).Unique().Required(),
		edge.To("participants", Profile.Type),
		edge.From("workouts", Workout.Type).Ref("competition"),
		edge.From("workout_data", WorkoutData.Type).Ref("competition"),
		edge.From("results", CompetitionResult.Type).Ref("competition"),
	}
}
