CREATE TABLE "todos" (
  "id" integer PRIMARY KEY,
  "text" varchar,
  "completed" boolean,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);