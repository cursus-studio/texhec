package gridpkg

import (
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/grid/internal/service"
	relationpkg "engine/modules/relation/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

type config struct {
	chunkSize grid.ChunkSize
	tileSize  float32
}

func NewConfig() *config {
	return &config{
		grid.NewChunkSize(5),
		100,
	}
}
func (c *config) GetTileSize() float32          { return c.tileSize }
func (c *config) SetTileSize(s float32)         { c.tileSize = s }
func (c *config) GetChunkSize() grid.ChunkSize  { return c.chunkSize }
func (c *config) SetChunkSize(s grid.ChunkSize) { c.chunkSize = s }

type Config interface {
	GetTileSize() float32
	SetTileSize(float32)
	GetChunkSize() grid.ChunkSize
	SetChunkSize(grid.ChunkSize)
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[grid.Coords],
		typeregistrypkg.PkgT[grid.ChunkCoordsComponent],
		typeregistrypkg.PkgT[grid.ClickEvent],
		typeregistrypkg.PkgT[grid.HoverEvent],

		relationpkg.MapRelationPkg(
			func(w ecs.World) ecs.DirtySet {
				set := ecs.NewDirtySet()
				ecs.GetComponentArray[grid.ChunkCoordsComponent](w).AddDirtySet(set)
				return set
			},
			func(w ecs.World) func(entity ecs.EntityID) (indexType grid.ChunkCoordsComponent, ok bool) {
				arr := ecs.GetComponentArray[grid.ChunkCoordsComponent](w)
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
		config := ioc.Get[Config](c)
		return service.NewService(c, config.GetTileSize(), config.GetChunkSize())
	})
})

func PkgT[Tile grid.TileConstraint](b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[grid.ChunkComponent[Tile]],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) grid.ServiceT[Tile] {
		return service.NewServiceT[Tile](c)
	})
}
