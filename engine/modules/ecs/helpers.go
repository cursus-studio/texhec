package ecs

import (
	"engine/modules/ecs/internal/ecstypes"
)

// events
type RemoveEntityEvent struct{ Entity EntityID }

func NewRemoveEntityEvent(entity EntityID) RemoveEntityEvent {
	return RemoveEntityEvent{entity}
}

// event wrappers
type ApplyEntityEvent interface {
	ApplyEntity(entityEmitting EntityID) (event any)
}

// systems
type SystemRegister interface{ Register() error }
type systemRegister func() error

func (s systemRegister) Register() error              { return s() }
func NewSystemRegister(l func() error) SystemRegister { return systemRegister(l) }
func RegisterSystems(systems ...SystemRegister) []error {
	errs := []error{}
	for _, system := range systems {
		if system == nil {
			continue
		}
		if err := system.Register(); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func ComponentComparator[Component any]() func(c1, c2 Component) bool {
	return ecstypes.ComponentComparator[Component]()
}
