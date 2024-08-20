package routes

import (
	handlers "gihub.com/dyxgou/server/internals/handlers/passwords"
	"github.com/go-chi/chi/v5"
)

func PasswordRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/create", handlers.CreatePassword)
	r.Get("/all", handlers.GetAllPasswords)
	r.Get("/all/{id}", handlers.GetUserPasswords)

	return r
}
