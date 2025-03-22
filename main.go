package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"

	"api/internal/server"
	"api/internal/storage"
	"api/internal/worker"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		slog.Warn("No .env file was found; using environment variables.", slog.Any("error", err))
	}

	storage.InitDataStore()
	go worker.StartWorkers()

	// Shutdown-chan~~
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChan

		if err := storage.GlobalDataStore.Save(); err != nil {
			slog.Error("Error saving data store on shutdown", slog.Any("error", err))
		} else {
			slog.Info("Data store saved successfully on shutdown")
		}

		os.Exit(0)
	}()

	server.NewRouter()
}
