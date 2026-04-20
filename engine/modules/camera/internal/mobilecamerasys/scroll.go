package mobilecamerasys

import (
	"engine"
	"engine/modules/camera"
	"engine/services/ecs"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type scrollSystem struct {
	engine.EngineWorld `inject:""`
}

func NewScrollSystem(c ioc.Dic) camera.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*scrollSystem](c)
		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (s *scrollSystem) Listen(event sdl.MouseWheelEvent) {
	if event.Y == 0 {
		return
	}

	var mul = float32(math.Pow(10, float64(event.Y)/50))

	mousePos := s.Window().GetMousePos()

	for _, cameraEntity := range s.Camera().Mobile().GetEntities() {
		ortho, ok := s.Camera().Ortho().Get(cameraEntity)
		if !ok {
			continue
		}

		pos, _ := s.Transform().AbsolutePos().Get(cameraEntity)
		rot, _ := s.Transform().AbsoluteRotation().Get(cameraEntity)

		rayBefore := s.Camera().ShootRay(cameraEntity, mousePos)

		// apply zoom
		ortho.Zoom *= mul
		if limits, ok := s.Camera().Limits().Get(cameraEntity); ok {
			ortho.Zoom = max(min(ortho.Zoom, limits.MaxZoom), limits.MinZoom)
		}

		s.Camera().Ortho().Set(cameraEntity, ortho)

		// read after
		rayAfter := s.Camera().ShootRay(cameraEntity, mousePos)

		// apply transform
		pos.Pos = pos.Pos.Add(rayBefore.Pos.Sub(rayAfter.Pos))

		rotationDifference := mgl32.QuatBetweenVectors(rayBefore.Direction, rayAfter.Direction)
		rot.Rotation = rotationDifference.Mul(rot.Rotation)

		s.Transform().AbsolutePos().Set(cameraEntity, pos)
		s.Transform().AbsoluteRotation().Set(cameraEntity, rot)
	}
}
