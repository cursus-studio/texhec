package internal

import (
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/player"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/services/ecs"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type placeholder struct{}

type service struct {
	engine.World `inject:"1"`
	Tile         tile.Service   `inject:"1"`
	Player       player.Service `inject:"1"`

	component   ecs.ComponentsArray[deploy.Component]
	link        ecs.ComponentsArray[deploy.LinkComponent]
	placeholder ecs.ComponentsArray[placeholder]
}

func NewService(c ioc.Dic) deploy.Service {
	s := ioc.GetServices[*service](c)

	s.component = ecs.GetComponentsArray[deploy.Component](s.World)
	s.link = ecs.GetComponentsArray[deploy.LinkComponent](s.World)
	s.placeholder = ecs.GetComponentsArray[placeholder](s.World)

	events.Listen(s.EventsBuilder, s.Unselect)
	events.Listen(s.EventsBuilder, s.Execute)
	events.Listen(s.EventsBuilder, s.Preview)
	events.Listen(s.EventsBuilder, s.Select)

	return s
}

func (s *service) Component() ecs.ComponentsArray[deploy.Component] { return s.component }
func (s *service) Link() ecs.ComponentsArray[deploy.LinkComponent]  { return s.link }

func (s *service) Deploy(
	blueprint,
	owner ecs.EntityID,
	coords grid.Coords,
) error {
	// check can place:

	// - is position occuped
	pos := tile.NewPos(coords.Coords())
	size, _ := s.Tile.Size().Get(blueprint)
	obstruction, _ := s.Tile.Obstruction().Get(blueprint)
	aabb := tile.NewAABB(pos, size)
	if s.Tile.IsOccupied(aabb, obstruction.Obstruction) {
		return tile.ErrPositionIsOccupied
	}

	// place
	deployed := s.Prototype.Clone(blueprint)
	s.Hierarchy.SetParent(deployed, s.Scene.Scene())

	s.Player.Owner().Set(deployed, player.NewOwner(owner))
	s.Tile.Deployed().Set(deployed, tile.NewDeployed())
	s.Inputs.LeftClick().Set(deployed, inputs.NewLeftClick(tile.NewClickEntityEvent()))
	s.Tile.Pos().Set(deployed, pos)
	events.Emit(s.Events, ui.HideUiEvent{})
	return nil
}

func (s *service) Unselect(e ui.HideUiEvent) {
	for _, entity := range s.placeholder.GetEntities() {
		s.RemoveEntity(entity)
	}
}
func (s *service) Select(e deploy.SelectEvent) {
	events.Emit(s.Events, tile.NewSelectEvent(deploy.NewPreviewEvent(e.By, e.Blueprint)))
}
func (s *service) Preview(e deploy.PreviewEvent) {
	for _, entity := range s.placeholder.GetEntities() {
		s.RemoveEntity(entity)
	}
	placeholderEntity := s.Prototype.Clone(e.Blueprint)
	s.Hierarchy.SetParent(placeholderEntity, s.Scene.Scene())
	s.Tile.Layer().Set(placeholderEntity, tile.NewLayer(definitions.PlaceholderLayer))

	pos := tile.NewPos(e.Coords.Coords())
	s.Tile.Pos().Set(placeholderEntity, pos)
	s.placeholder.Set(placeholderEntity, placeholder{})
	size, _ := s.Tile.Size().Get(e.Blueprint)

	{ // check can place:
		// - is position occupied
		blueprintObstruction, _ := s.Tile.Obstruction().Get(e.Blueprint)
		aabb := tile.NewAABB(pos, size)
		if s.Tile.IsOccupied(aabb, blueprintObstruction.Obstruction) {
			goto cannotPlace
		}

		// place
		s.Render.Color().Set(placeholderEntity, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
		s.Inputs.LeftClick().Set(placeholderEntity, inputs.NewLeftClick(deploy.NewExecuteEvent(e.By, e.Blueprint).ApplyCoords(e.Coords)))
		return
	}
cannotPlace:
	s.Render.Color().Set(placeholderEntity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
}
func (s *service) Execute(e deploy.ExecuteEvent) {
	// remove placeholder entities
	for _, entity := range s.placeholder.GetEntities() {
		s.RemoveEntity(entity)
	}

	// pay
	// ...

	// check can place
	pos := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile.Size().Get(e.Blueprint)
	blueprintObstruction, _ := s.Tile.Obstruction().Get(e.Blueprint)

	obstructionGridEntity := s.Tile.ObstructionGrid().GetEntities()[0]
	obstructed, ok := s.Tile.ObstructionGrid().Get(obstructionGridEntity)
	if !ok {
		s.Logger.Warn(tile.ErrPositionIsOccupied)
		return
	}
	aabb := tile.NewAABB(pos, size)
	for _, coords := range aabb.Tiles {
		index, ok := obstructed.GetIndex(coords.Coords())
		if !ok {
			s.Logger.Warn(tile.ErrPositionIsOccupied)
			return
		}
		coordsObstruction := obstructed.GetTile(index)
		if blueprintObstruction.Obstruction&coordsObstruction != 0 {
			s.Logger.Warn(tile.ErrPositionIsOccupied)
			return
		}
	}

	// place
	deployed := s.Prototype.Clone(e.Blueprint)
	s.Hierarchy.SetParent(deployed, s.Scene.Scene())
	if owner, ok := s.Player.Owner().Get(e.By); ok {
		s.Player.Owner().Set(deployed, owner)
	}
	s.Tile.Deployed().Set(deployed, tile.NewDeployed())
	s.Inputs.LeftClick().Set(deployed, inputs.NewLeftClick(tile.NewClickEntityEvent()))
	s.Tile.Pos().Set(deployed, tile.NewPos(e.Coords.Coords()))
	events.Emit(s.Events, ui.HideUiEvent{})
}
