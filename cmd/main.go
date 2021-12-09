package main

import (
	"github.com/assetto-corsa-web/accweb/internal/app"
	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

const configFile = "config.yml"

func main() {
	c := cfg.Load(configFile)

	sM := server_manager.New(c.ConfigPath, c.ACC.ServerPath, c.ACC.ServerExe)

	logrus.Info("initializing accweb")
	if err := sM.Bootstrap(); err != nil {
		logrus.WithError(err).Fatal("failed to bootstrap accweb")
	}

	logrus.WithField("addr", c.Webserver.Host).Info("initializing web server")
	app.StartServer(c, sM)
}
