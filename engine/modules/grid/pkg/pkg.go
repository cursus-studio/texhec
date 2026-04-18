package gridpkg

import (
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/grid/internal/gridcollider"
	"engine/modules/grid/internal/service"
	"engine/services/codec"
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
		ioc.Register(b, func(c ioc.Dic) grid.Service[Tile] {
			return service.NewService[Tile](c)
		})

		ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
			b.
				// components
				Register(grid.SquareGridComponent[Tile]{})
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
