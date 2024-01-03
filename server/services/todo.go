package services

import (
	"context"
	"log"
	"time"

	"github.com/ysrckr/frontendmentor-todo-app/database"
	"github.com/ysrckr/frontendmentor-todo-app/modals"
)

func createTimedContext(duration time.Duration) (context.Context, context.CancelFunc) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*duration))

	return ctx, cancel
}

func SelectAllTodos() ([]modals.Todo, error) {

	todos := []modals.Todo{}

	ctx, cancel := createTimedContext(200)

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

	ctx, cancel := createTimedContext(200)

	defer cancel()

	err := database.DB.Db.SelectContext(ctx, &todos, "SELECT * FROM todos where completed=$1", status)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return todos, err
}

func InsertATodo(todo modals.Todo) (int, error) {
	ctx, cancel := createTimedContext(200)

	defer cancel()

	todoQuery := `INSERT INTO todos (text, completed) VALUES ($1, $2) RETURNING id`

	lastInsertedId := 0

	err := database.DB.Db.QueryRowContext(ctx, todoQuery, todo.Text, todo.Completed).Scan(&lastInsertedId)
	if err != nil {
		return 0, err
	}

	return lastInsertedId, nil
}

func UpdateATodoStatusWithOpposite(id int) (int, error) {
	ctx, cancel := createTimedContext(200)

	defer cancel()

	todoQuery := `UPDATE todos SET completed = NOT completed WHERE id = $1 RETURNING id`

	lastUpdatedId := 0

	err := database.DB.Db.QueryRowContext(ctx, todoQuery, id).Scan(&lastUpdatedId)
	if err != nil {
		return 0, err
	}

	return lastUpdatedId, nil

}

func RemoveATodo(id int) error {
	ctx, cancel := createTimedContext(200)

	defer cancel()

	todoQuery := `DELETE FROM todos WHERE id = $1`

	err := database.DB.Db.QueryRowContext(ctx, todoQuery, id).Err()

	return err
}
