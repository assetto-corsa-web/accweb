package instance

import (
	"sort"
	"time"
)

type ServerState string

const (
	ServerStateOffline       ServerState = "offline"
	ServerStateStarting      ServerState = "starting"
	ServerStateNotRegistered ServerState = "not_registered"
	ServerStateOnline        ServerState = "online"
)

// DriverState contains the information about a single driver
type DriverState struct {
	ConnectionID int    `json:"-"`
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
	Laps               []*LapState    `json:"laps"`
	CurrLap            LapState       `json:"currLap"`
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
	CarID       int          `json:"carID"`
	DriverIndex int          `json:"driverIndex"`
	Car         *CarState    `json:"-"`
	Driver      *DriverState `json:"-"`
	LapTimeMS   int          `json:"lapTimeMS"`
	TimestampMS int          `json:"timestampMS"`
	Flags       int          `json:"flags"`
	S1          string       `json:"s1"`
	S2          string       `json:"s2"`
	S3          string       `json:"s3"`
	Fuel        int          `json:"fuel"`
	HasCut      bool         `json:"hasCut"`
	InLap       bool         `json:"inLap"`
	OutLap      bool         `json:"outLap"`
	SessionOver bool         `json:"sessionOver"`
}

type ServerChat struct {
	Timestamp time.Time `json:"ts"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
}

type LiveState struct {
	ServerState      ServerState       `json:"serverState"`
	NrClients        int               `json:"nrClients"`
	Track            string            `json:"track"`
	SessionType      string            `json:"sessionType"`
	SessionPhase     string            `json:"sessionPhase"`
	SessionRemaining int               `json:"sessionRemaining"`
	Cars             map[int]*CarState `json:"cars"`
	UpdatedAt        time.Time         `json:"updatedAt"`
	Chats            []ServerChat      `json:"chats"`

	// drivers waiting to be assigned to a car, key: ConnectionID
	connections map[int]*DriverState
}

func NewLiveState() *LiveState {
	return &LiveState{
		ServerState: ServerStateOffline,
		Cars:        map[int]*CarState{},
		connections: map[int]*DriverState{},
		UpdatedAt:   time.Now().UTC(),
		Chats:       []ServerChat{},
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

func (l *LiveState) setSessionState(t, p string, r int) {
	oldType := l.SessionType
	l.SessionType = t
	l.SessionPhase = p
	l.SessionRemaining = r

	if t != oldType {
		l.advanceSession()
	}
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
			car.Laps = []*LapState{}
			car.CurrLap = LapState{}
		}
	}
	l.recalculatePositions()
}

func (l *LiveState) addNewCar(carID, raceNumber, carModel int) {
	car := l.Cars[carID]

	if car == nil {
		car = &CarState{
			CarID:    carID,
			Position: len(l.Cars) + 1,
			Drivers:  []*DriverState{},
			Laps:     []*LapState{},
		}

		l.Cars[carID] = car
	}

	car.CarModel = carModel
	car.RaceNumber = raceNumber
}

func (l *LiveState) handshake(carID, connectionID int) {
	d := l.connections[connectionID]
	if d == nil {
		return
	}

	car := l.Cars[carID]
	if car == nil {
		return
	}

	d.car = car
	car.Drivers = append(car.Drivers, d)

	if car.CurrentDriver == nil {
		car.CurrentDriver = d
	}
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
	l.setSessionState("", "", 0)
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

	if lap.Flags == 0 && (lap.Car.BestLapMS <= 0 || lap.LapTimeMS < lap.Car.BestLapMS) {
		lap.Car.BestLapMS = lap.LapTimeMS
	}

	lap.Car.Laps = append(lap.Car.Laps, lap)

	l.recalculatePositions()
}

func (l *LiveState) setCurrLapState(lap LapState) {
	lap.Car.LastLapTimestampMS = lap.TimestampMS
	lap.Car.CurrLap = lap
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

func isEmpty(v string) bool {
	return v == ""
}

func cmpPositionMostDistance(a, b *CarState) bool {
	if a.NrLaps != b.NrLaps {
		return a.NrLaps > b.NrLaps
	}

	if isEmpty(a.CurrLap.S3) != isEmpty(b.CurrLap.S3) {
		return isEmpty(a.CurrLap.S3)
	}

	if isEmpty(a.CurrLap.S2) != isEmpty(b.CurrLap.S2) {
		return !isEmpty(a.CurrLap.S2)
	}

	if isEmpty(a.CurrLap.S1) != isEmpty(b.CurrLap.S1) {
		return !isEmpty(a.CurrLap.S1)
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

func (l *LiveState) addChat(name, message string) {
	l.Chats = append(l.Chats, ServerChat{
		Timestamp: time.Now().UTC(),
		Name:      name,
		Message:   message,
	})

	nrMsg := 30

	t := len(l.Chats)

	if t > nrMsg {
		l.Chats = l.Chats[t-nrMsg : t]
	}
}
