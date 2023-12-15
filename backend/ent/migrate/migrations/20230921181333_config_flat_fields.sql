-- Modify "app_configs" table
ALTER TABLE "app_configs" DROP COLUMN "app_config", ADD COLUMN "auto_sync_workouts" boolean NOT NULL DEFAULT true;
-- Create "activities" table
CREATE TABLE "activities" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "public" boolean NOT NULL DEFAULT true, PRIMARY KEY ("id"));
-- Create "posts" table
CREATE TABLE "posts" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "public" boolean NOT NULL DEFAULT true, PRIMARY KEY ("id"));
