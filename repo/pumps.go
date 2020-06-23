package repo

import (
	"context"
	"github.com/tidepool-org/devices/api"
	"github.com/tidepool-org/devices/config"
)

type PumpsRepo interface {
	GetById(ctx context.Context, id string) (*api.Pump, error)
	List(context.Context) ([]*api.Pump, error)
}

func NewPumpsRepo(devices *config.Devices) PumpsRepo {
	pumps := make(map[string]*api.Pump)
	for _, p := range devices.Pumps {
		model := pumpConfigToProto(p)
		pumps[model.GetId()] = &model
	}
	return &inMemoryPumpsRepo{pumps: pumps}
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

func pumpConfigToProto(pump *config.Pump) api.Pump {
	return api.Pump{
		Id:            pump.ID,
		DisplayName:   pump.DisplayName,
		Manufacturers: pump.Manufacturers,
		Model:         pump.Model,
	}
}
