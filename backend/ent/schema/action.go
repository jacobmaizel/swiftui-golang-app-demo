package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// Action holds the schema definition for the Action entity.
type Action struct {
	ent.Schema
}

func (Action) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the Action.
func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
	}
}

// Edges of the Action.
func (Action) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sender", Profile.Type).Unique(),
	}
}
