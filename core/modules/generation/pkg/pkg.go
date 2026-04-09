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

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) *internal.Config {
		return internal.NewConfig()
	})
	ioc.RegisterSingleton(b, func(c ioc.Dic) generation.Service {
		return internal.NewService(c)
	})

	ioc.WrapService(b, func(c ioc.Dic, r registry.Service) {
		type World struct {
			Config *internal.Config `inject:"1"`
			Logger logger.Logger    `inject:"1"`
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
}
