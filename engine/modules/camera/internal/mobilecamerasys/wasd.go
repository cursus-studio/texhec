package mobilecamerasys

import (
	"engine"
	"engine/modules/camera"
	"engine/services/ecs"
	"engine/services/frames"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type wasdMoveSystem struct {
	engine.EngineWorld `inject:""`

	cameraSpeed float32
}

func NewWasdSystem(c ioc.Dic, cameraSpeed float32) camera.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*wasdMoveSystem](c)
		s.cameraSpeed = cameraSpeed
		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (s *wasdMoveSystem) Listen(event frames.FrameEvent) {
	var moveVerticaly float32 = 0
	var moveHorizontaly float32 = 0
	{
		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_W] == 1 || keys[sdl.SCANCODE_UP] == 1 {
			moveVerticaly = 1
		}
		if keys[sdl.SCANCODE_S] == 1 || keys[sdl.SCANCODE_DOWN] == 1 {
			moveVerticaly = -1
		}

		if keys[sdl.SCANCODE_A] == 1 || keys[sdl.SCANCODE_LEFT] == 1 {
			moveHorizontaly = -1
		}
		if keys[sdl.SCANCODE_D] == 1 || keys[sdl.SCANCODE_RIGHT] == 1 {
			moveHorizontaly = 1
		}
	}

	{
		moveHorizontaly *= float32(event.Delta.Milliseconds()) * s.cameraSpeed
		moveVerticaly *= float32(event.Delta.Milliseconds()) * s.cameraSpeed
	}

	for _, camera := range s.Camera().Mobile().GetEntities() {
		pos, _ := s.Transform().AbsolutePos().Get(camera)
		ortho, ok := s.Camera().Ortho().Get(camera)
		if !ok {
			continue
		}

		pos.Pos = mgl32.Vec3{
			pos.Pos.X() + moveHorizontaly/ortho.Zoom,
			pos.Pos.Y() + moveVerticaly/ortho.Zoom,
			pos.Pos.Z(),
		}
		s.Transform().AbsolutePos().Set(camera, pos)
	}
}
