package tilesystem

import (
	"core/game"
	"core/modules/tile"
	"core/modules/ui"
	"engine/modules/inputs"
	"engine/modules/loop"
	"engine/modules/transform"
	"engine/services/ecs"
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"golang.org/x/exp/constraints"
)

var invSpeedTable [256]tile.Coord

type system struct {
	game.GameWorld `inject:""`

	tileSize      transform.SizeComponent
	selectedEvent *tile.SelectEvent
}

func NewSystem(c ioc.Dic) ecs.SystemRegister {
	for i := 1; i < 256; i++ {
		invSpeedTable[i] = 1. / tile.Coord(i)
	}
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		s.tileSize = s.Tile().GetTileSize()
		s.selectedEvent = nil

		//

		s.Tile().Pos().OnRemove(s.OnTilePosRemove)

		s.Tile().Pos().OnUpsert(s.OnTilePosSizeRotUpsert)
		s.Tile().Size().OnUpsert(s.OnTilePosSizeRotUpsert)
		s.Tile().Rot().OnUpsert(s.OnTilePosSizeRotUpsert)

		s.Obstruction().Deployed().OnUpsert(s.OnDeployUpsert)

		//

		events.Listen(s.EventsBuilder(), s.StepOnTick)
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

func (s *system) OnDeployUpsert(entity ecs.EntityID) {
	s.Inputs().Stack().Set(entity, inputs.StackComponent{})
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

var (
	rotLeft  = tile.NewRot(mgl32.DegToRad(90))
	rotRight = tile.NewRot(mgl32.DegToRad(270))
	rotUp    = tile.NewRot(mgl32.DegToRad(0))
	rotDown  = tile.NewRot(mgl32.DegToRad(180))
)

func abs[Number constraints.Float | constraints.Integer](n Number) Number { return max(-n, n) }

func (s *system) StepOnTick(e loop.TickEvent) {
	entities := s.Tile().Step().GetEntities()
	{
		cp := make([]ecs.EntityID, len(entities))
		copy(cp, entities)
		entities = cp
	}
	for _, entity := range entities {
		step, ok := s.Tile().Step().Get(entity)
		if !ok {
			continue
		}
		pos, ok := s.Tile().Pos().Get(entity)
		if !ok {
			s.Tile().Step().Remove(entity)
			s.Logger().Log(tile.ErrPositionAndSpeedIsRequiredToStep)
			continue
		}
		speed, ok := s.Tile().Speed().Get(entity)
		if !ok {
			s.Tile().Step().Remove(entity)
			s.Logger().Log(tile.ErrPositionAndSpeedIsRequiredToStep)
			continue
		}
		arrived := tile.Coord(step.X) == pos.X && tile.Coord(step.Y) == pos.Y
		if arrived {
			s.Tile().Step().Remove(entity)
			continue
		}
		size, _ := s.Tile().Size().Get(entity)
		obstruction, _ := s.Obstruction().Component().Get(entity)
		aligned, isFirstStep := pos.Aligned()
		if isFirstStep && !s.Obstruction().CanStep(aligned, size, obstruction, step) {
			s.Tile().Step().Remove(entity)
			s.Logger().Log(tile.ErrInvalidStep)
			continue
		}

		// move
		var rot tile.RotComponent
		stepSpeed := invSpeedTable[speed.InvSpeed]
		if tile.Coord(step.X) > pos.X {
			pos.X = min(pos.X+stepSpeed, tile.Coord(step.X))
			rot = rotRight
		} else if tile.Coord(step.X) < pos.X {
			pos.X = max(pos.X-stepSpeed, tile.Coord(step.X))
			rot = rotLeft
		} else if tile.Coord(step.Y) > pos.Y {
			pos.Y = min(pos.Y+stepSpeed, tile.Coord(step.Y))
			rot = rotUp
		} else if tile.Coord(step.Y) < pos.Y {
			pos.Y = max(pos.Y-stepSpeed, tile.Coord(step.Y))
			rot = rotDown
		} else {
			s.Logger().Log(fmt.Errorf("tile system isn't able to handle StepComponent"))
		}
		const epsilon tile.Coord = 1e-3
		if abs(tile.Coord(step.X)-pos.X) < epsilon {
			pos.X = tile.Coord(step.X)
		}
		if abs(tile.Coord(step.Y)-pos.Y) < epsilon {
			pos.Y = tile.Coord(step.Y)
		}
		s.Tile().Pos().Set(entity, pos)

		if isFirstStep {
			s.Tile().Rot().Set(entity, rot)
		}

		arrived = tile.Coord(step.X) == pos.X && tile.Coord(step.Y) == pos.Y
		if arrived {
			s.Tile().Step().Remove(entity)
		}
	}
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
	grid, ok := s.Tile().TileGrid().Get(e.Grid)
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
