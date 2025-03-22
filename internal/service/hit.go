package service

import (
	"api/internal/model"
	"api/internal/storage"
)

func IncrementHitCounter() model.Success {
	hitsData := storage.GlobalDataStore.Get("hits")

	var hits uint32
	if hitsData != nil {
		hits = hitsData.(uint32)
	}

	storage.GlobalDataStore.Set("hits", hits+1)

	return model.Success{
		Success: true,
	}
}
