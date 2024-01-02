package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ysrckr/frontendmentor-todo-app/services"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := services.SelectAllTodos()

	todosJson, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "JSON Parse Error", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")

	w.Write([]byte(todosJson))
}
