package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"

	"api/internal/server"
	"api/internal/service"
	"api/internal/storage"
	"api/internal/worker"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		slog.Warn("No .env file was found; using environment variables", slog.Any("error", err))
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		storage.InitDataStore()
	}()

	wg.Wait()
	service.LoadComputerStatTotals()
	go worker.StartWorkers()

	// Shutdown-chan~~
	// this joke sucks
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

	server.StartPrometheusServer()
	server.NewRouter()
}
