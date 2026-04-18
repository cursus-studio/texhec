package metadatapkg

import (
	"engine/modules/metadata"
	"engine/modules/metadata/internal"
	"engine/modules/registry"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) metadata.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, r registry.Service) {
		service := ioc.Get[metadata.Service](c)
		r.Register("name", func(entity ecs.EntityID, structTagValue string) {
			service.Name().Set(entity, metadata.NewName(structTagValue))
		})
		r.Register("description", func(entity ecs.EntityID, structTagValue string) {
			service.Description().Set(entity, metadata.NewDescription(structTagValue))
		})
	})
})
