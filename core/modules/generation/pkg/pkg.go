package generationpkg

import (
	"core/modules/generation"
	"core/modules/generation/internal"
	"engine/modules/registry"
	"engine/services/ecs"
	"engine/services/logger"
	"strconv"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) *internal.Config {
		return internal.NewConfig()
	})
	ioc.Register(b, func(c ioc.Dic) generation.Service {
		return internal.NewService(c)
	})

	ioc.Wrap(b, func(c ioc.Dic, r registry.Service) {
		type World struct {
			Config *internal.Config `inject:""`
			Logger logger.Logger    `inject:""`
		}
		r.Register("generate", func(entity ecs.EntityID, structTagValue string) {
			world := ioc.GetServices[World](c)
			chance, err := strconv.Atoi(structTagValue)
			world.Logger.Warn(err)
			if err == nil {
				world.Config.AddChance(entity, chance)
			}
		})
	})
})
