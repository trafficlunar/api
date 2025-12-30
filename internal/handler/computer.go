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
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// If computer disconnected and reconnected within 45 seconds,
// continue current session (keep uptime at the same value)
var lastSeen time.Time
var computerDataMutex sync.RWMutex
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// todo: change for security
		return true
	},
}

func HandleComputerWebSocket(w http.ResponseWriter, r *http.Request) {
	// Check if user is authorized
	if r.Header.Get("Authorization") != os.Getenv("WEBSOCKET_PASSWORD") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// Upgrade to WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error when upgrading WebSocket connection", slog.Any("error", err))
		return
	}
	defer conn.Close()

	slog.Info("WebSocket connection established")

	// Respond to client pings
	conn.SetPingHandler(func(appData string) error {
		slog.Info("Recieved ping from client")
		return conn.WriteControl(
			websocket.PongMessage,
			[]byte(appData),                // echo back the same data
			time.Now().Add(10*time.Second), // deadline
		)
	})

	computerDataMutex.Lock()
	service.LoadComputerStatTotals()

	service.ComputerData.Online = true

	// If computer has not been seen in the last 45s, start new session
	if time.Since(lastSeen) > 45*time.Second {
		service.ComputerData.UptimeStart = int(time.Now().Unix())
	} else {
		slog.Info("Continuing existing uptime session")
	}
	computerDataMutex.Unlock()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("WebSocket connection closed", slog.Any("error", err))

			computerDataMutex.Lock()
			service.ComputerData.Online = false

			// Calculate uptime only if we have a valid start time
			if service.ComputerData.UptimeStart > 0 {
				// Calculate uptime
				sessionUptime := int(time.Now().Unix()) - service.ComputerData.UptimeStart

				// Get current total uptime from data store
				totalUptimeData := storage.GlobalDataStore.Get("uptime")
				var totalUptime float64
				if totalUptimeData != nil {
					totalUptime = totalUptimeData.(float64)
				}

				// Add to totals
				service.ComputerData.Totals.Uptime = totalUptime + float64(sessionUptime)
				storage.GlobalDataStore.Set("uptime", service.ComputerData.Totals.Uptime)
			}
			computerDataMutex.Unlock()

			// Check after 45s that computer is still offline, if so set to -1 (computer is offline)
			time.AfterFunc(45*time.Second, func() {
				computerDataMutex.Lock()
				defer computerDataMutex.Unlock()

				if !service.ComputerData.Online {
					service.ComputerData.UptimeStart = -1
				}
			})

			lastSeen = time.Now()
			break
		}

		var clientMessage model.ComputerWebSocketMessage
		if err := json.Unmarshal(message, &clientMessage); err != nil {
			slog.Error("Error unmarshalling JSON", slog.Any("error", err))
			continue
		}

		worker.QueuedClientMessage = clientMessage
		slog.Info("Received message", slog.Any("message", clientMessage))

		// Add to totals
		keysData := storage.GlobalDataStore.Get("keys")
		clicksData := storage.GlobalDataStore.Get("clicks")

		var keys float64
		var clicks float64

		// Convert values (if any) to float64
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

	computerDataMutex.RLock()
	defer computerDataMutex.RUnlock()

	json.NewEncoder(w).Encode(service.ComputerData)
}
