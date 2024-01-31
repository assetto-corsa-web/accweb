package event

type EventInstanceLiveNewDriver struct {
	EventBase
	EventInstanceBase

	Name     string `json:"name"`
	PlayerID string `json:"playerID"`

	CarID      int `json:"carID"`
	RaceNumber int `json:"raceNumber"`
	CarModel   int `json:"carModel"`
}

func EmmitEventInstanceLiveNewDriver(eib EventInstanceBase, name, pId string, cId, rn, cm int) {
	Emmit(EventInstanceLiveNewDriver{
		EventBase:         eventBase("instance_live_new_driver"),
		EventInstanceBase: eib,
		Name:              name,
		PlayerID:          pId,
		CarID:             cId,
		RaceNumber:        rn,
		CarModel:          cm,
	})
}

type EventInstanceLiveRemoveConnection struct {
	EventBase
	EventInstanceBase

	Name     string `json:"name"`
	PlayerID string `json:"playerID"`
}

func EmmitEventInstanceLiveRemoveConnection(eib EventInstanceBase, name, pId string) {
	Emmit(EventInstanceLiveRemoveConnection{
		EventBase:         eventBase("instance_live_remove_connection"),
		EventInstanceBase: eib,
		Name:              name,
		PlayerID:          pId,
	})
}
