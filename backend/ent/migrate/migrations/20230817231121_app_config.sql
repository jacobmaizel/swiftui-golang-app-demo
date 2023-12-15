-- Create "app_configs" table
CREATE TABLE "app_configs" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "app_config" jsonb NOT NULL, PRIMARY KEY ("id"));
-- Modify "profiles" table
ALTER TABLE "profiles" ADD COLUMN "app_config_profile" uuid NULL, ADD CONSTRAINT "profiles_app_configs_profile" FOREIGN KEY ("app_config_profile") REFERENCES "app_configs" ("id") ON DELETE SET NULL;
-- Create index "profiles_app_config_profile_key" to table: "profiles"
CREATE UNIQUE INDEX "profiles_app_config_profile_key" ON "profiles" ("app_config_profile");
