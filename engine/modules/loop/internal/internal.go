package internal

import (
	"engine"
	"engine/modules/loop"
	"time"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	Running            bool

	TickDuration,
	FrameDuration time.Duration

	LastFrameTime time.Time
	TickProgress  time.Duration

	Ticker *time.Ticker
}

func NewService(c ioc.Dic) loop.Service {
	s := ioc.GetServices[*service](c)
	s.Running = false

	// TickDuration is initialized lazily
	// FrameDuration is initalized lazily

	// LastFrameTime is initialized lazily
	s.TickProgress = 0

	s.Ticker = nil

	events.Listen(s.EventsBuilder(), func(loop.StopEvent) {
		s.Running = false
	})
	events.Listen(s.EventsBuilder(), func(e loop.ConfigureEvent) {
		s.TickDuration = time.Second / time.Duration(e.TPS)
		s.FrameDuration = time.Second / time.Duration(e.FPS)
		prev := s.Ticker
		s.Ticker = time.NewTicker(s.FrameDuration)
		if prev != nil {
			prev.Stop()
		}
	})

	return s
}

func (s *service) Run(e loop.ConfigureEvent) {
	s.Configure(e)
	if s.Running {
		return
	}

	s.Running = true
	s.LastFrameTime = s.Clock().Now()

	for s.Running {
		<-s.Ticker.C
		currentTime := s.Clock().Now()

		delta := currentTime.Sub(s.LastFrameTime)
		s.LastFrameTime = currentTime

		frameEvent := loop.FrameEvent{Delta: delta}
		s.TickProgress += delta
		for s.TickProgress > s.TickDuration {
			s.TickProgress -= s.TickDuration
			events.Emit(s.Events(), loop.TickEvent{Delta: s.TickDuration})
		}
		events.Emit(s.Events(), frameEvent)
	}
}

func (s *service) Stop()                           { events.Emit(s.Events(), loop.StopEvent{}) }
func (s *service) Configure(e loop.ConfigureEvent) { events.Emit(s.Events(), e) }

func (s *service) Stats() loop.Stats { return s }

// stats interface
func (s *service) FrameBudget() time.Duration {
	return s.FrameDuration
}
func (s *service) FrameBudgetLeft() time.Duration {
	return max(0, s.LastFrameTime.Add(s.FrameDuration).Sub(s.Clock().Now()))
}
