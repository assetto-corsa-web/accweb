package server_manager

import (
	"errors"
	"io/ioutil"
	"path"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server"
)

var (
	ErrCantCreateConfigDir = errors.New("cant create accweb config dir")
)

type Config struct {
	ConfigBaseDir string
	AccServerPath string
	AccServerExe  string
}

type Service struct {
	config  *Config
	servers map[string]*server.Server
}

// LoadAll .
func (s *Service) LoadAll() error {
	if err := helper.CreateIfNotExists(s.config.ConfigBaseDir, 0755); err != nil {
		return helper.WrapErrors(ErrCantCreateConfigDir, err)
	}

	dir, err := ioutil.ReadDir(s.config.ConfigBaseDir)
	if err != nil {
		return err
	}

	// reset servers attribute
	s.servers = make(map[string]*server.Server, len(dir))

	for _, entry := range dir {
		if !entry.IsDir() {
			continue
		}

		srv, err := server.LoadServerFromPath(path.Join(s.config.ConfigBaseDir, entry.Name()))
		if err != nil {
			return err
		}

		s.servers[srv.GetID()] = srv
	}

	return nil
}

func (s *Service) AutoStart() error {
	for _, s := range s.servers {
		if !s.Cfg.AutoStart {
			continue
		}

		if err := s.Start(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) Bootstrap() error {
	if err := s.LoadAll(); err != nil {
		return err
	}

	if err := s.AutoStart(); err != nil {
		return err
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
