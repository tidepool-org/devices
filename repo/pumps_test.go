package repo

import (
	"context"
	"github.com/tidepool-org/devices/config"
	"testing"
)

var TestPumpData = []*config.Pump{
	&config.Pump{config.Device{
		ID:            "1",
		DisplayName:   "Test device 1",
		Manufacturers: []string{"Tidepool"},
		Model:         "Model 1",
	}},
	&config.Pump{config.Device{
		ID:            "2",
		DisplayName:   "Test device 2",
		Manufacturers: []string{"Tidepool"},
		Model:         "Model 2",
	}},
}

func TestInMemoryPumpsRepo_GetById(t *testing.T) {
	repo := NewPumpsRepo(&config.Devices{
		Pumps: TestPumpData,
	})

	ctx := context.Background()

	t.Run("Get by id works when ids exist", func(t *testing.T) {
		first := "1"
		second := "2"

		firstResult, err := repo.GetById(ctx, first)
		if err != nil {
			t.Errorf("unexpected error occurred while getting pump by id: %v", err)
		} else if firstResult == nil {
			t.Errorf("did not find device with id %v", first)
		} else {
			if firstResult.Id != first {
				t.Errorf("expected pump with id %v, got %v", first, firstResult.Id)
			}
		}

		secondResult, err := repo.GetById(ctx, second)
		if err != nil {
			t.Errorf("unexpected error occurred while getting pump by id: %v", err)
		} else if secondResult == nil {
			t.Errorf("did not find device with id %v", second)
		} else {
			if secondResult.Id != second {
				t.Errorf("expected pump with id %v, got %v", second, secondResult.Id)
			}
		}
	})

	t.Run("Get by return nil when id doesn't exist", func(t *testing.T) {
		id := "invalid-id"

		result, err := repo.GetById(ctx, id)
		if err != nil {
			t.Errorf("unexpected error occurred while getting pump by id: %v", err)
			t.FailNow()
		}
		
		if result != nil {
			t.Errorf("unexpected non nil id for non-existent id: %v", result)
		}
	})
}

func TestInMemoryPumpsRepo_List(t *testing.T) {
	repo := NewPumpsRepo(&config.Devices{
		Pumps: TestPumpData,
	})

	ctx := context.Background()
	t.Run("List returns all pumps", func(t *testing.T) {
		list, err := repo.List(ctx)
		if err != nil {
			t.Errorf("unexpected error occurred while getting the list with all pumps: %v", err)
			t.FailNow()
		}

		if len(list) != len(TestPumpData) {
			t.Errorf("Expected %v items , got %v", len(TestPumpData), len(list))
		}
	})
}
