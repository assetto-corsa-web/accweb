package app

import (
	"net/http"

	"github.com/assetto-corsa-web/accweb/internal/pkg/instance"

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

	ServerState  instance.ServerState `json:"serverState"`
	NrClients    int                  `json:"nrClients"`
	SessionType  string               `json:"sessionType"`
	SessionPhase string               `json:"sessionPhase"`
}

func buildListServerItem(srv *instance.Instance) ListServerItem {
	return ListServerItem{
		ID:           srv.GetID(),
		Name:         srv.AccCfg.Settings.ServerName,
		IsRunning:    srv.IsRunning(),
		ProcessID:    srv.GetProcessID(),
		UdpPort:      srv.AccCfg.Configuration.UdpPort,
		TcpPort:      srv.AccCfg.Configuration.TcpPort,
		Track:        srv.AccCfg.Event.Track,
		ServerState:  srv.Live.ServerState,
		NrClients:    srv.Live.NrClients,
		SessionType:  srv.Live.SessionType,
		SessionPhase: srv.Live.SessionPhase,
	}
}

// ListServers Handle the list all ACC dedicated servers
// @Summary List all ACC dedicated servers
// @Schemes
// @Description List all ACC dedicated servers
// @Tags servers
// @Accept json
// @Produce json
// @Success 200 {object} []ListServerItem{}
// @Router /servers [get]
// @Security JWT
func (h *Handler) ListServers(c *gin.Context) {
	list := h.sm.GetServers()
	res := []ListServerItem{}
	for _, srv := range list {
		res = append(res, buildListServerItem(srv))
	}

	c.JSON(http.StatusOK, res)
}

// StopAllServers Stops all running ACC dedicated servers
// @Summary Stops all running ACC dedicated servers
// @Schemes
// @Description Stops all running ACC dedicated servers
// @Tags servers
// @Accept json
// @Produce json
// @Success 200
// @Router /servers/stop-all [post]
// @Security JWT
func (h *Handler) StopAllServers(c *gin.Context) {
	if err := h.sm.StopAll(); err != nil {
		logrus.WithError(err).Error("failed during stoping all servers")
	}

	c.JSON(http.StatusOK, nil)
}
