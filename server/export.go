package server

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func ExportServer(id int, withPasswords bool) ([]byte, error) {
	server := GetServerById(id, withPasswords)

	if server == nil {
		return nil, ServerNotFound
	}

	buf := new(bytes.Buffer)
	archive := zip.NewWriter(buf)

	if err := addFileToZip(archive, server.Configuration, configurationJsonName); err != nil {
		return nil, err
	}

	if err := addFileToZip(archive, server.Settings, settingsJsonName); err != nil {
		return nil, err
	}

	if err := addFileToZip(archive, server.Event, eventJsonName); err != nil {
		return nil, err
	}

	if err := addFileToZip(archive, server.Entrylist, entrylistJsonName); err != nil {
		return nil, err
	}

	if err := archive.Close(); err != nil {
		logrus.WithError(err).Error("Error closing zip file")
		return nil, err
	}

	return buf.Bytes(), nil
}

func addFileToZip(archive *zip.Writer, config interface{}, filename string) error {
	configJson, err := json.Marshal(config)

	if err != nil {
		logrus.WithError(err).Error("Error marshalling server settings")
		return err
	}

	file, err := archive.Create(filename)

	if err != nil {
		logrus.WithError(err).Error("Error creating zip file")
		return err
	}

	if _, err := file.Write(configJson); err != nil {
		logrus.WithError(err).Error("Error writing zip file")
		return err
	}

	return nil
}
