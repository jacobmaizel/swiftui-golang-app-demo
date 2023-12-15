package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// Invite holds the schema definition for the Invite entity.
type Invite struct {
	ent.Schema
}

func (Invite) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the Invite.
func (Invite) Fields() []ent.Field {
	return []ent.Field{
		// pending, accepted, declined
		field.String("status").Default("pending"),
	}
}

// Edges of the Invite.
func (Invite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sender", Profile.Type).Unique().Required(),
		edge.To("receiver", Profile.Type).Unique().Required(),

		edge.To("squad", Squad.Type).Unique(),
		edge.To("competition", Competition.Type).Unique(),
		edge.To("workout", Workout.Type).Unique(),
	}
}
