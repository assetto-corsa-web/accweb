package instance

import (
	"regexp"
	"strconv"

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
			logrus.WithField("line", matches[0]).WithField("attr", matches[1:]).Info("log handled")
			matcher.handler(s, matches)
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
		newLogMatcher(`Removing dead connection (\d+)`, handleDeadConnection),
		//newLogMatcher(`^car (\d+) has no driving connection anymore, will remove it$`, handleLogger),
		newLogMatcher(`^Purging car_id (\d+)$`, handleCarPurge),
		newLogMatcher(`^Lap carId (\d+), driverId (\d+), lapTime (\d+):(\d+):(\d+), timestampMS (\d+)\.\d+, flags: (.*?), S1 (\d+:\d+:\d+), S2 (\d+:\d+:\d+), S3 (\d+:\d+:\d+), fuel (\d+)\.\d+(, hasCut )?(, InLap )?(, OutLap )?(, SessionOver)?$`, handleLap),
		newLogMatcher(`^\s*Car (\d+) Pos (\d+)$`, handleGridPosition),
	}
}

func toInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return value
}

//             1            2          3  4  5                  6                       7
// Lap carId 1005, driverId 0, lapTime 1:53:895, timestampMS 52610019.000000, flags: 8808693760,
//       8            9           10           11           12     13      14      15
// S1 0:36:280, S2 0:40:037, S3 0:37:577, fuel 40.000000, HasCut, InLap, OutLap, SessionOver
func handleLap(l *LiveState, p []string) {
	c := l.Cars[toInt(p[1])]
	if c == nil {
		return
	}

	dIdx := toInt(p[2])
	if len(c.Drivers) < dIdx+1 {
		return
	}

	d := c.Drivers[dIdx]
	if d == nil {
		return
	}

	lap := LapState{
		CarID:       c.CarID,
		DriverIndex: dIdx,
		Car:         c,
		Driver:      d,
		LapTimeMS:   toInt(p[3])*60000 + toInt(p[4])*1000 + toInt(p[5]),
		TimestampMS: toInt(p[6]),
		Flags:       0,
		S1:          p[8],
		S2:          p[9],
		S3:          p[10],
		Fuel:        toInt(p[11]),
		HasCut:      p[12] != "",
		InLap:       p[13] != "",
		OutLap:      p[14] != "",
		SessionOver: p[15] != "",
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

	l.setLapState(&lap)
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
	old := s.SessionType

	s.setSessionState(p[3], p[2])

	if old != p[3] {
		s.advanceSession()
	}
}

func handleNewConnection(l *LiveState, p []string) {
	l.addNewConnection(toInt(p[1]), p[2], p[3], toInt(p[4]))
}

func handleNewCar(l *LiveState, p []string) {
	l.addNewCar(toInt(p[1]), toInt(p[3]), toInt(p[2]))
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
