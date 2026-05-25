package main

import (
	"log/slog"
	"net/http"
	"time"
	"github.com/Ericles-Miller/SearchInfoIMDBsMovies/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
	}

	slog.Info("All system offline")
}

func run() error {

	db := make(map[string]api.User)	

	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout: 5 *
			time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
		IdleTimeout:  time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
