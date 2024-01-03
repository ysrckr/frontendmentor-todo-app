package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/ysrckr/frontendmentor-todo-app/controllers"
)

func InitialiseServer(PORT string) {
	r := chi.NewRouter()

	apiRouter := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		// AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apiRouter.Route("/todos", func(r chi.Router) {
		r.Get("/", controllers.GetAllTodos)
	})

	r.Mount("/api/v1", apiRouter)

	log.Printf("Listening on http://localhost:%s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
