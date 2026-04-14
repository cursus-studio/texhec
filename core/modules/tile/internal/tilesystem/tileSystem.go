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

func (s *system) OnTick(e frames.TickEvent) {
	entities := s.Tile.Destination().GetEntities()
	{
		cp := make([]ecs.EntityID, len(entities))
		copy(cp, entities)
		entities = cp
	}
	for _, entity := range entities {
		destination, ok := s.Tile.Destination().Get(entity)
		if !ok {
			continue
		}
		pos, ok := s.Tile.Pos().Get(entity)
		if !ok {
			s.Tile.Destination().Remove(entity)
			s.Logger.Warn(tile.ErrPositionAndSpeedIsRequiredToMoveToDestination)
			continue
		}
		speed, ok := s.Tile.Speed().Get(entity)
		if !ok {
			s.Tile.Destination().Remove(entity)
			s.Logger.Warn(tile.ErrPositionAndSpeedIsRequiredToMoveToDestination)
			continue
		}
		arrived := destination.X == grid.Coord(pos.X) && destination.Y == grid.Coord(pos.Y)
		if arrived {
			s.Tile.Destination().Remove(entity)
			continue
		}
		justStartedMoving := tile.Coord(int(pos.X)) == pos.X && tile.Coord(int(pos.Y)) == pos.Y
		if justStartedMoving { // check can move to destination
			isValidDestination := abs(destination.X-grid.Coord(pos.X))+abs(destination.Y-grid.Coord(pos.Y)) == 1
			if !isValidDestination {
				s.Tile.Destination().Remove(entity)
				s.Logger.Warn(tile.ErrInvalidDestination)
				continue
			}

			// is destination occupied
			size, _ := s.Tile.Size().Get(entity)
			obstruction, _ := s.Tile.Obstruction().Get(entity)
			if grid.Coord(pos.X) != destination.X && s.Tile.IsOccupied(
				tile.NewAABB(
					tile.NewPos(destination.X+size.X-1, destination.Y),
					tile.NewSize(1, size.Y),
				),
				obstruction.Obstruction,
			) {
				s.Logger.Warn(tile.ErrPositionIsOccupied)
				s.Tile.Destination().Remove(entity)
				continue
			}
			if grid.Coord(pos.Y) != destination.Y && s.Tile.IsOccupied(
				tile.NewAABB(
					tile.NewPos(destination.X, destination.Y+size.Y-1),
					tile.NewSize(size.X, 1),
				),
				obstruction.Obstruction,
			) {
				s.Logger.Warn(tile.ErrPositionIsOccupied)
				s.Tile.Destination().Remove(entity)
				continue
			}
		}

		// move
		step := invSpeedTable[speed.InvSpeed]
		if destination.X > grid.Coord(pos.X) {
			pos.X = min(pos.X+step, tile.Coord(destination.X))
		} else if destination.X < grid.Coord(pos.X) {
			pos.X = max(pos.X-step, tile.Coord(destination.X))
		}
		if destination.Y > grid.Coord(pos.Y) {
			pos.Y = min(pos.Y+step, tile.Coord(destination.Y))
		} else if destination.Y < grid.Coord(pos.Y) {
			pos.Y = max(pos.Y-step, tile.Coord(destination.Y))
		}
		s.Tile.Pos().Set(entity, pos)

		arrived = destination.X == grid.Coord(pos.X) && destination.Y == grid.Coord(pos.Y)
		if arrived {
			s.Tile.Destination().Remove(entity)
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
