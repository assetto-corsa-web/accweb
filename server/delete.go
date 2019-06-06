package server

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
)

func DeleteServer(id int) error {
	if err := StopServer(id); err != nil {
		return err
	}

	server := GetServerById(id, true)

	if server == nil {
		return ServerNotFound
	}

	if err := os.RemoveAll(filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(server.Id))); err != nil {
		logrus.WithError(err).Error("Error deleting server directory")
		return err
	}

	removeServer(id)
	return nil
}
