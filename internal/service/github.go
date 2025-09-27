package service

import (
	"api/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func GetGitHubData() []model.GitHubData {
	data := []model.GitHubData{}

	projects := strings.Split(os.Getenv("GITHUB_PROJECTS"), ",")
	client := &http.Client{}

	// Go through every project specified in GITHUB_PROJECTS
	for _, project := range projects {
		url := fmt.Sprintf("https://api.github.com/repos/%s", project)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			slog.Error("Error creating request", slog.Any("error", err))
			continue
		}

		// Add authorization header
		req.Header.Set("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))

		res, err := client.Do(req)
		if err != nil {
			slog.Error("Error requesting GitHub API", slog.Any("error", err))
			continue
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			slog.Error("Error reading body", slog.Any("error", err))
			continue
		}

		var apiResponse model.GitHubData
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			slog.Error("Error unmarshalling JSON", slog.Any("error", err))
			continue
		}

		data = append(data, model.GitHubData{
			Name:     apiResponse.Name,
			Stars:    apiResponse.Stars,
			Language: apiResponse.Language,
		})
	}

	return data
}
