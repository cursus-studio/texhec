package internal

import (
	"engine"
	"engine/modules/delay"
	"engine/modules/loop"
	"slices"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	delayed            []*delay.DelayedEvent
}

func NewService(c ioc.Dic) delay.Service {
	s := ioc.GetServices[*service](c)
	events.Listen(s.EventsBuilder(), s.ListenDelayed)
	return s
}

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.ListenFrame)
	return nil
}

func (s *service) Delay(e delay.DelayedEvent) { events.Emit(s.Events(), e) }
func (s *service) ListenDelayed(e delay.DelayedEvent) {
	insIdx, _ := slices.BinarySearchFunc(s.delayed, &e, func(a, b *delay.DelayedEvent) int {
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

		events.EmitAny(s.Events(), event.Event.ApplyDelay(-event.Duration))
		toOld++
	}

	s.delayed = s.delayed[toOld:]
}
