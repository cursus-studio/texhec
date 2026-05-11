package transitionimpl

import (
	"engine/services/ecs"
)

type Builder interface {
	Register(ecs.SystemRegister)
	Build() ecs.SystemRegister
}

type builder struct {
	systems []ecs.SystemRegister
}

func NewBuilder() Builder {
	return &builder{
		systems: nil,
	}
}

func (b *builder) Register(system ecs.SystemRegister) {
	b.systems = append(b.systems, system)
}

//

func (b *builder) Build() ecs.SystemRegister {
	systems := b.systems
	return ecs.NewSystemRegister(func() error {
		for _, system := range systems {
			if err := system.Register(); err != nil {
				return err
			}
		}
		return nil
	})
}
