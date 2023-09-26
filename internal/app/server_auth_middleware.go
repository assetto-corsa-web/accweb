package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ACCWebAuthLevel int

const (
	ACCWebAuthLevel_Mod ACCWebAuthLevel = iota
	ACCWebAuthLevel_Adm
)

func ACCWebAuthMiddleware(lvl ACCWebAuthLevel) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := GetUserFromClaims(c)

		if lvl == ACCWebAuthLevel_Mod && (!u.Mod && !u.Admin) {
			c.JSON(http.StatusForbidden, gin.H{"msg": errors.New("ximbro")})
			c.Abort()
			return
		}

		if lvl == ACCWebAuthLevel_Adm && !u.Admin {
			c.JSON(http.StatusForbidden, gin.H{"msg": errors.New("ximbro")})
			c.Abort()
			return
		}

		c.Next()
	}
}
