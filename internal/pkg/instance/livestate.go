package instance

import (
	"sort"
	"strings"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
)

type ServerState string

const (
	ServerStateOffline       ServerState = "offline"
	ServerStateStarting      ServerState = "starting"
	ServerStateStoping       ServerState = "stoping"
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

func (ds *DriverState) ToEILDB() event.EventInstanceLiveDriverBase {
	return event.EventInstanceLiveDriverBase{
		Name:     ds.Name,
		PlayerID: ds.PlayerID,
	}
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
	S1MS        int          `json:"s1MS"`
	S2          string       `json:"s2"`
	S2MS        int          `json:"s2MS"`
	S3          string       `json:"s3"`
	S3MS        int          `json:"s3MS"`
	Fuel        int          `json:"fuel"`
	HasCut      bool         `json:"hasCut"`
	InLap       bool         `json:"inLap"`
	OutLap      bool         `json:"outLap"`
	SessionOver bool         `json:"sessionOver"`
}

func (l LapState) IsValid() bool {
	return l.Flags == 0 && !l.HasCut && !l.InLap && !l.OutLap && !l.SessionOver
}

type ServerChat struct {
	Timestamp time.Time `json:"ts"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
}

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
	History          []ServerHistory   `json:"history"`
	historyId        int32

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
		History:     []ServerHistory{},
	}
}

func (l *LiveState) SetServerState(s ServerState) {
	l.ServerState = s
}

func (l *LiveState) SetNrClients(nr int) {
	l.NrClients = nr
}

func (l *LiveState) SetTrack(t string) {
	l.Track = t
}

func (l *LiveState) SetSessionState(t, p string, r int) {
	oldType := l.SessionType
	oldPhase := l.SessionPhase
	l.SessionType = t
	l.SessionPhase = p

	if r >= 0 {
		l.SessionRemaining = r
	}

	if t != oldType || p != oldPhase {
		l.AddHistory("session", ServerHistorySessionChange{
			SessionType:      l.SessionType,
			SessionPhase:     l.SessionPhase,
			SessionRemaining: l.SessionRemaining,
		})
	}

	if t != oldType {
		l.AdvanceSession()
	}
}

func (l *LiveState) AddNewConnection(connID int, name, playerID string, carModel int) {
	l.connections[connID] = &DriverState{
		ConnectionID: connID,
		Name:         name,
		PlayerID:     playerID,
		carModel:     carModel,
	}
}

func (l *LiveState) AdvanceSession() {
	for _, car := range l.Cars {
		if len(car.Drivers) == 0 {
			l.PurgeCar(car.CarID)
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

func (l *LiveState) AddNewCar(carID, raceNumber, carModel int) {
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

func (l *LiveState) Handshake(carID, connectionID int) {
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

func (l *LiveState) RemoveConnection(id int) {
	d, ok := l.connections[id]
	if !ok {
		return
	}

	if d.car != nil {
		d.car.removeDriver(d)
	}

	delete(l.connections, id)
}

func (l *LiveState) PurgeCar(id int) {
	delete(l.Cars, id)
}

func (l *LiveState) ServerOffline() {
	l.SetServerState(ServerStateOffline)
	for _, car := range l.Cars {
		l.PurgeCar(car.CarID)
	}
	l.SetNrClients(0)
	l.SetTrack("")
	l.SetSessionState("", "", 0)
	l.connections = map[int]*DriverState{}
}

func (l *LiveState) SetCarPosition(carID, pos int) {
	if car, ok := l.Cars[carID]; ok {
		car.Position = pos
	}
}

func (l *LiveState) SetLapState(lap *LapState) {
	lap.Car.NrLaps++
	lap.Car.Fuel = lap.Fuel
	lap.Car.LastLapMS = lap.LapTimeMS
	lap.Car.LastLapTimestampMS = lap.TimestampMS

	if lap.IsValid() && (lap.Car.BestLapMS <= 0 || lap.LapTimeMS < lap.Car.BestLapMS) {
		lap.Car.BestLapMS = lap.LapTimeMS
	}

	lap.Car.Laps = append(lap.Car.Laps, lap)

	l.recalculatePositions()
}

func (l *LiveState) SetCurrLapState(lap LapState) {
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

func (l *LiveState) AddChat(name, message string) {
	// skip /admin message
	if len(message) > 6 && strings.ToLower(message[0:6]) == "/admin" {
		return
	}

	l.AddHistory("chat", ServerHistoryChat{
		Name:    name,
		Message: message,
	})

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

func (l *LiveState) AddHistory(t string, data any) {
	l.historyId++
	l.History = append(l.History, ServerHistory{
		ID:        l.historyId,
		Timestamp: time.Now().UTC(),
		Type:      t,
		Data:      data,
	})

	nrMsg := 200

	tt := len(l.History)

	if tt > nrMsg {
		l.History = l.History[tt-nrMsg : tt]
	}
}

func (l *LiveState) AddDamage(carId int) {
	car := l.GetCar(carId)
	if car == nil {
		return
	}

	l.AddHistory("damage", ServerHistoryDamage{
		CarID:      car.CarID,
		RaceNumber: car.RaceNumber,
		CarModel:   car.CarModel,
		Name:       car.CurrentDriver.Name,
		PlayerID:   car.CurrentDriver.PlayerID,
	})
}
