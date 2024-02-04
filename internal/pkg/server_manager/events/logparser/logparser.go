package logparser

import (
	"regexp"
	"strconv"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
	"github.com/assetto-corsa-web/accweb/internal/pkg/instance"
	"github.com/sirupsen/logrus"
)

type matcherHandler func(*instance.Instance, *instance.LiveState, []string)

type logMatcher struct {
	er      *regexp.Regexp
	handler matcherHandler
}

func newLogMatcher(er string, h matcherHandler) *logMatcher {
	return &logMatcher{
		er:      regexp.MustCompile(er),
		handler: h,
	}
}

type logParser struct {
	matchers []*logMatcher
}

func newLogParser() *logParser {
	return &logParser{
		matchers: makeLogMatchers(),
	}
}

func (l *logParser) processLine(i *instance.Instance, line string) {
	for _, matcher := range l.matchers {
		matches := matcher.er.FindStringSubmatch(line)
		if matches != nil {
			// logrus.WithField("line", matches[0]).WithField("attr", matches[1:]).Debug("log handled")
			matcher.handler(i, i.Live, matches)
			i.Live.UpdatedAt = time.Now().UTC()
		}
	}
}

func makeLogMatchers() []*logMatcher {
	return []*logMatcher{
		newLogMatcher(`^Server starting with version (\d+)$`, handleServerStarting),
		newLogMatcher(`RegisterToLobby TCP connection failed`, handleLobbyConnectionFailed),
		newLogMatcher(`RegisterToLobby succeeded`, handleLobbyConnectionSucceeded),
		newLogMatcher(`^(\d+) client\(s\) online$`, handleNrClientsOnline),
		newLogMatcher(`^Track (\w+) was set and updated$`, handleTrack),
		newLogMatcher(`^Detected sessionPhase <([A-Za-z ]+)> -> <([A-Za-z ]+)> \(([A-Za-z ]+)\)$`, handleSessionPhaseChanged),
		newLogMatcher(`^Resetting race weekend$`, handleResettingRace),
		newLogMatcher(`^New connection request: id (\d+) (.+) (S\d+) on car model (\d+)$`, handleNewConnection),
		newLogMatcher(`^Creating new car connection: carId (\d+), carModel (\d+), raceNumber #(\d+)$`, handleNewCar),
		newLogMatcher(`^Sent handshake response for car (\d+) connection (\d+) with`, handleHandshake),
		newLogMatcher(`Removing dead connection (\d+)`, handleDeadConnection),
		newLogMatcher(`^Purging car_id (\d+)$`, handleCarPurge),
		newLogMatcher(`Lap carId (\d+), driverId (\d+), lapTime (\d+):(\d+):(\d+), timestampMS (\d+)\.\d+, flags: (.*?)(, S1 (\d+:\d+:\d+))(, S2 (\d+:\d+:\d+))(, S3 (\d+:\d+:\d+)), fuel (\d+)\.\d+(, hasCut )?(, InLap )?(, OutLap )?(, SessionOver)?`, handleLap),
		newLogMatcher(`Lap  ?carId (\d+), driverId (\d+), lapTime (35791):(23):(647), timestampMS (\d+)\.\d+, flags: (.*?)(, S1 (\d+:\d+:\d+))?(, S2 (\d+:\d+:\d+))?(, S3 (\d+:\d+:\d+))?, fuel (\d+)\.\d+(, hasCut )?(, InLap )?(, OutLap )?(, SessionOver)?`, handleCurrLap),
		newLogMatcher(`^\s*Car (\d+) Pos (\d+)$`, handleGridPosition),
		newLogMatcher(`^CHAT (.*?): (.*)$`, handleChat),
		newLogMatcher(`^Updated leaderboard for \d+ clients \(([A-Za-z ]+)-<([-A-Za-z ]+)> (\d+) min\)$`, handleSessionUpdate),
		newLogMatcher(`Updated \d+ clients with new damage zones for car (\d+)$`, handleNewDamage),
	}
}

func toInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return value
}

var timeEr = regexp.MustCompile(`(\d+):(\d+):(\d+)`)

func timeToMs(ts string) int {
	p := timeEr.FindStringSubmatch(ts)
	if p == nil {
		return 0
	}

	return toInt(p[1])*60000 + toInt(p[2])*1000 + toInt(p[3])
}

func toLap(l *instance.LiveState, p []string) *instance.LapState {
	c := l.Cars[toInt(p[1])]
	if c == nil {
		logrus.WithFields(logrus.Fields{
			"carID":    toInt(p[1]),
			"rawCarID": p[1],
			"track":    l.Track,
		}).Warn("car not found while building lap")
		return nil
	}

	dIdx := toInt(p[2])
	if len(c.Drivers) < dIdx+1 {
		logrus.WithFields(logrus.Fields{
			"driverIndex":    dIdx,
			"rawDriverIndex": p[2],
			"track":          l.Track,
		}).Warn("driver not found while building lap")
		return nil
	}

	d := c.Drivers[dIdx]
	if d == nil {
		logrus.WithFields(logrus.Fields{
			"driverIndex":    dIdx,
			"rawDriverIndex": p[2],
			"track":          l.Track,
		}).Warn("driver index not found while building lap")
		return nil
	}

	lap := &instance.LapState{
		CarID:       c.CarID,
		DriverIndex: dIdx,
		Car:         c,
		Driver:      d,
		LapTimeMS:   toInt(p[3])*60000 + toInt(p[4])*1000 + toInt(p[5]),
		TimestampMS: toInt(p[6]),
		Flags:       0,
		S1:          p[9],
		S1MS:        timeToMs(p[9]),
		S2:          p[11],
		S2MS:        timeToMs(p[11]),
		S3:          p[13],
		S3MS:        timeToMs(p[13]),
		Fuel:        toInt(p[14]),
		HasCut:      p[15] != "",
		InLap:       p[16] != "",
		OutLap:      p[17] != "",
		SessionOver: p[18] != "",
	}

	if lap.HasCut {
		lap.Flags += 1
	}

	if lap.OutLap {
		lap.Flags += 4
	}

	if lap.InLap {
		lap.Flags += 8
	}

	if lap.SessionOver {
		lap.Flags += 1024
	}

	return lap
}

