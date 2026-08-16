package attack

import (
	"core/modules/reach"
	"engine/modules/ecs"
	"engine/modules/transition"
	"errors"
)

var (
	ErrCannotAttackEnemyOutOfReach error = errors.New("attack: enemy is out of reach and cannot follow him")
)

type Health uint32

type TargetComponent struct {
	Entity ecs.EntityID
}
type HealthComponent struct {
	Health Health
}
type DamageComponent struct {
	Damage Health
}

func NewTarget(target ecs.EntityID) TargetComponent {
	return TargetComponent{target}
}
func NewHealth(health Health) HealthComponent {
	return HealthComponent{health}
}
func NewDamage(damage Health) DamageComponent {
	return DamageComponent{damage}
}

func (HealthComponent) Smooth() {}
func (c1 HealthComponent) Lerp(c2 HealthComponent, mix32 float32) HealthComponent {
	return HealthComponent{transition.LerpInt(c1.Health, c2.Health, mix32)}
}

//

type Service interface {
	ecs.SystemRegister
	Reach() reach.ServiceT[TargetComponent]

	Target() ecs.ComponentArray[TargetComponent]
	Health() ecs.ComponentArray[HealthComponent]
	Damage() ecs.ComponentArray[DamageComponent]

	FullHealth(ecs.EntityID) (HealthComponent, bool)
}
