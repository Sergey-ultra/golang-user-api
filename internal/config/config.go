package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"todo/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true" json:"is_debug,omitempty"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port" json:"type,omitempty"`
		BindIp string `yaml:"bind_ip" env-default:"127.0.0.1" json:"bind_ip,omitempty"`
		Port   string `yaml:"port" env-default:"8080" json:"port,omitempty"`
	} `yaml:"listen" json:"listen"`
	MongoDB struct {
		Host       string `yaml:"host" json:"host,omitempty"`
		Port       string `yaml:"port" json:"port,omitempty"`
		Database   string `yaml:"database" json:"database,omitempty"`
		AuthDb     string `yaml:"auth_db" json:"auth_db,omitempty"`
		Username   string `yaml:"username" json:"username,omitempty"`
		Password   string `yaml:"password" json:"password,omitempty"`
		Collection string `yaml:"collection" json:"collection,omitempty"`
	} `yaml:"mongodb" json:"mongodb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
