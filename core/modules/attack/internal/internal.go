package internal

import (
	"core/game"
	"core/modules/attack"
	"core/modules/obstruction"
	"core/modules/pathfind"
	"core/modules/reach"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/loop"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`
	ReachService   reach.ServiceT[attack.TargetComponent] `inject:""`
	target         ecs.ComponentArray[attack.TargetComponent]
}

func NewService(c ioc.Dic) attack.Service {
	s := ioc.GetServices[*service](c)
	s.target = ecs.GetComponentArray[attack.TargetComponent](s.World())
	return s
}

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.OnTick)
	return nil
}

func (s *service) Reach() reach.ServiceT[attack.TargetComponent] {
	return s.ReachService
}

func (s *service) Target() ecs.ComponentArray[attack.TargetComponent] {
	return s.target
}

func (s *service) OnTick(event loop.TickEvent) {
	for _, entity := range s.target.GetEntities() {
		target, ok := s.target.Get(entity)
		if !ok {
			continue
		}
		if s.Reach().Reaches(entity, target.Entity) {
			s.Pathfind().Target().Remove(entity)
			s.Attack().Target().Remove(entity)
			s.World().RemoveEntity(target.Entity)
			continue
		}
		if _, ok := s.Pathfind().Speed().Get(entity); !ok {
			s.Logger().Warn(attack.ErrCannotAttackEnemyOutOfReach)
			s.Attack().Target().Remove(entity)
			continue
		}
		pos, ok := s.Tile().Pos().Get(target.Entity)
		if !ok {
			continue
		}
		size, _ := s.Tile().Size().Get(target.Entity)
		reach, _ := s.Reach().Component().Get(entity)
		obstructionComp, _ := s.Obstruction().Component().Get(entity)
		coords := s.Attack().Reach().TilesFrom(pos, size, reach)
		for _, coord := range coords {
			aabb := obstruction.NewAABB(tile.NewPos(coord.Coords()), size)
			collisions := s.Obstruction().Collisions(aabb, obstructionComp.Obstruction)
			if len(collisions) != 0 {
				continue
			}
			s.Pathfind().Target().Set(entity, pathfind.NewTarget(coord))
			break
		}
	}
}
