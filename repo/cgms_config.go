package repo

import (
	"github.com/tidepool-org/devices/api"
	"github.com/tidepool-org/devices/config"
)

func CgmsConfigToProtoModels(cgms []*config.CGM) []*api.Cgm {
	models := make([]*api.Cgm, len(cgms))
	for i, cgm := range cgms {
		models[i] = CgmConfigToProto(cgm)
	}
	return models
}

func CgmConfigToProto(cgm *config.CGM) *api.Cgm {
	return &api.Cgm{
		Id:            cgm.ID,
		DisplayName:   cgm.DisplayName,
		Manufacturers: cgm.Manufacturers,
		Model:         cgm.Model,
	}
}
