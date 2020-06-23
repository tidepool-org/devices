package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DevicesConfigFilename string `envconfig:"TIDEPOOL_DEVICES_CONFIG_FILENAME" required:"true"`
	ServerPort            uint16 `envconfig:"TIDEPOOL_SERVER_PORT" default:"50051" required:"true"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadFromEnv() error {
	return envconfig.Process("", c)
}
