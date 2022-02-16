package instance

import (
	"regexp"
	"strconv"
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
	}
}

func toInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return value
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
