package instance

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
)

const (
	accDedicatedServerFile = "accServer.exe"
	accCfgDir              = "cfg"
	accServerLogDir        = "log"
	accServerLogFile       = "server.log"
)

var (
	ErrServerCantBeRunning = errors.New("server instance cant be running to perform this action")
	ErrServerDirIsInvalid  = errors.New("server directory is invalid")
)

type Instance struct {
	Path   string
	Cfg    AccWebConfigJson
	AccCfg AccConfigFiles

	cmd *exec.Cmd
}

func (s *Instance) GetID() string {
	return s.Cfg.ID
}

func (s *Instance) Start() error {
	if s.IsRunning() {
		return ErrServerCantBeRunning
	}

	if err := s.CheckDirectory(); err != nil {
		return err
	}

	// Copy config files to cfg dir
	if err := helper.CreateIfNotExists(path.Join(s.Path, accCfgDir), 0755); err != nil {
		return err
	}

	fileList := []string{
		configurationJsonName,
		settingsJsonName,
		eventJsonName,
		eventRulesJsonName,
		entrylistJsonName,
		bopJsonName,
		assistRulesJsonName,
	}

	for _, name := range fileList {
		if err := helper.Copy(path.Join(s.Path, name), path.Join(s.Path, accCfgDir, name)); err != nil {
			return err
		}
	}

	command := "." + string(filepath.Separator) + accDedicatedServerFile
	var args []string

	if runtime.GOOS == "linux" {
		command = "wine"
		args = []string{"accDedicatedServerFile"}
	}

	cmd := exec.Command(command, args...)
	cmd.Dir = s.Path

	s.cmd = cmd

	if err := s.cmd.Start(); err != nil {
		return err
	}

	logrus.WithField("server_id", s.GetID()).WithField("pid", s.GetProcessID()).Info("acc server started")

	go func(cmd *exec.Cmd) {
		// wait for shutdown or crash
		if err := cmd.Wait(); err != nil {
			logrus.WithError(err).Error("Error when server stopped")
		}
	}(cmd)

	return nil
}

func (s *Instance) Stop() error {
	if !s.IsRunning() {
		return nil
	}

	if err := s.cmd.Process.Signal(os.Interrupt); err != nil {
		if err := s.cmd.Process.Kill(); err != nil {
			return err
		}
	}
	logrus.WithField("server_id", s.GetID()).Info("acc server stopped")

	return nil
}

func (s *Instance) GetProcessID() int {
	if s.IsRunning() {
		return s.cmd.Process.Pid
	}

	return 0
}

func (s *Instance) Save() error {
	if s.IsRunning() {
		return ErrServerCantBeRunning
	}

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

func (s *Instance) CheckDirectory() error {
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

func (s *Instance) CheckServerExeMd5Sum() (bool, error) {
	sum, err := helper.CheckMd5Sum(path.Join(s.Path, accDedicatedServerFile))
	if err != nil {
		return false, err
	}

	r := false

	if s.Cfg.Md5Sum != sum {
		s.Cfg.Md5Sum = sum
		s.Cfg.SetUpdateAt()
		r = true
	}

	return r, nil
}

func (s *Instance) UpdateAccServerExe(srcFile string) (bool, error) {
	if s.IsRunning() {
		return false, ErrServerCantBeRunning
	}

	localFile := path.Join(s.Path, accDedicatedServerFile)

	if helper.Exists(localFile) {
		_ = os.Remove(localFile)
	}

	if err := helper.Copy(srcFile, localFile); err != nil {
		return false, err
	}

	if err := os.Chmod(localFile, 0755); err != nil {
		return false, err
	}

	return s.CheckServerExeMd5Sum()
}

func (s *Instance) IsRunning() bool {
	return s.cmd != nil && s.cmd.Process != nil && s.cmd.Process.Pid > 0 && s.cmd.ProcessState == nil
}

func (s *Instance) GetAccServerLogs() ([]byte, error) {
	logFilePath := path.Join(s.Path, accServerLogDir, accServerLogFile)
	if !helper.Exists(logFilePath) {
		return nil, errors.New("server log file doesn't exists")
	}

	return os.ReadFile(logFilePath)
}

func (s *Instance) ExportConfigFilesToZip() ([]byte, error) {
	fileList := map[string]interface{}{
		configurationJsonName: &s.AccCfg.Configuration,
		settingsJsonName:      &s.AccCfg.Settings,
		eventJsonName:         &s.AccCfg.Event,
		eventRulesJsonName:    &s.AccCfg.EventRules,
		entrylistJsonName:     &s.AccCfg.Entrylist,
		bopJsonName:           &s.AccCfg.Bop,
		assistRulesJsonName:   &s.AccCfg.AssistRules,
	}

	buf := new(bytes.Buffer)
	archive := zip.NewWriter(buf)

	for filename, obj := range fileList {
		l := logrus.WithField("filename", filename)

		contentData, err := helper.Encode(obj)
		if err != nil {
			l.WithError(err).Error("Error encoding config information")
			return nil, err
		}

		file, err := archive.Create(filename)
		if err != nil {
			l.WithError(err).Error("Error creating zip file")
			return nil, err
		}

		if _, err := file.Write(contentData); err != nil {
			l.WithError(err).Error("Error writing zip file")
			return nil, err
		}
	}

	if err := archive.Close(); err != nil {
		logrus.WithError(err).Error("Error closing zip file")
		return nil, err
	}

	return buf.Bytes(), nil
}
