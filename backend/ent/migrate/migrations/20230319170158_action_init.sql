-- Create "actions" table
CREATE TABLE "actions" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "title" character varying NOT NULL, "action_sender" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "actions_profiles_sender" FOREIGN KEY ("action_sender") REFERENCES "profiles" ("id") ON DELETE SET NULL);
