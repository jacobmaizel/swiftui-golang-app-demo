-- Modify "workouts" table
ALTER TABLE "workouts" ADD COLUMN "healthkit_workout_start_date" timestamptz NULL, ADD COLUMN "healthkit_workout_end_date" timestamptz NULL;
