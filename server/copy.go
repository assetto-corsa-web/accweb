package server

import (
	"errors"
)

func CopyServerSettings(id int) error {
	server := GetServerById(id)

	if server == nil {
		return errors.New("Server not found")
	}

	server.Id = 0
	server.Settings.ServerName += " (copy)"

	if err := SaveServerSettings(server); err != nil {
		return err
	}

	return nil
}
