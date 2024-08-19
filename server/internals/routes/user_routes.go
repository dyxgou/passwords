package routes

import (
	"gihub.com/dyxgou/server/internals/handlers"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/create", handlers.CreateUser)
	r.Get("/all", handlers.GetAllUsers)

	return r
}
