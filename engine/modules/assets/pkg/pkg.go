package assetspkg

import (
	"engine/modules/assets"
	"engine/modules/assets/internal"
	"engine/modules/registry"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

// can later append many parent directories
var parentDirectory = "assets/"

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	if len(parentDirectory) != 0 && parentDirectory[len(parentDirectory)-1] != '/' {
		parentDirectory += "/"
	}

	ioc.Register(b, func(c ioc.Dic) assets.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, registry registry.Service) {
		registry.Register("path", func(entity ecs.EntityID, structTagValue string) {
			assetsService := ioc.Get[assets.Service](c)
			path := assets.NewPath(parentDirectory + structTagValue)
			assetsService.Path().Set(entity, path)
		})
	})
})
