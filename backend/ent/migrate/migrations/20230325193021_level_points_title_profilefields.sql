-- Modify "profiles" table
ALTER TABLE "profiles" ADD COLUMN "level" bigint NOT NULL DEFAULT 1, ADD COLUMN "points" bigint NOT NULL DEFAULT 0, ADD COLUMN "title" character varying NOT NULL DEFAULT 'Novice';
