package schema

import (
	"entgo.io/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema/mixins"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

func (Post) Mixin() []ent.Mixin {
	return append(mixins.BaseMixin(), mixins.PrivacyMixin{})
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return nil
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
