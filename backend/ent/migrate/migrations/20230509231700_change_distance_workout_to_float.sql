-- Modify "workouts" table
ALTER TABLE workouts ALTER COLUMN distance TYPE double precision USING distance::double precision;
