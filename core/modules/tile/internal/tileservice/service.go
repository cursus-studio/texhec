package tileservice

import (
	"core/modules/definitions"
	"core/modules/tile"
	"engine"
	"engine/modules/grid"
	"engine/modules/relation"
	"engine/modules/transform"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.World           `inject:"1"`
	TileGridService        grid.Service[tile.ID]          `inject:"1"`
	ObstructionGridService grid.Service[tile.Obstruction] `inject:"1"`
	TileTypeRelation       relation.Service[tile.ID]      `inject:"1"`

	tile ecs.ComponentsArray[tile.TypeComponent]

	pos         ecs.ComponentsArray[tile.PosComponent]
	size        ecs.ComponentsArray[tile.SizeComponent]
	rot         ecs.ComponentsArray[tile.RotComponent]
	layer       ecs.ComponentsArray[tile.LayerComponent]
	obstruction ecs.ComponentsArray[tile.ObstructionComponent]
	deployed    ecs.ComponentsArray[tile.DeployedComponent]
}

func NewService(c ioc.Dic) tile.Service {
	s := ioc.GetServices[*service](c)
	s.tile = ecs.GetComponentsArray[tile.TypeComponent](s.World)

	s.pos = ecs.GetComponentsArray[tile.PosComponent](s.World)
	s.size = ecs.GetComponentsArray[tile.SizeComponent](s.World)
	s.rot = ecs.GetComponentsArray[tile.RotComponent](s.World)
	s.layer = ecs.GetComponentsArray[tile.LayerComponent](s.World)
	s.obstruction = ecs.GetComponentsArray[tile.ObstructionComponent](s.World)
	s.deployed = ecs.GetComponentsArray[tile.DeployedComponent](s.World)

	s.size.SetEmpty(tile.NewSize(1, 1))
	s.layer.SetEmpty(tile.NewLayer(definitions.TileLayer))

	return s
}

func (t *service) TileType() ecs.ComponentsArray[tile.TypeComponent] {
	return t.tile
}
func (t *service) TileGrid() ecs.ComponentsArray[grid.SquareGridComponent[tile.ID]] {
	return t.TileGridService.Component()
}
func (t *service) ObstructionGrid() ecs.ComponentsArray[grid.SquareGridComponent[tile.Obstruction]] {
	return t.ObstructionGridService.Component()
}
func (t *service) GetTileType(id tile.ID) (ecs.EntityID, bool) {
	return t.TileTypeRelation.Get(id)
}

func (t *service) Pos() ecs.ComponentsArray[tile.PosComponent]                 { return t.pos }
func (t *service) Size() ecs.ComponentsArray[tile.SizeComponent]               { return t.size }
func (t *service) Rot() ecs.ComponentsArray[tile.RotComponent]                 { return t.rot }
func (t *service) Layer() ecs.ComponentsArray[tile.LayerComponent]             { return t.layer }
func (t *service) Obstruction() ecs.ComponentsArray[tile.ObstructionComponent] { return t.obstruction }
func (t *service) Deployed() ecs.ComponentsArray[tile.DeployedComponent]       { return t.deployed }

func (t *service) GetPos(coords grid.Coords) transform.PosComponent {
	size := t.GetTileSize().Size
	return transform.NewPos(
		size.X()*(float32(coords.X)+.5),
		size.Y()*(float32(coords.Y)+.5),
		size.Z(),
	)
}
func (t *service) GetTileSize() transform.SizeComponent {
	return transform.NewSize(100, 100, 1)
}
