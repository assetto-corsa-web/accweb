package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/assetto-corsa-web/accweb/internal/app"
	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager/events"
	"github.com/sirupsen/logrus"
)

const configFile = "config.yml"

func main() {
	c := cfg.Load(configFile)

	sM := server_manager.New(c)

	logrus.Info("accweb: checking for secrets...")
	helper.GenerateTokenKeysIfNotPresent(c.Auth.PublicKeyPath, c.Auth.PrivateKeyPath)

	logrus.Info("accweb: initializing...")

	events.InitializeAll(sM)

	if err := sM.Bootstrap(); err != nil {
		logrus.WithError(err).Fatal("failed to bootstrap accweb")
	}

	logrus.WithField("addr", c.Webserver.Host).Info("initializing web server")
	go app.StartServer(c, sM)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	logrus.Info("Gracefully terminating...")
	sM.StopAll()
}
