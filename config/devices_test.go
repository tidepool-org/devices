package config

import (
	"strings"
	"testing"
)

func TestDevicesConfig_LoadFromFile(t *testing.T) {
	cfg := NewDevicesConfig()
	err := cfg.LoadFromFile("../devices.yaml")
	if err != nil {
		t.Fatalf("unexpected error occurred while loading config from file: %v", err)
	}

	t.Run("Config has four pumps", func(t *testing.T) {
		pumpCount := len(cfg.Devices.Pumps)
		if pumpCount != 4 {
			t.Fatalf("expected 4 pump in config, got %v", pumpCount)
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
					t.Fatalf("expected coastal pump with id %v, but did not find it in config", expectedCoastalId)
				}
			})

			if coastal == nil {
				t.FailNow()
			}

			t.Run("Display name equals 'Coastal'", func(t *testing.T) {
				expected := "Coastal"
				if coastal.DisplayName != expected {
					t.Fatalf("expected display name to equal %v, but got %v", expected, coastal.DisplayName)
				}
			})

			t.Run("Model equals 'Coastal'", func(t *testing.T) {
				expected := "Coastal"
				if coastal.Model != expected {
					t.Fatalf("expected model to equal %v, but got %v", expected, coastal.Model)
				}
			})

			t.Run("Manufacturers consists of 'Coastal'", func(t *testing.T) {
				expected := "Coastal"
				if len(coastal.Manufacturers) != 1 || coastal.Manufacturers[0] != expected {
					t.Fatalf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(coastal.Manufacturers, ","))
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
					t.Fatalf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(palmtree.Manufacturers, ","))
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
					maxSegments := int32(48)
					isExpected(t, palmtree.GuardRails.InsulinSensitivity, GuardRail{
						Units:        "mg/dL",
						DefaultValue: nil,
						MaxSegments:  &maxSegments,
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
						Units:       "U/h",
						MaxSegments: &maxSegments,
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
					maxSegments := int32(48)
					isExpected(t, palmtree.GuardRails.CarbohydrateRatio, GuardRail{
						Units:        "g/U",
						DefaultValue: nil,
						MaxSegments:  &maxSegments,
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
					maxSegments := int32(48)
					isExpected(t, palmtree.GuardRails.CorrectionRange, GuardRail{
						Units:       "mg/dL",
						MaxSegments: &maxSegments,
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

		t.Run("Palmtree Pump Demo", func(t *testing.T) {
			expectedPalmtreeId := "e6d9afc8-2642-4d95-a2b5-58929e44e105"
			var palmtree *Pump
			for _, p := range cfg.Devices.Pumps {
				if p.ID == expectedPalmtreeId {
					palmtree = p
					break
				}
			}

			t.Run("Exists", func(t *testing.T) {
				if palmtree == nil {
					t.Fatalf("expected palmtree pump with id %v, but did not find it in config", expectedPalmtreeId)
				}
			})

			t.Run("Display name equals 'Palmtree Pump Demo'", func(t *testing.T) {
				expected := "Palmtree Pump Demo"
				if palmtree.DisplayName != expected {
					t.Errorf("expected display name to equal %v, but got %v", expected, palmtree.DisplayName)
					t.FailNow()
				}
			})

			t.Run("Model equals 'Palmtree Demo'", func(t *testing.T) {
				expected := "Palmtree Demo"
				if palmtree.Model != expected {
					t.Errorf("expected model to equal %v, but got %v", expected, palmtree.Model)
					t.FailNow()
				}
			})

			t.Run("Manufacturers consists of 'Palmtree'", func(t *testing.T) {
				expected := "Palmtree"
				if len(palmtree.Manufacturers) != 1 || palmtree.Manufacturers[0] != expected {
					t.Fatalf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(palmtree.Manufacturers, ","))
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
					maxSegments := int32(48)
					isExpected(t, palmtree.GuardRails.InsulinSensitivity, GuardRail{
						Units:        "mg/dL",
						DefaultValue: nil,
						MaxSegments:  &maxSegments,
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
						Units:       "U/h",
						MaxSegments: &maxSegments,
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
					maxSegments := int32(48)
					isExpected(t, palmtree.GuardRails.CarbohydrateRatio, GuardRail{
						Units:        "g/U",
						DefaultValue: nil,
						MaxSegments:  &maxSegments,
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
					maxSegments := int32(48)
					isExpected(t, palmtree.GuardRails.CorrectionRange, GuardRail{
						Units:       "mg/dL",
						MaxSegments: &maxSegments,
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

		t.Run("Pump Simulator", func(t *testing.T) {
			expectedSimulatorId := "aff19260-fca9-4efc-9b81-dcd39f695979"
			var simulator *Pump
			for _, p := range cfg.Devices.Pumps {
				if p.ID == expectedSimulatorId {
					simulator = p
					break
				}
			}

			t.Run("Exists", func(t *testing.T) {
				if simulator == nil {
					t.Errorf("expected simulator pump with id %v, but did not find it in config", expectedSimulatorId)
					t.FailNow()
				}
			})

			if simulator == nil {
				t.FailNow()
			}

			t.Run("Display name equals 'Pump Simulator'", func(t *testing.T) {
				expected := "Pump Simulator"
				if simulator.DisplayName != expected {
					t.Errorf("expected display name to equal %v, but got %v", expected, simulator.DisplayName)
					t.FailNow()
				}
			})

			t.Run("Model equals 'Pump Simulator'", func(t *testing.T) {
				expected := "Pump Simulator"
				if simulator.Model != expected {
					t.Errorf("expected model to equal %v, but got %v", expected, simulator.Model)
					t.FailNow()
				}
			})

			t.Run("Manufacturers consists of 'Tidepool'", func(t *testing.T) {
				expected := "Tidepool"
				if len(simulator.Manufacturers) != 1 || simulator.Manufacturers[0] != expected {
					t.Fatalf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(simulator.Manufacturers, ","))
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
					isExpected(t, simulator.GuardRails.GlucoseSafetyLimit, GuardRail{
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
					maxSegments := int32(48)
					isExpected(t, simulator.GuardRails.InsulinSensitivity, GuardRail{
						Units:        "mg/dL",
						DefaultValue: nil,
						MaxSegments:  &maxSegments,
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
					isExpected(t, simulator.GuardRails.BasalRates, GuardRail{
						Units:       "U/h",
						MaxSegments: &maxSegments,
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
					maxSegments := int32(48)
					isExpected(t, simulator.GuardRails.CarbohydrateRatio, GuardRail{
						Units:        "g/U",
						DefaultValue: nil,
						MaxSegments:  &maxSegments,
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
					isExpected(t, simulator.GuardRails.BasalRateMaximum, GuardRail{
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
					isExpected(t, simulator.GuardRails.BolusAmountMaximum, GuardRail{
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
					maxSegments := int32(48)
					isExpected(t, simulator.GuardRails.CorrectionRange, GuardRail{
						Units:       "mg/dL",
						MaxSegments: &maxSegments,
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
					isExpected(t, simulator.GuardRails.WorkoutCorrectionRange, GuardRail{
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
					isExpected(t, simulator.GuardRails.WorkoutCorrectionRange, GuardRail{
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

	t.Run("Config has a three cgms", func(t *testing.T) {
		cgmCount := len(cfg.Devices.CGMs)
		if cgmCount != 3 {
			t.Fatalf("expected 3 cgm in config, got %v", cgmCount)
		}

		t.Run("Mock CGM", func(t *testing.T) {
			expectedMockCGMId := "c97bd194-5e5e-44c1-9629-4cb87be1a4c9"
			var mock *CGM
			for _, p := range cfg.Devices.CGMs {
				if p.ID == expectedMockCGMId {
					mock = p
					break
				}
			}

			t.Run("Exists", func(t *testing.T) {
				if mock == nil {
					t.Fatalf("expected mock cgm with id %v, but did not find it in config", expectedMockCGMId)
				}
			})

			if mock == nil {
				t.FailNow()
			}

			t.Run("Display name equals 'Mock CGM'", func(t *testing.T) {
				expected := "Mock CGM"
				if mock.DisplayName != expected {
					t.Fatalf("expected display name to equal %v, but got %v", expected, mock.DisplayName)
				}
			})

			t.Run("Model equals 'Mock'", func(t *testing.T) {
				expected := "Mock"
				if mock.Model != expected {
					t.Fatalf("expected model to equal %v, but got %v", expected, mock.Model)
				}
			})

			t.Run("Manufacturers consists of 'Tidepool'", func(t *testing.T) {
				expected := "Tidepool"
				if len(mock.Manufacturers) != 1 || mock.Manufacturers[0] != expected {
					t.Fatalf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(mock.Manufacturers, ","))
				}
			})

		})

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
					t.Fatalf("expected dexcom g6 pod with id %v, but did not find it in config", expectedDexomG6Id)
				}
			})

			if g6 == nil {
				t.FailNow()
			}

			t.Run("Display name equals 'Dexcom G6'", func(t *testing.T) {
				expected := "Dexcom G6"
				if g6.DisplayName != expected {
					t.Fatalf("expected display name to equal %v, but got %v", expected, g6.DisplayName)
				}
			})

			t.Run("Model equals 'G6'", func(t *testing.T) {
				expected := "G6"
				if g6.Model != expected {
					t.Fatalf("expected model to equal %v, but got %v", expected, g6.Model)
				}
			})

			t.Run("Manufacturers consists of 'Dexcom'", func(t *testing.T) {
				expected := "Dexcom"
				if len(g6.Manufacturers) != 1 || g6.Manufacturers[0] != expected {
					t.Fatalf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(g6.Manufacturers, ","))
				}
			})

		})

		t.Run("Dexcom G6 Demo", func(t *testing.T) {
			expectedDexomG6DemoId := "80eb646f-18c0-4799-a1bf-3a0e1a5390d9"
			var g6demo *CGM
			for _, p := range cfg.Devices.CGMs {
				if p.ID == expectedDexomG6DemoId {
					g6demo = p
					break
				}
			}

			t.Run("Exists", func(t *testing.T) {
				if g6demo == nil {
					t.Fatalf("expected dexcom g6 pod with id %v, but did not find it in config", expectedDexomG6DemoId)
				}
			})

			t.Run("Display name equals 'Dexcom G6 Demo'", func(t *testing.T) {
				expected := "Dexcom G6 Demo"
				if g6demo.DisplayName != expected {
					t.Fatalf("expected display name to equal %v, but got %v", expected, g6demo.DisplayName)
				}
			})

			t.Run("Model equals 'G6 Demo'", func(t *testing.T) {
				expected := "G6 Demo"
				if g6demo.Model != expected {
					t.Fatalf("expected model to equal %v, but got %v", expected, g6demo.Model)
				}
			})

			t.Run("Manufacturers consists of 'Dexcom'", func(t *testing.T) {
				expected := "Dexcom"
				if len(g6demo.Manufacturers) != 1 || g6demo.Manufacturers[0] != expected {
					t.Fatalf("expected manufacturers equal [%v], but got [%v]", expected, strings.Join(g6demo.Manufacturers, ","))
				}
			})

		})
	})
}
