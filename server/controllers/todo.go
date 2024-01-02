package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ysrckr/frontendmentor-todo-app/services"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := services.SelectAllTodos()

	fmt.Println(todos)

	todosJson, err := json.Marshal(todos)
	if err != nil {
		fmt.Fprintf(w, "Parse error")
	}

	w.Write([]byte(todosJson))
}
