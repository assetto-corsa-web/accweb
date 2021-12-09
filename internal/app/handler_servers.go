package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListServerItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsRunning bool   `json:"is_running"`
	ProcessID int    `json:"process_id"`
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
