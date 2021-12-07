package server

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
)

var (
	ErrServerCantBeRunning = errors.New("server instance cant be running to perform this action")
	ErrServerIsNotRunning  = errors.New("server is not running")
	ErrServerDirIsInvalid  = errors.New("server directory is invalid")
)

type Server struct {
	Path   string
	Cfg    AccWebConfigJson
	AccCfg AccConfigFiles

	cmd *exec.Cmd
}

func (s *Server) GetID() string {
	return s.Cfg.ID
}

func (s *Server) Start() error {
	if s.isRunning() {
		return ErrServerCantBeRunning
	}

	if err := s.CheckDirectory(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	if !s.isRunning() {
		return ErrServerIsNotRunning
	}
	return nil
}

func (s *Server) GetProcessID() int {
	if s.isRunning() {
		return s.cmd.Process.Pid
	}

	return 0
}

func (s *Server) Save() error {
	if s.isRunning() {
		return ErrServerCantBeRunning
	}

	s.Cfg.UpdatedAt = time.Now().UTC()

	fileList := map[string]interface{}{
		accwebConfigJsonName:  &s.Cfg,
		configurationJsonName: &s.AccCfg.Configuration,
		settingsJsonName:      &s.AccCfg.Settings,
		eventJsonName:         &s.AccCfg.Event,
		eventRulesJsonName:    &s.AccCfg.EventRules,
		entrylistJsonName:     &s.AccCfg.Entrylist,
		bopJsonName:           &s.AccCfg.Bop,
		assistRulesJsonName:   &s.AccCfg.AssistRules,
	}

	for filename, cfg := range fileList {
		if err := helper.SaveToPath(s.Path, filename, cfg); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) CheckDirectory() error {
	fileList := []string{
		accwebConfigJsonName,
		configurationJsonName,
		settingsJsonName,
		eventJsonName,
		eventRulesJsonName,
		entrylistJsonName,
		bopJsonName,
		assistRulesJsonName,
		accDedicatedServerFile,
	}

	for _, filename := range fileList {
		p := path.Join(s.Path, filename)
		if !helper.Exists(p) {
			return fmt.Errorf("%w - missing '%s'", ErrServerDirIsInvalid, p)
		}
	}

	return nil
}

func (s *Server) isRunning() bool {
	return s.cmd != nil && s.cmd.Process != nil && s.cmd.Process.Pid != 0
}
