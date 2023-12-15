package schema

import (
	// "context"
	// "fmt"

	"context"
	"fmt"

	"entgo.io/ent"
	// "entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/edge"
	// _ "github.com/jacobmaizel/swiftui-golang-app-demo/ent/runtime"

	// "github.com/jacobmaizel/swiftui-golang-app-demo/ent/hook"
	// _ "github.com/jacobmaizel/swiftui-golang-app-demo/ent/runtime"

	gen "github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/redis"

	// "entgo.io/ent/hook"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

func (Profile) Mixin() []ent.Mixin {
	return append(mixins.BaseMixin(), mixins.PrivacyMixin{})
}

func (Profile) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {

			return ent.MutateFunc(

				func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					// Pre mutation action.
					// prof_mut := m.(*gen.ProfileMutation)
					if m.Op() == ent.OpDeleteOne || m.Op() == ent.OpDelete {
						return next.Mutate(ctx, m)
					}

					v, err := next.Mutate(ctx, m)

					updated_profile := v.(*gen.Profile)

					// log.Println("updated profile:", updated_profile)

					// Post mutation action.
					fmt.Println("new value:", v)

					redis.IndexProfileDoc(ctx, updated_profile)
					return v, err

				},
			)
		},
	}
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name"),
		field.String("last_name"),
		field.String("picture").Optional(),

		field.Time("birthday").Nillable().Optional(),
		field.Bool("onboarding_completed").Default(false),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("profile").Unique().Required(),

		edge.To("fcm_tokens", FcmToken.Type),
		edge.To("notification_preferences", NotificationPreferences.Type).Unique(),
		edge.From("app_config", AppConfig.Type).Ref("profile").Unique(),

		edge.To("notifications", Notification.Type),

		edge.From("competitions", Competition.Type).Ref("participants"),

		edge.From("actions", Action.Type).Ref("sender"),

		edge.From("squad", Squad.Type).Ref("members"),
		edge.From("squads_owned", Squad.Type).Ref("owner"),

		edge.From("invites", Invite.Type).Ref("receiver"),
		edge.From("invites_sent", Invite.Type).Ref("sender"),

		edge.From("workouts", Workout.Type).Ref("leader"),

		edge.From("goals", Goal.Type).Ref("profile"),

		edge.From("workout_data", WorkoutData.Type).Ref("profile"),

		edge.From("competition_results", CompetitionResult.Type).Ref("profile"),
	}
}
