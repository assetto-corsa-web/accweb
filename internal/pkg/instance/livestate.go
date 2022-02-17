package instance

import "sort"

type ServerState string

const (
	ServerStateOffline       ServerState = "offline"
	ServerStateStarting      ServerState = "starting"
	ServerStateNotRegistered ServerState = "not_registered"
	ServerStateOnline        ServerState = "online"
)

// DriverState contains the information about a single driver
type DriverState struct {
	ConnectionID int    `json:"connectionID"`
	Name         string `json:"name"`
	PlayerID     string `json:"playerID"`

	car      *CarState
	carModel int
}

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
}

func (c *CarState) removeDriver(d *DriverState) {
	if c.CurrentDriver != nil && c.CurrentDriver.PlayerID == d.PlayerID {
		c.CurrentDriver = nil
	}

	k := -1
	for i, driver := range c.Drivers {
		if driver.PlayerID == d.PlayerID {
			k = i
			break
		}
	}

	if k == -1 {
		return
	}

	//copy(c.Drivers[k:], c.Drivers[:k+1])
	copy(c.Drivers[k:], c.Drivers[k+1:])
	c.Drivers = c.Drivers[:len(c.Drivers)-1]
}

type LapState struct {
	CarID       int
	DriverIndex int
	Car         *CarState
	Driver      *DriverState
	LapTimeMS   int
	TimestampMS int
	Flags       int
	S1          string
	S2          string
	S3          string
	Fuel        int
	HasCut      bool
	InLap       bool
	OutLap      bool
	SessionOver bool
}

type LiveState struct {
	ServerState  ServerState       `json:"serverState"`
	NrClients    int               `json:"nrClients"`
	Track        string            `json:"track"`
	SessionType  string            `json:"sessionType"`
	SessionPhase string            `json:"sessionPhase"`
	Cars         map[int]*CarState `json:"cars"`

	// drivers waiting to be assigned to a car, key: ConnectionID
	connections map[int]*DriverState
}

func newLiveState() *LiveState {
	return &LiveState{
		ServerState: ServerStateOffline,
		Cars:        map[int]*CarState{},
		connections: map[int]*DriverState{},
	}
}

func (l *LiveState) setServerState(s ServerState) {
	l.ServerState = s
}

func (l *LiveState) setNrClients(nr int) {
	l.NrClients = nr
}

func (l *LiveState) setTrack(t string) {
	l.Track = t
}

func (l *LiveState) setSessionState(t, p string) {
	l.SessionType = t
	l.SessionPhase = p
}

func (l *LiveState) addNewConnection(connID int, name, playerID string, carModel int) {
	l.connections[connID] = &DriverState{
		ConnectionID: connID,
		Name:         name,
		PlayerID:     playerID,
		carModel:     carModel,
	}
}

func (l *LiveState) advanceSession() {
	for _, car := range l.Cars {
		if len(car.Drivers) == 0 {
			l.purgeCar(car.CarID)
		} else {
			car.Fuel = 0
			car.NrLaps = 0
			car.BestLapMS = 0
			car.LastLapMS = 0
			car.LastLapTimestampMS = 0
		}
	}
	l.recalculatePositions()
}

func (l *LiveState) addNewCar(carID, raceNumber, carModel int) {
	c := CarState{
		CarID:      carID,
		RaceNumber: raceNumber,
		CarModel:   carModel,
	}

	car := l.Cars[carID]

	if car == nil {
		car = &CarState{
			CarID:    carID,
			Position: len(l.Cars) + 1,
			Drivers:  []*DriverState{},
		}
	}

	car.CarModel = c.CarModel
	car.RaceNumber = c.RaceNumber

	for _, d := range l.connections {
		if d.car != nil {
			continue
		}

		if d.carModel != c.CarModel {
			continue
		}

		d.car = car
		car.Drivers = append(car.Drivers, d)

		if car.CurrentDriver == nil {
			car.CurrentDriver = d
		}
		break
	}

	l.Cars[carID] = car
}

func (l *LiveState) removeConnection(id int) {
	d, ok := l.connections[id]
	if !ok {
		return
	}

	if d.car != nil {
		d.car.removeDriver(d)
	}

	delete(l.connections, id)
}

func (l *LiveState) purgeCar(id int) {
	delete(l.Cars, id)
}

func (l *LiveState) serverOffline() {
	l.setServerState(ServerStateOffline)
	for _, car := range l.Cars {
		l.purgeCar(car.CarID)
	}
	l.setNrClients(0)
	l.setTrack("")
	l.setSessionState("", "")
	l.connections = map[int]*DriverState{}
}

func (l *LiveState) setCarPosition(carID, pos int) {
	if car, ok := l.Cars[carID]; ok {
		car.Position = pos
	}
}

func (l *LiveState) setLapState(lap *LapState) {
	lap.Car.NrLaps++
	lap.Car.Fuel = lap.Fuel
	lap.Car.LastLapMS = lap.LapTimeMS
	lap.Car.LastLapTimestampMS = lap.TimestampMS

	if lap.Flags == 8808693760 && (lap.Car.BestLapMS <= 0 || lap.LapTimeMS < lap.Car.BestLapMS) {
		lap.Car.BestLapMS = lap.LapTimeMS
	}

	l.recalculatePositions()
}

func cmpPositionFastestLap(a, b *CarState) bool {
	if a.BestLapMS > 0 {
		if b.BestLapMS > 0 { // Both a and b have a lap
			return a.BestLapMS < b.BestLapMS
		} else { // Only a has a lap
			return true
		}
	} else {
		if b.BestLapMS > 0 { // Only b has a lap
			return false
		} else { // Neither a nor b has a lap
			return a.Position < b.Position
		}
	}
}

func cmpPositionMostDistance(a, b *CarState) bool {
	if a.NrLaps != b.NrLaps {
		return a.NrLaps > b.NrLaps
	}
	if a.LastLapTimestampMS != b.LastLapTimestampMS {
		return a.LastLapTimestampMS < b.LastLapTimestampMS
	}
	return a.Position < b.Position
}

func (l *LiveState) recalculatePositions() {
	cars := make([]*CarState, 0, len(l.Cars))
	for _, car := range l.Cars {
		cars = append(cars, car)
	}

	sort.Slice(cars, func(i, j int) bool {
		if l.SessionType == "Race" {
			return cmpPositionMostDistance(cars[i], cars[j])
		}
		return cmpPositionFastestLap(cars[i], cars[j])
	})

	for i := 0; i < len(cars); i++ {
		if cars[i].Position != i+1 {
			cars[i].Position = i + 1
		}
	}
}
