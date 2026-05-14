package internal

import (
	"core/modules/pathfind"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/modules/loop"
	"engine/services/ecs"
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

var invSpeedTable [256]tile.Coord

var (
	rotLeft  = tile.NewRot(mgl32.DegToRad(90))
	rotRight = tile.NewRot(mgl32.DegToRad(270))
	rotUp    = tile.NewRot(mgl32.DegToRad(0))
	rotDown  = tile.NewRot(mgl32.DegToRad(180))
)

func (s *service) StepOnTick(e loop.TickEvent) {
	entities := s.Pathfind().Step().GetEntities()
	{
		cp := make([]ecs.EntityID, len(entities))
		copy(cp, entities)
		entities = cp
	}
	for _, entity := range entities {
		step, ok := s.Pathfind().Step().Get(entity)
		if !ok {
			continue
		}
		pos, ok := s.Tile().Pos().Get(entity)
		if !ok {
			s.Pathfind().Step().Remove(entity)
			s.Logger().Log(tile.ErrPositionAndSpeedIsRequiredToStep)
			continue
		}
		speed, ok := s.Pathfind().Speed().Get(entity)
		if !ok {
			s.Pathfind().Step().Remove(entity)
			s.Logger().Log(tile.ErrPositionAndSpeedIsRequiredToStep)
			continue
		}
		arrived := tile.Coord(step.X) == pos.X && tile.Coord(step.Y) == pos.Y
		if arrived {
			s.Pathfind().Step().Remove(entity)
			continue
		}
		size, _ := s.Tile().Size().Get(entity)
		obstruction, _ := s.Obstruction().Component().Get(entity)
		aligned, isFirstStep := pos.Aligned()
		if isFirstStep && !s.Pathfind().CanStep(aligned, size, obstruction, step) {
			s.Pathfind().Step().Remove(entity)
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
			s.Pathfind().Step().Remove(entity)
		}
	}
}

func (s *service) PathfindOnTick(e loop.TickEvent) {
	for _, entity := range s.Target().GetEntities() {
		if _, ok := s.Pathfind().Step().Get(entity); ok {
			continue
		}

		from, ok := s.Tile().Pos().Get(entity)
		if !ok {
			s.Logger().Log(tile.ErrInvalidPosition)
			return
		}
		target, _ := s.Target().Get(entity)
		fromAligned, _ := from.Aligned()
		to := tile.NewPos(target.Coords.Coords())
		toAligned, _ := to.Aligned()
		if toAligned == fromAligned {
			s.Target().Remove(entity)
			continue
		}
		size, _ := s.Tile().Size().Get(entity)
		obstruction, _ := s.Obstruction().Component().Get(entity)
		fromCoords, _ := from.Aligned()
		toCoords, _ := to.Aligned()
		path, ok := s.findPath(fromCoords, toCoords, size, obstruction)
		if !ok {
			s.Logger().Log(pathfind.ErrInvalidPath)
			continue
		}
		step := pathfind.NewStep(grid.Coord(path[0].X), grid.Coord(path[0].Y))
		for !s.Pathfind().CanStep(fromCoords, size, obstruction, step) {
			path, _ = s.findPath(fromCoords, toCoords, size, obstruction)
			step = pathfind.NewStep(grid.Coord(path[0].X), grid.Coord(path[0].Y))
		}
		s.Pathfind().Step().Set(entity, step)
	}
}
