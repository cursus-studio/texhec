package gridpkg

import (
	codecpkg "engine/modules/codec/pkg"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/grid/internal/gridcollider"
	"engine/modules/grid/internal/service"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type config[Tile grid.TileConstraint] struct {
	hoverEvent func(ecs.EntityID, grid.Index) any
}

func NewConfig[Tile grid.TileConstraint](
	hoverEvent func(ecs.EntityID, grid.Index) any,
) config[Tile] {
	return config[Tile]{hoverEvent}
}

func Pkg[Tile grid.TileConstraint](config config[Tile]) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		pkgs := []ioc.Pkg{
			codecpkg.PkgT[grid.SquareGridComponent[Tile]],

			prototypepkg.PkgT[grid.SquareGridComponent[Tile]],
		}
		for _, pkg := range pkgs {
			pkg(b)
		}
		ioc.Register(b, func(c ioc.Dic) grid.Service[Tile] {
			return service.NewService[Tile](c)
		})

		if config.hoverEvent == nil {
			return
		}
		ioc.Wrap(b, func(c ioc.Dic, collider collider.Service) {
			policy := gridcollider.NewColliderWithPolicy[Tile](
				c,
				config.hoverEvent,
			)
			collider.AddRayFallThroughPolicy(policy)
		})
	})
}
