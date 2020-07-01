package repo

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidepool-org/devices/api"
	"github.com/tidepool-org/devices/config"
)

func PumpsConfigToProtoModels(pumps []*config.Pump) ([]*api.Pump, error) {
	models := make([]*api.Pump, len(pumps))
	for i, pump := range pumps {
		model, err := PumpConfigToProto(pump)
		if err != nil {
			return nil, err
		}
		models[i] = model
	}

	return models, nil
}

func PumpConfigToProto(cfg *config.Pump) (*api.Pump, error) {
	guardRails, err := GuardRailsConfigToProto(cfg.GuardRails)
	if err != nil {
		return nil, err
	}

	return &api.Pump{
		Id:            cfg.ID,
		DisplayName:   cfg.DisplayName,
		Manufacturers: cfg.Manufacturers,
		Model:         cfg.Model,
		GuardRails:    &guardRails,
	}, nil
}

func GuardRailsConfigToProto(cfg config.GuardRails) (guardRails api.GuardRails, err error) {
	suspendThreshold := &api.SuspendThresholdGuardRail{}
	guardRails.SuspendThreshold = suspendThreshold
	if err = PopulateSuspendThresholdFromConfig(cfg.SuspendThreshold, suspendThreshold); err != nil {
		return
	}
	insulinSensitivity := &api.InsulinSensitivityGuardRail{}
	guardRails.InsulinSensitivity = insulinSensitivity
	if err = PopulateInsulinSensitivityFromConfig(cfg.InsulinSensitivity, insulinSensitivity); err != nil {
		return
	}
	basalRates := &api.BasalRatesGuardRail{}
	guardRails.BasalRates = basalRates
	if err = PopulateBasalRatesFromConfig(cfg.BasalRates, basalRates); err != nil {
		return
	}
	carbohydrateRatio := &api.CarbohydrateRatioGuardRail{}
	guardRails.CarbohydrateRatio = carbohydrateRatio
	if err = PopulateCarbohydrateRatioFromConfig(cfg.CarbohydrateRatio, carbohydrateRatio); err != nil {
		return
	}
	basalRateMaximum := &api.BasalRateMaximumGuardRail{}
	guardRails.BasalRateMaximum = basalRateMaximum
	if err = PopulateBasalRateMaximumFromConfig(cfg.BasalRateMaximum, basalRateMaximum); err != nil {
		return
	}
	bolusAmountMaximum := &api.BolusAmountMaximumGuardRail{}
	guardRails.BolusAmountMaximum = bolusAmountMaximum
	if err = PopulateBolusAmountMaximumFromConfig(cfg.BolusAmountMaximum, bolusAmountMaximum); err != nil {
		return
	}
	correctionRange := &api.CorrectionRangeGuardRail{}
	guardRails.CorrectionRange = correctionRange
	if err = PopulateCorrectionRangeFromConfig(cfg.CorrectionRange, correctionRange); err != nil {
		return
	}

	return
}

func PopulateSuspendThresholdFromConfig(cfg config.GuardRail, guardRail *api.SuspendThresholdGuardRail) error {
	if cfg.Units != "mg/dL" {
		return errors.New(fmt.Sprintf("unrecognized blood glucose unit %v", cfg.Units))
	}

	guardRail.Units = api.BloodGlucoseUnits_MilligramsPerDeciliter
	guardRail.RecommendedBounds = &api.RecommendedBounds{}
	guardRail.AbsoluteBounds = &api.AbsoluteBounds{}

	if err := PopulateRecommendedBoundsFromConfig(cfg.RecommendedBounds, guardRail.RecommendedBounds); err != nil {
		return err
	}
	if err := PopulateAbsoluteBoundsFromFirstConfigValue(cfg.AbsoluteBounds, guardRail.AbsoluteBounds); err != nil {
		return err
	}

	return nil
}

func PopulateInsulinSensitivityFromConfig(cfg config.GuardRail, guardRail *api.InsulinSensitivityGuardRail) error {
	if cfg.Units != "mg/dL" {
		return errors.New(fmt.Sprintf("unrecognized blood glucose unit %v", cfg.Units))
	}

	guardRail.Units = api.BloodGlucoseUnits_MilligramsPerDeciliter
	guardRail.RecommendedBounds = &api.RecommendedBounds{}
	guardRail.AbsoluteBounds = &api.AbsoluteBounds{}

	if err := PopulateRecommendedBoundsFromConfig(cfg.RecommendedBounds, guardRail.RecommendedBounds); err != nil {
		return err
	}
	if err := PopulateAbsoluteBoundsFromFirstConfigValue(cfg.AbsoluteBounds, guardRail.AbsoluteBounds); err != nil {
		return err
	}

	return nil
}

func PopulateBasalRatesFromConfig(cfg config.GuardRail, guardRail *api.BasalRatesGuardRail) error {
	if cfg.Units != "U/h" {
		return errors.New(fmt.Sprintf("unrecognized basal rate unit %v", cfg.Units))
	}
	if cfg.DefaultValue == nil {
		return errors.New(fmt.Sprintf("defaul value cannot be nil"))
	}

	guardRail.Units = api.BasalRateUnits_UnitsPerHour
	guardRail.DefaultValue = *cfg.DefaultValue
	guardRail.AbsoluteBounds = make([]*api.AbsoluteBounds, len(cfg.AbsoluteBounds))

	if err := PopulateAbsoluteBoundsArrayFromConfig(cfg.AbsoluteBounds, guardRail.AbsoluteBounds); err != nil {
		return err
	}

	return nil
}

