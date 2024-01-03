package main

import (
	_ "github.com/lib/pq"

	"github.com/ysrckr/frontendmentor-todo-app/database"
)

var allowedOrigins = []string{"http://localhost:3000"}
var allowedMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}

func main() {

	Env.SetEnviromentVars()

	database.DB.Initialise(Env.DB, Env.DBUSER, Env.DBPASSWORD, Env.DBNAME)

	InitialiseServer(Env.PORT, allowedOrigins, allowedMethods)
}
