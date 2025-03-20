package worker

import (
	"api/internal/model"
	"api/internal/service"
	"log/slog"
	"time"
)

var QueuedClientMessage model.ComputerWebSocketMessage

func StartComputerWorker() {
	slog.Info("Starting computer worker...")

	for range time.Tick(1 * time.Second) {
		now := time.Now()

		if now.Second() == 0 {
			if !service.ComputerData.Online {
				service.AddComputerData(model.ComputerWebSocketMessage{
					Cpu: 0,
					Ram: 0,
				})
			} else {
				service.AddComputerData(QueuedClientMessage)
			}
		}
	}
}
