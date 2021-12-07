package server

import (
	"errors"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
)

const (
	accwebConfigJsonName   = "accwebConfig.json"
	configurationJsonName  = "configuration.json"
	settingsJsonName       = "settings.json"
	eventJsonName          = "event.json"
	eventRulesJsonName     = "eventRules.json"
	entrylistJsonName      = "entrylist.json"
	bopJsonName            = "bop.json"
	assistRulesJsonName    = "assistRules.json"
	configVersion          = 1
	accDedicatedServerFile = "accServer.exe"
)

// LoadServerFromPath load the server configuration data based on baseDir and returns a Server instance
func LoadServerFromPath(baseDir string) (*Server, error) {
	s := &Server{Path: baseDir}

	if cfg, err := loadAccWebConfig(baseDir); err != nil {
		return nil, err
	} else {
		s.Cfg = *cfg
	}

	fileList := map[string]interface{}{
		configurationJsonName: &s.AccCfg.Configuration,
		settingsJsonName:      &s.AccCfg.Settings,
		eventJsonName:         &s.AccCfg.Event,
		eventRulesJsonName:    &s.AccCfg.EventRules,
		entrylistJsonName:     &s.AccCfg.Entrylist,
		bopJsonName:           &s.AccCfg.Bop,
		assistRulesJsonName:   &s.AccCfg.AssistRules,
	}

	for filename, cfg := range fileList {
		if err := helper.LoadFromPath(baseDir, filename, cfg); err != nil {
			return nil, err
		}
	}

	setConfigVersion(&s.AccCfg)

	sum, err := helper.CheckMd5Sum(path.Join(baseDir, accDedicatedServerFile))
	if err != nil {
		return nil, err
	}

	if s.Cfg.Md5Sum != sum {
		s.Cfg.Md5Sum = sum
		s.Cfg.UpdatedAt = time.Now().UTC()

		if err := helper.SaveToPath(baseDir, accwebConfigJsonName, s.Cfg); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func loadAccWebConfig(baseDir string) (*AccWebConfigJson, error) {
	var cfg AccWebConfigJson
	if err := helper.LoadFromPath(baseDir, accwebConfigJsonName, &cfg); err != nil {
		// For backward compatibility when the config file don't exist,create a new one
		if errors.Is(err, syscall.ENOENT) {
			cfg = AccWebConfigJson{
				ID:        filepath.Base(baseDir),
				Md5Sum:    "",
				AutoStart: false,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}
		} else {
			return nil, err
		}
	}

	return &cfg, nil
}

func setConfigVersion(settings *AccConfigFiles) {
	settings.Configuration.ConfigVersion = configVersion
	settings.Settings.ConfigVersion = configVersion
	settings.Event.ConfigVersion = configVersion
	settings.EventRules.ConfigVersion = configVersion
	settings.Entrylist.ConfigVersion = configVersion
	settings.Bop.ConfigVersion = configVersion
	settings.AssistRules.ConfigVersion = configVersion
}
