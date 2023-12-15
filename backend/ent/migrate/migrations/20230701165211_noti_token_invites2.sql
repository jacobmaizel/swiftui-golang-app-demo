-- Modify "profiles" table
ALTER TABLE "profiles" ADD COLUMN "public" boolean NOT NULL DEFAULT true;
-- Create "notification_preferences" table
CREATE TABLE "notification_preferences" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "settings" jsonb NOT NULL, "profile_notification_preferences" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "notification_preferences_profiles_notification_preferences" FOREIGN KEY ("profile_notification_preferences") REFERENCES "profiles" ("id") ON DELETE NO ACTION);
-- Create index "notification_preferences_profile_notification_preferences_key" to table: "notification_preferences"
CREATE UNIQUE INDEX "notification_preferences_profile_notification_preferences_key" ON "notification_preferences" ("profile_notification_preferences");
-- Create "workout_data" table
CREATE TABLE "workout_data" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "average_heart_rate" character varying NULL, "distance" double precision NULL, "duration" character varying NULL, "energy_burned" character varying NULL, "workout_workout_data" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "workout_data_workouts_workout_data" FOREIGN KEY ("workout_workout_data") REFERENCES "workouts" ("id") ON DELETE NO ACTION);
-- Create index "workout_data_workout_workout_data_key" to table: "workout_data"
CREATE UNIQUE INDEX "workout_data_workout_workout_data_key" ON "workout_data" ("workout_workout_data");
-- Create "workout_route_data" table
CREATE TABLE "workout_route_data" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "workout_workout_route_data" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "workout_route_data_workouts_workout_route_data" FOREIGN KEY ("workout_workout_route_data") REFERENCES "workouts" ("id") ON DELETE NO ACTION);
-- Create index "workout_route_data_workout_workout_route_data_key" to table: "workout_route_data"
CREATE UNIQUE INDEX "workout_route_data_workout_workout_route_data_key" ON "workout_route_data" ("workout_workout_route_data");
-- Modify "squads" table
ALTER TABLE "squads" ADD COLUMN "public" boolean NOT NULL DEFAULT true;
-- Modify "competitions" table
ALTER TABLE "competitions" ADD COLUMN "public" boolean NOT NULL DEFAULT true, ADD COLUMN "scheduled" boolean NOT NULL DEFAULT false, ADD COLUMN "workout_types" jsonb NULL;
-- Modify "invites" table
ALTER TABLE "invites" DROP CONSTRAINT "invites_squads_squad", ALTER COLUMN "invite_squad" DROP NOT NULL, ADD COLUMN "invite_competition" uuid NULL, ADD COLUMN "invite_workout" uuid NULL, ADD CONSTRAINT "invites_squads_squad" FOREIGN KEY ("invite_squad") REFERENCES "squads" ("id") ON DELETE SET NULL, ADD CONSTRAINT "invites_competitions_competition" FOREIGN KEY ("invite_competition") REFERENCES "competitions" ("id") ON DELETE SET NULL, ADD CONSTRAINT "invites_workouts_workout" FOREIGN KEY ("invite_workout") REFERENCES "workouts" ("id") ON DELETE SET NULL;
-- Create "fcm_tokens" table
CREATE TABLE "fcm_tokens" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "token" character varying NOT NULL, "profile_fcm_tokens" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "fcm_tokens_profiles_fcm_tokens" FOREIGN KEY ("profile_fcm_tokens") REFERENCES "profiles" ("id") ON DELETE NO ACTION);
-- Create "notifications" table
CREATE TABLE "notifications" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "title" character varying NOT NULL, "body" character varying NOT NULL, "sent" timestamptz NOT NULL, "opened" timestamptz NULL, "data" jsonb NOT NULL, "profile_notifications" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "notifications_profiles_notifications" FOREIGN KEY ("profile_notifications") REFERENCES "profiles" ("id") ON DELETE NO ACTION);
