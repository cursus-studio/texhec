package systems

import (
	"engine"
	"engine/modules/render"
	"engine/services/ecs"
	"engine/services/frames"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type renderSystem struct {
	engine.EngineWorld `inject:""`
}

func NewRenderSystem(c ioc.Dic) render.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*renderSystem](c)
		events.ListenE(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (s *renderSystem) Listen(args frames.FrameEvent) error {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	cameras := s.Camera().OrderedCameras()
	for _, camera := range cameras {

		gl.Clear(gl.DEPTH_BUFFER_BIT)
		gl.Viewport(s.Camera().GetViewport(camera))

		events.Emit(s.Events(), render.RenderEvent{
			Camera: camera,
		})
	}

	s.Window().Window().GLSwap()

	return nil
}
