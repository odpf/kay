package config

import (
	"errors"

	"github.com/odpf/kay/pkg/pgsx"
	"github.com/odpf/salt/config"
)

type Config struct {
	// configuration version
	Version int         `yaml:"version"`
	Log     string      `yaml:"log"`
	App     AppConfig   `yaml:"app"`
	DB      pgsx.Config `yaml:"db"`
}

type AppConfig struct {
	Port int    `yaml:"port" mapstructure:"port" default:"8080"`
	Host string `yaml:"host" mapstructure:"host" default:"127.0.0.1"`
}

func Load(configFile string) (*Config, error) {
	conf := &Config{}

	var options = []config.LoaderOption{config.WithFile(configFile)}

	l := config.NewLoader(options...)
	if err := l.Load(conf); err != nil {
		if errors.As(err, &config.ConfigFileNotFoundError{}) {
			return nil, errors.New("config file not found")
		} else {
			return nil, err
		}
	}
	return conf, nil
}
