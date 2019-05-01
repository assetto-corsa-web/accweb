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
	configurationJson = "configuration.json"
	settingsJson      = "settings.json"
	eventJson         = "event.json"
	configVersion     = 1
)

type ServerSettings struct {
	Id            int               `json:"id"`
	Configuration ConfigurationJson `json:"basic"`
	Settings      SettingsJson      `json:"settings"`
	Event         EventJson         `json:"event"`
}

type ConfigurationJson struct {
	UdpPort       int `json:"udpPort"`
	TcpPort       int `json:"tcpPort"`
	MaxClients    int `json:"maxClients"`
	ConfigVersion int `json:"configVersion"`
}

type SettingsJson struct {
	ServerName              string `json:"serverName"`
	Password                string `json:"password"`
	AdminPassword           string `json:"adminPassword"`
	TrackMedalsRequirement  int    `json:"trackMedalsRequirement"`
	SafetyRatingRequirement int    `json:"safetyRatingRequirement"`
	ConfigVersion           int    `json:"configVersion"`
}

type EventJson struct {
	Track                     string            `json:"track"`
	EventType                 string            `json:"eventType"`
	PreRaceWaitingTimeSeconds int               `json:"preRaceWaitingTimeSeconds"`
	SessionOverTimeSeconds    int               `json:"sessionOverTimeSeconds"`
	AmbientTemp               int               `json:"ambientTemp"`
	TrackTemp                 int               `json:"trackTemp"`
	CloudLevel                float64           `json:"cloudLevel"`
	Rain                      float64           `json:"rain"`
	WeatherRandomness         int               `json:"weatherRandomness"`
	ConfigVersion             int               `json:"configVersion"`
	Sessions                  []SessionSettings `json:"sessions"`
}

type SessionSettings struct {
	HourOfDay              int    `json:"hourOfDay"`
	DayOfWeekend           int    `json:"dayOfWeekend"`
	TimeMultiplier         int    `json:"timeMultiplier"`
	SessionType            string `json:"sessionType"`
	SessionDurationMinutes int    `json:"sessionDurationMinutes"`
}

func SaveServerSettings(settings *ServerSettings) error {
	setConfigVersion(settings)
	dir, err := getConfigDirectory()

	if err != nil {
		return err
	}

	if err := saveConfigToFile(settings.Configuration, dir, configurationJson); err != nil {
		return err
	}

	if err := saveConfigToFile(settings.Settings, dir, settingsJson); err != nil {
		return err
	}

	if err := saveConfigToFile(settings.Event, dir, eventJson); err != nil {
		return err
	}

	return nil
}

func setConfigVersion(settings *ServerSettings) {
	settings.Configuration.ConfigVersion = configVersion
	settings.Settings.ConfigVersion = configVersion
	settings.Event.ConfigVersion = configVersion
}

func getConfigDirectory() (string, error) {
	dir := filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(int(time.Now().Unix())))
	err := os.MkdirAll(dir, 0777)

	if err != nil {
		logrus.WithField("err", err).Error("Error creating configuration directory")
	}

	return dir, err
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
