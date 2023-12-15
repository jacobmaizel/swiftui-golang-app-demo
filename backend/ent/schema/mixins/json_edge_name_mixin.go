package mixins

import (
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type JsonEdgeNameMixin struct {
	mixin.Schema
}

func (JsonEdgeNameMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{edge.Annotation{
		StructTag: `json:"related"`,
	},

		field.Annotation{
			StructTag: map[string]string{
				"onboarding_completed": `json:"onboarding_completed"`,
			},
		},

		field.Annotation{
			StructTag: map[string]string{
				"public": `json:"public"`,
			},
		},

		field.Annotation{
			StructTag: map[string]string{
				"scheduled": `json:"scheduled"`,
			},
		},
	}
}
