package systems

import (
	"engine"
	"engine/modules/inputs"
	"engine/services/ecs"
	"engine/services/frames"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type sys struct {
	engine.EngineWorld `inject:""`
	Closed             bool
}

func NewQuitSystem(c ioc.Dic) inputs.ShutdownSystem {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*sys](c)
		events.Listen(s.EventsBuilder(), s.Listen)
		events.Listen(s.EventsBuilder(), s.ListenFrame)
		return nil
	})
}

func (s *sys) Listen(inputs.QuitEvent) {
	s.Closed = true
}

func (s *sys) ListenFrame(frames.FrameEvent) {
	if s.Closed {
		s.Runtime().Stop()
	}
}
