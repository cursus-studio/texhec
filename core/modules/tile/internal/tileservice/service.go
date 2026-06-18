package tileservice

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/modules/interactions"
	"engine/modules/relation"
	"engine/modules/transform"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld   `inject:""`
	TileGridService  grid.ServiceT[tile.ID]    `inject:""`
	TileTypeRelation relation.Service[tile.ID] `inject:""`

	CoordsInteractionService       interactions.InteractionService[tile.CoordsInteraction]       `inject:""`
	ObjectInteractionService       interactions.InteractionService[tile.ObjectInteraction]       `inject:""`
	SourceObjectInteractionService interactions.InteractionService[tile.SourceObjectInteraction] `inject:""`

	ecs.SystemRegister
	renderer ecs.SystemRegister

	tile ecs.ComponentsArray[tile.Component]

	pos   ecs.ComponentsArray[tile.PosComponent]
	size  ecs.ComponentsArray[tile.SizeComponent]
	rot   ecs.ComponentsArray[tile.RotComponent]
	layer ecs.ComponentsArray[tile.LayerComponent]

	coordsCursor      ecs.ComponentsArray[tile.CoordsCursorComponent]
	coordsCursorRange ecs.ComponentsArray[tile.CoordsCursorRangeComponent]

	tileSize float32
}

func NewService(c ioc.Dic, system, renderer ecs.SystemRegister, tileSize float32) tile.Service {
	s := ioc.GetServices[*service](c)
	s.SystemRegister = system
	s.renderer = renderer

	s.tile = ecs.GetComponentsArray[tile.Component](s.World())

	s.pos = ecs.GetComponentsArray[tile.PosComponent](s.World())
	s.size = ecs.GetComponentsArray[tile.SizeComponent](s.World())
	s.rot = ecs.GetComponentsArray[tile.RotComponent](s.World())
	s.layer = ecs.GetComponentsArray[tile.LayerComponent](s.World())

	s.coordsCursor = ecs.GetComponentsArray[tile.CoordsCursorComponent](s.World())
	s.coordsCursorRange = ecs.GetComponentsArray[tile.CoordsCursorRangeComponent](s.World())

	s.size.SetEmpty(tile.NewSize(1, 1))
	s.layer.SetEmpty(tile.NewLayer(definitions.TileLayer))

	events.Listen(s.EventsBuilder(), s.OnClickEntityRenderFeatures)

	events.Listen(s.EventsBuilder(), s.OnTileHover)
	events.Listen(s.EventsBuilder(), s.OnTileClick)
	s.CoordsInteractionService.Interaction().OnUpsert(s.OnCoordsInteractionUpsert)

	s.ObjectInteractionService.MissingInteraction().OnUpsert(s.OnMissingObjectInteractionUpsert)
	s.ObjectInteractionService.MissingInteraction().OnUpsert(s.OnMissingObjectInteractionRemove)
	s.ObjectInteractionService.Interaction().OnUpsert(s.OnObjectInteractionUpsert)
	events.Listen(s.EventsBuilder(), s.OnClickEntitySelect)

	s.SourceObjectInteractionService.MissingInteraction().OnUpsert(s.OnMissingSourceObjectInteractionUpsert)
	s.SourceObjectInteractionService.MissingInteraction().OnRemove(s.OnMissingSourceObjectInteractionRemove)
	s.SourceObjectInteractionService.Interaction().OnMod(s.OnSourceObjectMod)

	s.tileSize = tileSize

	return s
}

func (s *service) Renderer() ecs.SystemRegister { return s.renderer }

func (s *service) Component() ecs.ComponentsArray[tile.Component] {
	return s.tile
}
func (s *service) Grid() grid.ServiceT[tile.ID] { return s.TileGridService }
func (s *service) GetTile(id tile.ID) (ecs.EntityID, bool) {
	return s.TileTypeRelation.Get(id)
}

func (s *service) Pos() ecs.ComponentsArray[tile.PosComponent]     { return s.pos }
func (s *service) Size() ecs.ComponentsArray[tile.SizeComponent]   { return s.size }
func (s *service) Rot() ecs.ComponentsArray[tile.RotComponent]     { return s.rot }
func (s *service) Layer() ecs.ComponentsArray[tile.LayerComponent] { return s.layer }

func (s *service) CoordsCursor() ecs.ComponentsArray[tile.CoordsCursorComponent] {
	return s.coordsCursor
}
func (s *service) CoordsCursorRange() ecs.ComponentsArray[tile.CoordsCursorRangeComponent] {
	return s.coordsCursorRange
}

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
