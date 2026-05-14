package tileservice

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/modules/relation"
	"engine/modules/transform"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld   `inject:""`
	TileGridService  grid.Service[tile.ID]     `inject:""`
	TileTypeRelation relation.Service[tile.ID] `inject:""`

	ecs.SystemRegister
	renderer ecs.SystemRegister

	tile ecs.ComponentsArray[tile.TypeComponent]

	pos   ecs.ComponentsArray[tile.PosComponent]
	size  ecs.ComponentsArray[tile.SizeComponent]
	rot   ecs.ComponentsArray[tile.RotComponent]
	layer ecs.ComponentsArray[tile.LayerComponent]

	speed ecs.ComponentsArray[tile.SpeedComponent]
	step  ecs.ComponentsArray[tile.StepComponent]
}

func NewService(c ioc.Dic, system, renderer ecs.SystemRegister) tile.Service {
	s := ioc.GetServices[*service](c)
	s.SystemRegister = system
	s.renderer = renderer

	s.tile = ecs.GetComponentsArray[tile.TypeComponent](s.World())

	s.pos = ecs.GetComponentsArray[tile.PosComponent](s.World())
	s.size = ecs.GetComponentsArray[tile.SizeComponent](s.World())
	s.rot = ecs.GetComponentsArray[tile.RotComponent](s.World())
	s.layer = ecs.GetComponentsArray[tile.LayerComponent](s.World())

	s.speed = ecs.GetComponentsArray[tile.SpeedComponent](s.World())
	s.step = ecs.GetComponentsArray[tile.StepComponent](s.World())

	s.size.SetEmpty(tile.NewSize(1, 1))
	s.layer.SetEmpty(tile.NewLayer(definitions.TileLayer))

	return s
}

func (s *service) Renderer() ecs.SystemRegister { return s.renderer }

func (s *service) TileType() ecs.ComponentsArray[tile.TypeComponent] {
	return s.tile
}
func (s *service) TileGrid() ecs.ComponentsArray[grid.SquareGridComponent[tile.ID]] {
	return s.TileGridService.Component()
}
func (s *service) GetTileType(id tile.ID) (ecs.EntityID, bool) {
	return s.TileTypeRelation.Get(id)
}

func (s *service) Pos() ecs.ComponentsArray[tile.PosComponent]     { return s.pos }
func (s *service) Size() ecs.ComponentsArray[tile.SizeComponent]   { return s.size }
func (s *service) Rot() ecs.ComponentsArray[tile.RotComponent]     { return s.rot }
func (s *service) Layer() ecs.ComponentsArray[tile.LayerComponent] { return s.layer }

func (s *service) Speed() ecs.ComponentsArray[tile.SpeedComponent] { return s.speed }
func (s *service) Step() ecs.ComponentsArray[tile.StepComponent]   { return s.step }

// NewBiomAsset in other file

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
