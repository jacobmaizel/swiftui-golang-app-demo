-- Modify "workout_data" table
ALTER TABLE "workout_data" ADD COLUMN "healthkit_workout_id" character varying NULL, ADD COLUMN "healthkit_workout_start_date" timestamptz NULL, ADD COLUMN "healthkit_workout_end_date" timestamptz NULL, ADD COLUMN "source" character varying NOT NULL DEFAULT 'trainton';
-- Modify "workouts" table
ALTER TABLE "workouts" DROP COLUMN "healthkit_workout_id", DROP COLUMN "workout_profile", DROP COLUMN "healthkit_workout_start_date", DROP COLUMN "healthkit_workout_end_date", DROP COLUMN "source", ADD COLUMN "workout_leader" uuid NULL, ADD CONSTRAINT "workouts_profiles_leader" FOREIGN KEY ("workout_leader") REFERENCES "profiles" ("id") ON DELETE SET NULL;
