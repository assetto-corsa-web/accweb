package app

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

var ErrForbidden = errors.New("access denied")

type ACCWebAuthLevel int

const (
	ACCWebAuthLevel_Mod ACCWebAuthLevel = iota
	ACCWebAuthLevel_Adm
)

func ACCWebAuthMiddleware(lvl ACCWebAuthLevel) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			user, err := echo.ContextGet[*jwt.Token](c, identityKey)
			if err != nil {
				return echo.ErrUnauthorized.Wrap(err)
			}

			u, ok := user.Claims.(*User)
			if !ok {
				return echo.ErrForbidden
			}

			if lvl == ACCWebAuthLevel_Mod && (!u.Mod && !u.Admin) {
				return echo.ErrForbidden
			}

			if lvl == ACCWebAuthLevel_Adm && !u.Admin {
				return echo.ErrForbidden
			}

			return next(c)
		}
	}
}
