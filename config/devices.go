package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

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
	Device     `yaml:",inline"`
	GuardRails GuardRails `yaml:"guard_rails" validate:"required"`
}

type GuardRails struct {
	GlucoseSafetyLimit         GuardRail `yaml:"glucose_safety_limit" validate:"required"`
	InsulinSensitivity         GuardRail `yaml:"insulin_sensitivity" validate:"required"`
	BasalRates                 GuardRail `yaml:"basal_rates" validate:"required"`
	CarbohydrateRatio          GuardRail `yaml:"carbohydrate_ratio" validate:"required"`
	BasalRateMaximum           GuardRail `yaml:"basal_rate_maximum" validate:"required"`
	BolusAmountMaximum         GuardRail `yaml:"bolus_amount_maximum" validate:"required"`
	CorrectionRange            GuardRail `yaml:"correction_range" validate:"required"`
	PreprandialCorrectionRange GuardRail `yaml:"preprandial_correction_range" validate:"required"`
	WorkoutCorrectionRange     GuardRail `yaml:"workout_correction_range" validate:"required"`
}

type GuardRail struct {
	Units             string             `yaml:"units" validate:"required"`
	DefaultValue      *FixedDecimal      `yaml:"default"`
	AbsoluteBounds    []*AbsoluteBounds  `yaml:"absolute_bounds" validate:"required"`
	RecommendedBounds *RecommendedBounds `yaml:"recommended_bounds"`
	MaxSegments       *int32             `yaml:"max_segments"`
}

type AbsoluteBounds struct {
	Bounds `yaml:",inline"`

	Increment *FixedDecimal `yaml:"increment" validate:"gt=0"`
}

type RecommendedBounds struct {
	Bounds `yaml:",inline"`
}

type Bounds struct {
	Minimum *FixedDecimal `yaml:"min"`
	Maximum *FixedDecimal `yaml:"max"`
}

type FixedDecimal struct {
	Units int32
	Nanos int32
}

func (f *FixedDecimal) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	if err := ParseFixedDecimal(value, f); err != nil {
		return err
	}

	return nil
}

func ParseFixedDecimal(value string, decimal *FixedDecimal) error {
	split := strings.Split(value, ".")
	if len(split) != 1 && len(split) != 2 {
		return errors.New(fmt.Sprintf("invalid fixed decimal value %v", value))
	}

	var units, nanos int32
	units, err := parseFixedDecimalUnits(split[0])
	if err != nil {
		return err
	}
	if len(split) > 1 {
		nanos, err = parseFixedDecimalNanos(split[1])
		if err != nil {
			return err
		}
	}

	decimal.Units = units
	decimal.Nanos = nanos

	if units < 0 {
		decimal.Nanos = decimal.Nanos * -1
	}

	return nil
}

func parseFixedDecimalUnits(value string) (int32, error) {
	if len(value) == 0 {
		return 0, nil
	}
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

func parseFixedDecimalNanos(value string) (int32, error) {
	if len(value) == 0 {
		return 0, nil
	}
	if len(value) > 9 {
		return 0, errors.New("nanos must be 9 digits or less")
	}
	// pad with trailing zeroes
	value = value + strings.Repeat("0", 9-len(value))
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
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
