package instance

import (
	"regexp"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type matcherHandler func(*LiveState, []string)

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

func (l *logParser) processLine(s *LiveState, line string) {
	for _, matcher := range l.matchers {
		matches := matcher.er.FindStringSubmatch(line)
		if matches != nil {
			logrus.WithField("line", matches[0]).WithField("attr", matches[1:]).Debug("log handled")
			matcher.handler(s, matches)
			s.UpdatedAt = time.Now().UTC()
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
	}
}

func toInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return value
}

func toLap(l *LiveState, p []string) *LapState {
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

	lap := &LapState{
		CarID:       c.CarID,
		DriverIndex: dIdx,
		Car:         c,
		Driver:      d,
		LapTimeMS:   toInt(p[3])*60000 + toInt(p[4])*1000 + toInt(p[5]),
		TimestampMS: toInt(p[6]),
		Flags:       0,
		S1:          p[9],
		S2:          p[11],
		S3:          p[13],
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

//             1            2          3  4  5                  6                       7
// Lap carId 1005, driverId 0, lapTime 1:53:895, timestampMS 52610019.000000, flags: 8808693760,
//       9            11           13           14          15     16      17      18
// S1 0:36:280, S2 0:40:037, S3 0:37:577, fuel 40.000000, HasCut, InLap, OutLap, SessionOver
func handleLap(l *LiveState, p []string) {
	lap := toLap(l, p)
	if lap == nil {
		return
	}

	l.setLapState(lap)
}

func handleCurrLap(l *LiveState, p []string) {
	lap := toLap(l, p)
	if lap == nil {
		return
	}

	l.setCurrLapState(*lap)
}

func handleServerStarting(s *LiveState, _ []string) {
	s.setServerState(ServerStateNotRegistered)
}

func handleLobbyConnectionFailed(s *LiveState, _ []string) {
	s.setServerState(ServerStateNotRegistered)
}

func handleLobbyConnectionSucceeded(s *LiveState, _ []string) {
	s.setServerState(ServerStateOnline)
}

func handleNrClientsOnline(s *LiveState, p []string) {
	s.setNrClients(toInt(p[1]))
}

func handleTrack(s *LiveState, p []string) {
	s.setTrack(p[1])
}

func handleSessionPhaseChanged(s *LiveState, p []string) {
	s.setSessionState(p[3], p[2], 0)
}

func handleSessionUpdate(s *LiveState, p []string) {
	s.setSessionState(p[1], p[2], toInt(p[3]))
}

func handleNewConnection(l *LiveState, p []string) {
	l.addNewConnection(toInt(p[1]), p[2], p[3], toInt(p[4]))
}

func handleNewCar(l *LiveState, p []string) {
	l.addNewCar(toInt(p[1]), toInt(p[3]), toInt(p[2]))
}

func handleHandshake(l *LiveState, p []string) {
	l.handshake(toInt(p[1]), toInt(p[2]))
}

func handleDeadConnection(l *LiveState, p []string) {
	l.removeConnection(toInt(p[1]))
}

func handleCarPurge(l *LiveState, p []string) {
	l.purgeCar(toInt(p[1]))
}

func handleGridPosition(l *LiveState, p []string) {
	l.setCarPosition(toInt(p[1]), toInt(p[2]))
}

func handleResettingRace(l *LiveState, _ []string) {
	l.advanceSession()
}

func handleChat(l *LiveState, p []string) {
	l.addChat(p[1], p[2])
}
