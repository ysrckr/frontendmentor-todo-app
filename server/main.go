package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/ysrckr/frontendmentor-todo-app/database"
)

func main() {
  godotenv.Load("../.env")

  PORT := os.Getenv("API_PORT")
  DB := os.Getenv("DB")
  DBUSER := os.Getenv("DB_USER")
  DBPASSWORD := os.Getenv("DB_PASSWORD")
  DBNAME := os.Getenv("DB_NAME")

  database.DB.Initialise(DB, DBUSER, DBPASSWORD, DBNAME)

  InitialiseServer(PORT)
}
