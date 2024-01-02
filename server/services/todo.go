package services

import (
	"log"

	"github.com/ysrckr/frontendmentor-todo-app/database"
	"github.com/ysrckr/frontendmentor-todo-app/modals"
)

func SelectAllTodos() []modals.Todo {

	rows, err := database.DB.Db.Queryx("SELECT * FROM todos")
	if err != nil {
		log.Fatalln(err)
	}

	var todos []modals.Todo

	for rows.Next() {
		var todo modals.Todo
		err = rows.StructScan(&todo)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)

	}

	return todos
}
