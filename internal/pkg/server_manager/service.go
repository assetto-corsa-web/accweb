package server_manager

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server"
)

var (
	ErrCantCreateConfigDir = errors.New("cant create accweb config dir")
	ErrServerNotFound      = errors.New("server not found")
)

type Config struct {
	ConfigBaseDir   string
	AccServerPath   string
	AccServerExe    string
	AccServerMd5Sum string
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

func (s *Service) GetAccServerExeMd5Sum() error {
	sum, err := helper.CheckMd5Sum(path.Join(s.config.AccServerPath, s.config.AccServerExe))
	if err != nil {
		return err
	}

	if s.config.AccServerMd5Sum != sum {
		s.config.AccServerMd5Sum = sum
	}

	return nil
}

func (s *Service) UpdateServersServerExeFile() error {
	for _, srv := range s.servers {
		if srv.Cfg.Md5Sum == s.config.AccServerMd5Sum {
			continue
		}

		if err := srv.UpdateAccServerExe(path.Join(s.config.AccServerPath, s.config.AccServerExe)); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) Bootstrap() error {
	if err := s.GetAccServerExeMd5Sum(); err != nil {
		return err
	}

	if err := s.LoadAll(); err != nil {
		return err
	}

	if err := s.UpdateServersServerExeFile(); err != nil {
		return err
	}

	if err := s.AutoStart(); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetServerByID(id string) (*server.Server, error) {
	if srv, ok := s.servers[id]; ok {
		return srv, nil
	}
	return nil, ErrServerNotFound
}

func (s *Service) Create(accConfig *server.AccConfigFiles) (*server.Server, error) {
	id := string(time.Now().Unix())
	baseDir := path.Join(s.config.ConfigBaseDir, id)

	if err := helper.CreateIfNotExists(baseDir, 0755); err != nil {
		return nil, err
	}

	srv := server.Server{
		Path: baseDir,
		Cfg: server.AccWebConfigJson{
			ID:        id,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		AccCfg: *accConfig,
	}

	if err := srv.UpdateAccServerExe(path.Join(s.config.AccServerPath, s.config.AccServerExe)); err != nil {
		return nil, err
	}

	if err := srv.Save(); err != nil {
		return nil, err
	}

	s.servers[id] = &srv

	return &srv, nil
}

func (s *Service) Delete(id string) error {
	srv, err := s.GetServerByID(id)
	if err != nil {
		return err
	}

	if err := srv.Stop(); err != nil {
		return nil
	}

	if err := os.RemoveAll(srv.Path); err != nil {
		return err
	}

	delete(s.servers, id)

	return nil
}

func (s *Service) Duplicate(srcId string) (*server.Server, error) {
	srcSrv, err := s.GetServerByID(srcId)
	if err != nil {
		return nil, err
	}

	cfg := srcSrv.AccCfg
	cfg.Settings.ServerName += " (COPY)"

	return s.Create(&cfg)
}

func (s *Service) Export() {

}
