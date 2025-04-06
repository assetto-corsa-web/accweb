package instance

import (
	"sync"

	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
)

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

	drvLock sync.RWMutex
}

func (c *CarState) ToEILCB() event.EventInstanceLiveCarBase {
	return event.EventInstanceLiveCarBase{
		CarID:      c.CarID,
		RaceNumber: c.RaceNumber,
		CarModel:   c.CarModel,
	}
}

func (c *CarState) addDriver(d *DriverState) {
	c.drvLock.Lock()
	defer c.drvLock.Unlock()

	if d == nil {
		return
	}

	c.Drivers = append(c.Drivers, d)
	d.car = c

	if c.CurrentDriver == nil {
		c.CurrentDriver = d
	}
}

func (c *CarState) removeDriver(d *DriverState) {
	c.drvLock.Lock()
	defer c.drvLock.Unlock()

	if d == nil {
		return
	}

	if c.CurrentDriver != nil && c.CurrentDriver.ConnectionID == d.ConnectionID {
		c.CurrentDriver = nil
	}

	dd := []*DriverState{}

	if len(c.Drivers) > 1 {
		for _, driver := range c.Drivers {
			if driver.ConnectionID == d.ConnectionID {
				continue
			}
			dd = append(dd, driver)
		}
	}

	c.Drivers = dd
}

func (c *CarState) LenDrivers() int {
	c.drvLock.RLock()
	defer c.drvLock.RUnlock()

	return len(c.Drivers)
}

func (c *CarState) RangeDrivers(f func(i int, d *DriverState) bool) {
	c.drvLock.RLock()
	defer c.drvLock.RUnlock()

	for i, driver := range c.Drivers {
		if ok := f(i, driver); !ok {
			break
		}
	}
}
