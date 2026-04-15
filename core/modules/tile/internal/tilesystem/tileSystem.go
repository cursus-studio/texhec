package tilesystem

import (
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/inputs"
	"engine/modules/transform"
	"engine/services/ecs"
	"engine/services/frames"
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var invSpeedTable [256]tile.Coord

type system struct {
	engine.World `inject:"1"`
	Ui           ui.Service   `inject:"1"`
	Tile         tile.Service `inject:"1"`

	dirtyDeployedSet  ecs.DirtySet
	dirtyTransformSet ecs.DirtySet
	tileSize          transform.SizeComponent
	selectedEvent     *tile.SelectEvent
}

func NewSystem(c ioc.Dic) tile.System {
	for i := 1; i < 256; i++ {
		invSpeedTable[i] = 1. / tile.Coord(i)
	}
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		s.tileSize = s.Tile.GetTileSize()
		s.dirtyDeployedSet = ecs.NewDirtySet()
		s.dirtyTransformSet = ecs.NewDirtySet()
		s.selectedEvent = nil

		//
		s.Tile.Pos().AddDirtySet(s.dirtyTransformSet)
		s.Tile.Size().AddDirtySet(s.dirtyTransformSet)
		s.Tile.Rot().AddDirtySet(s.dirtyTransformSet)

		s.Transform.PivotPoint().AddDependency(s.Tile.Pos())
		s.Transform.Pos().AddDependency(s.Tile.Pos())
		s.Transform.Size().AddDependency(s.Tile.Size())
		s.Transform.Rotation().AddDependency(s.Tile.Rot())

		s.Transform.PivotPoint().BeforeGet(s.BeforeTransformGet)
		s.Transform.Pos().BeforeGet(s.BeforeTransformGet)
		s.Transform.Size().BeforeGet(s.BeforeTransformGet)
		s.Transform.Rotation().BeforeGet(s.BeforeTransformGet)

		//

		s.Tile.Deployed().AddDirtySet(s.dirtyDeployedSet)
		s.Inputs.Stack().AddDependency(s.Tile.Deployed())

		s.Inputs.Stack().BeforeGet(s.BeforeStackGet)

		//

		events.Listen(s.EventsBuilder, s.OnTick)
		events.Listen(s.EventsBuilder, s.OnUnselect)
		events.Listen(s.EventsBuilder, s.OnSelect)
		events.Listen(s.EventsBuilder, s.OnHover)
		return nil
	})
}

func (s *system) BeforeStackGet() {
	for _, entity := range s.dirtyDeployedSet.Get() {
		if _, ok := s.Tile.Deployed().Get(entity); !ok {
			s.Inputs.Stack().Remove(entity)
			continue
		}

		s.Inputs.Stack().Set(entity, inputs.StackComponent{})
	}
}

func (s *system) BeforeTransformGet() {
	for _, entity := range s.dirtyTransformSet.Get() {
		pos, ok := s.Tile.Pos().Get(entity)
		if !ok {
			s.Transform.Size().Remove(entity)
			s.Inputs.Stack().Remove(entity)
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

var (
	rotLeft  = tile.NewRot(mgl32.DegToRad(90))
	rotRight = tile.NewRot(mgl32.DegToRad(270))
	rotUp    = tile.NewRot(mgl32.DegToRad(0))
	rotDown  = tile.NewRot(mgl32.DegToRad(180))
)

func (s *system) OnTick(e frames.TickEvent) {
	entities := s.Tile.Step().GetEntities()
	{
		cp := make([]ecs.EntityID, len(entities))
		copy(cp, entities)
		entities = cp
	}
	for _, entity := range entities {
		step, ok := s.Tile.Step().Get(entity)
		if !ok {
			continue
		}
		pos, ok := s.Tile.Pos().Get(entity)
		if !ok {
			s.Tile.Step().Remove(entity)
			s.Logger.Warn(tile.ErrPositionAndSpeedIsRequiredToStep)
			continue
		}
		speed, ok := s.Tile.Speed().Get(entity)
		if !ok {
			s.Tile.Step().Remove(entity)
			s.Logger.Warn(tile.ErrPositionAndSpeedIsRequiredToStep)
			continue
		}
		arrived := tile.Coord(step.X) == pos.X && tile.Coord(step.Y) == pos.Y
		if arrived {
			s.Tile.Step().Remove(entity)
			continue
		}
		isFirstStep := pos.IsAligned()
		if isFirstStep && !s.Tile.CanStep(entity, step) {
			s.Tile.Step().Remove(entity)
			s.Logger.Warn(tile.ErrInvalidStep)
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
			s.Logger.Warn(fmt.Errorf("tile system isn't able to handle StepComponent"))
		}
		s.Tile.Pos().Set(entity, pos)

		if isFirstStep {
			s.Tile.Rot().Set(entity, rot)
		}

		arrived = tile.Coord(step.X) == pos.X && tile.Coord(step.Y) == pos.Y
		if arrived {
			s.Tile.Step().Remove(entity)
		}
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
