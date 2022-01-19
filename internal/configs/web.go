package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type LogFormat string

type WebConfig struct {
	LogLevel  logrus.Level
	LogFormat LogFormat

	HttpHost string `env:"HTTP_HOST" envDefault:"0.0.0.0"`
	HttpPort int    `env:"HTTP_PORT" envDefault:"8080"`
}

func NewWebConfig() (*WebConfig, error) {
	config := new(WebConfig)
	var err error

	err = env.Parse(config)
	if err != nil {
		return nil, err
	}

	config.LogLevel, err = logrus.ParseLevel("debug")
	if err != nil {
		return nil, fmt.Errorf("unknown log level")
	}

	config.LogFormat = "json"
	if err != nil {
		return nil, fmt.Errorf("unknown log level")
	}

	return config, err
}
