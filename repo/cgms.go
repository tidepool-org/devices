package repo

import (
	"context"
	"github.com/tidepool-org/devices/api"
)

type CgmsRepo interface {
	GetById(ctx context.Context, id string) (*api.Cgm, error)
	List(context.Context) ([]*api.Cgm, error)
}

func NewCgmsRepo(cgms []*api.Cgm) CgmsRepo {
	cgmsIdMap := make(map[string]*api.Cgm)
	for _, cgm := range cgms {
		cgmsIdMap[cgm.GetId()] = cgm
	}
	return &inMemoryCgmsRepo{cgms: cgmsIdMap}
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

