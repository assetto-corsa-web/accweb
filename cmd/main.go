package main

import (
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

const configFile = "config.yml"

func main() {
	c := cfg.Load(configFile)

	sM := server_manager.New(c.ConfigPath, c.ACC.ServerPath, c.ACC.ServerExe)

	if err := sM.Bootstrap(); err != nil {
		logrus.WithError(err).Fatal("failed to bootstrap accweb")
	}

	time.Sleep(10 * time.Second)
	if err := sM.StopAll(); err != nil {
		logrus.WithError(err).Fatal("failed to stop all acc servers")
	}
}
