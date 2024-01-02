package services

import (
	"github.com/ysrckr/frontendmentor-todo-app/database"
	"github.com/ysrckr/frontendmentor-todo-app/modals"
)

func SelectAllTodos() []modals.Todo {
	todos := []modals.Todo{}

	database.DB.Db.Select(&todos, "SELECT * FROM todos")

	return todos
}
