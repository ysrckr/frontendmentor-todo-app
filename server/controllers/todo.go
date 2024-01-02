package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ysrckr/frontendmentor-todo-app/services"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	completed := r.URL.Query().Get("completed")
	var status bool
	if completed != "" && completed == "true" || completed == "false" {

		if completed == "true" {
			status = true
		}

		if completed == "false" {
			status = false
		}
    
		todos := services.SelectAllTodosWithCompletedStatus(status)

		todosJson, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, "JSON Parse Error", http.StatusInternalServerError)
		}

		w.Header().Add("Content-Type", "application/json")

		w.Write([]byte(todosJson))
		return
	}

	todos := services.SelectAllTodos()

	todosJson, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "JSON Parse Error", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")

	w.Write([]byte(todosJson))
}
