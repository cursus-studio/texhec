package internal

import (
	"core/modules/definitions"
	"core/modules/deploy"
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
	Tile         tile.Service `inject:"1"`

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

func (s *service) Deploy(blueprint ecs.EntityID, coords grid.Coords) {
	// perform verification can you deploy by someone
	// if you cannot than do a flip
	// if deploy.By ? {
	//   log warning (this shouldn't be a button)
	// }
	// do not pay because this is performed by system

	deployed := s.Prototype.Clone(blueprint)
	s.Hierarchy.SetParent(deployed, s.Scene.Scene())

	s.Tile.Deployed().Set(deployed, tile.NewDeployed())
	s.Tile.Pos().Set(deployed, tile.NewPos(coords.Coords()))
	events.Emit(s.Events, ui.HideUiEvent{})
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

	s.Tile.Pos().Set(placeholderEntity, tile.NewPos(e.Coords.Coords()))
	s.placeholder.Set(placeholderEntity, placeholder{})

	{ // check can place and place
		obstructionGridEntity := s.Tile.ObstructionGrid().GetEntities()[0]
		obstructed, ok := s.Tile.ObstructionGrid().Get(obstructionGridEntity)
		if !ok {
			goto cannotPlace
		}
		index, ok := obstructed.GetIndex(e.Coords.Coords())
		if !ok {
			goto cannotPlace
		}
		blueprintObstruction, _ := s.Tile.Obstruction().Get(e.Blueprint)
		coordsObstruction := obstructed.GetTile(index)
		if blueprintObstruction.Obstruction&coordsObstruction != 0 {
			goto cannotPlace
		}
		s.Tile.ObstructionGrid().Set(obstructionGridEntity, obstructed)
		s.Render.Color().Set(placeholderEntity, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
		s.Inputs.LeftClick().Set(placeholderEntity, inputs.NewLeftClick(deploy.NewExecuteEvent(e.By, e.Blueprint).ApplyCoords(e.Coords)))
		s.Tile.Layer().Set(placeholderEntity, tile.NewLayer(definitions.PlaceholderLayer))
		return
	}
cannotPlace:
	s.Render.Color().Set(placeholderEntity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
}
func (s *service) Execute(e deploy.ExecuteEvent) {
	for _, entity := range s.placeholder.GetEntities() {
		s.RemoveEntity(entity)
	}
	// perform verification can you deploy by someone
	// if you cannot than do a flip
	// if deploy.By ? {
	//   log warning (this shouldn't be a button)
	// }
	// pay and perform everything

	deployed := s.Prototype.Clone(e.Blueprint)
	s.Hierarchy.SetParent(deployed, s.Scene.Scene())

	s.Tile.Deployed().Set(deployed, tile.NewDeployed())
	s.Tile.Pos().Set(deployed, tile.NewPos(e.Coords.Coords()))
	events.Emit(s.Events, ui.HideUiEvent{})
}
