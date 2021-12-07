package server

import (
	"errors"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/file"
)

const (
	accwebConfigJsonName  = "accwebConfig.json"
	configurationJsonName = "configuration.json"
	settingsJsonName      = "settings.json"
	eventJsonName         = "event.json"
	eventRulesJsonName    = "eventRules.json"
	entrylistJsonName     = "entrylist.json"
	bopJsonName           = "bop.json"
	assistRulesJsonName   = "assistRules.json"
	configVersion         = 1
)

// LoadServerFromPath load the server configuration data based on baseDir and returns a Server instance
func LoadServerFromPath(baseDir string) (*Server, error) {
	s := &Server{Path: baseDir}

	if cfg, err := loadAccWebConfig(baseDir); err != nil {
		return nil, err
	} else {
		s.Cfg = *cfg
		s.ID = s.Cfg.ID
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
		if err := file.LoadFromPath(baseDir, filename, cfg); err != nil {
			return nil, err
		}
	}

	setConfigVersion(&s.AccCfg)

	// TODO check current md5Sum of the exe

	return s, nil
}

func loadAccWebConfig(baseDir string) (*AccWebConfigJson, error) {
	var cfg AccWebConfigJson
	if err := file.LoadFromPath(baseDir, accwebConfigJsonName, &cfg); err != nil {
		// For backward compatibility when the config file don't exist,create a new one
		if errors.Is(err, syscall.ENOENT) {
			id, err := strconv.Atoi(filepath.Base(baseDir))
			if err != nil {
				return nil, err
			}

			cfg = AccWebConfigJson{
				ID:        id,
				Md5Sum:    "", // TODO
				AutoStart: false,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}

			// TODO: save the file
		} else {
			return nil, err
		}
	}

	return &cfg, nil
}

func setConfigVersion(settings *accConfigFiles) {
	settings.Configuration.ConfigVersion = configVersion
	settings.Settings.ConfigVersion = configVersion
	settings.Event.ConfigVersion = configVersion
	settings.EventRules.ConfigVersion = configVersion
	settings.Entrylist.ConfigVersion = configVersion
	settings.Bop.ConfigVersion = configVersion
	settings.AssistRules.ConfigVersion = configVersion
}
