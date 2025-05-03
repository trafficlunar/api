package storage

import (
	"encoding/json"
	"log/slog"
	"os"
	"sync"
)

type DataStore struct {
	Data  map[string]any
	Mutex sync.Mutex
}

var GlobalDataStore = &DataStore{
	Data: make(map[string]any),
}

func InitDataStore() *DataStore {
	GlobalDataStore = &DataStore{
		Data: make(map[string]any),
	}

	file, err := os.Open("./data/data.json")
	if err != nil {
		if os.IsNotExist(err) {
			slog.Warn("Data store file not found; creating new file")

			GlobalDataStore.Save()
			return GlobalDataStore
		}
		slog.Error("Could not load data store file!", slog.Any("error", err))
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&GlobalDataStore.Data); err != nil {
		slog.Error("Failed to decode data store file", slog.Any("error", err))
		return nil
	}

	slog.Info("Data store loaded")
	return GlobalDataStore
}

func (store *DataStore) Get(key string) any {
	store.Mutex.Lock()
	defer store.Mutex.Unlock()
	return store.Data[key]
}

func (store *DataStore) Set(key string, value any) {
	store.Mutex.Lock()
	defer store.Mutex.Unlock()
	store.Data[key] = value
}

func (store *DataStore) Save() error {
	store.Mutex.Lock()
	defer store.Mutex.Unlock()

	os.Mkdir("./data/", os.ModePerm)
	file, err := os.Create("./data/data.json")
	if err != nil {
		slog.Error("Could not create data store file", slog.Any("error", err))
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(store.Data); err != nil {
		slog.Error("Failed to encode data store", slog.Any("error", err))
		return err
	}

	return nil
}
