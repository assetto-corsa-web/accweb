package server_manager

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
	"github.com/assetto-corsa-web/accweb/internal/pkg/instance"
)

var (
	ErrCantCreateConfigDir = errors.New("cant create accweb config dir")
	ErrServerNotFound      = errors.New("server not found")
	ErrServerAlreadyExists = errors.New("server already exists")
)

type Config struct {
	ConfigBaseDir   string
	AccServerPath   string
	AccServerExe    string
	AccServerMd5Sum string
}

func (c Config) AccServerFullPath() string {
	return path.Join(c.AccServerPath, c.AccServerExe)
}

type Service struct {
	config  *Config
	servers map[string]*instance.Instance
	lock    sync.Mutex
}

func New(baseDir, accServerPath, accServerExe string) *Service {
	return &Service{config: &Config{
		ConfigBaseDir: baseDir,
		AccServerPath: accServerPath,
		AccServerExe:  accServerExe,
	}}
}

// LoadAll .
func (s *Service) LoadAll() error {
	if err := helper.CreateIfNotExists(s.config.ConfigBaseDir, 0755); err != nil {
		return helper.WrapErrors(ErrCantCreateConfigDir, err)
	}

	dir, err := os.ReadDir(s.config.ConfigBaseDir)
	if err != nil {
		return err
	}

	//filter out all tmp folders used by google drive
	var newdirs []fs.DirEntry

	for index := range dir {
		logrus.Info("accweb1: " + dir[index].Name() + " - " + fmt.Sprint(index))
		if dir[index].IsDir() && !strings.Contains(dir[index].Name(), ".tmp.") {
			logrus.Info("accweb2: " + dir[index].Name())
			newdirs = append(newdirs, dir[index])
		}
	}
	dir = newdirs

	// reset servers attribute
	s.servers = make(map[string]*instance.Instance, len(dir))

	for _, entry := range dir {
		if !entry.IsDir() {
			continue
		}

		srv, err := instance.LoadServerFromPath(path.Join(s.config.ConfigBaseDir, entry.Name()))
		if err != nil {
			return err
		}

		s.servers[srv.GetID()] = srv
	}

	return nil
}

func (s *Service) AutoStart() error {
	for _, s := range s.servers {
		if !s.Cfg.Settings.AutoStart {
			continue
		}

		if err := s.Start(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) StopAll() error {
	var wg sync.WaitGroup
	for _, s := range s.servers {
		wg.Add(1)
		go func(s *instance.Instance, wg *sync.WaitGroup) {
			defer wg.Done()
			if err := s.Stop(); err != nil {
				logrus.WithError(err).Error("server stopped with an error")
			}
		}(s, &wg)
	}
	wg.Wait()

	return nil
}

func (s *Service) GetAccServerExeMd5Sum() error {
	sum, err := helper.CheckMd5Sum(s.config.AccServerFullPath())
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
		if err := s.updateAccServerExeIfDifferent(srv); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) updateAccServerExeIfDifferent(srv *instance.Instance) error {
	if srv.Cfg.Md5Sum == s.config.AccServerMd5Sum {
		return nil
	}

	if ok, err := srv.UpdateAccServerExe(s.config.AccServerFullPath()); err != nil {
		return err
	} else if ok {
		srv.Cfg.SetUpdateAt()

		if err := srv.Save(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) Bootstrap() error {
	if err := s.GetAccServerExeMd5Sum(); err != nil {
		return err
	}
	logrus.WithField("md5sum", s.config.AccServerMd5Sum).Info("boot: calculating acc dedicated server md5sum")

	if err := s.LoadAll(); err != nil {
		return err
	}
	logrus.WithField("total", len(s.servers)).Info("boot: loaded all configured acc servers")

	logrus.Info("boot: checking for outdated acc server instances")
	if err := s.UpdateServersServerExeFile(); err != nil {
		return err
	}

	logrus.Info("boot: auto starting acc server instances")
	if err := s.AutoStart(); err != nil {
		return err
	}

	return nil
}

func (s *Service) addServer(srv *instance.Instance) error {
	if _, ok := s.servers[srv.GetID()]; ok {
		return ErrServerAlreadyExists
	}

	s.lock.Lock()
	s.servers[srv.GetID()] = srv
	s.lock.Unlock()

	return nil
}

func (s *Service) delServer(srv *instance.Instance) error {
	if _, ok := s.servers[srv.GetID()]; !ok {
		return ErrServerNotFound
	}

	s.lock.Lock()
	delete(s.servers, srv.GetID())
	s.lock.Unlock()

	return nil
}

func (s *Service) GetServers() map[string]*instance.Instance {
	return s.servers
}

func (s *Service) GetServerByID(id string) (*instance.Instance, error) {
	if srv, ok := s.servers[id]; ok {
		return srv, nil
	}
	return nil, ErrServerNotFound
}

func (s *Service) Create(accConfig *instance.AccConfigFiles, accWebSettings instance.AccWebSettingsJson) (*instance.Instance, error) {
	id := strconv.FormatInt(time.Now().Unix(), 10)
	baseDir := path.Join(s.config.ConfigBaseDir, id)

	if err := helper.CreateIfNotExists(baseDir, 0755); err != nil {
		return nil, err
	}

	srv := instance.Instance{
		Path: baseDir,
		Cfg: instance.AccWebConfigJson{
			ID:        id,
			Settings:  accWebSettings,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		AccCfg: *accConfig,
		Live:   instance.NewLiveState(),
	}

	if _, err := srv.UpdateAccServerExe(s.config.AccServerFullPath()); err != nil {
		return nil, err
	}

	if err := srv.Save(); err != nil {
		return nil, err
	}

	if err := s.addServer(&srv); err != nil {
		return nil, err
	}

	return &srv, nil
}

func (s *Service) Delete(id string) error {
	srv, err := s.GetServerByID(id)
	if err != nil {
		return err
	}

	if err := srv.Stop(); err != nil {
		return err
	}

	if err := os.RemoveAll(srv.Path); err != nil {
		return err
	}

	if err := s.delServer(srv); err != nil {
		return err
	}

	return nil
}

func (s *Service) Duplicate(srcId string) (*instance.Instance, error) {
	srcSrv, err := s.GetServerByID(srcId)
	if err != nil {
		return nil, err
	}

	cfg := srcSrv.AccCfg
	cfg.Settings.ServerName += " (COPY)"

	return s.Create(&cfg, srcSrv.Cfg.Settings)
}

func (s *Service) Start(id string) error {
	srv, err := s.GetServerByID(id)
	if err != nil {
		return err
	}

	if err := s.GetAccServerExeMd5Sum(); err != nil {
		return err
	}

	if err := s.updateAccServerExeIfDifferent(srv); err != nil {
		return err
	}

	if err := srv.Start(); err != nil {
		return err
	}

	return nil
}
