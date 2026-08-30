package internal

import (
	"core/game"
	"core/modules/attack"
	"core/modules/reach"
	"engine/modules/ecs"

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
	if linkEntity, ok := s.Tile().GetLink(entity); ok {
		return s.health.Get(linkEntity)
	}
	return s.zeroHealth, false
}
