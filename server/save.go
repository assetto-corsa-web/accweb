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

	if settings.Id == 0 {
		settings.Id = id
		serverList = append(serverList, *settings)
	} else {
		for i, server := range serverList {
			if server.Id == settings.Id {
				serverList[i] = *settings
				break
			}
		}
	}

	return nil
}

func setConfigVersion(settings *ServerSettings) {
	settings.Configuration.ConfigVersion = configVersion
	settings.Settings.ConfigVersion = configVersion
	settings.Event.ConfigVersion = configVersion
}

func getConfigDirectoryAndID(id int) (string, int, error) {
	// create new ID for new server or use existing one
	if id == 0 {
		id = int(time.Now().Unix())
	}

	dir := filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(id))
	err := os.MkdirAll(dir, 0777)

	if err != nil {
		logrus.WithField("err", err).Error("Error creating configuration directory")
	}

	return dir, id, err
}

func saveConfigToFile(config interface{}, dir, name string) error {
	data, err := json.Marshal(config)

	if err != nil {
		logrus.WithField("err", err).Error("Error marshalling server configuration")
		return err
	}

	path := filepath.Join(dir, name)
	logrus.WithField("path", path).Debug("Saving server configuration file")

	if err := ioutil.WriteFile(path, data, 0666); err != nil {
		logrus.WithField("err", err).Error("Error saving server configuration file")
		return err
	}

	return nil
}
