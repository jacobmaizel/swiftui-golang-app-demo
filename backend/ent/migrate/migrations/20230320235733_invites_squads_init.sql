-- Create "squads" table
CREATE TABLE "squads" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "title" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "squads_title_key" to table: "squads"
CREATE UNIQUE INDEX "squads_title_key" ON "squads" ("title");
-- Create "squad_members" table
CREATE TABLE "squad_members" ("squad_id" uuid NOT NULL, "profile_id" uuid NOT NULL, PRIMARY KEY ("squad_id", "profile_id"), CONSTRAINT "squad_members_squad_id" FOREIGN KEY ("squad_id") REFERENCES "squads" ("id") ON DELETE CASCADE, CONSTRAINT "squad_members_profile_id" FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id") ON DELETE CASCADE);
-- Create "invites" table
CREATE TABLE "invites" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'pending', "invite_sender" uuid NOT NULL, "invite_receiver" uuid NOT NULL, "invite_squad" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "invites_profiles_sender" FOREIGN KEY ("invite_sender") REFERENCES "profiles" ("id") ON DELETE NO ACTION, CONSTRAINT "invites_profiles_receiver" FOREIGN KEY ("invite_receiver") REFERENCES "profiles" ("id") ON DELETE NO ACTION, CONSTRAINT "invites_squads_squad" FOREIGN KEY ("invite_squad") REFERENCES "squads" ("id") ON DELETE NO ACTION);
