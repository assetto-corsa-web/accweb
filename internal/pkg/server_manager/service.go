package server_manager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"syscall"

	"github.com/assetto-corsa-web/accweb/internal/pkg/server"
)

var (
	ErrCantCreateConfigDir = errors.New("cant create accweb config dir")
)

type Config struct {
	ConfigDir string
}

type Service struct {
	config  *Config
	servers map[int]*server.Server
}

// LoadAll .
func (s *Service) LoadAll() error {
	dir, err := ioutil.ReadDir(s.config.ConfigDir)
	if err != nil {
		if errors.Is(err, syscall.ENOENT) {
			if err := os.MkdirAll(s.config.ConfigDir, 0755); err != nil {
				return fmt.Errorf("%e (err: %w)", ErrCantCreateConfigDir, err)
			}
		}
		return err
	}

	for _, entry := range dir {
		if !entry.IsDir() {
			continue
		}

		srv, err := server.LoadServerFromPath(path.Join(s.config.ConfigDir, entry.Name()))
		if err != nil {
			return err
		}

		s.servers[srv.ID] = srv
	}

	return nil
}

func (s *Service) Create() {

}

func (s *Service) Delete() {

}

func (s *Service) Duplicate() {

}

func (s *Service) Export() {

}
