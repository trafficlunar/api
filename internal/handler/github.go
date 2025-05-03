package handler

import (
	"api/internal/worker"
	"encoding/json"
	"net/http"
)

func HandleGetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worker.GitHubData)
}
