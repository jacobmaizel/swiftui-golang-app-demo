-- Modify "workouts" table
ALTER TABLE "workouts" ADD COLUMN "location_type" character varying NOT NULL DEFAULT 'indoor';
