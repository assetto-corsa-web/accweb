package app

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ListServerItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsRunning bool   `json:"isRunning"`
	ProcessID int    `json:"pid"`
	UdpPort   int    `json:"udpPort"`
	TcpPort   int    `json:"tcpPort"`
	Track     string `json:"track"`
}

func (h *Handler) ListServers(c *gin.Context) {
	u := GetUserFromClaims(c)

	logrus.WithField("foo", u).Info("aeeew")

	list := h.sm.GetServers()
	res := []ListServerItem{}
	for id, srv := range list {
		res = append(res, ListServerItem{
			ID:        id,
			Name:      srv.AccCfg.Settings.ServerName,
			IsRunning: srv.IsRunning(),
			ProcessID: srv.GetProcessID(),
			UdpPort:   srv.AccCfg.Configuration.UdpPort,
			TcpPort:   srv.AccCfg.Configuration.TcpPort,
			Track:     srv.AccCfg.Event.Track,
		})
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) StopAllServers(c *gin.Context) {
	if err := h.sm.StopAll(); err != nil {
		// todo
	}

	c.JSON(http.StatusOK, nil)
}
