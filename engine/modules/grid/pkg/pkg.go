package gridpkg

import (
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/grid/internal/gridcollider"
	"engine/modules/grid/internal/service"
	relationpkg "engine/modules/relation/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type config struct {
	chunkSize grid.ChunkSize
}

func NewConfig() *config {
	return &config{grid.NewChunkSize(5)}
}
func (c *config) GetChunkSize() grid.ChunkSize  { return c.chunkSize }
func (c *config) SetChunkSize(s grid.ChunkSize) { c.chunkSize = s }

type Config interface {
	GetChunkSize() grid.ChunkSize
	SetChunkSize(grid.ChunkSize)
}

//

type configT[Tile grid.TileConstraint] struct {
	hoverEvent func(ecs.EntityID, grid.Coords) any
}

func NewConfigT[Tile grid.TileConstraint]() *configT[Tile] {
	return &configT[Tile]{nil}
}
func (c *configT[Tile]) GetHoverEvent() func(ecs.EntityID, grid.Coords) any { return c.hoverEvent }
func (c *configT[Tile]) SetHoverEvent(hoverEvent func(ecs.EntityID, grid.Coords) any) {
	c.hoverEvent = hoverEvent
}

type ConfigT[Tile grid.TileConstraint] interface {
	GetHoverEvent() func(ecs.EntityID, grid.Coords) any
	SetHoverEvent(func(ecs.EntityID, grid.Coords) any)
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		relationpkg.MapRelationPkg(
			func(w ecs.World) ecs.DirtySet {
				set := ecs.NewDirtySet()
				ecs.GetComponentsArray[grid.ChunkCoordsComponent](w).AddDirtySet(set)
				return set
			},
			func(w ecs.World) func(entity ecs.EntityID) (indexType grid.ChunkCoordsComponent, ok bool) {
				arr := ecs.GetComponentsArray[grid.ChunkCoordsComponent](w)
				return func(entity ecs.EntityID) (indexType grid.ChunkCoordsComponent, ok bool) {
					return arr.Get(entity)
				}
			},
		),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) Config {
		return NewConfig()
	})
	ioc.Register(b, func(c ioc.Dic) grid.Service {
		return service.NewService(c, ioc.Get[Config](c).GetChunkSize())
	})
})

func PkgT[Tile grid.TileConstraint]() ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		pkgs := []ioc.Pkg{
			typeregistrypkg.PkgT[grid.ChunkComponent[Tile]],
		}
		for _, pkg := range pkgs {
			pkg(b)
		}
		ioc.Register(b, func(c ioc.Dic) ConfigT[Tile] {
			return NewConfigT[Tile]()
		})
		ioc.Register(b, func(c ioc.Dic) grid.ServiceT[Tile] {
			return service.NewServiceT[Tile](c)
		})

		ioc.Wrap(b, func(c ioc.Dic, collider collider.Service) {
			config := ioc.Get[ConfigT[Tile]](c)
			if config.GetHoverEvent() == nil {
				return
			}
			policy := gridcollider.NewColliderWithPolicy[Tile](
				c,
				config.GetHoverEvent(),
			)
			collider.AddRayFallThroughPolicy(policy)
		})
	})
}
