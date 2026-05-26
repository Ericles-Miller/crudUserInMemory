package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Ericles-Miller/crudUserInMemory/application"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error to make marshal de json", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(data); err != nil {
		slog.Error("error to send the response", "error", err)
		return
	}
}

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

func HandleCreateUser(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userBody application.UserBody

		if error := json.NewDecoder(r.Body).Decode(&userBody); error != nil {
			sendJSON(w, Response{Error: "Invalid request body"}, http.StatusBadRequest)
			return
		}

		createUser, err := app.Insert(userBody)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		sendJSON(w, Response{Data: createUser}, http.StatusCreated)
	}
}

func HandleGetUser(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			sendJSON(w, Response{Error: "id is required"}, http.StatusBadRequest)
			return
		}

		parsedId, err := uuid.Parse(id)
		if err != nil {
			sendJSON(w, Response{Error: "invalid id"}, http.StatusBadRequest)
			return
		}

		user, ok := app.FindById(parsedId)
		if !ok {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func HandleGetAllUsers(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users := app.FindAll()

		sendJSON(w, Response{Data: users}, http.StatusOK)
	}
}

func HandleDeleteUser(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			sendJSON(w, Response{Error: "id is required"}, http.StatusBadRequest)
			return
		}

		parsedId, err := uuid.Parse(id)
		if err != nil {
			sendJSON(w, Response{Error: "invalid id"}, http.StatusBadRequest)
			return
		}

		user, ok := app.Delete(parsedId)

		if !ok {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)

	}
}

func HandleUpdateUser(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			sendJSON(w, Response{Error: "id is required"}, http.StatusBadRequest)
			return
		}

		parsedId, err := uuid.Parse(id)
		if err != nil {
			sendJSON(w, Response{Error: "invalid id"}, http.StatusBadRequest)
			return
		}

		var userBody application.UserBody

		if error := json.NewDecoder(r.Body).Decode(&userBody); error != nil {
			sendJSON(w, Response{Error: "Invalid request body"}, http.StatusBadRequest)
			return
		}

		user, ok, err := app.Update(parsedId, userBody)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		if !ok {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
