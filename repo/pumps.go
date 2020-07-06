package repo

import (
	"context"
	"github.com/tidepool-org/devices/api"
)

type PumpsRepo interface {
	GetById(ctx context.Context, id string) (*api.Pump, error)
	List(context.Context) ([]*api.Pump, error)
}

func NewPumpsRepo(pumps []*api.Pump) PumpsRepo {
	pumpsIdMap := make(map[string]*api.Pump)
	for _, p := range pumps {
		pumpsIdMap[p.GetId()] = p
	}
	return &inMemoryPumpsRepo{pumps: pumpsIdMap}
}

type inMemoryPumpsRepo struct {
	pumps map[string]*api.Pump
}

func (i *inMemoryPumpsRepo) GetById(ctx context.Context, id string) (*api.Pump, error) {
	p, ok := i.pumps[id]
	if !ok {
		return nil, nil
	}

	return p, nil
}

func (i *inMemoryPumpsRepo) List(ctx context.Context) ([]*api.Pump, error) {
	list := make([]*api.Pump, len(i.pumps))
	idx := 0
	for _, p := range i.pumps {
		list[idx] = p
		idx++
	}
	return list, nil
}
