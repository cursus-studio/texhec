package systems

import (
	"engine"
	"engine/modules/loop"
	"engine/modules/render"
	"engine/services/ecs"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type errorLogger struct {
	engine.EngineWorld `inject:""`
}

func NewErrorLogger(c ioc.Dic) render.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*errorLogger](c)
		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (logger *errorLogger) Listen(args loop.FrameEvent) {
	if err := logger.Render().Error(); err != nil {
		logger.Logger().Warn(fmt.Errorf("opengl error: %s", err))
	}
}
