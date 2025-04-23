package instance

import (
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
	"github.com/sirupsen/logrus"
)

type ServerState string

const (
	ServerStateOffline       ServerState = "offline"
	ServerStateStarting      ServerState = "starting"
	ServerStateStoping       ServerState = "stoping"
	ServerStateNotRegistered ServerState = "not_registered"
	ServerStateOnline        ServerState = "online"
)

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
	History          []ServerHistory   `json:"history"`
	historyId        int32

	// drivers waiting to be assigned to a car, key: ConnectionID
	connections map[int]*DriverState

	carsLock sync.RWMutex
	connLock sync.RWMutex
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

func (l *LiveState) GetCar(cId int) *CarState {
	l.carsLock.RLock()
	defer l.carsLock.RUnlock()

	if c, ok := l.Cars[cId]; ok {
		return c
	}

	logrus.WithFields(logrus.Fields{
		"carID": cId,
	}).Error("car not found in GetCar")

	return nil
}

func (l *LiveState) GetDriver(connId int) *DriverState {
	l.connLock.RLock()
	defer l.connLock.RUnlock()

	if c, ok := l.connections[connId]; ok {
		return c
	}

	logrus.WithFields(logrus.Fields{
		"connID": connId,
	}).Error("driver not found in GetDriver")

	return nil
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

func (l *LiveState) SetSession(new string) {
	l.SessionType = new
	l.AdvanceSession()
}

func (l *LiveState) SetSessionState(p string, r int) {
	oldPhase := l.SessionPhase
	l.SessionPhase = p

	if r >= 0 {
		l.SessionRemaining = r
	}

	if p != oldPhase {
		l.AddHistory("session", ServerHistorySessionChange{
			SessionType:      l.SessionType,
			SessionPhase:     l.SessionPhase,
			SessionRemaining: l.SessionRemaining,
		})
	}
}

func (l *LiveState) AddNewConnection(connID int, name, playerID string, carModel int) {
	l.connLock.Lock()
	defer l.connLock.Unlock()

	l.connections[connID] = &DriverState{
		ConnectionID: connID,
		Name:         name,
		PlayerID:     playerID,
		carModel:     carModel,
	}
}

func (l *LiveState) AdvanceSession() {
	l.carsLock.Lock()
	defer l.carsLock.Unlock()

	for _, car := range l.Cars {
		if car.LenDrivers() == 0 {
			delete(l.Cars, car.CarID)
			continue
		}

		car.Fuel = 0
		car.NrLaps = 0
		car.BestLapMS = 0
		car.LastLapMS = 0
		car.LastLapTimestampMS = 0
		car.Laps = []*LapState{}
		car.CurrLap = LapState{}
	}

	l.recalculatePositions()
}

func (l *LiveState) AddNewCar(carID, raceNumber, carModel int) {
	l.carsLock.Lock()
	defer l.carsLock.Unlock()

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

func (l *LiveState) Handshake(carID, connId int) {
	l.carsLock.Lock()
	defer l.carsLock.Unlock()

	l.connLock.Lock()
	defer l.connLock.Unlock()

	d := l.connections[connId]
	if d == nil {
		logrus.WithFields(logrus.Fields{
			"connID": connId,
			"carID":  carID,
		}).Error("connection not found in Handshake")
		return
	}

	car := l.Cars[carID]
	if car == nil {
		logrus.WithFields(logrus.Fields{
			"connID": connId,
			"carID":  carID,
		}).Error("car not found in Handshake")
		return
	}

	d.car = car
	car.addDriver(d)
}

func (l *LiveState) RemoveConnection(connId int) {
	l.connLock.Lock()
	defer l.connLock.Unlock()

	d, ok := l.connections[connId]
	if !ok {
		logrus.WithFields(logrus.Fields{
			"connID": connId,
		}).Error("connection not found in RemoveConnection")
		return
	}

	if d.car != nil {
		d.car.removeDriver(d)
	}

	delete(l.connections, connId)
}

func (l *LiveState) PurgeCar(id int) {
	l.carsLock.Lock()
	defer l.carsLock.Unlock()

	delete(l.Cars, id)
}

func (l *LiveState) ServerOffline() {
	l.connLock.Lock()
	defer l.connLock.Unlock()

	l.SetServerState(ServerStateOffline)
	for _, car := range l.Cars {
		l.PurgeCar(car.CarID)
	}
	l.SetNrClients(0)
	l.SetTrack("")
	l.SetSessionState("", 0)
	l.connections = map[int]*DriverState{}
}

func (l *LiveState) SetCarPosition(carID, pos int) {
	l.carsLock.Lock()
	defer l.carsLock.Unlock()

	if car, ok := l.Cars[carID]; ok {
		car.Position = pos
		return
	}

	logrus.WithFields(logrus.Fields{
		"carID": carID,
		"pos":   pos,
	}).Error("car not found in SetCarPosition")
}

func (l *LiveState) SetLapState(lap *LapState) {
	l.carsLock.Lock()
	defer l.carsLock.Unlock()

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
	l.carsLock.Lock()
	defer l.carsLock.Unlock()

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
		logrus.WithFields(logrus.Fields{
			"carID": carId,
		}).Error("car not found in AddDamage")
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

func (l *LiveState) ToEIC() map[int]event.EventCarState {
	l.carsLock.RLock()
	defer l.carsLock.RUnlock()

	cars := map[int]event.EventCarState{}

	for k, c := range l.Cars {
		dd := make([]event.EventInstanceLiveDriverBase, c.LenDrivers())

		c.RangeDrivers(func(i int, d *DriverState) bool {
			dd[i] = d.ToEILDB()
			return true
		})

		ll := make([]event.EventLapState, len(c.Laps))
		for i, l := range c.Laps {
			ll[i] = l.ToEILS()
		}

		cars[k] = event.EventCarState{
			RaceNumber:         c.RaceNumber,
			CarModel:           c.CarModel,
			Drivers:            dd,
			CurrentDriver:      c.CurrentDriver.ToEILDB(),
			Fuel:               c.Fuel,
			Position:           c.Position,
			NrLaps:             c.NrLaps,
			BestLapMS:          c.BestLapMS,
			LastLapMS:          c.LastLapMS,
			LastLapTimestampMS: c.LastLapTimestampMS,
			Laps:               ll,
			CurrLap:            c.CurrLap.ToEILS(),
		}
	}

	return cars
}
