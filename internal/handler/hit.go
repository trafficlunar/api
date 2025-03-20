package handler

import (
	"api/internal/service"
	"encoding/json"
	"net/http"
)

func HandleGetHitCounter(w http.ResponseWriter, r *http.Request) {
	data := service.GetHitCounter()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HandlePatchHitCounter(w http.ResponseWriter, r *http.Request) {
	data := service.IncrementHitCounter()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
