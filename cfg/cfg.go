package cfg

import (
	"github.com/sirupsen/logrus"
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
	Dev        bool      `yaml:"dev"`
	Loglevel   string    `yaml:"loglevel"`
	ConfigPath string    `yaml:"config_path"`
	Webserver  Webserver `yaml:"webserver"`
	CORS       CORS      `yaml:"cors"`
	Auth       Auth      `yaml:"auth"`
	ACC        ACC       `yaml:"acc"`
}

type Webserver struct {
	Host         string `yaml:"host"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	TLS          bool   `yaml:"tls"`
	Cert         string `yaml:"cert"`
	PrivateKey   string `yaml:"private_key"`
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

type ACC struct {
	ServerPath string `yaml:"server_path"`
	ServerExe  string `yaml:"server_exe"`
}

// Load loads the application config from config.yml.
func Load() {
	loadConfigYml()

	if out, err := yaml.Marshal(config); err == nil {
		logrus.Info("\n" + string(out))
	}
}

func loadConfigYml() {
	logrus.Info("Loading configuration file")
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		logrus.WithField("err", err).Fatal("Error loading configuration file")
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		logrus.WithField("err", err).Fatal("Error loading parsing configuration file")
	}
}

// Get returns the application configuration.
func Get() *Config {
	return &config
}
