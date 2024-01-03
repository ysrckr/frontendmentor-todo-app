package services

import (
	"context"
	"log"
	"time"

	"github.com/ysrckr/frontendmentor-todo-app/database"
	"github.com/ysrckr/frontendmentor-todo-app/modals"
)

func SelectAllTodos() ([]modals.Todo, error) {

	todos := []modals.Todo{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*200))

	defer cancel()

	err := database.DB.Db.SelectContext(ctx, &todos, "SELECT * FROM todos")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return todos, err
}

func SelectAllTodosWithCompletedStatus(status bool) ([]modals.Todo, error) {
	todos := []modals.Todo{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*200))

	defer cancel()

	err := database.DB.Db.SelectContext(ctx, &todos, "SELECT * FROM todos where completed=$1", status)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return todos, err
}
