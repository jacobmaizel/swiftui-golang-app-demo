-- Modify "competitions" table
ALTER TABLE "competitions" DROP COLUMN "category", DROP COLUMN "level", DROP COLUMN "tags", DROP COLUMN "group_style", DROP COLUMN "official";
-- Modify "profiles" table
ALTER TABLE "profiles" DROP COLUMN "level", DROP COLUMN "points", DROP COLUMN "title";
-- Modify "workouts" table
ALTER TABLE "workouts" DROP COLUMN "average_heart_rate", DROP COLUMN "distance", DROP COLUMN "duration", DROP COLUMN "energy_burned";
