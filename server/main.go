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

	database.DB.Initialise("postgres", "admin", "admin", "todoapp")

	InitialiseServer(PORT)
}
