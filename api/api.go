package api

import (
	"net/http"

	"github.com/Ericles-Miller/crudUserInMemory/application"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewHandler(app *application.Application) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/users", HandleCreateUser(app))
	r.Get("/api/users/{id}", HandleGetUser(app))
	r.Get("/api/users", HandleGetAllUsers(app))
	r.Delete("/api/users/{id}", HandleDeleteUser(app))
	r.Put("/api/users/{id}", HandleUpdateUser(app))

	return r
}

