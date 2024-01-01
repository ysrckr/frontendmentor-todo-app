CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "ip" varchar UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "todos" (
  "id" integer PRIMARY KEY,
  "text" varchar,
  "completed" boolean,
  "user_id" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "users" USING HASH ("ip");

ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
