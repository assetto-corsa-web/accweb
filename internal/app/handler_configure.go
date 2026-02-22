package app

import (
	"net/http"
	"strconv"

	"github.com/assetto-corsa-web/accweb/internal/pkg/instance"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/labstack/echo/v5"
)

// ListGlobalAdmins List all global admins
// @Summary List all global admins
// @Schemes
// @Description List all global admins
// @Tags servers
// @Accept json
// @Produce json
// @Success 200 {object} instance.AccwebGlobalEntrylistJson{}
// @Failure 500 {object} AccWError
// @Router /configure/global-entrylist [get]
// @Security JWT
func (h *Handler) ListGlobalAdmins(c *echo.Context) error {
	var data instance.AccwebGlobalEntrylistJson
	if err := h.sm.LoadGlobalEntry(server_manager.GlobalListCtxEntry, &data); err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	return c.JSON(http.StatusOK, data)
}

// SaveGlobalAdmins Save global admins
// @Summary Save global admins
// @Schemes
// @Description Save global admins
// @Tags servers
// @Accept json
// @Produce json
// @Param payload body instance.AccwebGlobalEntrylistJson true "Global admins payload"
// @Success 200
// @Failure 400 {object} AccWError
// @Failure 500 {object} AccWError
// @Router /configure/global-entrylist [post]
// @Security JWT
func (h *Handler) SaveGlobalAdmins(c *echo.Context) error {
	var payload instance.AccwebGlobalEntrylistJson
	if err := c.Bind(&payload); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}

	if err := h.sm.SaveGlobalEntry(server_manager.GlobalListCtxEntry, payload); err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	return c.JSON(http.StatusOK, nil)
}

// ListGlobalBans List all global bans
// @Summary List all global bans
// @Schemes
// @Description List all global bans
// @Tags servers
// @Accept json
// @Produce json
// @Success 200 {object} instance.AccwebGlobalBanlistJson{}
// @Failure 500 {object} AccWError
// @Router /configure/global-ban [get]
// @Security JWT
func (h *Handler) ListGlobalBans(c *echo.Context) error {
	var data instance.AccwebGlobalBanlistJson
	err := h.sm.LoadGlobalEntry(server_manager.GlobalEntryCtxBan, &data)
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	return c.JSON(http.StatusOK, data)
}

// SaveGlobalBans Save global bans
// @Summary Save global bans
// @Schemes
// @Description Save global bans
// @Tags servers
// @Accept json
// @Produce json
// @Param payload body instance.AccwebGlobalBanEntryJson true "Global ban entry payload"
// @Success 200
// @Failure 400 {object} AccWError
// @Failure 500 {object} AccWError
// @Router /configure/global-ban [post]
// @Security JWT
func (h *Handler) SaveGlobalBans(c *echo.Context) error {
	var entry instance.AccwebGlobalBanEntryJson
	if err := c.Bind(&entry); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}

	var data instance.AccwebGlobalBanlistJson
	err := h.sm.LoadGlobalEntry(server_manager.GlobalEntryCtxBan, &data)
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	for _, e := range data.Entries {
		if e.PlayerId == entry.PlayerId {
			return echo.ErrBadRequest.Wrap(err)
		}
	}

	data.Entries = append(data.Entries, entry)

	if err := h.sm.SaveGlobalEntry(server_manager.GlobalEntryCtxBan, data); err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	return c.JSON(http.StatusOK, nil)
}

// EnableToggleGlobalBans Toggle global bans enabled status
// @Summary Toggle global bans enabled status
// @Schemes
// @Description Toggle global bans enabled status
// @Tags servers
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 {object} AccWError
// @Router /configure/global-ban/enable-toggle [post]
// @Security JWT
func (h *Handler) EnableToggleGlobalBans(c *echo.Context) error {
	var data instance.AccwebGlobalBanlistJson
	err := h.sm.LoadGlobalEntry(server_manager.GlobalEntryCtxBan, &data)
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	data.Enabled = !data.Enabled

	if err := h.sm.SaveGlobalEntry(server_manager.GlobalEntryCtxBan, data); err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	return c.JSON(http.StatusOK, nil)
}

// RemoveGlobalBans Remove a global ban entry
// @Summary Remove a global ban entry
// @Schemes
// @Description Remove a global ban entry by ID
// @Tags servers
// @Accept json
// @Produce json
// @Param id path int true "Ban entry ID"
// @Success 200
// @Failure 400 {object} AccWError
// @Failure 500 {object} AccWError
// @Router /configure/global-ban/{id} [delete]
// @Security JWT
func (h *Handler) RemoveGlobalBans(c *echo.Context) error {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}

	var data instance.AccwebGlobalBanlistJson
	if err := h.sm.LoadGlobalEntry(server_manager.GlobalEntryCtxBan, &data); err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	if id >= len(data.Entries) {
		return echo.ErrBadRequest.Wrap(err)
	}

	data.Entries = append(data.Entries[:id], data.Entries[id+1:]...)

	if err := h.sm.SaveGlobalEntry(server_manager.GlobalEntryCtxBan, data); err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}

	return c.JSON(http.StatusOK, nil)
}
