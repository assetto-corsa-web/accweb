package app

import (
	"log"
	"net/http"
	"os"

	"github.com/assetto-corsa-web/accweb/frontend"
	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/sirupsen/logrus"
)

// @title           ACCWeb API documentation
// @description     ACCweb api documentation
// @termsOfService  http://swagger.io/terms/
// @version         1.19

// @contact.name   ACCWeb project
// @contact.url    https://github.com/assetto-corsa-web/accweb

// @license.name  MIT
// @license.url   https://github.com/assetto-corsa-web/accweb/blob/master/LICENSE

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

const identityKey = "user_name"

type Handler struct {
	sm *server_manager.Service
}

func StartServer(config *cfg.Config, sM *server_manager.Service) {
	e := echo.New()

	if !config.Dev {
		e.Use(middleware.Recover())
	} else {
		e.Use(middleware.RequestLogger())
	}

	// setup CORS
	e.Use(middleware.CORS(config.CORS.Origins))

	e.RouteNotFound("/*", func(c *echo.Context) error {
		return c.Redirect(http.StatusFound, "/")
	})

	// setup routers
	setupRouters(e, sM, config)

	// Starting HTTP Server
	if config.Webserver.TLS {
		if err := http.ListenAndServeTLS(config.Webserver.Host, config.Webserver.Cert, config.Webserver.PrivateKey, e); err != nil {
			logrus.WithError(err).Fatal("failed to start http server with TLS")
		}
	} else {
		if err := e.Start(config.Webserver.Host); err != nil {
			logrus.WithError(err).Fatal("failed to start http server")
		}
	}
}

func setupRouters(e *echo.Echo, sM *server_manager.Service, config *cfg.Config) {
	h := Handler{sm: sM}

	if config.Dev {
		basedir := "frontend/dist"
		e.File("/", basedir+"/index.html")
		e.Static("/assets", basedir+"/assets")
		e.Static("/public", basedir+"/public")
	} else {
		e.StaticFS("/", echo.MustSubFS(frontend.Content, "dist"))
		e.StaticFS("/assets", echo.MustSubFS(frontend.Content, "dist/assets"))
		e.StaticFS("/public", echo.MustSubFS(frontend.Content, "dist/public"))
	}

	authMW := setupAuthRouters(config)

	e.POST("/api/login", h.LoginHandler)

	api := e.Group("/api")
	api.Use(authMW)

	api.GET("/logout", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{})
	})
	api.GET("/token", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{})
	})

	api.GET("/servers", h.ListServers)
	api.GET("/metadata", h.Metadata)
	api.GET("/instance/:id", h.GetInstance)
	api.GET("/instance/:id/logs", h.GetInstanceLogs)
	api.GET("/instance/:id/live", h.GetInstanceLiveState)
	api.GET("/instance/:id/export", h.ExportInstance)
	api.GET("/instance/:id/results", h.GetInstanceResultList)
	api.GET("/instance/:id/results/:resultId", h.GetInstanceResultContent)

	// moderator level
	mod := api.Group("", ACCWebAuthMiddleware(ACCWebAuthLevel_Mod))
	{
		mod.POST("/servers/stop-all", h.StopAllServers)
		mod.POST("/instance/:id/start", h.StartInstance)
		mod.POST("/instance/:id/stop", h.StopInstance)
	}

	// // admin level
	adm := api.Group("", ACCWebAuthMiddleware(ACCWebAuthLevel_Adm))
	{
		adm.POST("/instance", h.NewInstance)
		adm.POST("/instance/:id", h.SaveInstance)
		adm.DELETE("/instance/:id", h.DeleteInstance)
		adm.POST("/instance/:id/clone", h.CloneInstance)

		adm.GET("/configure/global-entrylist", h.ListGlobalAdmins)
		adm.POST("/configure/global-entrylist", h.SaveGlobalAdmins)

		adm.GET("/configure/global-ban", h.ListGlobalBans)
		adm.POST("/configure/global-ban", h.SaveGlobalBans)
		adm.POST("/configure/global-ban/enable-toggle", h.EnableToggleGlobalBans)
		adm.DELETE("/configure/global-ban/:id", h.RemoveGlobalBans)

		adm.GET("/configure/user", h.ListUsers)
		adm.POST("/configure/user", h.AddUser)
		adm.PUT("/configure/user/:username", h.UpdateUser)
		adm.DELETE("/configure/user/:username", h.DeleteUser)
	}

}

type User struct {
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	Admin    bool   `json:"admin"`
	Mod      bool   `json:"mod"`
	ReadOnly bool   `json:"read_only"`

	jwt.RegisteredClaims
}

func setupAuthRouters(config *cfg.Config) echo.MiddlewareFunc {
	publicKey, err := os.ReadFile(config.Auth.PublicKeyPath)
	if err != nil {
		log.Fatal("Could not read public key:", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatal("Could not parse public key:", err)
	}

	mdw := echojwt.WithConfig(echojwt.Config{
		ContextKey:    identityKey,
		SigningKey:    pubKey,
		SigningMethod: "RS256",
		TokenLookup:   "header:Authorization:Bearer ,query:token",
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(User)
		},
	})

	return mdw
}
