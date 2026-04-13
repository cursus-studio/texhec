package tileui

import (
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/transform"
	"engine/services/ecs"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type system struct {
	engine.World `inject:"1"`
	Ui           ui.Service   `inject:"1"`
	Tile         tile.Service `inject:"1"`

	dirtySet      ecs.DirtySet
	tileSize      transform.SizeComponent
	selectedEvent *tile.SelectEvent
}

func NewSystem(c ioc.Dic) tile.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		s.tileSize = s.Tile.GetTileSize()
		s.dirtySet = ecs.NewDirtySet()
		s.selectedEvent = nil

		s.Tile.Pos().AddDirtySet(s.dirtySet)
		s.Tile.Size().AddDirtySet(s.dirtySet)
		s.Tile.Rot().AddDirtySet(s.dirtySet)

		s.Transform.PivotPoint().AddDependency(s.Tile.Pos())
		s.Transform.Pos().AddDependency(s.Tile.Pos())
		s.Transform.Size().AddDependency(s.Tile.Size())
		s.Transform.Rotation().AddDependency(s.Tile.Rot())

		s.Transform.PivotPoint().BeforeGet(s.BeforeGet)
		s.Transform.Pos().BeforeGet(s.BeforeGet)
		s.Transform.Size().BeforeGet(s.BeforeGet)
		s.Transform.Rotation().BeforeGet(s.BeforeGet)

		events.Listen(s.EventsBuilder, s.OnUnselect)
		events.Listen(s.EventsBuilder, s.OnSelect)
		events.Listen(s.EventsBuilder, s.OnHover)
		return nil
	})
}

func (s *system) BeforeGet() {
	for _, entity := range s.dirtySet.Get() {
		pos, ok := s.Tile.Pos().Get(entity)
		if !ok {
			s.Transform.Size().Remove(entity)
			continue
		}
		size, _ := s.Tile.Size().Get(entity)
		rot, _ := s.Tile.Rot().Get(entity)
		layer, _ := s.Tile.Layer().Get(entity)

		transformPos := transform.NewPos(
			s.tileSize.Size.X()*float32(pos.X),
			s.tileSize.Size.Y()*float32(pos.Y),
			float32(layer.Z),
		)
		transformSize := transform.NewSize(
			s.tileSize.Size[0]*float32(size.X),
			s.tileSize.Size[1]*float32(size.Y),
			s.tileSize.Size[2],
		)
		transformRot := transform.NewRotation(rot.Quat())

		s.Transform.PivotPoint().Set(entity, transform.NewPivotPoint(0, 0, .5))
		s.Transform.Pos().Set(entity, transformPos)
		s.Transform.Size().Set(entity, transformSize)
		s.Transform.Rotation().Set(entity, transformRot)
	}
}

func (s *system) OnUnselect(e ui.HideUiEvent) {
	s.selectedEvent = nil
}

func (s *system) OnSelect(e tile.SelectEvent) {
	s.selectedEvent = &e
}

func (s *system) OnHover(e tile.HoverEvent) {
	if s.selectedEvent == nil {
		return
	}
	grid, ok := s.Tile.TileGrid().Get(e.Grid)
	if !ok {
		s.Logger.Warn(fmt.Errorf("grid doesn't exist"))
		return
	}
	coords := grid.GetCoords(e.Tile)
	if event, ok := s.selectedEvent.HoverEvent.(tile.ApplyCoordsEvent); ok {
		s.selectedEvent.HoverEvent = event.ApplyCoords(coords)
	}
	events.EmitAny(s.Events, s.selectedEvent.HoverEvent)
}
