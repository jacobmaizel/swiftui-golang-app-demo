-- modify "competitions" table
ALTER TABLE "competitions" ADD COLUMN "competition_owner" uuid NOT NULL, ADD CONSTRAINT "competitions_profiles_owner" FOREIGN KEY ("competition_owner") REFERENCES "profiles" ("id") ON DELETE NO ACTION;
