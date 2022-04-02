package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName         string `envconfig:"APP_NAME" default:"{{cookiecutter.project_name}}"`
	Version         string `default:"1.0.0"`
	Debug           bool   `envconfig:"DEBUG" default:"false"`
	Port            int    `envconfig:"PORT" default:"8000"`
	QiscusAppID     string `envconfig:"QISCUS_APP_ID" required:"true"`
	QiscusSecretKey string `envconfig:"QISCUS_SECRET_KEY" required:"true"`
	QismoBaseURL    string `envconfig:"QISMO_BASE_URL" default:"https://multichannel.qiscus.com"`
	QiscusBaseURL   string `envconfig:"QISCUS_BASE_URL" default:"https://api.qiscus.com"`
}

var AppConfig Config

func DefineConfig() (*Config, error) {
	if err := envconfig.Process("", &AppConfig); err != nil {
		return nil, err
	}

	return &AppConfig, nil
}

func (c *Config) GetHTTPUserAgent() string {
	return fmt.Sprintf("%s-v%s", c.AppName, c.Version)
}
