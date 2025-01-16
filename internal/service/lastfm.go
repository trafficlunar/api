package service

import (
	"backend/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func GetLastFMData() model.LastFMData {
	data := model.LastFMData{
		Song:    "api error",
		Artist:  "???",
		Image:   "/missing.webp",
		Url:     "https://www.last.fm/user/" + os.Getenv("LASTFM_USERNAME"),
		Playing: false,
	}

	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=1", os.Getenv("LASTFM_USERNAME"), os.Getenv("LASTFM_API_KEY"))
	res, err := http.Get(url)
	if err != nil {
		slog.Error("Error requesting last.fm API", slog.Any("error", err))
		return data
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error reading body", slog.Any("error", err))
		return data
	}

	var lastfmJSON model.LastFMAPI

	err = json.Unmarshal(body, &lastfmJSON)
	if err != nil {
		slog.Error("Error unmarshalling JSON", slog.Any("error", err))
		return data
	}

	lastfmData := lastfmJSON.RecentTracks.TrackList[0]

	if lastfmData.Attributes != nil {
		data.Playing = true
	}

	data.Song = lastfmData.Name
	data.Artist = lastfmData.Artist.Text
	data.Image = lastfmData.Image[2].Text
	data.Url = lastfmData.Url

	return data
}
