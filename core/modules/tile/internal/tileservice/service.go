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
	s.obstruction.SetEmpty(tile.NewObstruction(definitions.LowlandObstruction))

	return s
}

func (s *service) TileType() ecs.ComponentsArray[tile.TypeComponent] {
	return s.tile
}
func (s *service) TileGrid() ecs.ComponentsArray[grid.SquareGridComponent[tile.ID]] {
	return s.TileGridService.Component()
}
func (s *service) ObstructionGrid() ecs.ComponentsArray[grid.SquareGridComponent[tile.Obstruction]] {
	return s.ObstructionGridService.Component()
}
func (s *service) GetTileType(id tile.ID) (ecs.EntityID, bool) {
	return s.TileTypeRelation.Get(id)
}

func (s *service) Pos() ecs.ComponentsArray[tile.PosComponent]                 { return s.pos }
func (s *service) Size() ecs.ComponentsArray[tile.SizeComponent]               { return s.size }
func (s *service) Rot() ecs.ComponentsArray[tile.RotComponent]                 { return s.rot }
func (s *service) Layer() ecs.ComponentsArray[tile.LayerComponent]             { return s.layer }
func (s *service) Obstruction() ecs.ComponentsArray[tile.ObstructionComponent] { return s.obstruction }
func (s *service) Deployed() ecs.ComponentsArray[tile.DeployedComponent]       { return s.deployed }

func (s *service) GetPos(coords grid.Coords) transform.PosComponent {
	size := s.GetTileSize().Size
	return transform.NewPos(
		size.X()*(float32(coords.X)+.5),
		size.Y()*(float32(coords.Y)+.5),
		size.Z(),
	)
}
func (s *service) GetTileSize() transform.SizeComponent {
	return transform.NewSize(100, 100, 1)
}

func (s *service) IsOccupied(aabb tile.AABB, obstruction tile.Obstruction) bool {
	obstructionGridEntity := s.ObstructionGrid().GetEntities()[0]
	obstructed, ok := s.ObstructionGrid().Get(obstructionGridEntity)
	if !ok {
		return true
	}
	for _, coords := range aabb.Tiles {
		index, ok := obstructed.GetIndex(coords.Coords())
		if !ok {
			return true
		}
		coordsObstruction := obstructed.GetTile(index)
		if obstruction&coordsObstruction != 0 {
			return true
		}
	}
	return false
}
