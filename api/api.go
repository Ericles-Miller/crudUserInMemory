package api

import (
	"net/http"
	"github.com/Ericles-Miller/crudUserInMemory/application"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)


func NewHandler() http.Handler {
	app := application.New()
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// declare endpoint here
	r.Post("/api/users", HandleCreateUser(app))
	r.Get("/api/users/{id}", HandleGetUser(app))
	r.Get("/api/users", HandleGetAllUsers(app))
	r.Delete("/api/users/{id}", HandleDeleteUser(app))
	r.Put("/api/users/{id}", HandleUpdateUser(app))

	return r
}

