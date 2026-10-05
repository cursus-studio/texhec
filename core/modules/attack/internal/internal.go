package internal

import (
	"core/game"
	"core/modules/actions"
	"core/modules/attack"
	"core/modules/reach"
	"engine/modules/ecs"
	"engine/modules/uuid"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`
	ReachService   reach.ServiceT[attack.TargetComponent] `inject:""`
	target         ecs.ComponentArray[attack.TargetComponent]
	health         ecs.ComponentArray[attack.HealthComponent]
	damage         ecs.ComponentArray[attack.DamageComponent]

	dirtyHealth ecs.DirtySet
	zeroHealth  attack.HealthComponent
}

func NewService(c ioc.Dic) attack.Service {
	s := ioc.GetServices[*service](c)
	s.target = ecs.GetComponentArray[attack.TargetComponent](s.World())
	s.health = ecs.GetComponentArray[attack.HealthComponent](s.World())
	s.damage = ecs.GetComponentArray[attack.DamageComponent](s.World())

	s.dirtyHealth = ecs.NewDirtySet()
	s.health.AddDirtySet(s.dirtyHealth)
	return s
}

func (s *service) Reach() reach.ServiceT[attack.TargetComponent] { return s.ReachService }

func (s *service) Target() ecs.ComponentArray[attack.TargetComponent] { return s.target }
func (s *service) Health() ecs.ComponentArray[attack.HealthComponent] { return s.health }
func (s *service) Damage() ecs.ComponentArray[attack.DamageComponent] { return s.damage }

func (s *service) FullHealth(entity ecs.EntityID) (attack.HealthComponent, bool) {
	if linkEntity, ok := s.Tile().Blueprint().Get(entity); ok {
		return s.health.Get(linkEntity)
	}
	return s.zeroHealth, false
}

func (s *service) CanAttack(entityUUID actions.Entity) error {
	entity, ok := s.UUID().Entity(uuid.UUID(entityUUID))
	if !ok {
		return s.Actions().CanAttack().Error()
	}
	if _, ok := s.Damage().Get(entity); !ok {
		return s.Actions().CanAttack().Error()
	}
	return nil
}
func (s *service) CanBeAttacked(attackerUUID, attackedUUID actions.Entity) error {
	if err := s.CanAttack(attackerUUID); err != nil {
		return s.Actions().CanBeAttacked().Error()
	}
	attacker, ok := s.UUID().Entity(uuid.UUID(attackerUUID))
	if !ok {
		return s.Actions().CanBeAttacked().Error()
	}
	attacked, ok := s.UUID().Entity(uuid.UUID(attackedUUID))
	if !ok {
		return s.Actions().CanBeAttacked().Error()
	}
	attackerRegion, _ := s.Pathfind().EntityRegion(attacker)
	attackedRegion, ok := s.Pathfind().EntityRegion(attacked)
	if !ok || attackerRegion != attackedRegion {
		return s.Actions().CanBeAttacked().Error()
	}
	return nil
}
