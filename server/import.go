package server

import (
	"encoding/json"
	"io"

	"github.com/sirupsen/logrus"

	"golang.org/x/text/transform"
)

func ImportServer(configuration, settings, event, eventRules, entrylist, bop, assistRules io.ReadSeeker) error {
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

func importFile(reader io.ReadSeeker, config interface{}, filename string) error {
	r := transform.NewReader(reader, utf16Encoding.NewDecoder().Transformer)

	if err := json.NewDecoder(r).Decode(config); err != nil {
		if _, er := reader.Seek(0, io.SeekStart); er != nil {
			return er
		}

		// trying decode non ut16 content
		if err2 := json.NewDecoder(reader).Decode(config); err2 != nil {
			logrus.WithError(err).
				WithField("file", filename).
				WithField("err_non_utf16", err2).
				Error("Error unmarshalling configuration file JSON on import")
			return err
		} else {
			return nil
		}
	}

	return nil
}
