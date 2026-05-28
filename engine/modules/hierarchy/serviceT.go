package hierarchy

import "engine/services/ecs"

type InheritComponent[Component any] struct{}

func NewInherit[Component any]() InheritComponent[Component] { return InheritComponent[Component]{} }

type ServiceT[Component any] interface {
	Inherit() ecs.ComponentsArray[InheritComponent[Component]]
}
