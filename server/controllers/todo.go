package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ysrckr/frontendmentor-todo-app/modals"
	"github.com/ysrckr/frontendmentor-todo-app/services"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	completed := r.URL.Query().Get("completed")

	var err error
	var status bool

	if completed != "" {

		if completed != "true" && completed != "false" {
			http.Error(w, "Wrong Query", http.StatusBadRequest)
			return
		}

		status, err = strconv.ParseBool(completed)
		if err != nil {
			http.Error(w, "Wrong Query", http.StatusBadRequest)
			return
		}

		todos, err := services.SelectAllTodosWithCompletedStatus(status)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error:%s", err), http.StatusInternalServerError)
		}

		todosJson, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, "JSON Parse Error", http.StatusInternalServerError)
		}

		w.Header().Add("Content-Type", "application/json")

		w.Write([]byte(todosJson))
		return
	}

	todos, err := services.SelectAllTodos()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error:%s", err), http.StatusInternalServerError)
	}

	todosJson, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "JSON Parse Error", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")

	w.Write([]byte(todosJson))
}

func CreateATodo(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	var todo modals.Todo

	err := json.NewDecoder(body).Decode(&todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error:%s", err), http.StatusInternalServerError)
	}

	id, err := services.InsertATodo(todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error:%s", err), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "Id:%d", id)
}

func ToggleTodoStatus(w http.ResponseWriter, r *http.Request) {
	var result int
	var err error
	if id := chi.URLParam(r, "id"); id != "" {
		idx, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error:%s", err), http.StatusInternalServerError)
		}
		result, err = services.UpdateATodoStatusWithOpposite(int(idx))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error:%s", err), http.StatusInternalServerError)
		}

	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error:%s", err), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Id:%d", result)
}
