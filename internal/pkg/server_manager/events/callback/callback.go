package callback

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

var sM *server_manager.Service
var client *http.Client
var cacheEvents = map[string]struct{}{}
var validEvents = map[string]struct{}{
	"instance_started":                    {},
	"instance_stopped":                    {},
	"instance_live_new_driver":            {},
	"instance_live_remove_connection":     {},
	"instance_live_new_lap":               {},
	"instance_live_session_phase_changed": {},
	"instance_live_new_damage_zone":       {},
}

func Register(sm *server_manager.Service) {
	if !sm.Config().Callback.Enabled {
		return
	}

	sM = sm

	for _, v := range sM.Config().Callback.Events {
		cacheEvents[v] = struct{}{}
	}

	client = &http.Client{
		Timeout: 100 * time.Millisecond,
	}

	event.Register(handleEvent)
}

func shouldProcess(all bool, val string) bool {
	if _, ok := validEvents[val]; !ok {
		return false
	}

	if all {
		return true
	}

	_, ok := cacheEvents[val]
	return ok
}

func handleEvent(ev event.Eventer) {
	cb := sM.Config().Callback
	info := ev.GetInfo()

	if !shouldProcess(cb.AllEvents, info.Name) {
		return
	}

	buf, err := json.Marshal(ev)

	if err != nil {
		logrus.WithError(err).Error("failed to build callback payload.")
		return
	}

	wg := sync.WaitGroup{}
	hdrs := http.Header{}
	for h, v := range cb.Headers {
		hdrs.Add(h, v)
	}

	wg.Add(len(cb.Urls))
	ts := time.Now()

	for _, url := range cb.Urls {
		go func(wg *sync.WaitGroup, url string) {
			defer wg.Done()

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(buf))
			if err != nil {
				logrus.WithError(err).Error("failed to build request.")
				return
			}

			req.Header = hdrs

			if _, err := client.Do(req); err != nil {
				logrus.WithError(err).Warn("failed to request callback.")
			}
		}(&wg, url)
	}

	wg.Wait()
	logrus.Debug("callback handled in " + time.Since(ts).String())
}
