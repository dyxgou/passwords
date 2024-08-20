package main

import (
	"log"
	"net/http"

	"gihub.com/dyxgou/server/internals/handlers"
	"gihub.com/dyxgou/server/internals/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	type response struct {
		Message string `json:"message"`
	}

	r.Get("/healtz", handlers.ReadinessHandler)
	r.Get("/", handlers.EntryHandler)
	r.Get("/error", handlers.ErrorHandler)

	r.Mount("/users", routes.UserRoutes())
	r.Mount("/passwords", routes.PasswordRoutes())

	log.Println("Server running")
	http.ListenAndServe(":5000", r)
}
