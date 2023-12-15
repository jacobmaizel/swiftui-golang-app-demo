-- Modify "goals" table
ALTER TABLE "goals" ADD COLUMN "status" character varying NOT NULL DEFAULT 'active';
