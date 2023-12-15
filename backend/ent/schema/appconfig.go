package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// AppConfig holds the schema definition for the AppConfig entity.
type AppConfig struct {
	ent.Schema
}

func (AppConfig) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the NotificationPreferences.
func (AppConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("auto_sync_workouts").Default(true),
	}
}

// Edges of the NotificationPreferences.
func (AppConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Unique().Required(),
	}
}
