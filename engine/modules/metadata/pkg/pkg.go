package metadatapkg

import (
	"engine/modules/metadata"
	"engine/modules/metadata/internal"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/registry"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		prototypepkg.PkgT[metadata.LinkComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) metadata.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, r registry.Service) {
		service := ioc.Get[metadata.Service](c)
		r.Register("name", func(entity ecs.EntityID, structTagValue string) {
			service.Name().Set(entity, metadata.NewName(structTagValue))
			service.Link().Set(entity, metadata.NewLink(entity))
		})
		r.Register("description", func(entity ecs.EntityID, structTagValue string) {
			service.Description().Set(entity, metadata.NewDescription(structTagValue))
			service.Link().Set(entity, metadata.NewLink(entity))
		})
	})
})
