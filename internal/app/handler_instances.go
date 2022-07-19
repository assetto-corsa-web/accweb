package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/assetto-corsa-web/accweb/internal/pkg/instance"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/gin-gonic/gin"
)

type ExtraAccSettings struct {
	PasswordIsEmpty          bool `json:"passwordIsEmpty"`
	AdminPasswordIsEmpty     bool `json:"adminPasswordIsEmpty"`
	SpectatorPasswordIsEmpty bool `json:"spectatorPasswordIsEmpty"`
}

type InstancePayload struct {
	ID               string                      `json:"id"`
	Path             string                      `json:"path"`
	IsRunning        bool                        `json:"is_running"`
	PID              int                         `json:"pid"`
	Settings         instance.AccWebSettingsJson `json:"accWeb"`
	AccSettings      instance.AccConfigFiles     `json:"acc"`
	AccExtraSettings ExtraAccSettings            `json:"accExtraSettings"`
}

type InstanceOS struct {
	Name   string `json:"name"`
	NumCPU int    `json:"numCpu"`
}

type SaveInstancePayload struct {
	AccWeb           instance.AccWebSettingsJson `json:"accWeb"`
	Acc              instance.AccConfigFiles     `json:"acc"`
	AccExtraSettings ExtraAccSettings            `json:"accExtraSettings"`
}

func NewInstancePayload(srv *instance.Instance) InstancePayload {
	res := InstancePayload{
		ID:          srv.GetID(),
		Path:        srv.Path,
		IsRunning:   srv.IsRunning(),
		PID:         srv.GetProcessID(),
		Settings:    srv.Cfg.Settings,
		AccSettings: srv.AccCfg,
	}

	res.AccExtraSettings.PasswordIsEmpty = res.AccSettings.Settings.Password == ""
	res.AccSettings.Settings.Password = ""

	res.AccExtraSettings.AdminPasswordIsEmpty = res.AccSettings.Settings.AdminPassword == ""
	res.AccSettings.Settings.AdminPassword = ""

	res.AccExtraSettings.SpectatorPasswordIsEmpty = res.AccSettings.Settings.SpectatorPassword == ""
	res.AccSettings.Settings.SpectatorPassword = ""

	return res
}

// GetInstance Get instance information
// @Summary Get acc instance information
// @Description Get acc instance information
// @Tags instances
// @Accept json
// @Produce json
// @Success 200 {object} InstancePayload
// @Failure 404
// @Param id path int true "Instance ID"
// @Router /instance/{id} [get]
// @Security JWT
func (h *Handler) GetInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	res := NewInstancePayload(srv)

	c.JSON(http.StatusOK, res)
}

// NewInstance Create new instance information
// @Summary Create new acc instance information
// @Description Create new acc instance information
// @Tags instances
// @Accept json
// @Produce json
// @Success 200 {object} InstancePayload
// @Failure 400  {object} AccWError
// @Failure 500  {object} AccWError
// @Param instance body SaveInstancePayload true "Instance data"
// @Router /instance [post]
// @Security JWT
func (h *Handler) NewInstance(c *gin.Context) {
	var json SaveInstancePayload
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, newAccWError(err.Error()))
		return
	}

	srv, err := h.sm.Create(&json.Acc, json.AccWeb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newAccWError(err.Error()))
		return
	}

	res := NewInstancePayload(srv)

	c.JSON(http.StatusCreated, res)
}

// SaveInstance Saves instance information
// @Summary Saves acc instance information
// @Description Saves acc instance information
// @Tags instances
// @Accept json
// @Produce json
// @Success 200 {object} InstancePayload
// @Failure 404
// @Failure 400 {object} AccWError
// @Failure 500 {object} AccWError
// @Param id path int true "Instance ID"
// @Param instance body SaveInstancePayload true "Instance data"
// @Router /instance/{id} [post]
// @Security JWT
func (h *Handler) SaveInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	var json SaveInstancePayload
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, newAccWError(err.Error()))
		return
	}

	if json.AccExtraSettings.PasswordIsEmpty {
		json.Acc.Settings.Password = ""
	} else if json.Acc.Settings.Password == "" {
		json.Acc.Settings.Password = srv.AccCfg.Settings.Password
	}

	if json.AccExtraSettings.SpectatorPasswordIsEmpty {
		json.Acc.Settings.SpectatorPassword = ""
	} else if json.Acc.Settings.SpectatorPassword == "" {
		json.Acc.Settings.SpectatorPassword = srv.AccCfg.Settings.SpectatorPassword
	}

	if json.AccExtraSettings.AdminPasswordIsEmpty {
		json.Acc.Settings.AdminPassword = ""
	} else if json.Acc.Settings.AdminPassword == "" {
		json.Acc.Settings.AdminPassword = srv.AccCfg.Settings.AdminPassword
	}

	if err := srv.CanSaveSettings(json.AccWeb, json.Acc); err != nil {
		c.JSON(http.StatusBadRequest, newAccWError(err.Error()))
		return
	}

	srv.AccCfg = json.Acc
	srv.Cfg.Settings = json.AccWeb

	if err := srv.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, newAccWError(err.Error()))
		return
	}

	res := NewInstancePayload(srv)

	c.JSON(http.StatusCreated, res)
}

