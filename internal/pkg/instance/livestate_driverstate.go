package instance

import "github.com/assetto-corsa-web/accweb/internal/pkg/event"

// DriverState contains the information about a single driver
type DriverState struct {
	ConnectionID int    `json:"-"`
	Name         string `json:"name"`
	PlayerID     string `json:"playerID"`

	car      *CarState
	carModel int
}

func (ds *DriverState) ToEILDB() event.EventInstanceLiveDriverBase {
	if ds == nil {
		return event.EventInstanceLiveDriverBase{}
	}

	return event.EventInstanceLiveDriverBase{
		Name:     ds.Name,
		PlayerID: ds.PlayerID,
	}
}
