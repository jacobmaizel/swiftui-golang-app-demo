package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// FcmToken holds the schema definition for the FcmToken entity.
type FcmToken struct {
	ent.Schema
}

func (FcmToken) Mixin() []ent.Mixin {
	return mixins.BaseMixin()
}

// Fields of the FcmToken.
func (FcmToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token"),
	}
}

// Edges of the FcmToken.
func (FcmToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("fcm_tokens").Unique().Required(),
	}
}