//	1            2          3  4  5                  6                       7
//
// Lap carId 1005, driverId 0, lapTime 1:53:895, timestampMS 52610019.000000, flags: 8808693760,
//
//	9            11           13           14          15     16      17      18
//
// S1 0:36:280, S2 0:40:037, S3 0:37:577, fuel 40.000000, HasCut, InLap, OutLap, SessionOver
func handleLap(i *instance.Instance, l *instance.LiveState, p []string) {
	lap := toLap(l, p)
	if lap == nil {
		return
	}

	l.SetLapState(lap)

	event.EmmitEventInstanceLiveNewLap(
		i.ToEIB(),
		lap.Driver.ToEILDB(),
		lap.Car.ToEILCB(),
		lap.LapTimeMS, lap.TimestampMS, lap.Flags, lap.Fuel,
		lap.S1, lap.S2, lap.S3,
		lap.HasCut, lap.InLap, lap.OutLap, lap.SessionOver,
	)
}

func handleCurrLap(i *instance.Instance, l *instance.LiveState, p []string) {
	lap := toLap(l, p)
	if lap == nil {
		return
	}

	l.SetCurrLapState(*lap)
}

func handleServerStarting(i *instance.Instance, s *instance.LiveState, _ []string) {
	s.SetServerState(instance.ServerStateNotRegistered)
}

func handleLobbyConnectionFailed(i *instance.Instance, s *instance.LiveState, _ []string) {
	s.SetServerState(instance.ServerStateNotRegistered)
}

func handleLobbyConnectionSucceeded(i *instance.Instance, s *instance.LiveState, _ []string) {
	s.SetServerState(instance.ServerStateOnline)
}

func handleNrClientsOnline(i *instance.Instance, s *instance.LiveState, p []string) {
	s.SetNrClients(toInt(p[1]))
}

func handleTrack(i *instance.Instance, s *instance.LiveState, p []string) {
	s.SetTrack(p[1])
}

func handleSessionPhaseChanged(i *instance.Instance, s *instance.LiveState, p []string) {
	s.SetSessionState(p[3], p[2], -1)

	event.EmmitEventInstanceLiveSessionPhaseChanged(
		i.ToEIB(),
		s.SessionType, s.SessionPhase, s.SessionRemaining,
	)
}

func handleSessionUpdate(i *instance.Instance, s *instance.LiveState, p []string) {
	s.SetSessionState(p[1], p[2], toInt(p[3]))
}

func handleNewConnection(i *instance.Instance, l *instance.LiveState, p []string) {
	l.AddNewConnection(toInt(p[1]), p[2], p[3], toInt(p[4]))
}

func handleNewCar(i *instance.Instance, l *instance.LiveState, p []string) {
	l.AddNewCar(toInt(p[1]), toInt(p[3]), toInt(p[2]))
}

func handleHandshake(i *instance.Instance, l *instance.LiveState, p []string) {
	carId := toInt(p[1])
	connId := toInt(p[2])
	l.Handshake(carId, connId)

	drv := l.GetDriver(connId)
	car := l.GetCar(carId)

	if drv == nil || car == nil {
		return
	}

	event.EmmitEventInstanceLiveNewDriver(i.ToEIB(), drv.Name, drv.PlayerID, car.CarID, car.RaceNumber, car.CarModel)
}

func handleDeadConnection(i *instance.Instance, l *instance.LiveState, p []string) {
	connId := toInt(p[1])

	if drv := l.GetDriver(connId); drv != nil {
		event.EmmitEventInstanceLiveRemoveConnection(i.ToEIB(), drv.Name, drv.PlayerID)
	}

	l.RemoveConnection(toInt(p[1]))
}

func handleCarPurge(i *instance.Instance, l *instance.LiveState, p []string) {
	l.PurgeCar(toInt(p[1]))
}

func handleGridPosition(i *instance.Instance, l *instance.LiveState, p []string) {
	l.SetCarPosition(toInt(p[1]), toInt(p[2]))
}

func handleResettingRace(i *instance.Instance, l *instance.LiveState, _ []string) {
	l.AdvanceSession()
}

func handleChat(i *instance.Instance, l *instance.LiveState, p []string) {
	l.AddChat(p[1], p[2])
}

func handleNewDamage(i *instance.Instance, l *instance.LiveState, p []string) {
	l.AddDamage(toInt(p[1]))

	if car := l.GetCar(toInt(p[1])); car != nil {
		event.EmmitEventInstanceLiveNewDamageZone(
			i.ToEIB(),
			car.CurrentDriver.ToEILDB(),
			car.ToEILCB(),
		)
	}
}
