package transformservice

import (
	"engine/modules/ecs"
	"engine/modules/transform"
)

type absolutePosArray struct {
	t *service
	ecs.ComponentArray[transform.AbsolutePosComponent]
}

func (s *absolutePosArray) Set(entity ecs.EntityID, absolutePos transform.AbsolutePosComponent) {
	size, _ := s.t.AbsoluteSizeArray.Get(entity)
	pos := transform.NewPos(absolutePos.Pos.
		Sub(s.t.GetRelativeParentPos(entity)).
		Sub(s.t.GetPivotPos(entity, size)).Elem())

	s.t.PosArray.Set(entity, pos)
}

//

type absoluteSizeArray struct {
	s *service
	ecs.ComponentArray[transform.AbsoluteSizeComponent]
}

func (s *absoluteSizeArray) Set(entity ecs.EntityID, absoluteSize transform.AbsoluteSizeComponent) {
	parentSize := s.s.GetRelativeParentSize(entity)
	size := transform.NewSize(
		absoluteSize.Size[0]/parentSize[0],
		absoluteSize.Size[1]/parentSize[1],
		absoluteSize.Size[2]/parentSize[2],
	)

	s.s.SizeArray.Set(entity, size)
}

//

type absoluteRotationArray struct {
	s *service
	ecs.ComponentArray[transform.AbsoluteRotationComponent]
}

func (s *absoluteRotationArray) Set(entity ecs.EntityID, absoluteRot transform.AbsoluteRotationComponent) {
	rot := transform.NewRotation(absoluteRot.Rotation.
		Mul(s.s.GetRelativeParentRotation(entity).Inverse()))

	s.s.RotationArray.Set(entity, rot)
}
