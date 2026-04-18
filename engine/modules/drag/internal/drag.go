package internal

import (
	"engine"
	"engine/modules/drag"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type s struct {
	engine.EngineWorld `inject:""`
}

func NewSystem(c ioc.Dic) drag.System {
	return ioc.GetServices[s](c)
}

func (s s) Register() error {
	events.Listen(s.EventsBuilder(), func(event drag.DraggableEvent) {
		entity := event.Entity

		pos, _ := s.Transform().AbsolutePos().Get(entity)
		rot, _ := s.Transform().AbsoluteRotation().Get(entity)

		fromRay := s.Camera().ShootRay(event.Drag.Camera, event.Drag.From)
		toRay := s.Camera().ShootRay(event.Drag.Camera, event.Drag.To)

		posDiff := toRay.Pos.Sub(fromRay.Pos)
		pos.Pos = pos.Pos.Add(posDiff)
		s.Transform().AbsolutePos().Set(entity, pos)

		rotDiff := mgl32.QuatBetweenVectors(toRay.Direction, fromRay.Direction)
		rot.Rotation = rot.Rotation.Mul(rotDiff)
		s.Transform().AbsoluteRotation().Set(entity, rot)
	})
	return nil
}
