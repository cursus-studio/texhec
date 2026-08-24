package internal

import (
	"core/modules/attack"
	"core/modules/obstruction"
	"core/modules/pathfind"
	"core/modules/tile"
	"engine/modules/loop"

	"github.com/ogiusek/events"
)

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.OnTick)
	events.Listen(s.EventsBuilder(), s.OnFrame)
	return nil
}

func (s *service) OnTick(event loop.TickEvent) {
	for _, entity := range s.target.GetEntities() {
		damage, ok := s.Attack().Damage().Get(entity)
		if !ok {
			continue
		}
		damage.Damage = attack.Health(event.Delta.Seconds() * float64(damage.Damage))
		target, ok := s.target.Get(entity)
		if !ok {
			continue
		}
		targetHealth, ok := s.Attack().Health().Get(target.Entity)
		if !ok {
			continue
		}
		if s.Reach().Reaches(entity, target.Entity) {
			targetHealth.Health -= min(targetHealth.Health, damage.Damage)
			s.Attack().Health().Set(target.Entity, targetHealth)
			if targetHealth.Health != 0 {
				continue
			}
			s.Pathfind().Target().Remove(entity)
			s.Attack().Target().Remove(entity)
			// s.World().RemoveEntity(target.Entity)
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

func (s *service) OnFrame(loop.FrameEvent) {
	for _, entity := range s.dirtyHealth.Get() {
		if health, ok := s.health.Get(entity); !ok || health.Health != 0 {
			continue
		}
		s.World().RemoveEntity(entity)
	}
}
