package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// CompetitionResult holds the schema definition for the CompetitionResult entity.
type CompetitionResult struct {
	ent.Schema
}

// Mixin of the CompetitionResult.
func (CompetitionResult) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the CompetitionResult.
func (CompetitionResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("place"),
	}
}

// Edges of the CompetitionResult.
func (CompetitionResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("competition", Competition.Type).Unique().Required(),
		edge.To("profile", Profile.Type).Unique(),
		edge.To("squad", Squad.Type).Unique(),
	}
}
