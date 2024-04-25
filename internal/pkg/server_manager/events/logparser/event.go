package logparser

import (
	"strings"

	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

var sM *server_manager.Service
var parser *logParser

func Register(sm *server_manager.Service) {
	sM = sm
	parser = newLogParser()
	event.Register(handleEvent)
}

func handleEvent(data event.Eventer) {
	switch ev := data.(type) {
	case event.EventInstanceOutput:
		i, err := sM.GetServerByID(ev.InstanceId)
		if err != nil {
			logrus.WithError(err).Error("instance not found")
			return
		}

		parser.processLine(i, strings.TrimSpace(string(ev.Output)))
	}
}
