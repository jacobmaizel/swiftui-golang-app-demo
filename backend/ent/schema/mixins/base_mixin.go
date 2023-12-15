package mixins

import (
	"entgo.io/ent"
)

func BaseMixin() []ent.Mixin {
	return []ent.Mixin{
		IdMixin{},
		JsonEdgeNameMixin{},
		CreateUpdateTimeMixin{},
	}
}
