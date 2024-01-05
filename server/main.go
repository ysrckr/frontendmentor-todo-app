package main

import (
	"os"

	_ "github.com/lib/pq"

	"github.com/ysrckr/frontendmentor-todo-app/database"
)

var allowedMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}

func main() {
	if len(os.Args) > 1 {
		Env.SetEnviromentMode(os.Args[1])
	}

	Env.SetEnviromentVars()

	database.DB.Initialise(Env.DB, Env.DBUSER, Env.DBPASSWORD, Env.DBNAME)

	InitialiseServer(Env.PORT, Env.ALLOWEDORIGINS, allowedMethods)
}
