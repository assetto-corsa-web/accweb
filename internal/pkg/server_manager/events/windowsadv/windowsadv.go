package windowsadv

import (
	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

var sM *server_manager.Service

func Register(sm *server_manager.Service) {
	sM = sm
	event.Register(handleEvent)
}

func handleEvent(data event.Eventer) {
	switch ev := data.(type) {
	case event.EventInstanceBeforeStart:
		i, err := sM.GetServerByID(ev.InstanceId)
		if err != nil {
			logrus.Error("instance not found")
		}

		if !hasAdvancedWindowsConfig(i) {
			return
		}

		startWithAdvWindows(i)

	case event.EventInstanceStopped:
		i, err := sM.GetServerByID(ev.InstanceId)
		if err != nil {
			logrus.Error("instance not found")
		}

		if !hasAdvancedWindowsConfig(i) {
			return
		}

		stopWithAdvWindows(i)
	}
}
