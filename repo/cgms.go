package repo

import (
	"context"
	"github.com/tidepool-org/devices/api"
	"github.com/tidepool-org/devices/config"
)

type CgmsRepo interface {
	GetById(ctx context.Context, id string) (*api.Cgm, error)
	List(context.Context) ([]*api.Cgm, error)
}

func NewCgmsRepo(devices *config.Devices) CgmsRepo {
	cgms := make(map[string]*api.Cgm)
	for _, p := range devices.CGMs {
		model := cgmConfigToProto(p)
		cgms[model.GetId()] = &model
	}
	return &inMemoryCgmsRepo{cgms: cgms}
}

type inMemoryCgmsRepo struct {
	cgms map[string]*api.Cgm
}

func (i *inMemoryCgmsRepo) GetById(ctx context.Context, id string) (*api.Cgm, error) {
	p, ok := i.cgms[id]
	if !ok {
		return nil, nil
	}

	return p, nil
}

func (i *inMemoryCgmsRepo) List(ctx context.Context) ([]*api.Cgm, error) {
	list := make([]*api.Cgm, len(i.cgms))
	idx := 0
	for _, p := range i.cgms {
		list[idx] = p
		idx++
	}
	return list, nil
}

func cgmConfigToProto(cgm *config.CGM) api.Cgm {
	return api.Cgm{
		Id:            cgm.ID,
		DisplayName:   cgm.DisplayName,
		Manufacturers: cgm.Manufacturers,
		Model:         cgm.Model,
	}
}
