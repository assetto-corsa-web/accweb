package server

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

func DeleteServer(id int) error {
	if err := StopServer(id); err != nil {
		return err
	}

	server := GetServerById(id)

	if server == nil {
		return errors.New("Server not found")
	}

	if err := os.RemoveAll(filepath.Join(os.Getenv("ACCWEB_CONFIG_PATH"), strconv.Itoa(server.Id))); err != nil {
		return err
	}

	for i, s := range serverList {
		if s.Id == id {
			serverList = append(serverList[:i], serverList[i+1:]...)
			break
		}
	}

	return nil
}
