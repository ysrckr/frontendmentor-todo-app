package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/ysrckr/frontendmentor-todo-app/controllers"
)

func InitialiseServer(PORT string, AllowedOrigins, AllowMethods []string) {
	r := chi.NewRouter()

	apiRouter := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/health-check"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: AllowedOrigins,
		// AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   AllowMethods,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apiRouter.Route("/todos", func(r chi.Router) {
		r.Get("/", controllers.GetAllTodos)
		r.Post("/", controllers.CreateATodo)
		r.Patch("/{id}", controllers.ToggleTodoStatus)
		r.Delete("/{id}", controllers.DeleteATodo)
	})

	r.Mount("/api/v1", apiRouter)

	log.Printf("Listening on http://localhost:%s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
