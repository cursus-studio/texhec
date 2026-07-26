package tileservice

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/interactions"
	"engine/modules/relation"
	"engine/modules/transform"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld   `inject:""`
	TileGridService  grid.ServiceT[tile.ID]    `inject:""`
	TileTypeRelation relation.Service[tile.ID] `inject:""`

	CoordsInteractionService    interactions.InteractionService[tile.CoordsInteraction]    `inject:""`
	ObjectInteractionService    interactions.InteractionService[tile.ObjectInteraction]    `inject:""`
	BlueprintInteractionService interactions.InteractionService[tile.BlueprintInteraction] `inject:""`

	ecs.SystemRegister
	renderer ecs.SystemRegister

	tile ecs.ComponentArray[tile.Component]

	pos   ecs.ComponentArray[tile.PosComponent]
	size  ecs.ComponentArray[tile.SizeComponent]
	rot   ecs.ComponentArray[tile.RotComponent]
	layer ecs.ComponentArray[tile.LayerComponent]

	canDeploy    ecs.ComponentArray[tile.CanDeployComponent]
	coordsCursor ecs.ComponentArray[tile.CoordsCursorComponent]
	coordsAnchor ecs.ComponentArray[tile.CoordsAnchorComponent]

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

	s.canDeploy = ecs.GetComponentArray[tile.CanDeployComponent](s.World())
	s.coordsCursor = ecs.GetComponentArray[tile.CoordsCursorComponent](s.World())
	s.coordsAnchor = ecs.GetComponentArray[tile.CoordsAnchorComponent](s.World())

	s.size.SetEmpty(tile.NewSize(1, 1))
	s.layer.SetEmpty(tile.NewLayer(definitions.TileLayer))

	s.CoordsInteractionService.MissingPreview().OnUpsert(s.OnCoordsMissingUpsert)
	s.CoordsInteractionService.StatePreview().OnUpsert(s.OnCoordsStateUpsert)
	events.Listen(s.EventsBuilder(), s.OnTileHover)
	events.Listen(s.EventsBuilder(), s.OnTileClick)

	s.ObjectInteractionService.MissingPreview().OnUpsert(s.OnObjectMissingUpsert)
	s.ObjectInteractionService.StatePreview().OnUpsert(s.OnObjectStateUpsert)
	events.Listen(s.EventsBuilder(), s.OnClickObject)

	s.BlueprintInteractionService.MissingPreview().OnUpsert(s.OnBlueprintMissingUpsert)
	s.BlueprintInteractionService.MissingPreview().OnRemove(s.OnBlueprintMissingRemove)
	s.BlueprintInteractionService.StatePreview().OnUpsert(s.OnBlueprintStateUpsert)
	events.Listen(s.EventsBuilder(), s.OnClickBlueprint)

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

func (s *service) CanDeploy() ecs.ComponentArray[tile.CanDeployComponent] {
	return s.canDeploy
}
func (s *service) CoordsAnchor() ecs.ComponentArray[tile.CoordsAnchorComponent] {
	return s.coordsAnchor
}
func (s *service) CoordsCursor() ecs.ComponentArray[tile.CoordsCursorComponent] {
	return s.coordsCursor
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
