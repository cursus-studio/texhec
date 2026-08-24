package definitionspkg

import (
	"core/modules/definitions"
	"core/modules/definitions/internal"
	"core/modules/obstruction"
	_ "image/png"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) definitions.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, service obstruction.Service) {
		_ = service.Obstructions().Add(definitions.AirspaceObstruction)
		_ = service.Obstructions().Add(definitions.WaterObstruction)
		_ = service.Obstructions().Add(definitions.LowlandObstruction)
	})
})
