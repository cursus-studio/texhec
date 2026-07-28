package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/transition"
	"slices"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	easing             ecs.ComponentArray[transition.EasingComponent]
	easingFunction     ecs.ComponentArray[transition.EasingFunctionComponent]
	register           ecs.SystemRegister

	delayed []*transition.DelayedEvent
}

func NewService(c ioc.Dic, register ecs.SystemRegister) transition.Service {
	s := ioc.GetServices[*service](c)
	s.easing = ecs.GetComponentArray[transition.EasingComponent](s.World())
	s.easingFunction = ecs.GetComponentArray[transition.EasingFunctionComponent](s.World())
	s.register = register

	return s
}

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.ListenDelayed)
	events.Listen(s.EventsBuilder(), s.ListenFrame)
	return s.register.Register()
}
func (s *service) Easing() ecs.ComponentArray[transition.EasingComponent] {
	return s.easing
}
func (s *service) EasingFunction() ecs.ComponentArray[transition.EasingFunctionComponent] {
	return s.easingFunction
}

//

func (s *service) ListenDelayed(e transition.DelayedEvent) {
	insIdx, _ := slices.BinarySearchFunc(s.delayed, &e, func(a, b *transition.DelayedEvent) int {
		return int(a.Duration - b.Duration)
	})

	s.delayed = slices.Insert(s.delayed, insIdx, &e)
}

func (s *service) ListenFrame(e loop.FrameEvent) {
	toOld := 0
	for _, event := range s.delayed {
		event.Duration -= e.Delta
		if event.Duration > 0 {
			continue
		}

		events.EmitAny(s.Events(), event.Event)
		toOld++
	}

	s.delayed = s.delayed[toOld:]
}
