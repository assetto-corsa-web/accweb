package callback

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

var sM *server_manager.Service
var client *http.Client
var cacheEvents = map[string]struct{}{}

func Register(sm *server_manager.Service) {
	if !sm.Config().Callback.Enable {
		return
	}

	sM = sm

	for _, v := range sM.Config().Callback.Events {
		cacheEvents[v] = struct{}{}
	}

	client = &http.Client{
		Timeout: 500 * time.Millisecond,
	}

	event.Register(handleEvent)
}

func handleEvent(ev event.Eventer) {
	cb := sM.Config().Callback
	info := ev.GetInfo()

	if _, ok := cacheEvents[info.Name]; !ok {
		return
	}

	buf := bytes.NewBuffer(nil)

	if err := json.NewEncoder(buf).Encode(ev); err != nil {
		logrus.WithError(err).Error("failed to build callback payload.")
		return
	}

	req, err := http.NewRequest(http.MethodPost, cb.Url, buf)
	if err != nil {
		logrus.WithError(err).Error("failed to build request.")
		return
	}

	for h, v := range cb.Headers {
		req.Header.Add(h, v)
	}

	go func(req *http.Request) {
		_, err = client.Do(req)
		if err != nil {
			logrus.WithError(err).Error("failed to request callback.")
		}
	}(req)

}
