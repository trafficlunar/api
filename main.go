package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"

	"api/internal/server"
	"api/internal/worker"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		slog.Warn("No .env file was found; using environment variables.", slog.Any("error", err))
	}

	go worker.StartWorkers()
	server.NewRouter()
}