// DeleteInstance Delete instance
// @Summary Delete acc instance
// @Description Delete acc instance
// @Tags instances
// @Accept json
// @Produce json
// @Success 200
// @Failure 404
// @Failure 500 {object} AccWError
// @Param id path int true "Instance ID"
// @Router /instance/{id} [delete]
// @Security JWT
func (h *Handler) DeleteInstance(c *gin.Context) {
	id := c.Param("id")

	if err := h.sm.Delete(id); err != nil {
		if errors.Is(err, server_manager.ErrServerNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusInternalServerError, newAccWError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}

// StartInstance Starts acc instance
// @Summary Starts acc instance
// @Description Starts acc instance
// @Tags instances
// @Accept json
// @Produce json
// @Success 200
// @Failure 404
// @Failure 400 {object} AccWError
// @Failure 500 {object} AccWError
// @Param id path int true "Instance ID"
// @Router /instance/{id}/start [post]
// @Security JWT
func (h *Handler) StartInstance(c *gin.Context) {
	if err := h.sm.Start(c.Param("id")); err != nil {
		if errors.Is(err, server_manager.ErrServerNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		if errors.Is(err, instance.ErrServerCantBeRunning) {
			c.JSON(http.StatusBadRequest, newAccWError(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, newAccWError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}

// StopInstance Stops acc instance
// @Summary Stops acc instance
// @Description Stops acc instance
// @Tags instances
// @Accept json
// @Produce json
// @Success 200
// @Failure 404
// @Failure 400 {object} AccWError
// @Failure 500 {object} AccWError
// @Param id path int true "Instance ID"
// @Router /instance/{id}/stop [post]
// @Security JWT
func (h *Handler) StopInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if err := srv.Stop(); err != nil {
		c.JSON(http.StatusInternalServerError, newAccWError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}

// CloneInstance Clones acc instance
// @Summary Clones acc instance
// @Description Clones acc instance
// @Tags instances
// @Accept json
// @Produce json
// @Success 200
// @Failure 404
// @Failure 500 {object} AccWError
// @Param id path int true "Instance ID"
// @Router /instance/{id}/clone [post]
// @Security JWT
func (h *Handler) CloneInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.Duplicate(id)
	if err != nil {
		if errors.Is(err, server_manager.ErrServerNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusInternalServerError, newAccWError(err.Error()))
		return
	}

	res := NewInstancePayload(srv)

	c.JSON(http.StatusOK, res)
}

type accWebInstanceLogs struct {
	ID   string `json:"id"`
	Logs string `json:"logs"`
}

// GetInstanceLogs Get acc instance logs
// @Summary Get acc instance logs
// @Description Get acc instance logs
// @Tags instances
// @Accept json
// @Produce json
// @Success 200 {object} accWebInstanceLogs
// @Failure 404
// @Failure 500 {object} AccWError
// @Param id path int true "Instance ID"
// @Router /instance/{id}/logs [get]
// @Security JWT
func (h *Handler) GetInstanceLogs(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	data, err := srv.GetAccServerLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, newAccWError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, accWebInstanceLogs{ID: srv.GetID(), Logs: string(data)})
}

// ExportInstance Get acc instance configuration files
// @Summary Get acc instance configuration files
// @Description Get acc instance configuration files
// @Tags instances
// @Accept json
// @Produce json
// @Success 200 string filedata "Zip file with all cofiguration files"
// @Failure 404
// @Failure 500 {object} AccWError
// @Param id path int true "Instance ID"
// @Router /instance/{id}/export [get]
// @Security JWT
func (h *Handler) ExportInstance(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	data, err := srv.ExportConfigFilesToZip()
	if err != nil {
		c.JSON(http.StatusInternalServerError, newAccWError(err.Error()))
		return
	}

	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"accweb_%s_cfg.zip\"", id))
	c.Data(http.StatusOK, "application/zip", data)
}

type LiveServerInstancePayload struct {
	ListServerItem
	Live *instance.LiveState `json:"live"`
}

// GetInstanceLiveState Get acc instance live information
// @Summary Get acc instance live information
// @Description Get acc instance live information
// @Tags instances
// @Accept json
// @Produce json
// @Success 200 {object} LiveServerInstancePayload
// @Failure 404
// @Param id path int true "Instance ID"
// @Router /instance/{id}/live [get]
// @Security JWT
func (h *Handler) GetInstanceLiveState(c *gin.Context) {
	id := c.Param("id")

	srv, err := h.sm.GetServerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, LiveServerInstancePayload{
		ListServerItem: buildListServerItem(srv),
		Live:           srv.Live,
	})
}
