package service

import (
	"api/internal/model"
	"time"
)

var ComputerData model.ComputerData = model.ComputerData{
	Online: false,
	Graph:  initializeGraphData(),
}

func initializeGraphData() []model.ComputerGraphData {
	graphData := make([]model.ComputerGraphData, 60)

	for i := 0; i < 60; i++ {
		graphData[i] = model.ComputerGraphData{
			Timestamp: time.Now().Truncate(1 * time.Minute).Add(time.Duration(-60+i) * time.Minute),
			Cpu:       0,
			Ram:       0,
		}
	}

	return graphData
}

func AddComputerData(clientMessage model.ComputerWebSocketMessage) {
	ComputerData.Graph = append(ComputerData.Graph, model.ComputerGraphData{
		Timestamp: time.Now().Truncate(time.Minute).Add(-time.Minute),
		Cpu:       int(clientMessage.Cpu),
		Ram:       int(clientMessage.Ram),
	})

	if len(ComputerData.Graph) > 60 {
		ComputerData.Graph = ComputerData.Graph[1:]
	}
}
