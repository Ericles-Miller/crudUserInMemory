package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Ericles-Miller/crudUserInMemory/application"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func HandleCreateUser(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body application.UserBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "Invalid request body"}, http.StatusBadRequest)
			return
		}

		user, err := app.Insert(r.Context(), body)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusCreated)
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

		user, err := app.FindById(r.Context(), parsedId)
		if errors.Is(err, application.ErrNotFound) {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}
		if err != nil {
			sendJSON(w, Response{Error: "failed to get user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func HandleGetAllUsers(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := app.FindAll(r.Context())
		if err != nil {
			sendJSON(w, Response{Error: "failed to get users"}, http.StatusInternalServerError)
			return
		}

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

		user, err := app.Delete(r.Context(), parsedId)
		if errors.Is(err, application.ErrNotFound) {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}
		if err != nil {
			sendJSON(w, Response{Error: "failed to delete user"}, http.StatusInternalServerError)
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

		var body application.UserBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "Invalid request body"}, http.StatusBadRequest)
			return
		}

		user, err := app.Update(r.Context(), parsedId, body)
		if errors.Is(err, application.ErrNotFound) {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
