package routes

import (
	handlers "gihub.com/dyxgou/server/internals/handlers/users"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/create", handlers.CreateUser)
	r.Get("/all", handlers.GetAllUsers)
	r.Get("/all/points", handlers.GetUsersSortedByPoints)
	r.Get("/{id}", handlers.GetUserById)
	r.Get("/name/{name}", handlers.GetUserByName)
	r.Get("/id/{name}", handlers.GetIdByName)
	r.Get("/nameid/{id}", handlers.GetNameById)
	r.Put("/points", handlers.AddUserPoints)
	r.Get("/points/{id}", handlers.GetUserPoints)

	return r
}
