-- Modify "workouts" table
ALTER TABLE "workouts" ADD COLUMN "status" character varying NOT NULL DEFAULT 'active';
