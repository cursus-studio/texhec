package hierarchy

import "engine/modules/ecs"

type InheritComponent[Component any] struct{}

func NewInherit[Component any]() InheritComponent[Component] { return InheritComponent[Component]{} }

type ServiceT[Component any] interface {
	Inherit() ecs.ComponentArray[InheritComponent[Component]]
}
