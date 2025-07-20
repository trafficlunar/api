package model

import "time"

type ComputerWebSocketMessage struct {
	Cpu    uint8  `json:"cpu"`
	Ram    uint8  `json:"ram"`
	Keys   uint16 `json:"keys"`
	Clicks uint16 `json:"clicks"`
}

type ComputerData struct {
	Online      bool                `json:"online"`
	UptimeStart int                 `json:"uptimeStart"`
	Totals      ComputerTotals      `json:"totals"`
	Graph       []ComputerGraphData `json:"graph"`
}

type ComputerTotals struct {
	Keys   float64 `json:"keys"`
	Clicks float64 `json:"clicks"`
}

type ComputerGraphData struct {
	Timestamp time.Time `json:"timestamp"`
	Cpu       int       `json:"cpu"`
	Ram       int       `json:"ram"`
	Keys      int       `json:"keys"`
	Clicks    int       `json:"clicks"`
}
