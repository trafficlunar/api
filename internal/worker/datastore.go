package worker

import (
	"api/internal/storage"
	"log/slog"
	"time"
)

func StartDataStoreWorker() {
	slog.Info("Starting data store worker...")

	for range time.Tick(1 * time.Minute) {
		if err := storage.GlobalDataStore.Save(); err != nil {
			slog.Error("Error saving data store", slog.Any("error", err))
		}
	}
}
