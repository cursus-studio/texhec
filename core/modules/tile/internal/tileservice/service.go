package tileservice

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/relation"
	"engine/modules/transform"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld   `inject:""`
	TileGridService  grid.ServiceT[tile.ID]    `inject:""`
	TileTypeRelation relation.Service[tile.ID] `inject:""`

	ecs.SystemRegister
	renderer ecs.SystemRegister

	tile ecs.ComponentArray[tile.Component]

	pos   ecs.ComponentArray[tile.PosComponent]
	size  ecs.ComponentArray[tile.SizeComponent]
	rot   ecs.ComponentArray[tile.RotComponent]
	layer ecs.ComponentArray[tile.LayerComponent]

	tileSize float32
}

func NewService(c ioc.Dic, system, renderer ecs.SystemRegister, tileSize float32) tile.Service {
	s := ioc.GetServices[*service](c)
	s.SystemRegister = system
	s.renderer = renderer

	s.tile = ecs.GetComponentArray[tile.Component](s.World())

	s.pos = ecs.GetComponentArray[tile.PosComponent](s.World())
	s.size = ecs.GetComponentArray[tile.SizeComponent](s.World())
	s.rot = ecs.GetComponentArray[tile.RotComponent](s.World())
	s.layer = ecs.GetComponentArray[tile.LayerComponent](s.World())

	s.size.SetEmpty(tile.NewSize(1, 1))
	s.layer.SetEmpty(tile.NewLayer(definitions.TileLayer))

	s.tileSize = tileSize

	return s
}

func (s *service) Renderer() ecs.SystemRegister { return s.renderer }

func (s *service) Component() ecs.ComponentArray[tile.Component] {
	return s.tile
}
func (s *service) Grid() grid.ServiceT[tile.ID] { return s.TileGridService }
func (s *service) GetTile(id tile.ID) (ecs.EntityID, bool) {
	return s.TileTypeRelation.Get(id)
}

func (s *service) Pos() ecs.ComponentArray[tile.PosComponent]     { return s.pos }
func (s *service) Size() ecs.ComponentArray[tile.SizeComponent]   { return s.size }
func (s *service) Rot() ecs.ComponentArray[tile.RotComponent]     { return s.rot }
func (s *service) Layer() ecs.ComponentArray[tile.LayerComponent] { return s.layer }

// NewBiomeAsset in other file

func (s *service) GetPos(coords grid.Coords) transform.PosComponent {
	size := s.GetTileSize().Size
	return transform.NewPos(
		size.X()*(float32(coords.X)+.5),
		size.Y()*(float32(coords.Y)+.5),
		size.Z(),
	)
}
func (s *service) GetTileSize() transform.SizeComponent {
	return transform.NewSize(s.tileSize, s.tileSize, 1)
}
