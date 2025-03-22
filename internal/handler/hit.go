package handler

import (
	"api/internal/service"
	"api/internal/storage"
	"encoding/json"
	"net/http"
)

func HandleGetHitCounter(w http.ResponseWriter, r *http.Request) {
	data := storage.GlobalDataStore.Get("hits")
	if data == nil {
		data = 0
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HandlePatchHitCounter(w http.ResponseWriter, r *http.Request) {
	data := service.IncrementHitCounter()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
