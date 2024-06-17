package cfg

import (
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var logLevel = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

var skipWine bool

type Config struct {
	Dev        bool      `yaml:"dev"`
	SkipWine   bool      `yaml:"skip_wine"`
	Loglevel   string    `yaml:"loglevel"`
	ConfigPath string    `yaml:"config_path"`
	Webserver  Webserver `yaml:"webserver"`
	CORS       CORS      `yaml:"cors"`
	Auth       Auth      `yaml:"auth"`
	ACC        ACC       `yaml:"acc"`
	Log        Log       `yaml:"log"`
	Callback   Callback  `yaml:"callback"`
}

func (c Config) AccServerFullPath() string {
	return path.Join(c.ACC.ServerPath, c.ACC.ServerExe)
}

type Webserver struct {
	Host       string `yaml:"host"`
	TLS        bool   `yaml:"tls"`
	Cert       string `yaml:"cert"`
	PrivateKey string `yaml:"private_key"`
}

type CORS struct {
	Origins  string `yaml:"origins"`
	Loglevel string `yaml:"loglevel"`
}

type Auth struct {
	PublicKeyPath     string         `yaml:"public_key_path"`
	PrivateKeyPath    string         `yaml:"private_key_path"`
	AdminPassword     string         `yaml:"admin_password"`
	ModeratorPassword string         `yaml:"moderator_password"`
	ReadOnlyPassword  string         `yaml:"read_only_password"`
	Timeout           *time.Duration `yaml:"timeout"`
}

type ACC struct {
	ServerPath string `yaml:"server_path"`
	ServerExe  string `yaml:"server_exe"`
}

type Log struct {
	WithTimestamp bool `yaml:"with_timestamp"`
}

type CallbackOld struct {
	Enabled   bool              `yaml:"enabled"`
	Urls      []string          `yaml:"urls"`
	Headers   map[string]string `yaml:"headers"`
	AllEvents bool              `yaml:"all_events"`
	Events    []string          `yaml:"events"`
}

type Callback struct {
	Enabled bool             `yaml:"enabled"`
	Clients []CallbackClient `yaml:"clients"`
}

type CallbackClient struct {
	Enabled *bool             `yaml:"enabled"`
	Url     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers"`
	Events  []string          `yaml:"events"`
}

// Load loads the application config from config.yml.
func Load(file string) *Config {
	data, err := os.ReadFile(file)
	if err != nil {
		logrus.WithError(err).Fatal("Error loading configuration file")
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		logrus.WithError(err).Fatal("Error loading parsing configuration file")
	}

	if l, ok := logLevel[config.Loglevel]; ok {
		logrus.SetLevel(l)
	}

	if config.Auth.PrivateKeyPath == "" {
		config.Auth.PrivateKeyPath = "secrets/token.private"
	}

	if config.Auth.PublicKeyPath == "" {
		config.Auth.PublicKeyPath = "secrets/token.public"
	}

	if config.Auth.Timeout == nil {
		m := 20 * time.Minute
		config.Auth.Timeout = &m
	}

	skipWine = config.SkipWine

	return &config
}

func SkipWine() bool {
	return skipWine
}
