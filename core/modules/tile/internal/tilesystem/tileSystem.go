package tilesystem

import (
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/transform"
	"engine/services/ecs"
	"engine/services/frames"
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"golang.org/x/exp/constraints"
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

func abs[Number constraints.Float | constraints.Integer](n Number) Number {
	return max(-n, n)
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
		justStartedMoving := tile.Coord(int(pos.X)) == pos.X && tile.Coord(int(pos.Y)) == pos.Y
		if justStartedMoving { // check can step
			isValidStep := abs(step.X-grid.Coord(pos.X))+abs(step.Y-grid.Coord(pos.Y)) == 1
			if !isValidStep {
				s.Tile.Step().Remove(entity)
				s.Logger.Warn(tile.ErrInvalidStep)
				continue
			}

			// is step destination occupied
			size, _ := s.Tile.Size().Get(entity)
			obstruction, _ := s.Tile.Obstruction().Get(entity)
			var aabbPos tile.PosComponent
			var aabbSize tile.SizeComponent

			// aabb size
			if grid.Coord(pos.X) != step.X {
				aabbSize = tile.NewSize(1, size.Y)
			} else if grid.Coord(pos.Y) != step.Y {
				aabbSize = tile.NewSize(size.X, 1)
			}
			// aabb pos
			if grid.Coord(pos.X) < step.X {
				aabbPos = tile.NewPos(step.X+size.X-1, step.Y)
			} else if grid.Coord(pos.Y) < step.Y {
				aabbPos = tile.NewPos(step.X, step.Y+size.Y-1)
			} else {
				aabbPos = tile.NewPos(step.Coords.Coords())
			}
			// perform is step destination occupied
			if s.Tile.IsOccupied(tile.NewAABB(aabbPos, aabbSize), obstruction.Obstruction) {
				s.Tile.Step().Remove(entity)
				s.Logger.Warn(tile.ErrPositionIsOccupied)
				continue
			}
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

		if justStartedMoving {
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
