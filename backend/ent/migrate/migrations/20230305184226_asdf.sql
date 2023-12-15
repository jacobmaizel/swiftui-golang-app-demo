-- create "competitions" table
CREATE TABLE "competitions" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "title" character varying NOT NULL, "description" character varying NOT NULL, "category" character varying NOT NULL, "level" character varying NOT NULL, "tags" character varying NULL, "start" timestamptz NOT NULL, "end" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "subject" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "users_subject_key" to table: "users"
CREATE UNIQUE INDEX "users_subject_key" ON "users" ("subject");
-- create "profiles" table
CREATE TABLE "profiles" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "first_name" character varying NOT NULL, "last_name" character varying NOT NULL, "birthday" timestamptz NULL, "picture" character varying NULL, "onboarding_completed" boolean NOT NULL DEFAULT false, "user_profile" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "profiles_users_profile" FOREIGN KEY ("user_profile") REFERENCES "users" ("id") ON DELETE NO ACTION);
-- create index "profiles_user_profile_key" to table: "profiles"
CREATE UNIQUE INDEX "profiles_user_profile_key" ON "profiles" ("user_profile");
