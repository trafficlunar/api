package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"api/internal/handler"
)

func getAllowedOrigins() []string {
	allowedOrigins := []string{"https://trafficlunar.net", "https://trafficlunar.nekoweb.org/"}

	if os.Getenv("DEVELOPMENT_MODE") == "true" {
		allowedOrigins = append(allowedOrigins, "http://localhost:4321")
	}

	return allowedOrigins
}

func NewRouter() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "application/json"))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httprate.LimitByRealIP(32, time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: getAllowedOrigins(),
		AllowedMethods: []string{"GET", "PATCH"},
		AllowedHeaders: []string{"Content-Type"},
		MaxAge:         300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"name": "trafficlunar's api",
			"url":  "https://github.com/trafficlunar/api",
		})
	})

	r.Get("/hit", handler.HandleGetHitCounter)
	r.With(httprate.LimitByRealIP(1, time.Hour)).Patch("/hit", handler.HandlePatchHitCounter)
	r.Get("/song", handler.HandleGetCurrentlyPlaying)
	r.Get("/projects", handler.HandleGetProjects)
	r.Get("/computer", handler.HandleComputerGraphData)
	r.Get("/computer/ws", handler.HandleComputerWebSocket)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8888"
	}

	slog.Info("Starting server", slog.Any("port", port))
	http.ListenAndServe(":"+port, r)
}
