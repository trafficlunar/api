package service

import (
	"api/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func GetRobloxGrowARobloxianLikesCount() int {
	url := fmt.Sprintf("https://games.roblox.com/v1/games/votes?universeIds=%s", os.Getenv("GROWAROBLOXIAN_UNIVERSE_ID"))
	res, err := http.Get(url)
	if err != nil {
		slog.Error("Error requesting Roblox votes API", slog.Any("error", err))
		return 0
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error reading body", slog.Any("error", err))
		return 0
	}

	var apiData model.RobloxGameVotesAPI
	err = json.Unmarshal(body, &apiData)
	if err != nil {
		slog.Error("Error unmarshalling JSON", slog.Any("error", err))
		return 0
	}

	if apiData.Data == nil {
		slog.Warn("No data returned from Roblox votes API")
		return 0
	}

	return apiData.Data[0].Upvotes
}
