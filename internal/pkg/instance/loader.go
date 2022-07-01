package instance

import (
	"errors"
	"math"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
)

var DefaultCoreAffinity = uint(math.Pow(2, float64(runtime.NumCPU()))) - 1

// LoadServerFromPath load the server configuration data based on baseDir and returns a Instance instance
func LoadServerFromPath(baseDir string) (*Instance, error) {
	s := &Instance{Path: baseDir, Live: NewLiveState()}

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
				Settings:  AccWebSettingsJson{},
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}
		} else {
			return nil, err
		}
	}

	changes := false

	// For backward compatibility
	if cfg.AutoStart && !cfg.Settings.AutoStart {
		cfg.AutoStart = false
		cfg.Settings.AutoStart = true
		changes = true
	}

	if cfg.Settings.AdvWindowsCfg != nil {
		if cfg.Settings.AdvWindowsCfg.CoreAffinity == 0 || cfg.Settings.AdvWindowsCfg.CoreAffinity > DefaultCoreAffinity {
			cfg.Settings.AdvWindowsCfg.CoreAffinity = DefaultCoreAffinity
			changes = true
		}
	}

	if changes {
		if err := helper.SaveToPath(baseDir, accwebConfigJsonName, cfg); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func SetConfigVersion(settings *AccConfigFiles) {
	settings.Configuration.ConfigVersion = configVersion
	settings.Settings.ConfigVersion = configVersion
	settings.Event.ConfigVersion = configVersion
	settings.EventRules.ConfigVersion = configVersion
	settings.Entrylist.ConfigVersion = configVersion
	settings.Bop.ConfigVersion = configVersion
	settings.AssistRules.ConfigVersion = configVersion
}
