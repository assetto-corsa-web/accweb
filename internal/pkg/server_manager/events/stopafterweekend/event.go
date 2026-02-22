package stopafterweekend

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
	case event.EventInstanceLive:
		if _, ok := ev.Data.(event.EventInstanceLiveResetingRaceWeekend); !ok {
			return
		}

		i, err := sM.GetServerByID(ev.InstanceId)
		if err != nil {
			logrus.WithError(err).Error("instance not found")
			return
		}

		if i.Cfg.StopAfterWeekend {
			i.Cfg.StopAfterWeekend = false
			i.Stop()
		}

	case event.EventInstanceStopped:
		i, err := sM.GetServerByID(ev.InstanceId)
		if err != nil {
			logrus.WithError(err).Error("instance not found")
			return
		}

		i.Cfg.StopAfterWeekend = false

	case event.EventInstanceBeforeStart:
		i, err := sM.GetServerByID(ev.InstanceId)
		if err != nil {
			logrus.WithError(err).Error("instance not found")
			return
		}

		i.Cfg.StopAfterWeekend = false
	}
}
