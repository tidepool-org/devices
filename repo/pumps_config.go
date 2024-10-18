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
	glucodeSafetyLimit := &api.GlucoseSafetyLimitGuardRail{}
	guardRails.GlucoseSafetyLimit = glucodeSafetyLimit
	if err = PopulateGlucoseSafetyLimitFromConfig(cfg.GlucoseSafetyLimit, glucodeSafetyLimit); err != nil {
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
	preprandialCorrectionRange := &api.CorrectionRangeGuardRail{}
	guardRails.PreprandialCorrectionRange = preprandialCorrectionRange
	if err = PopulateCorrectionRangeFromConfig(cfg.PreprandialCorrectionRange, preprandialCorrectionRange); err != nil {
		return
	}
	workoutCorrectionRange := &api.CorrectionRangeGuardRail{}
	guardRails.WorkoutCorrectionRange = workoutCorrectionRange
	if err = PopulateCorrectionRangeFromConfig(cfg.WorkoutCorrectionRange, workoutCorrectionRange); err != nil {
		return
	}

	return
}

func PopulateGlucoseSafetyLimitFromConfig(cfg config.GuardRail, guardRail *api.GlucoseSafetyLimitGuardRail) error {
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
	guardRail.MaxSegments = cfg.MaxSegments

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

	guardRail.Units = api.BasalRateUnits_UnitsPerHour
	if cfg.DefaultValue != nil {
		guardRail.DefaultValue = &api.FixedDecimal{
			Units: cfg.DefaultValue.Units,
			Nanos: cfg.DefaultValue.Nanos,
		}
	}
	guardRail.AbsoluteBounds = make([]*api.AbsoluteBounds, len(cfg.AbsoluteBounds))
	guardRail.MaxSegments = cfg.MaxSegments

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
	guardRail.MaxSegments = cfg.MaxSegments

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
		return errors.New(fmt.Sprintf("default value cannot be nil"))
	}

	guardRail.Units = api.BasalRateUnits_UnitsPerHour
	guardRail.DefaultValue = &api.FixedDecimal{
		Units: cfg.DefaultValue.Units,
		Nanos: cfg.DefaultValue.Nanos,
	}
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

	guardRail.Units = api.BolusUnits_Units
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
	guardRail.AbsoluteBounds = &api.AbsoluteBounds{}
	guardRail.MaxSegments = cfg.MaxSegments

	if err := PopulateAbsoluteBoundsFromFirstConfigValue(cfg.AbsoluteBounds, guardRail.AbsoluteBounds); err != nil {
		return err
	}

	if cfg.RecommendedBounds != nil {
		guardRail.RecommendedBounds = &api.RecommendedBounds{}
		if err := PopulateRecommendedBoundsFromConfig(cfg.RecommendedBounds, guardRail.RecommendedBounds); err != nil {
			return err
		}
	}

	return nil
}

func PopulateRecommendedBoundsFromConfig(bounds *config.RecommendedBounds, recommendedBounds *api.RecommendedBounds) error {
	if bounds == nil {
		return errors.New("recommended bounds cannot be empty")
	}

	recommendedBounds.Minimum = &api.FixedDecimal{
		Units: bounds.Minimum.Units,
		Nanos: bounds.Minimum.Nanos,
	}
	recommendedBounds.Maximum = &api.FixedDecimal{
		Units: bounds.Maximum.Units,
		Nanos: bounds.Maximum.Nanos,
	}
	return nil
}

func PopulateAbsoluteBoundsFromFirstConfigValue(bounds []*config.AbsoluteBounds, absoluteBounds *api.AbsoluteBounds) error {
	if len(bounds) == 0 {
		return errors.New("absolute bounds is empty")
	}

	absoluteBounds.Minimum = &api.FixedDecimal{
		Units: bounds[0].Minimum.Units,
		Nanos: bounds[0].Minimum.Nanos,
	}
	absoluteBounds.Maximum = &api.FixedDecimal{
		Units: bounds[0].Maximum.Units,
		Nanos: bounds[0].Maximum.Nanos,
	}
	absoluteBounds.Increment = &api.FixedDecimal{
		Units: bounds[0].Increment.Units,
		Nanos: bounds[0].Increment.Nanos,
	}
	return nil
}

func PopulateAbsoluteBoundsArrayFromConfig(bounds []*config.AbsoluteBounds, absoluteBounds []*api.AbsoluteBounds) error {
	if len(bounds) == 0 {
		return errors.New("absolute bounds is empty")
	}

	for i, b := range bounds {
		result := &api.AbsoluteBounds{
			Minimum: &api.FixedDecimal{
				Units: b.Minimum.Units,
				Nanos: b.Minimum.Nanos,
			},
			Maximum: &api.FixedDecimal{
				Units: b.Maximum.Units,
				Nanos: b.Maximum.Nanos,
			},
			Increment: &api.FixedDecimal{
				Units: b.Increment.Units,
				Nanos: b.Increment.Nanos,
			},
		}
		absoluteBounds[i] = result
	}

	return nil
}
