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

	clientTimeout := 100 * time.Millisecond
	if cbCfg.Timeout != nil {
		clientTimeout = *cbCfg.Timeout
	}

	httpClient = &http.Client{
		Timeout: clientTimeout,
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

	if _, ok := validEvents[info.Name]; !ok {
		return
	}

	buf, err := json.Marshal(ev)
	if err != nil {
		logrus.WithError(err).Error("failed to build callback payload.")
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(clients))

	for _, c := range clients {
		go func(c *client) {
			defer wg.Done()

			if !c.shouldProcess(info.Name) {
				return
			}

			ts := time.Now()

			logrus.Debug("processing callback client " + c.Url + " event " + info.Name + " with body: " + string(buf))
			logrus.Debug("callback handled in " + time.Since(ts).String())

			c.process(buf)
		}(c)
	}

	wg.Wait()
}
