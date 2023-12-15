-- Modify "workouts" table
ALTER TABLE "workouts" ADD COLUMN "source" character varying NOT NULL DEFAULT 'trainton';
