package service

import (
	"api/internal/model"
	"api/internal/storage"
	"time"
)

var ComputerData model.ComputerData = model.ComputerData{
	Online:      false,
	UptimeStart: -1,
	Totals: model.ComputerTotals{
		Uptime: 0,
		Keys:   0,
		Clicks: 0,
	},
	Graph: initializeGraphData(),
}

func LoadComputerStatTotals() {
	uptimeData := storage.GlobalDataStore.Get("uptime")
	keysData := storage.GlobalDataStore.Get("keys")
	clicksData := storage.GlobalDataStore.Get("clicks")

	var uptime float64
	var keys float64
	var clicks float64

	if uptimeData != nil {
		uptime = uptimeData.(float64)
	}
	if keysData != nil {
		keys = keysData.(float64)
	}
	if clicksData != nil {
		clicks = clicksData.(float64)
	}

	ComputerData.Totals = model.ComputerTotals{
		Uptime: uptime,
		Keys:   keys,
		Clicks: clicks,
	}
}

func initializeGraphData() []model.ComputerGraphData {
	graphData := make([]model.ComputerGraphData, 60)

	for i := range 60 {
		graphData[i] = model.ComputerGraphData{
			Timestamp: time.Now().Truncate(1 * time.Minute).Add(time.Duration(-60+i) * time.Minute),
			Cpu:       0,
			Ram:       0,
			Keys:      0,
			Clicks:    0,
		}
	}

	return graphData
}

func AddComputerData(clientMessage model.ComputerWebSocketMessage) {
	ComputerData.Graph = append(ComputerData.Graph, model.ComputerGraphData{
		Timestamp: time.Now().Truncate(time.Minute).Add(-time.Minute),
		Cpu:       int(clientMessage.Cpu),
		Ram:       int(clientMessage.Ram),
		Keys:      int(clientMessage.Keys),
		Clicks:    int(clientMessage.Clicks),
	})

	if len(ComputerData.Graph) > 60 {
		ComputerData.Graph = ComputerData.Graph[1:]
	}
}
