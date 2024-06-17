package callback

import (
	"bytes"
	"net/http"

	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/sirupsen/logrus"
)

type client struct {
	Enabled      bool
	Url          string
	Headers      http.Header
	Events       []string
	AllEvents    bool
	cachedEvents map[string]struct{}
}

func newClient(cfgClient cfg.CallbackClient) *client {
	ce := map[string]struct{}{}
	for _, v := range cfgClient.Events {
		ce[v] = struct{}{}
	}

	if len(ce) == 0 {
		ce = validEvents
	}

	hdrs := http.Header{}
	for h, v := range cfgClient.Headers {
		hdrs.Add(h, v)
	}

	return &client{
		Enabled:      cfgClient.Enabled == nil || *cfgClient.Enabled,
		Url:          cfgClient.Url,
		Headers:      hdrs,
		AllEvents:    len(cfgClient.Events) == 0,
		Events:       cfgClient.Events,
		cachedEvents: ce,
	}
}

func (c *client) shouldProcess(ev string) bool {
	if !c.Enabled {
		return false
	}

	if _, ok := validEvents[ev]; !ok {
		return false
	}

	if _, ok := c.cachedEvents[ev]; !ok {
		return false
	}

	return true
}

func (c *client) process(buf []byte) {
	req, err := http.NewRequest(http.MethodPost, c.Url, bytes.NewReader(buf))
	if err != nil {
		logrus.WithError(err).Error("failed to build request.")
		return
	}

	req.Header = c.Headers

	if _, err := httpClient.Do(req); err != nil {
		logrus.WithError(err).Warn("failed to request callback.")
	}
}