func PopulateCarbohydrateRatioFromConfig(cfg config.GuardRail, guardRail *api.CarbohydrateRatioGuardRail) error {
	if cfg.Units != "g/U" {
		return errors.New(fmt.Sprintf("unrecognized carbohydrate ratio unit %v", cfg.Units))
	}

	guardRail.Units = api.CarbohydrateRatioUnits_GramsPerUnit
	guardRail.RecommendedBounds = &api.RecommendedBounds{}
	guardRail.AbsoluteBounds = &api.AbsoluteBounds{}

	if err := PopulateRecommendedBoundsFromConfig(cfg.RecommendedBounds, guardRail.RecommendedBounds); err != nil {
		return err
	}
	if err := PopulateAbsoluteBoundsFromFirstConfigValue(cfg.AbsoluteBounds, guardRail.AbsoluteBounds); err != nil {
		return err
	}

	return nil
}

func PopulateBasalRateMaximumFromConfig(cfg config.GuardRail, guardRail *api.BasalRateMaximumGuardRail) error {
	if cfg.Units != "U/h" {
		return errors.New(fmt.Sprintf("unrecognized basal rate unit %v", cfg.Units))
	}
	if cfg.DefaultValue == nil {
		return errors.New(fmt.Sprintf("defaul value cannot be nil"))
	}

	guardRail.Units = api.BasalRateUnits_UnitsPerHour
	guardRail.DefaultValue = *cfg.DefaultValue
	guardRail.AbsoluteBounds = &api.AbsoluteBounds{}

	if err := PopulateAbsoluteBoundsFromFirstConfigValue(cfg.AbsoluteBounds, guardRail.AbsoluteBounds); err != nil {
		return err
	}

	return nil
}

func PopulateBolusAmountMaximumFromConfig(cfg config.GuardRail, guardRail *api.BolusAmountMaximumGuardRail) error {
	if cfg.Units != "U" {
		return errors.New(fmt.Sprintf("unrecognized bolus amount unit %v", cfg.Units))
	}
	if cfg.DefaultValue == nil {
		return errors.New(fmt.Sprintf("defaul value cannot be nil"))
	}

	guardRail.Units = api.BolusUnits_Units
	guardRail.DefaultValue = *cfg.DefaultValue
	guardRail.RecommendedBounds = &api.RecommendedBounds{}
	guardRail.AbsoluteBounds = &api.AbsoluteBounds{}

	if err := PopulateRecommendedBoundsFromConfig(cfg.RecommendedBounds, guardRail.RecommendedBounds); err != nil {
		return err
	}
	if err := PopulateAbsoluteBoundsFromFirstConfigValue(cfg.AbsoluteBounds, guardRail.AbsoluteBounds); err != nil {
		return err
	}

	return nil
}

func PopulateCorrectionRangeFromConfig(cfg config.GuardRail, guardRail *api.CorrectionRangeGuardRail) error {
	if cfg.Units != "mg/dL" {
		return errors.New(fmt.Sprintf("unrecognized blood glucose unit %v", cfg.Units))
	}

	guardRail.Units = api.BloodGlucoseUnits_MilligramsPerDeciliter
	guardRail.RecommendedBounds = &api.RecommendedBounds{}
	guardRail.AbsoluteBounds = &api.AbsoluteBounds{}

	if err := PopulateRecommendedBoundsFromConfig(cfg.RecommendedBounds, guardRail.RecommendedBounds); err != nil {
		return err
	}
	if err := PopulateAbsoluteBoundsFromFirstConfigValue(cfg.AbsoluteBounds, guardRail.AbsoluteBounds); err != nil {
		return err
	}

	return nil
}

func PopulateRecommendedBoundsFromConfig(bounds *config.RecommendedBounds, recommendedBounds *api.RecommendedBounds) error {
	if bounds == nil {
		return errors.New("recommended bounds cannot be empty")
	}

	recommendedBounds.Minimum = *bounds.Minimum
	recommendedBounds.Maximum = *bounds.Maximum
	return nil
}

func PopulateAbsoluteBoundsFromFirstConfigValue(bounds []*config.AbsoluteBounds, absoluteBounds *api.AbsoluteBounds) error {
	if len(bounds) == 0 {
		return errors.New("absolute bounds is empty")
	}

	absoluteBounds.Minimum = *bounds[0].Minimum
	absoluteBounds.Maximum = *bounds[0].Maximum
	absoluteBounds.Increment = bounds[0].Increment
	return nil
}

func PopulateAbsoluteBoundsArrayFromConfig(bounds []*config.AbsoluteBounds, absoluteBounds []*api.AbsoluteBounds) error {
	if len(bounds) == 0 {
		return errors.New("absolute bounds is empty")
	}

	for i, b := range bounds {
		result := &api.AbsoluteBounds{
			Minimum:   *b.Minimum,
			Maximum:   *b.Maximum,
			Increment: b.Increment,
		}
		absoluteBounds[i] = result
	}

	return nil
}

