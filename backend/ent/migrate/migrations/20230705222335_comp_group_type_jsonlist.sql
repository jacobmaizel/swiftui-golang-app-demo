-- Modify "competitions" table
ALTER TABLE "competitions" DROP COLUMN "group_type", ADD COLUMN "participant_types" jsonb NULL;
