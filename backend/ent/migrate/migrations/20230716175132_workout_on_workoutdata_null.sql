-- Modify "workouts" table
ALTER TABLE "workouts" DROP COLUMN "workout_data_workout", DROP COLUMN "workout_route_data_workout";
-- Modify "workout_data" table
ALTER TABLE "workout_data" ADD COLUMN "workout_data_workout" uuid NULL, ADD CONSTRAINT "workout_data_workouts_workout" FOREIGN KEY ("workout_data_workout") REFERENCES "workouts" ("id") ON DELETE SET NULL;
-- Modify "workout_route_data" table
ALTER TABLE "workout_route_data" ADD COLUMN "workout_route_data_workout" uuid NOT NULL, ADD CONSTRAINT "workout_route_data_workouts_workout" FOREIGN KEY ("workout_route_data_workout") REFERENCES "workouts" ("id") ON DELETE NO ACTION;
