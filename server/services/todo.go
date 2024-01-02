package services

import (
	"log"

	"github.com/ysrckr/frontendmentor-todo-app/database"
	"github.com/ysrckr/frontendmentor-todo-app/modals"
)

func SelectAllTodos() []modals.Todo {

	todos := []modals.Todo{}

	err := database.DB.Db.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		log.Fatalln(err)
	}

	return todos
}

func SelectAllTodosWithCompletedStatus(status bool) []modals.Todo {
	todos := []modals.Todo{}

	err := database.DB.Db.Select(&todos, "SELECT * FROM todos where completed=$1", status)
	if err != nil {
		log.Fatalln(err)
	}

	return todos
}
