package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/assetto-corsa-web/accweb/internal/pkg/instance"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/gin-gonic/gin"
)

type InstancePayload struct {
	ID          string                    `json:"id"`
	Path        string                    `json:"path"`
	IsRunning   bool                      `json:"is_running"`
	PID         int                       `json:"pid"`
	Settings    instance.AccWebConfigJson `json:"accWeb"`
	AccSettings instance.AccConfigFiles   `json:"acc"`
}

type SaveInstancePayload struct {
	AccWeb instance.AccWebConfigJson `json:"accWeb"`
	Acc    instance.AccConfigFiles   `json:"acc"`
}

func NewInstancePayload(srv *instance.Instance) InstancePayload {
	return InstancePayload{
		ID:          srv.GetID(),
		Path:        srv.Path,
		IsRunning:   srv.IsRunning(),
		PID:         srv.GetProcessID(),
		Settings:    srv.Cfg,
		AccSettings: srv.AccCfg,
	}
}

func (h *Handler) GetInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	res := NewInstancePayload(srv)

	res.AccSettings.Settings.AdminPassword = ""
	res.AccSettings.Settings.Password = ""
	res.AccSettings.Settings.SpectatorPassword = ""

	c.JSON(http.StatusOK, res)
}

func (h *Handler) NewInstance(c *gin.Context) {
	var json SaveInstancePayload
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	srv, err := h.sm.Create(&json.Acc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := NewInstancePayload(srv)

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) SaveInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	var json SaveInstancePayload
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Acc.Settings.Password == "" {
		json.Acc.Settings.Password = srv.AccCfg.Settings.Password
	}

	if json.Acc.Settings.SpectatorPassword == "" {
		json.Acc.Settings.SpectatorPassword = srv.AccCfg.Settings.SpectatorPassword
	}

	if json.Acc.Settings.AdminPassword == "" {
		json.Acc.Settings.AdminPassword = srv.AccCfg.Settings.AdminPassword
	}

	srv.AccCfg = json.Acc
	srv.Cfg.AutoStart = json.AccWeb.AutoStart

	if err := srv.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := NewInstancePayload(srv)

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) DeleteInstance(c *gin.Context) {
	id := c.Param("id")

	if err := h.sm.Delete(id); err != nil {
		if errors.Is(err, server_manager.ErrServerNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) StartInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if err := srv.Start(); err != nil {
		if errors.Is(err, instance.ErrServerCantBeRunning) {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) StopInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if err := srv.Stop(); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) CloneInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.Duplicate(id)
	if err != nil {
		if errors.Is(err, server_manager.ErrServerNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	res := NewInstancePayload(srv)

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetInstanceLogs(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	data, err := srv.GetAccServerLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": srv.GetID(), "logs": string(data)})
}

func (h *Handler) ExportInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	data, err := srv.ExportConfigFilesToZip()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"accweb_%s_cfg.zip\"", id))
	c.Data(http.StatusOK, "application/zip", data)
}
