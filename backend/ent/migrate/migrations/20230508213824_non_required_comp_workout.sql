-- Modify "workouts" table
ALTER TABLE "workouts" DROP CONSTRAINT "workouts_competitions_competition", ALTER COLUMN "workout_competition" DROP NOT NULL, ADD CONSTRAINT "workouts_competitions_competition" FOREIGN KEY ("workout_competition") REFERENCES "competitions" ("id") ON DELETE SET NULL;
