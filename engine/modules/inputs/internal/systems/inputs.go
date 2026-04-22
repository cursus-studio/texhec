package systems

import (
	"engine"
	"engine/modules/inputs"
	"engine/modules/loop"
	"engine/services/ecs"
	mediainputs "engine/services/media/inputs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type inputsSystem struct {
	engine.EngineWorld `inject:""`
	Inputs             mediainputs.Api `inject:""`
}

func NewInputsSystem(c ioc.Dic) inputs.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*inputsSystem](c)
		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (system *inputsSystem) Listen(args loop.FrameEvent) {
	system.Inputs.Poll()
}
