package callback

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

var httpClient *http.Client
var validEvents = map[string]struct{}{
	"instance_started":                    {},
	"instance_stopped":                    {},
	"instance_live_new_driver":            {},
	"instance_live_remove_connection":     {},
	"instance_live_new_lap":               {},
	"instance_live_session_phase_changed": {},
	"instance_live_new_damage_zone":       {},
	"instance_live_reseting_race_weekend": {},
	"instance_live_session_changed":       {},
}
var clients []*client

func Register(sm *server_manager.Service) {
	cbCfg := sm.Config().Callback
	if !cbCfg.Enabled {
		return
	}

	setupClients(cbCfg)

	httpClient = &http.Client{
		Timeout: 100 * time.Millisecond,
	}

	event.Register(handleEvent)
}

func setupClients(cb cfg.Callback) {
	clients = []*client{}

	for _, c := range cb.Clients {
		clients = append(clients, newClient(c))
	}
}

func handleEvent(ev event.Eventer) {
	info := ev.GetInfo()

	buf, err := json.Marshal(ev)
	if err != nil {
		logrus.WithError(err).Error("failed to build callback payload.")
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(clients))

	ts := time.Now()

	for _, c := range clients {
		go func(c *client) {
			defer wg.Done()

			if !c.shouldProcess(info.Name) {
				return
			}

			c.process(buf)
		}(c)
	}

	wg.Wait()
	logrus.Debug("callback handled in " + time.Since(ts).String())
}
