package schema

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
	"github.com/jacobmaizel/swiftui-golang-app-demo/redis"
)

// Squad holds the schema definition for the Squad entity.
type Squad struct {
	ent.Schema
}

func (Squad) Mixin() []ent.Mixin {
	return append(mixins.BaseMixin(), mixins.PrivacyMixin{})
}

func (Squad) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {

			return ent.MutateFunc(

				func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					// Pre mutation action.
					// if it is a delete mutation, skip it and return the value

					if m.Op() == ent.OpDeleteOne || m.Op() == ent.OpDelete {
						return next.Mutate(ctx, m)
					}
					squad_mut := m.(*gen.SquadMutation)

					log.Println("squad mutation:", squad_mut)
					v, err := next.Mutate(ctx, m)

					updated_squad := v.(*gen.Squad)

					// log.Println("updated Squad:", updated_squad)

					// Post mutation action.
					fmt.Println("new value:", v)

					redis.IndexSquadDoc(ctx, updated_squad)
					return v, err

				},
			)
		},
	}
}

// Fields of the Squad.
func (Squad) Fields() []ent.Field {

	return []ent.Field{
		field.String("title").Unique(),
	}

}

// Edges of the Squad.
func (Squad) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", Profile.Type),
		edge.From("invites", Invite.Type).Ref("squad"),
		edge.From("competition_results", CompetitionResult.Type).Ref("squad"),
		edge.To("owner", Profile.Type).Unique(),
	}
}
