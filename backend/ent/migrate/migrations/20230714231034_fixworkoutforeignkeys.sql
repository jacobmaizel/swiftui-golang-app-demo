-- Modify "workout_data" table
ALTER TABLE "workout_data" DROP COLUMN "workout_workout_data";
-- Modify "workout_route_data" table
ALTER TABLE "workout_route_data" DROP COLUMN "workout_workout_route_data";
-- Modify "workouts" table
ALTER TABLE "workouts" ADD COLUMN "workout_data_workout" uuid NULL, ADD COLUMN "workout_route_data_workout" uuid NULL, ADD CONSTRAINT "workouts_workout_data_workout" FOREIGN KEY ("workout_data_workout") REFERENCES "workout_data" ("id") ON DELETE SET NULL, ADD CONSTRAINT "workouts_workout_route_data_workout" FOREIGN KEY ("workout_route_data_workout") REFERENCES "workout_route_data" ("id") ON DELETE SET NULL;
-- Create index "workouts_workout_data_workout_key" to table: "workouts"
CREATE UNIQUE INDEX "workouts_workout_data_workout_key" ON "workouts" ("workout_data_workout");
-- Create index "workouts_workout_route_data_workout_key" to table: "workouts"
CREATE UNIQUE INDEX "workouts_workout_route_data_workout_key" ON "workouts" ("workout_route_data_workout");
