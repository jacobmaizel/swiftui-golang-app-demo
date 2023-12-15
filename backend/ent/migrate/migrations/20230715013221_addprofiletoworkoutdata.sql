-- Modify "workouts" table
ALTER TABLE "workouts" DROP COLUMN "status";
-- Modify "workout_data" table
ALTER TABLE "workout_data" ADD COLUMN "location_type" character varying NOT NULL DEFAULT 'indoor', ADD COLUMN "workout_data_profile" uuid NOT NULL, ADD CONSTRAINT "workout_data_profiles_profile" FOREIGN KEY ("workout_data_profile") REFERENCES "profiles" ("id") ON DELETE NO ACTION;
