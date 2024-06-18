package config

import (
	"github.com/caarlos0/env/v11"
	"sync"
	"time"
)

const (
	Local      string = "local"
	Production        = "production"
)

type Config struct {
	IsDebug  bool   `env:"DEBUG" envDefault:"false"`
	Env      string `env:"ENV" envDefault:"production"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	Listen   struct {
		Host           string        `env:"HOST" envDefault:"0.0.0.0"`
		Port           string        `env:"PORT" envDefault:"8080"`
		RequestTimeout time.Duration `env:"REQUEST_TIMEOUT" envDefault:"5s"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := env.Parse(instance); err != nil {
			panic(err)
		}

		parsedCfg, err := env.ParseAs[Config]()
		if err != nil {
			panic(err)
		}

		instance = &parsedCfg
	})

	return instance
}
