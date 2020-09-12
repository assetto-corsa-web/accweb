package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

func ImportServer(configuration, settings, event, eventRules, entrylist, bop, assistRules io.Reader) error {
	server := new(ServerSettings)

	if err := importFile(configuration, &server.Configuration, configurationJsonName); err != nil {
		return err
	}

	if err := importFile(settings, &server.Settings, settingsJsonName); err != nil {
		return err
	}

	if err := importFile(event, &server.Event, eventJsonName); err != nil {
		return err
	}

	if err := importFile(eventRules, &server.EventRules, eventRulesJsonName); err != nil {
		return err
	}

	if err := importFile(entrylist, &server.Entrylist, entrylistJsonName); err != nil {
		return err
	}

	if err := importFile(bop, &server.Bop, bopJsonName); err != nil {
		return err
	}

	if err := importFile(assistRules, &server.AssistRules, assistRulesJsonName); err != nil {
		return err
	}

	if err := SaveServerSettings(server); err != nil {
		return err
	}

	return nil
}

func importFile(reader io.Reader, config interface{}, filename string) error {
	data, err := ioutil.ReadAll(reader)

	if err != nil {
		logrus.WithError(err).WithField("file", filename).Error("Error reading configuration file JSON on import")
		return err
	}

	if err := json.Unmarshal(data, config); err != nil {
		logrus.WithError(err).WithField("file", filename).Error("Error unmarshalling configuration file JSON on import")
		return err
	}

	return nil
}
