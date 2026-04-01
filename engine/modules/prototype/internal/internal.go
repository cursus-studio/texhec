package internal

import (
	"engine/modules/prototype"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type Service interface {
	Add(arr ecs.AnyComponentArray)
	prototype.Service
}

type service struct {
	World  ecs.World     `inject:"1"`
	Events events.Events `inject:"1"`

	arrays []ecs.AnyComponentArray
}

func NewService(c ioc.Dic) Service {
	s := ioc.GetServices[*service](c)
	return s
}

func (s *service) Add(array ecs.AnyComponentArray) {
	s.arrays = append(s.arrays, array)
}

func (s *service) Clone(cloned ecs.EntityID) ecs.EntityID {
	clone := s.World.NewEntity()
	for _, arr := range s.arrays {
		if comp, ok := arr.GetAny(cloned); ok {
			_ = arr.SetAny(clone, comp)
		}
	}
	events.Emit(s.Events, prototype.NewCloneEvent(cloned, clone))
	return clone
}
