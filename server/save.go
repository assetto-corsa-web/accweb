package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	configurationJsonName = "configuration.json"
	settingsJsonName      = "settings.json"
	eventJsonName         = "event.json"
	entrylistJsonName     = "entrylist.json"
	configVersion         = 1
)

func SaveServerSettings(settings *ServerSettings) error {
	setConfigVersion(settings)
	dir, id, err := getConfigDirectoryAndID(settings.Id)

	if err != nil {
		return err
	}

	if err := saveConfigToFile(settings.Configuration, dir, configurationJsonName); err != nil {
		return err
	}

	if err := saveConfigToFile(settings.Settings, dir, settingsJsonName); err != nil {
		return err
	}

	if err := saveConfigToFile(settings.Event, dir, eventJsonName); err != nil {
		return err
	}

	if err := saveConfigToFile(settings.Entrylist, dir, entrylistJsonName); err != nil {
		return err
	}

	if settings.Id == 0 {
		settings.Id = id
		addServer(settings)
		logrus.WithField("server", settings).Debug("Adding new server")
	} else {
		server := GetServerById(settings.Id, false)
		settings.PID = server.PID
		settings.Cmd = server.Cmd
		setServer(settings)
		logrus.WithField("server", settings).Debug("Updating existing server")
	}

	return nil
}

func setConfigVersion(settings *ServerSettings) {
	settings.Configuration.ConfigVersion = configVersion
	settings.Settings.ConfigVersion = configVersion
	settings.Event.ConfigVersion = configVersion
	settings.Entrylist.ConfigVersion = configVersion
}

func getConfigDirectoryAndID(id int) (string, int, error) {
	// create new ID for new server or use existing one
	if id == 0 {
		id = int(time.Now().Unix())
	}

	dir := filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id))

	if err := os.MkdirAll(dir, 0755); err != nil {
		logrus.WithField("err", err).Error("Error creating configuration directory")
		return "", 0, err
	}

	return dir, id, nil
}

func saveConfigToFile(config interface{}, dir, name string) error {
	data, err := json.Marshal(config)

	if err != nil {
		logrus.WithField("err", err).Error("Error marshalling server configuration")
		return err
	}

	path := filepath.Join(dir, name)
	logrus.WithField("path", path).Debug("Saving server configuration file")

	if err := ioutil.WriteFile(path, data, 0655); err != nil {
		logrus.WithField("err", err).Error("Error saving server configuration file")
		return err
	}

	return nil
}
