//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"reflect"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	// _ "github.com/jacobmaizel/swiftui-golang-app-demo/ent/runtime"
)

func main() {

	// names := entc.FeatureNames("sql/versioned-migration")
	/* opts:= []entc.Option{ */

	/* } */
	cfg := &gen.Config{
		Hooks: []gen.Hook{
			EnsureStructTag("json"),
			EnsureNoOmitEmptyonBoolField("omitempty"),
		},
		Features: []gen.Feature{
			gen.FeatureVersionedMigration,
			gen.FeatureIntercept,
			gen.FeatureModifier,
		},
	}

	schemaPath := "./schema"

	// if err := entc.Generate(schemaPath, cfg, entc.BuildTags("skiphooks")); err != nil {
	// 	log.Fatalf("failed running first ent codegen: %s", err)
	// }

	// if err := entc.Generate(schemaPath, cfg, entc.BuildTags("skippolicy")); err != nil {
	// 	log.Fatalf("failed running second ent codegen with hooks: %s", err)
	// }

	if err := entc.Generate(schemaPath, cfg); err != nil {
		log.Fatalf("failed running third ent codegen with hooks: %s", err)
	}

}

// EnsureStructTag ensures all fields in the graph have a specific tag name.
func EnsureStructTag(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					tag := reflect.StructTag(field.StructTag)
					if _, ok := tag.Lookup(name); !ok {
						return fmt.Errorf("struct tag %q is missing for field %s.%s", name, node.Name, field.Name)
					}
				}
			}
			return next.Generate(g)
		})
	}
}

// Ensures that false values for bools are also returned in queries.
// right now falsy go values like nil or false allow omitempty to remove it from the response altogether
// no bueno
func EnsureNoOmitEmptyonBoolField(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					tag := reflect.StructTag(field.StructTag)
					if node.TypeName() == "bool" {
						if _, ok := tag.Lookup("omitempty"); ok {
							log.Println("Found an omit empty on a bool")
						}
					}

				}
			}
			return next.Generate(g)
		})
	}
}
