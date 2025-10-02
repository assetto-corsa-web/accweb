package instance

import "time"

type ServerHistory struct {
	ID        int32     `json:"id"`
	Timestamp time.Time `json:"ts"`
	Type      string    `json:"type"`
	Data      any       `json:"data"`
}

type ServerHistoryChat struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type ServerHistoryDamage struct {
	CarID      int    `json:"carID"`
	RaceNumber int    `json:"raceNumber"`
	CarModel   int    `json:"carModel"`
	Name       string `json:"name"`
	PlayerID   string `json:"playerID"`
}

type ServerHistorySessionChange struct {
	SessionType      string `json:"sessionType"`
	SessionPhase     string `json:"sessionPhase"`
	SessionRemaining int    `json:"sessionRemaining"`
}

type ServerHistoryConnection struct {
	Name     string `json:"name"`
	PlayerID string `json:"playerID"`
}
