package cfg

import (
	"os"
	"path"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/sirupsen/logrus"
)

var logLevel = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

var skipWine bool

type Config struct {
	file       string
	Dev        bool      `yaml:"dev" json:"dev"`
	SkipWine   bool      `yaml:"skip_wine" json:"skip_wine"`
	Loglevel   string    `yaml:"loglevel" json:"loglevel"`
	ConfigPath string    `yaml:"config_path" json:"config_path"`
	Webserver  Webserver `yaml:"webserver" json:"webserver"`
	CORS       CORS      `yaml:"cors" json:"cors"`
	Auth       Auth      `yaml:"auth" json:"auth"`
	ACC        ACC       `yaml:"acc" json:"acc"`
	Log        Log       `yaml:"log" json:"log"`
	Callback   Callback  `yaml:"callback" json:"callback"`
}

func (c Config) AccServerFullPath() string {
	return path.Join(c.ACC.ServerPath, c.ACC.ServerExe)
}

func (c Config) Save() error {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	return os.WriteFile(c.file, data, 0644)
}

type Webserver struct {
	Host       string `yaml:"host" json:"host"`
	TLS        bool   `yaml:"tls" json:"tls"`
	Cert       string `yaml:"cert" json:"cert"`
	PrivateKey string `yaml:"private_key" json:"private_key"`
}

type CORS struct {
	Origins  string `yaml:"origins" json:"origins"`
	Loglevel string `yaml:"loglevel" json:"loglevel"`
}

type Auth struct {
	PublicKeyPath     string         `yaml:"public_key_path"`
	PrivateKeyPath    string         `yaml:"private_key_path"`
	AdminPassword     string         `yaml:"admin_password"`
	ModeratorPassword string         `yaml:"moderator_password"`
	ReadOnlyPassword  string         `yaml:"read_only_password"`
	Users             Users          `yaml:"users"`
	Timeout           *time.Duration `yaml:"timeout"`
}

type ACC struct {
	ServerPath string `yaml:"server_path"`
	ServerExe  string `yaml:"server_exe"`
}

type Log struct {
	WithTimestamp bool `yaml:"with_timestamp"`
}

type Callback struct {
	Enabled bool             `yaml:"enabled"`
	Timeout *time.Duration   `yaml:"timeout"`
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

	config.file = file

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
