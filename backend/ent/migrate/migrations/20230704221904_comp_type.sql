-- Modify "competitions" table
ALTER TABLE "competitions" ADD COLUMN "type" character varying NOT NULL DEFAULT 'official';
