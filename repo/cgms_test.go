package repo

import (
	"context"
	"github.com/tidepool-org/devices/api"
	"testing"
)

var TestCgmData = []*api.Cgm{
	{
		Id:            "1",
		DisplayName:   "Test device 1",
		Manufacturers: []string{"Tidepool"},
		Model:         "Model 1",
	},
	{
		Id:            "2",
		DisplayName:   "Test device 2",
		Manufacturers: []string{"Tidepool"},
		Model:         "Model 2",
	},
}

func TestInMemoryCgmsRepo_GetById(t *testing.T) {
	repo := NewCgmsRepo(TestCgmData)

	ctx := context.Background()
	t.Run("Get by id works when ids exist", func(t *testing.T) {
		first := "1"
		second := "2"

		firstResult, err := repo.GetById(ctx, first)
		if err != nil {
			t.Errorf("unexpected error occurred while getting cgm by id: %v", err)
		} else if firstResult == nil {
			t.Errorf("did not find device with id %v", first)
		} else {
			if firstResult.Id != first {
				t.Errorf("expected cgm with id %v, got %v", first, firstResult.Id)
			}
		}

		secondResult, err := repo.GetById(ctx, second)
		if err != nil {
			t.Errorf("unexpected error occurred while getting cgm by id: %v", err)
		} else if secondResult == nil {
			t.Errorf("did not find device with id %v", second)
		} else {
			if secondResult.Id != second {
				t.Errorf("expected cgm with id %v, got %v", second, secondResult.Id)
			}
		}
	})

	t.Run("Get by return nil when id doesn't exist", func(t *testing.T) {
		id := "invalid-id"

		result, err := repo.GetById(ctx, id)
		if err != nil {
			t.Errorf("unexpected error occurred while getting cgm by id: %v", err)
			t.FailNow()
		}
		
		if result != nil {
			t.Errorf("unexpected non nil id for non-existent id: %v", result)
		}
	})
}

func TestInMemoryCgmsRepo_List(t *testing.T) {
	repo := NewCgmsRepo(TestCgmData)

	ctx := context.Background()
	t.Run("List returns all cgms", func(t *testing.T) {
		list, err := repo.List(ctx)
		if err != nil {
			t.Errorf("unexpected error occurred while getting the list with all cgms: %v", err)
			t.FailNow()
		}

		if len(list) != len(TestCgmData) {
			t.Errorf("Expected %v items , got %v", len(TestCgmData), len(list))
		}
	})
}
