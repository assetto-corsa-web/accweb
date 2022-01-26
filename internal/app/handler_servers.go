package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logrus.WithError(err).Error("failed during stoping all servers")
	}

	c.JSON(http.StatusOK, nil)
}
