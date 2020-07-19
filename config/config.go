package config

import (
	"github.com/emvi/logbuch"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	configFile = "config.yml"
)

var (
	config Config
)

type Config struct {
	Loglevel string `yaml:"loglevel"`
	Server   Server `yaml:"server"`
	CORS     CORS   `yaml:"cors"`
	Auth     Auth   `yaml:"auth"`
}

type Server struct {
	Host         string `yaml:"host"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	TLS          bool   `yaml:"tls"`
	Cert         string `yaml:"cert"`
	PrivateKey   string `yaml:"private_key"`
	HotReload    bool   `yaml:"hot_reload"` // to reload templates in development mode
}

type CORS struct {
	Origins  string `yaml:"origins"`
	Loglevel string `yaml:"loglevel"`
}

type Auth struct {
	PublicKeyPath     string `yaml:"public_key_path"`
	PrivateKeyPath    string `yaml:"private_key_path"`
	AdminPassword     string `yaml:"admin_password"`
	ModeratorPassword string `yaml:"moderator_password"`
	ReadOnlyPassword  string `yaml:"read_only_password"`
}

// Load loads the application configuration.
func Load() {
	loadConfigYml()

	if out, err := yaml.Marshal(config); err == nil {
		logbuch.Info("\n" + string(out))
	}
}

// Get returns the application configuration.
func Get() *Config {
	return &config
}

func loadConfigYml() {
	logbuch.Info("Loading configuration file")
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		logbuch.Fatal("Error loading configuration file", logbuch.Fields{"err": err})
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		logbuch.Fatal("Error parsing configuration file", logbuch.Fields{"err": err})
	}
}
