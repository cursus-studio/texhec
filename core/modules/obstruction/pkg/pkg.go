package obstructionpkg

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/obstruction"
	"core/modules/obstruction/internal/service"
	"engine/modules/entityregistry"
	gridpkg "engine/modules/grid/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/services/ecs"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		gridpkg.PkgT(gridpkg.NewConfig[obstruction.Obstruction](nil)),
		typeregistrypkg.PkgT[obstruction.ObstructionComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) obstruction.Service {
		return service.NewService(c)
	})

	ioc.Wrap(b, func(c ioc.Dic, b entityregistry.Service) {
		world := ioc.GetServices[game.GameWorld](c)
		b.Register("obstruction", func(entity ecs.EntityID, structTagValue string) {
			var obstructionVal obstruction.Obstruction
			if strings.Contains(structTagValue, "water") {
				obstructionVal |= definitions.WaterObstruction
			}
			if strings.Contains(structTagValue, "lowland") {
				obstructionVal |= definitions.LowlandObstruction
			}
			if strings.Contains(structTagValue, "air") {
				obstructionVal |= definitions.AirspaceObstruction
			}
			world.Obstruction().Component().Set(entity, obstruction.NewObstruction(obstructionVal))
		})
	})
})
