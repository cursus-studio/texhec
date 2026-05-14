package tilesystem

import (
	"core/game"
	"core/modules/tile"
	"core/modules/ui"
	"engine/modules/transform"
	"engine/services/ecs"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type system struct {
	game.GameWorld `inject:""`

	tileSize      transform.SizeComponent
	selectedEvent *tile.SelectEvent
}

func NewSystem(c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		s.tileSize = s.Tile().GetTileSize()
		s.selectedEvent = nil

		//

		s.Tile().Pos().OnRemove(s.OnTilePosRemove)

		s.Tile().Pos().OnUpsert(s.OnTilePosSizeRotUpsert)
		s.Tile().Size().OnUpsert(s.OnTilePosSizeRotUpsert)
		s.Tile().Rot().OnUpsert(s.OnTilePosSizeRotUpsert)

		//

		events.Listen(s.EventsBuilder(), s.OnUnselect)
		events.Listen(s.EventsBuilder(), s.OnSelect)
		events.Listen(s.EventsBuilder(), s.OnHover)
		return nil
	})
}

func (s *system) OnTilePosRemove(entity ecs.EntityID) {
	s.Transform().Size().Remove(entity)
	s.Inputs().Stack().Remove(entity)
}

func (s *system) OnTilePosSizeRotUpsert(entity ecs.EntityID) {
	pos, ok := s.Tile().Pos().Get(entity)
	if !ok {
		return
	}
	size, _ := s.Tile().Size().Get(entity)
	rot, _ := s.Tile().Rot().Get(entity)
	layer, _ := s.Tile().Layer().Get(entity)

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

	s.Transform().PivotPoint().Set(entity, transform.NewPivotPoint(0, 0, .5))
	s.Transform().Pos().Set(entity, transformPos)
	s.Transform().Size().Set(entity, transformSize)
	s.Transform().Rotation().Set(entity, transformRot)
}

func (s *system) OnUnselect(ui.UnselectEvent[ui.ObjectComponent]) {
	s.selectedEvent = nil
}

func (s *system) OnSelect(e tile.SelectEvent) {
	s.selectedEvent = &e
}

func (s *system) OnHover(e tile.HoverEvent) {
	if s.selectedEvent == nil {
		return
	}
	grid, ok := s.Tile().Grid().Get(e.Grid)
	if !ok {
		s.Logger().Log(fmt.Errorf("grid doesn't exist"))
		return
	}
	coords := grid.GetCoords(e.Tile)
	if event, ok := s.selectedEvent.HoverEvent.(tile.ApplyCoordsEvent); ok {
		s.selectedEvent.HoverEvent = event.ApplyCoords(coords)
	}
	events.EmitAny(s.Events(), s.selectedEvent.HoverEvent)
}
