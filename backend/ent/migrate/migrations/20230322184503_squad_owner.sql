-- Modify "squads" table
ALTER TABLE "squads" ADD COLUMN "squad_owner" uuid NULL, ADD CONSTRAINT "squads_profiles_owner" FOREIGN KEY ("squad_owner") REFERENCES "profiles" ("id") ON DELETE SET NULL;
