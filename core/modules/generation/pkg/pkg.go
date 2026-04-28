package generationpkg

import (
	"core/game"
	"core/modules/generation"
	"core/modules/generation/internal"
	"engine/modules/entityregistry"
	"engine/services/ecs"
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

	ioc.Wrap(b, func(c ioc.Dic, r entityregistry.Service) {
		type World struct {
			game.GameWorld `inject:""`
			Config         *internal.Config `inject:""`
		}
		r.Register("generate", func(entity ecs.EntityID, structTagValue string) {
			world := ioc.GetServices[World](c)
			chance, err := strconv.Atoi(structTagValue)
			world.Logger().Log(err)
			if err == nil {
				world.Config.AddChance(entity, chance)
			}
		})
	})
})
