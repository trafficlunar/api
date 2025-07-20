package handler

import (
	"api/internal/model"
	"api/internal/service"
	"api/internal/storage"
	"api/internal/worker"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// todo: change for security
		return true
	},
}

func HandleComputerWebSocket(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != os.Getenv("WEBSOCKET_PASSWORD") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error when upgrading websocket connection", slog.Any("error", err))
		return
	}
	defer conn.Close()

	slog.Info("WebSocket connection established")
	service.ComputerData.Online = true
	service.ComputerData.UptimeStart = int(time.Now().Unix())

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("WebSocket connection closed by client", slog.Any("error", err))
			service.ComputerData.Online = false
			service.ComputerData.UptimeStart = 0
			break
		}

		var clientMessage model.ComputerWebSocketMessage
		if err := json.Unmarshal(message, &clientMessage); err != nil {
			slog.Error("Error unmarshalling JSON", slog.Any("error", err))
			continue
		}

		worker.QueuedClientMessage = clientMessage
		slog.Info("Recieved message", slog.Any("message", clientMessage))

		// Add to totals
		keysData := storage.GlobalDataStore.Get("keys")
		clicksData := storage.GlobalDataStore.Get("clicks")

		var keys float64
		var clicks float64

		if keysData != nil {
			keys = keysData.(float64)
		}
		if clicksData != nil {
			clicks = clicksData.(float64)
		}

		storage.GlobalDataStore.Set("keys", keys+float64(clientMessage.Keys))
		storage.GlobalDataStore.Set("clicks", clicks+float64(clientMessage.Clicks))

		service.LoadComputerStatTotals()
	}
}

func HandleComputerGraphData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.ComputerData)
}
