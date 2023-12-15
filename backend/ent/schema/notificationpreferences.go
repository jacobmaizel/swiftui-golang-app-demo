package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// NotificationPreferences holds the schema definition for the NotificationPreferences entity.
type NotificationPreferences struct {
	ent.Schema
}

func (NotificationPreferences) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the NotificationPreferences.
func (NotificationPreferences) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("settings", &shared.NotificationPreferenceSettings{}),
	}
}

// Edges of the NotificationPreferences.
func (NotificationPreferences) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("notification_preferences").Unique().Required(),
	}
}
