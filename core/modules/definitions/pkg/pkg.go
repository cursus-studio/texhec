package definitionspkg

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/definitions/internal"
	generationpkg "core/modules/generation/pkg"
	"core/modules/obstruction"
	"engine/modules/grid"
	"engine/modules/seed"
	_ "image/png"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		generationpkg.GeneratorPkg(definitions.CenterChunkGenerator, func(c ioc.Dic) func(grid.ChunkCoordsComponent) bool {
			world := ioc.Get[game.GameWorld](c)
			return func(g grid.ChunkCoordsComponent) bool {
				_, camera, _ := world.Lobby().GameCamera().GetSingleton()
				return g.X == g.Y && grid.Coord(camera.MaxChunksZoomOut/2) == g.X
			}
		}),
		generationpkg.GeneratorPkg(definitions.RandomChunkGenerator, func(c ioc.Dic) func(grid.ChunkCoordsComponent) bool {
			world := ioc.Get[game.GameWorld](c)
			return func(g grid.ChunkCoordsComponent) bool {
				_, s, _ := world.Seed().Seed().GetSingleton()
				// 1:100 chunks is special
				return (s.Seed*seed.Seed(g.X)*seed.Seed(g.Y))%100 == 0
			}
		}),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) definitions.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, service obstruction.Service) {
		_ = service.Obstructions().Add(definitions.AirspaceObstruction)
		_ = service.Obstructions().Add(definitions.WaterObstruction)
		_ = service.Obstructions().Add(definitions.LowlandObstruction)
	})
})
