package internal

import (
	"core/game"
	"core/modules/obstruction"
	"core/modules/pathfind"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"golang.org/x/exp/constraints"
)

type service struct {
	game.GameWorld `inject:""`

	target ecs.ComponentArray[pathfind.TargetComponent]
	speed  ecs.ComponentArray[pathfind.SpeedComponent]
	step   ecs.ComponentArray[pathfind.StepComponent]
}

func NewService(c ioc.Dic) pathfind.Service {
	s := ioc.GetServices[*service](c)
	s.target = ecs.GetComponentArray[pathfind.TargetComponent](s.World())
	s.speed = ecs.GetComponentArray[pathfind.SpeedComponent](s.World())
	s.step = ecs.GetComponentArray[pathfind.StepComponent](s.World())
	return s
}

func (s *service) Register() error {
	for i := 1; i < 256; i++ {
		invSpeedTable[i] = 1. / tile.Coord(i)
	}

	events.Listen(s.EventsBuilder(), s.FindPathFeature)
	events.Listen(s.EventsBuilder(), s.FindPath)
	events.Listen(s.EventsBuilder(), s.StepOnTick)
	events.Listen(s.EventsBuilder(), s.PathfindOnTick)
	return nil
}

func (s *service) Target() ecs.ComponentArray[pathfind.TargetComponent] { return s.target }
func (s *service) Speed() ecs.ComponentArray[pathfind.SpeedComponent]   { return s.speed }
func (s *service) Step() ecs.ComponentArray[pathfind.StepComponent]     { return s.step }

func abs[Number constraints.Float | constraints.Integer](n Number) Number { return max(-n, n) }

func (s *service) CanStep(
	pos grid.Coords,
	size tile.SizeComponent,
	obstructionComp obstruction.Component,
	step pathfind.StepComponent,
) bool {
	isValidStep := abs(step.X-pos.X)+abs(step.Y-pos.Y) == abs(grid.Coord(1))
	if !isValidStep {
		return false
	}

	// is step destination occupied
	var aabbPos tile.PosComponent
	var aabbSize tile.SizeComponent

	// aabb size
	if pos.X != step.X {
		aabbSize = tile.NewSize(1, size.Y)
	} else if pos.Y != step.Y {
		aabbSize = tile.NewSize(size.X, 1)
	}
	// aabb pos
	if pos.X < step.X {
		aabbPos = tile.NewPos(step.X+size.X-1, step.Y)
	} else if pos.Y < step.Y {
		aabbPos = tile.NewPos(step.X, step.Y+size.Y-1)
	} else {
		aabbPos = tile.NewPos(step.Coords.Coords())
	}
	// perform is step destination occupied
	if collisions := s.Obstruction().Collisions(obstruction.NewAABB(aabbPos, aabbSize), obstructionComp.Obstruction); len(collisions) != 0 {
		return false
	}
	return true
}
