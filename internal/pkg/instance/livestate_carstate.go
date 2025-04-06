package instance

import "github.com/assetto-corsa-web/accweb/internal/pkg/event"

// CarState represents the current state of a single car
type CarState struct {
	CarID              int            `json:"carID"`
	RaceNumber         int            `json:"raceNumber"`
	CarModel           int            `json:"carModel"`
	Drivers            []*DriverState `json:"drivers"`
	CurrentDriver      *DriverState   `json:"currentDriver"`
	Fuel               int            `json:"fuel"`
	Position           int            `json:"position"`
	NrLaps             int            `json:"nrLaps"`
	BestLapMS          int            `json:"bestLapMS"`
	LastLapMS          int            `json:"lastLapMS"`
	LastLapTimestampMS int            `json:"lastLapTimestampMS"`
	Laps               []*LapState    `json:"laps"`
	CurrLap            LapState       `json:"currLap"`
}

func (cs *CarState) ToEILCB() event.EventInstanceLiveCarBase {
	return event.EventInstanceLiveCarBase{
		CarID:      cs.CarID,
		RaceNumber: cs.RaceNumber,
		CarModel:   cs.CarModel,
	}
}

func (c *CarState) removeDriver(d *DriverState) {
	if c.CurrentDriver != nil && c.CurrentDriver.ConnectionID == d.ConnectionID {
		c.CurrentDriver = nil
	}

	k := -1
	for i, driver := range c.Drivers {
		if driver.ConnectionID == d.ConnectionID {
			k = i
			break
		}
	}

	if k == -1 {
		return
	}

	copy(c.Drivers[k:], c.Drivers[k+1:])
	c.Drivers = c.Drivers[:len(c.Drivers)-1]
}
