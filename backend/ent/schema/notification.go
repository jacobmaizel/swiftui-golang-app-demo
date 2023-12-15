package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

func (Notification) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("body"),
		field.Time("sent"),
		field.Time("opened").Nillable().Optional(),
		// field.String("type"),
		field.JSON("data", &shared.NotificationData{}).Optional(),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("notifications").Unique().Required(),
	}
}
