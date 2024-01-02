-- Active: 1704121805152@@127.0.0.1@5432@todoapp
CREATE TABLE "todos" (
  "id" SERIAL PRIMARY KEY,
  "text" varchar,
  "completed" boolean DEFAULT (false),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);