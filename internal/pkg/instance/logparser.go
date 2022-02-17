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
			handleLogger(s, matches)
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
		newLogMatcher(`^Resetting race weekend$`, handleLogger),
		newLogMatcher(`^New connection request: id (\d+) (.+) (S\d+) on car model (\d+)$`, handleNewConnection),
		newLogMatcher(`^Creating new car connection: carId (\d+), carModel (\d+), raceNumber #(\d+)$`, handleNewCar),
		newLogMatcher(`Removing dead connection (\d+)`, handleDeadConnection),
		//newLogMatcher(`^car (\d+) has no driving connection anymore, will remove it$`, handleLogger),
		newLogMatcher(`^Purging car_id (\d+)$`, handleCarPurge),
		newLogMatcher(`^Lap carId (\d+), driverId (\d+), lapTime (\d+):(\d+):(\d+), timestampMS (\d+)\.\d+, flags: (\d+), S1 (\d+:\d+:\d+), S2 ((\d+:\d+:\d+)), S3 (\d+:\d+:\d+), fuel (\d+)\.\d+(, hasCut )?(, InLap )?(, OutLap )?(, SessionOver)?$`, handleLogger),
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

func handleLogger(_ *LiveState, p []string) {
	logrus.WithField("line", p[0]).WithField("attr", p[1:]).Info("log handled")
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
