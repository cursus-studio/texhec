package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/loop"
	"slices"
	"time"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	emitOnFrameArray   ecs.ComponentArray[loop.EmitOnFrameComponent]
	emitOnTickArray    ecs.ComponentArray[loop.EmitOnTickComponent]
	lastTick           ecs.ComponentArray[loop.LastTickComponent]
	Running            bool

	TickDuration,
	FrameDuration time.Duration

	LastFrameTime time.Time
	TickProgress  time.Duration

	Ticker *time.Ticker
}

func NewService(c ioc.Dic) loop.Service {
	s := ioc.GetServices[*service](c)
	s.emitOnFrameArray = ecs.GetComponentArray[loop.EmitOnFrameComponent](s.World())
	s.emitOnTickArray = ecs.GetComponentArray[loop.EmitOnTickComponent](s.World())
	s.lastTick = ecs.GetComponentArray[loop.LastTickComponent](s.World())
	s.Running = false

	// TickDuration is initialized lazily
	// FrameDuration is initalized lazily

	// LastFrameTime is initialized lazily
	s.TickProgress = 0

	s.Ticker = nil

	events.Listen(s.EventsBuilder(), func(loop.StopEvent) {
		s.Running = false
	})

	events.Listen(s.EventsBuilder(), func(loop.TickEvent) {
		entity := s.LastTickEntity()
		lastTick, _ := s.lastTick.Get(entity)
		defer s.lastTick.Set(entity, lastTick)
		lastTick.UnixNano = time.Unix(0, lastTick.UnixNano).Add(s.TickDuration).UnixNano()
	})

	//

	// emit proper event on tick/frame
	events.Listen(s.EventsBuilder(), func(loop.FrameEvent) {
		arr := slices.Clone(s.emitOnFrameArray.GetEntities())
		for _, entity := range arr {
			if comp, ok := s.emitOnFrameArray.Get(entity); ok {
				events.EmitAny(s.Events(), comp.Event)
			}
			s.World().RemoveEntity(entity)
		}
	})
	events.Listen(s.EventsBuilder(), func(loop.TickEvent) {
		arr := slices.Clone(s.emitOnTickArray.GetEntities())
		for _, entity := range arr {
			if comp, ok := s.emitOnTickArray.Get(entity); ok {
				events.EmitAny(s.Events(), comp.Event)
			}
			s.World().RemoveEntity(entity)
		}
	})
	events.Listen(s.EventsBuilder(), func(event loop.EmitOnFrameEvent) { s.EmitOnFrame(event.Event) })
	events.Listen(s.EventsBuilder(), func(event loop.EmitOnTickEvent) { s.EmitOnTick(event.Event) })

	//

	events.Listen(s.EventsBuilder(), s.Configure)

	return s
}

func (s *service) LastTickEntity() ecs.EntityID {
	entities := s.lastTick.GetEntities()
	if len(entities) == 0 {
		entity := s.World().NewEntity()
		comp := loop.NewLastTick(0)
		s.lastTick.Set(entity, comp)
		return entity
	}
	for _, entity := range slices.Clone(entities[1:]) {
		s.lastTick.Remove(entity)
	}
	return entities[0]
}
func (s *service) LastTick() loop.LastTickComponent {
	entity := s.LastTickEntity()
	comp, _ := s.lastTick.Get(entity)
	return comp
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

func (s *service) Stop() { events.Emit(s.Events(), loop.StopEvent{}) }
func (s *service) Configure(e loop.ConfigureEvent) {
	s.TickDuration = time.Second / time.Duration(e.TPS)
	s.FrameDuration = time.Second / time.Duration(e.FPS)

	now := s.Clock().Now()
	s.TickProgress = now.Sub(now.Truncate(s.TickDuration))

	tickTime := now.Add(-s.TickProgress)
	entity := s.LastTickEntity()
	lastTick, _ := s.lastTick.Get(entity)
	defer s.lastTick.Set(entity, lastTick)
	lastTick.UnixNano = tickTime.UnixNano()

	prev := s.Ticker
	s.Ticker = time.NewTicker(s.FrameDuration)
	if prev != nil {
		prev.Stop()
	}
}

func (s *service) SyncToUnixNano(lastTickUnix int64) {
	entity := s.LastTickEntity()
	lastTick, _ := s.lastTick.Get(entity)
	defer s.lastTick.Set(entity, lastTick)
	lastTick.UnixNano = time.Unix(0, lastTick.UnixNano).Add(-s.TickDuration).UnixNano()
	for s.LastTick().UnixNano-lastTickUnix != 0 {
		events.Emit(s.Events(), loop.TickEvent{Delta: s.TickDuration})
	}
}

func (s *service) Stats() loop.Stats { return s }

// stats interface
func (s *service) FrameBudget() time.Duration {
	return s.FrameDuration
}
func (s *service) FrameBudgetLeft() time.Duration {
	return max(0, s.LastFrameTime.Add(s.FrameDuration).Sub(s.Clock().Now()))
}

func (s *service) EmitOnFrame(event any) {
	entity := s.World().NewEntity()
	s.emitOnFrameArray.Set(entity, loop.EmitOnFrameComponent{Event: event})
}
func (s *service) EmitOnTick(event any) {
	entity := s.World().NewEntity()
	s.emitOnTickArray.Set(entity, loop.EmitOnTickComponent{Event: event})
}
