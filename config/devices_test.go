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

	t.Run("Config has three pumps", func(t *testing.T) {
		pumpCount := len(cfg.Devices.Pumps)
		if pumpCount != 2 {
			t.Errorf("expected 2 pump in config, got %v", pumpCount)
			t.FailNow()
		}

		t.Run("Coastal", func(t *testing.T) {
			expectedCoastalId := "e4a46eda-02f9-4faf-b8f4-ef7b40d02e4f"
			var coastal *Pump
			for _, p := range cfg.Devices.Pumps {
				if p.ID == expectedCoastalId {
					coastal = p
					break
				}
			}

			t.Run("Exists", func(t *testing.T) {
				if coastal == nil {
					t.Errorf("expected coastal pump with id %v, but did not find it in config", expectedCoastalId)
					t.FailNow()
				}
			})

			if coastal == nil {
				t.FailNow()
			}

			t.Run("Display name equals 'Coastal'", func(t *testing.T) {
				expected := "Coastal"
				if coastal.DisplayName != expected {
					t.Errorf("expected display name to equal %v, but got %v", expected, coastal.DisplayName)
					t.FailNow()
				}
			})

			t.Run("Model equals 'Coastal'", func(t *testing.T) {
				expected := "Coastal"
				if coastal.Model != expected {
					t.Errorf("expected model to equal %v, but got %v", expected, coastal.Model)
					t.FailNow()
				}
			})

			t.Run("Manufacturers consists of 'Coastal'", func(t *testing.T) {
				expected := "Coastal"
				if len(coastal.Manufacturers) != 1 || coastal.Manufacturers[0] != expected {
					t.Errorf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(coastal.Manufacturers, ","))
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
					isExpected(t, coastal.GuardRails.GlucoseSafetyLimit, GuardRail{
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
					isExpected(t, coastal.GuardRails.InsulinSensitivity, GuardRail{
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
					isExpected(t, coastal.GuardRails.BasalRates, GuardRail{
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
					isExpected(t, coastal.GuardRails.CarbohydrateRatio, GuardRail{
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
					isExpected(t, coastal.GuardRails.BasalRateMaximum, GuardRail{
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
					isExpected(t, coastal.GuardRails.BolusAmountMaximum, GuardRail{
						Units: "U",
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 0, Nanos: 200000000},
									Maximum: &FixedDecimal{Units: 30},
								},
								Increment: &FixedDecimal{Nanos: 50000000},
							},
						},
						RecommendedBounds: &RecommendedBounds{
							Bounds{
								Minimum: &FixedDecimal{Units: 0, Nanos: 200000000},
								Maximum: &FixedDecimal{Units: 19, Nanos: 950000000},
							},
						},
					})
				})
				t.Run("Correction range is correct", func(t *testing.T) {
					isExpected(t, coastal.GuardRails.CorrectionRange, GuardRail{
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
					isExpected(t, coastal.GuardRails.WorkoutCorrectionRange, GuardRail{
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
					isExpected(t, coastal.GuardRails.WorkoutCorrectionRange, GuardRail{
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

		t.Run("Palmtree", func(t *testing.T) {
			expectedPalmtreeId := "c524b5b0-632e-4125-8f6a-df9532d8f6fe"
			var palmtree *Pump
			for _, p := range cfg.Devices.Pumps {
				if p.ID == expectedPalmtreeId {
					palmtree = p
					break
				}
			}

			t.Run("Exists", func(t *testing.T) {
				if palmtree == nil {
					t.Errorf("expected palmtree pump with id %v, but did not find it in config", expectedPalmtreeId)
					t.FailNow()
				}
			})

			if palmtree == nil {
				t.FailNow()
			}

			t.Run("Display name equals 'Palmtree'", func(t *testing.T) {
				expected := "Palmtree"
				if palmtree.DisplayName != expected {
					t.Errorf("expected display name to equal %v, but got %v", expected, palmtree.DisplayName)
					t.FailNow()
				}
			})

			t.Run("Model equals 'Palmtree'", func(t *testing.T) {
				expected := "Palmtree"
				if palmtree.Model != expected {
					t.Errorf("expected model to equal %v, but got %v", expected, palmtree.Model)
					t.FailNow()
				}
			})

			t.Run("Manufacturers consists of 'Palmtree'", func(t *testing.T) {
				expected := "Palmtree"
				if len(palmtree.Manufacturers) != 1 || palmtree.Manufacturers[0] != expected {
					t.Errorf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(palmtree.Manufacturers, ","))
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
					if expected.MaxSegments != nil {
						if result.MaxSegments == nil || *expected.MaxSegments != *result.MaxSegments {
							t.Errorf("expected %v got %v", *expected.MaxSegments, *result.MaxSegments)
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
					isExpected(t, palmtree.GuardRails.GlucoseSafetyLimit, GuardRail{
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
					isExpected(t, palmtree.GuardRails.InsulinSensitivity, GuardRail{
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
					maxSegments := int32(24)
					isExpected(t, palmtree.GuardRails.BasalRates, GuardRail{
						Units:        "U/h",
						DefaultValue: &FixedDecimal{Nanos: 50000000},
						MaxSegments:  &maxSegments,
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
					isExpected(t, palmtree.GuardRails.CarbohydrateRatio, GuardRail{
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
					isExpected(t, palmtree.GuardRails.BasalRateMaximum, GuardRail{
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
					isExpected(t, palmtree.GuardRails.BolusAmountMaximum, GuardRail{
						Units: "U",
						AbsoluteBounds: []*AbsoluteBounds{
							&AbsoluteBounds{
								Bounds: Bounds{
									Minimum: &FixedDecimal{Units: 0, Nanos: 200000000},
									Maximum: &FixedDecimal{Units: 30},
								},
								Increment: &FixedDecimal{Nanos: 50000000},
							},
						},
						RecommendedBounds: &RecommendedBounds{
							Bounds{
								Minimum: &FixedDecimal{Units: 0, Nanos: 200000000},
								Maximum: &FixedDecimal{Units: 19, Nanos: 950000000},
							},
						},
					})
				})
				t.Run("Correction range is correct", func(t *testing.T) {
					isExpected(t, palmtree.GuardRails.CorrectionRange, GuardRail{
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
								Maximum: &FixedDecimal{Units: 125},
							},
						},
					})
				})
				t.Run("Workout correction range is correct", func(t *testing.T) {
					isExpected(t, palmtree.GuardRails.WorkoutCorrectionRange, GuardRail{
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
					isExpected(t, palmtree.GuardRails.WorkoutCorrectionRange, GuardRail{
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
