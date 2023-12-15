package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type PrivacyMixin struct {
	mixin.Schema
}

func (PrivacyMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("public").Default(true),
	}
}
