package app

import (
	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	sm *server_manager.Service
}

func StartServer(config *cfg.Config, sM *server_manager.Service) {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	if !config.Dev {
		gin.SetMode(gin.ReleaseMode)
	}

	// setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.CORS.Origins}
	r.Use(cors.New(corsConfig))

	// setup routers
	h := Handler{sm: sM}

	api := r.Group("/api")

	api.GET("/servers", h.ListServers)
	api.POST("/servers/stopall", h.StopAllServers)

	api.GET("/instance/:id", h.ListServers)
	api.POST("/instance/:id", h.ListServers)
	api.DELETE("/instance/:id", h.ListServers)
	api.POST("/instance/:id/start", h.ListServers)
	api.POST("/instance/:id/stop", h.ListServers)
	api.POST("/instance/:id/clone", h.ListServers)

	// Starting HTTP Server
	if config.Webserver.TLS {
		if err := r.RunTLS(config.Webserver.Host, config.Webserver.Cert, config.Webserver.PrivateKey); err != nil {
			logrus.WithError(err).Fatal("failed to start http server with TLS")
		}
	} else {
		if err := r.Run(config.Webserver.Host); err != nil {
			logrus.WithError(err).Fatal("failed to start http server")
		}
	}
}
