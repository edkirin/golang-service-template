package cfg

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type configStruct struct {
	Site struct {
		Url string `yaml:"url"`
	}
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
		Name string `yaml:"name"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
	Application struct {
		LogPath      *string `yaml:"logPath"`
		DebugSQL     bool    `yaml:"debugSQL"`
		IsProduction bool    `yaml:"isProduction"`
	} `yaml:"application"`
}

var Config configStruct
var log *logrus.Logger

func processError(err error) {
	log.Error("Config file error: " + err.Error())
	os.Exit(2)
}

func readFile(cfgFile string, cfg *configStruct) {
	f, err := os.Open(cfgFile)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *configStruct) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func InitConfig(configFile string, logger *logrus.Logger) {
	log = logger
	readFile(configFile, &Config)
	readEnv(&Config)
}
