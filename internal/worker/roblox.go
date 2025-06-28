package worker

import (
	"api/internal/service"
	"log/slog"
	"time"
)

var GrowARobloxianLikes int

func StartRobloxWorker() {
	slog.Info("Starting Roblox worker...")
	GrowARobloxianLikes = service.GetRobloxGrowARobloxianLikesCount()

	for range time.Tick(1 * time.Minute) {
		GrowARobloxianLikes = service.GetRobloxGrowARobloxianLikesCount()
	}
}
