package config

import (
	"strings"
	"testing"
)

func TestDevicesConfig_LoadFromFile(t *testing.T) {
	cfg := NewDevicesConfig()
	err := cfg.LoadFromFile("../devices.yaml")
	if err != nil {
		t.Errorf("unexpected error occurred while loading config from file: %v", err)
		t.FailNow()
	}

	t.Run("Config has a single pump", func(t *testing.T) {
		pumpCount := len(cfg.Devices.Pumps)
		if pumpCount != 1 {
			t.Errorf("expected 1 pump in config, got %v", pumpCount)
			t.FailNow()
		}

		t.Run("Omnipod Horizon", func(t *testing.T) {
			expectedOmnipodId := "6678c377-928c-49b3-84c1-19e2dafaff8d"
			var omnipod *Pump
			for _, p := range cfg.Devices.Pumps {
				if p.ID == expectedOmnipodId {
					omnipod = p
					break
				}
			}

			t.Run("Exists", func(t *testing.T) {
				if omnipod == nil {
					t.Errorf("expected omnipod pod with id %v, but did not find it in config", expectedOmnipodId)
					t.FailNow()
				}
			})

			if omnipod == nil {
				t.FailNow()
			}

			t.Run("Display name equals 'Omnipod Horizon'", func(t *testing.T) {
				expected := "Omnipod Horizon"
				if omnipod.DisplayName != expected {
					t.Errorf("expected display name to equal %v, but got %v", expected, omnipod.DisplayName)
					t.FailNow()
				}
			})

			t.Run("Model equals 'Omnipod Horizon'", func(t *testing.T) {
				expected := "Omnipod Horizon"
				if omnipod.Model != expected {
					t.Errorf("expected model to equal %v, but got %v", expected, omnipod.Model)
					t.FailNow()
				}
			})

			t.Run("Manufacturers consists of 'Insulet'", func(t *testing.T) {
				expected := "Insulet"
				if len(omnipod.Manufacturers) != 1 || omnipod.Manufacturers[0] != expected {
					t.Errorf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(omnipod.Manufacturers, ","))
					t.FailNow()
				}
			})

			t.Run("Guard Rails", func(t *testing.T) {
				fixedDecimalsAreEqual := func(a, b FixedDecimal) bool {
					return a.Units == b.Units && a.Nanos == b.Nanos
				}
				isExpected := func(t *testing.T, result GuardRail, expected GuardRail) {
					if result.Units != expected.Units {
						t.Errorf("expected %v units got %v", expected.Units, result.Units)
					}
					if expected.DefaultValue != nil {
						if result.DefaultValue == nil || !fixedDecimalsAreEqual(*result.DefaultValue, *expected.DefaultValue) {
							t.Errorf("expected %v got %v", *expected.DefaultValue, *result.DefaultValue)
						}
					}
					if expected.AbsoluteBounds != nil {
						for i, b := range expected.AbsoluteBounds {
							expectedValue := *expected.AbsoluteBounds[i].Increment
							resultValue := *b.Increment
							if !fixedDecimalsAreEqual(expectedValue, resultValue) {
								t.Errorf("expected %v got %v", expectedValue, resultValue)
							}

							expectedValue = *expected.AbsoluteBounds[i].Minimum
							resultValue = *b.Minimum
							if !fixedDecimalsAreEqual(expectedValue, resultValue) {
								t.Errorf("expected %v got %v", expectedValue, resultValue)
							}

							expectedValue = *expected.AbsoluteBounds[i].Maximum
							resultValue = *b.Maximum
							if !fixedDecimalsAreEqual(expectedValue, resultValue) {
								t.Errorf("expected %v got %v", expectedValue, resultValue)
							}
						}
					}
					if expected.RecommendedBounds != nil {
						expectedValue := *expected.RecommendedBounds.Minimum
						resultValue := *result.RecommendedBounds.Minimum
						if !fixedDecimalsAreEqual(expectedValue, resultValue) {
							t.Errorf("expected %v got %v", expectedValue, resultValue)
						}

						expectedValue = *expected.RecommendedBounds.Maximum
						resultValue = *result.RecommendedBounds.Maximum
						if !fixedDecimalsAreEqual(expectedValue, resultValue) {
							t.Errorf("expected %v got %v", expectedValue, resultValue)
						}
					}
				}
				t.Run("Glucose safety limit is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.GlucoseSafetyLimit, GuardRail{
						Units:        "mg/dL",
						DefaultValue: nil,
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 67},
									Maximum: &FixedDecimal{Units: 110},
								},
								Increment: &FixedDecimal{Units: 1},
							},
						},
						RecommendedBounds: &RecommendedBounds{
							Bounds{
								Minimum: &FixedDecimal{Units: 74},
								Maximum: &FixedDecimal{Units: 80},
							},
						},
					})
				})
				t.Run("Insulin sensitivity is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.InsulinSensitivity, GuardRail{
						Units:        "mg/dL",
						DefaultValue: nil,
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 10},
									Maximum: &FixedDecimal{Units: 500},
								},
								Increment: &FixedDecimal{Units: 1},
							},
						},
						RecommendedBounds: &RecommendedBounds{
							Bounds{
								Minimum: &FixedDecimal{Units: 16},
								Maximum: &FixedDecimal{Units: 399},
							},
						},
					})
				})
				t.Run("Basal rates is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.BasalRates, GuardRail{
						Units:        "U/h",
						DefaultValue: &FixedDecimal{Nanos: 50000000},
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Nanos: 50000000},
									Maximum: &FixedDecimal{Units: 30},
								},
								Increment: &FixedDecimal{Nanos: 50000000},
							},
						},
						RecommendedBounds: nil,
					})
				})
				t.Run("Carbohydrate ratio is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.CarbohydrateRatio, GuardRail{
						Units:        "g/U",
						DefaultValue: nil,
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 2},
									Maximum: &FixedDecimal{Units: 150},
								},
								Increment: &FixedDecimal{Nanos: 10000000},
							},
						},
						RecommendedBounds: &RecommendedBounds{
							Bounds{
								Minimum: &FixedDecimal{Units: 4},
								Maximum: &FixedDecimal{Units: 28},
							},
						},
					})
				})
				t.Run("Basal rate maximum is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.BasalRateMaximum, GuardRail{
						Units:        "U/h",
						DefaultValue: &FixedDecimal{Units: 0, Nanos: 50000000},
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 0, Nanos: 50000000},
									Maximum: &FixedDecimal{Units: 30},
								},
								Increment: &FixedDecimal{Nanos: 50000000},
							},
						},
						RecommendedBounds: nil,
					})
				})
				t.Run("Bolus amount maximum is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.BolusAmountMaximum, GuardRail{
						Units: "U",
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 0, Nanos: 50000000},
									Maximum: &FixedDecimal{Units: 30},
								},
								Increment: &FixedDecimal{Nanos: 50000000},
							},
						},
						RecommendedBounds: &RecommendedBounds{
							Bounds{
								Minimum: &FixedDecimal{Units: 0, Nanos: 100000000},
								Maximum: &FixedDecimal{Units: 19, Nanos: 950000000},
							},
						},
					})
				})
				t.Run("Correction range is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.CorrectionRange, GuardRail{
						Units: "mg/dL",
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 87},
									Maximum: &FixedDecimal{Units: 180},
								},
								Increment: &FixedDecimal{Units: 1},
							},
						},
						RecommendedBounds: &RecommendedBounds{
							Bounds{
								Minimum: &FixedDecimal{Units: 100},
								Maximum: &FixedDecimal{Units: 115},
							},
						},
					})
				})
				t.Run("Workout correction range is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.WorkoutCorrectionRange, GuardRail{
						Units: "mg/dL",
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 87},
									Maximum: &FixedDecimal{Units: 250},
								},
								Increment: &FixedDecimal{Units: 1},
							},
						},
					})
				})
				t.Run("Pre-meal correction range is correct", func(t *testing.T) {
					isExpected(t, omnipod.GuardRails.WorkoutCorrectionRange, GuardRail{
						Units: "mg/dL",
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 87},
									Maximum: &FixedDecimal{Units: 180},
								},
								Increment: &FixedDecimal{Units: 1},
							},
						},
					})
				})
			})

		})
	})

	t.Run("Config has a single cgm", func(t *testing.T) {
		cgmCount := len(cfg.Devices.CGMs)
		if cgmCount != 1 {
			t.Errorf("expected 1 cgm in config, got %v", cgmCount)
			t.FailNow()
		}

		t.Run("Dexcom G6", func(t *testing.T) {
			expectedDexomG6Id := "d25c3f1b-a2e8-44e2-b3a3-fd07806fc245"
			var g6 *CGM
			for _, p := range cfg.Devices.CGMs {
				if p.ID == expectedDexomG6Id {
					g6 = p
					break
				}
			}

			t.Run("Exists", func(t *testing.T) {
				if g6 == nil {
					t.Errorf("expected dexcom g6 pod with id %v, but did not find it in config", expectedDexomG6Id)
					t.FailNow()
				}
			})

			if g6 == nil {
				t.FailNow()
			}

			t.Run("Display name equals 'Dexcom G6'", func(t *testing.T) {
				expected := "Dexcom G6"
				if g6.DisplayName != expected {
					t.Errorf("expected display name to equal %v, but got %v", expected, g6.DisplayName)
					t.FailNow()
				}
			})

			t.Run("Model equals 'G6'", func(t *testing.T) {
				expected := "G6"
				if g6.Model != expected {
					t.Errorf("expected model to equal %v, but got %v", expected, g6.Model)
					t.FailNow()
				}
			})

			t.Run("Manufacturers consists of 'Dexcom'", func(t *testing.T) {
				expected := "Dexcom"
				if len(g6.Manufacturers) != 1 || g6.Manufacturers[0] != expected {
					t.Errorf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(g6.Manufacturers, ","))
					t.FailNow()
				}
			})

		})
	})
}
