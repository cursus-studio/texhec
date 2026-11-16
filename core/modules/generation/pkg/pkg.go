package generationpkg

import (
	"core/game"
	"core/modules/generation"
	"core/modules/generation/internal"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	"engine/modules/grid"
	"math/bits"
	"runtime"
	"strconv"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) *internal.Config {
		world := ioc.Get[game.GameWorld](c)
		s := int(world.Grid().ChunkSize())
		s /= runtime.NumCPU()
		s = 1 << max(bits.Len(uint(s))-1, 1)
		return internal.NewConfig(s)
	})
	ioc.Register(b, func(c ioc.Dic) internal.Service {
		return internal.NewService(c)
	})
	ioc.Register(b, func(c ioc.Dic) generation.Service {
		return ioc.Get[internal.Service](c)
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
		r.Register("chunk_generate", func(entity ecs.EntityID, structTagValue string) {
			world := ioc.GetServices[World](c)
			world.Config.AddChunkGenerate(entity)
		})
	})
})

//

type isSpecialService struct {
	isSpecial func(grid.ChunkCoordsComponent) bool
}

func (s *isSpecialService) IsSpecial(coords grid.ChunkCoordsComponent) bool {
	return s.isSpecial(coords)
}

func GeneratorPkg(
	generator generation.ChunkGenerator,
	isSpecial func(c ioc.Dic) func(grid.ChunkCoordsComponent) bool,
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Wrap(b, func(c ioc.Dic, s internal.Service) {
			s.RegisterGenerator(generator, &isSpecialService{isSpecial(c)})
		})
	})
}
