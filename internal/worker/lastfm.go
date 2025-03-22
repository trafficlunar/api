package worker

import (
	"api/internal/model"
	"api/internal/service"
	"log/slog"
	"time"
)

var LastFMData model.LastFMData

func StartLastFMWorker() {
	slog.Info("Starting last.fm worker...")
	LastFMData = service.GetLastFMData()

	for range time.Tick(30 * time.Second) {
		LastFMData = service.GetLastFMData()
	}
}
