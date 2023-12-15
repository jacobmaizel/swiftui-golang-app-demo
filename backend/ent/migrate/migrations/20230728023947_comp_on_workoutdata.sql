-- Modify "workout_data" table
ALTER TABLE "workout_data" ADD COLUMN "workout_data_competition" uuid NULL, ADD CONSTRAINT "workout_data_competitions_competition" FOREIGN KEY ("workout_data_competition") REFERENCES "competitions" ("id") ON DELETE SET NULL;
-- Modify "workout_route_data" table
ALTER TABLE "workout_route_data" ADD COLUMN "workout_data_workout_route_data" uuid NULL, ADD CONSTRAINT "workout_route_data_workout_data_workout_route_data" FOREIGN KEY ("workout_data_workout_route_data") REFERENCES "workout_data" ("id") ON DELETE SET NULL;
-- Create index "workout_route_data_workout_data_workout_route_data_key" to table: "workout_route_data"
CREATE UNIQUE INDEX "workout_route_data_workout_data_workout_route_data_key" ON "workout_route_data" ("workout_data_workout_route_data");
