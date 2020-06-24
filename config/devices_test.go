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
