package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"
	"github.com/joho/godotenv"
	"github.com/Ericles-Miller/crudUserInMemory/api"
	"github.com/Ericles-Miller/crudUserInMemory/internal/database"

)

func main() {
	godotenv.Load()
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
	}

	slog.Info("All system offline")
}

func run() error {
	ctx := context.Background()

	pool, err := database.ConnectDB(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()

	if err := database.RunMigrations(pool); err != nil {
		return err
	}

	handler := api.NewHandler()

	s := http.Server{
		ReadTimeout:  5 * time.Second,
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
