package config

import (
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type DevicesConfig struct {
	Devices *Devices `yaml:"devices" validate:"required"`
}

type Devices struct {
	CGMs  []*CGM  `yaml:"cgms" validate:"required"`
	Pumps []*Pump `yaml:"pumps" validate:"required"`
}

type Device struct {
	ID            string   `yaml:"id" validate:"required,uuid"`
	DisplayName   string   `yaml:"display_name" validate:"required"`
	Manufacturers []string `yaml:"manufacturers" validate:"required"`
	Model         string   `yaml:"model" validate:"required"`
}

type CGM struct {
	Device `yaml:",inline"`
}

type Pump struct {
	Device `yaml:",inline"`
}

func NewDevicesConfig() *DevicesConfig {
	return &DevicesConfig{}
}

func (c *DevicesConfig) LoadFromFile(filepath string) error {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(f, c)
}

func (c *DevicesConfig) Validate(validate *validator.Validate) error {
	return validate.Struct(c)
}
