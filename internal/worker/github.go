package worker

import (
	"api/internal/model"
	"api/internal/service"
	"log/slog"
	"time"
)

var GitHubData []model.GitHubData

func StartGitHubWorker() {
	slog.Info("Starting GitHub worker...")
	GitHubData = service.GetGitHubData()

	for range time.Tick(24 * time.Hour) {
		GitHubData = service.GetGitHubData()
	}
}
