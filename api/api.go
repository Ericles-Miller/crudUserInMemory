package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"github.com/google/uuid"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Biography string `json:"biography"`
	Id uuid.UUID `json:"id"`
}

type UserBody struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Biography string `json:"biography"`
}

type Response struct {
	Data User `json:"data, omitempty"`
	Error string `json:"error, omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error ao fazer marshal de json", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error ao enviar a resposta", "error", err)
		return
	}
}

func NewHandler(db map[string]User) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// declare endpoint here 
	r.Post("/users", HandleCreateUser(db))

	return r 
}

func HandleCreateUser(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userBody UserBody

		if error := json.NewDecoder(r.Body).Decode(&userBody); error != nil {
			sendJSON(w, Response{Error: "Invalid request body"}, http.StatusBadRequest)
			return
		}

		if error := ValidateUser(userBody); error != nil {
			sendJSON(w, Response{Error: error.Error()}, http.StatusBadRequest)
			return
		}

		user := User {
			FirstName: userBody.FirstName,
			LastName: userBody.LastName,
			Biography: userBody.Biography,
			Id: uuid.New(),
		}
		db[user.Id.String()] = user
		sendJSON(w, Response{Data: user}, http.StatusCreated)
	}
}

func ValidateUser(user UserBody)  error {
	if user.FirstName == "" || user.LastName == "" {
		return fmt.Errorf("First Name and LastName is required")
	}

	if(len(user.FirstName) <=2 || len(user.LastName) <= 2) {
		return fmt.Errorf("First Name and LastName must be at least 3 characters long")
	}

	if user.Biography == "" {
		return fmt.Errorf("Biography is required")
	}
	return nil
}