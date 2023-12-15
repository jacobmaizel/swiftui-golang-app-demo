//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
)

const (
	dir = "ent/migrate/migrations"
)

func main() {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("creating migration directory: %v", err)
	}
	dir, err := atlas.NewLocalDir(dir)
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                          // provide migration directory
		schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
		schema.WithDialect(dialect.Postgres),
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithDropColumn(true),
	}
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}

	err = migrate.NamedDiff(ctx, "postgres://postgres:postgres@db:5432/postgres?sslmode=disable", os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
