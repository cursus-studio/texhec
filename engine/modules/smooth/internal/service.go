package internal

import (
	"engine"
	"engine/modules/loop"
	"engine/modules/smooth"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
}

func NewService(c ioc.Dic) smooth.Service {
	return ioc.GetServices[*service](c)
}

func (s *service) Start() ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		events.Listen(s.EventsBuilder(), func(tick loop.TickEvent) {
			events.Emit(s.Events(), FirstEvent(tick))
		})
		return nil
	})
}
func (s *service) Stop() ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		events.Listen(s.EventsBuilder(), func(tick loop.TickEvent) {
			events.Emit(s.Events(), LastEvent(tick))
		})
		return nil
	})
}
