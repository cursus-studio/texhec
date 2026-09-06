package transformservice

import (
	"engine/modules/ecs"
	"engine/modules/transform"

	"github.com/go-gl/mathgl/mgl32"
)

type save struct {
	entity ecs.EntityID
	pos    transform.AbsolutePosComponent
	rot    transform.AbsoluteRotationComponent
	size   transform.AbsoluteSizeComponent
}

func (s *service) Init() {
	arrays := []ecs.AnyComponentArray{
		s.AbsolutePosArray,
		s.AbsoluteRotationArray,
		s.AbsoluteSizeArray,
	}

	s.PosArray.SetEmpty(transform.PosComponent{Pos: mgl32.Vec3{0, 0, 0}})
	s.RotationArray.SetEmpty(s.defaultRot)
	s.SizeArray.SetEmpty(s.defaultSize)

	s.MaxSizeArray.SetEmpty(transform.NewMaxSize(0, 0, 0)) // 0 means not set
	s.MinSizeArray.SetEmpty(transform.NewMinSize(0, 0, 0)) // 0 means not set

	s.AspectRatioArray.SetEmpty(transform.NewAspectRatio(0, 0, 0, 0)) // 0 means not set
	s.PivotPointArray.SetEmpty(s.defaultPivot)

	s.InheritMaskArray.SetEmpty(transform.NewInherit(transform.RelativePos))
	s.ParentPivotPointArray.SetEmpty(s.defaultParentPivot)

	s.AbsolutePosArray.SetEmpty(transform.AbsolutePosComponent{Pos: mgl32.Vec3{0, 0, 0}})
	s.AbsoluteRotationArray.SetEmpty(transform.AbsoluteRotationComponent(s.defaultRot))
	s.AbsoluteSizeArray.SetEmpty(transform.AbsoluteSizeComponent(s.defaultSize))

	for _, arr := range arrays {
		arr.AddDependency(s.PosArray)
		arr.AddDependency(s.RotationArray)
		arr.AddDependency(s.SizeArray)

		arr.AddDependency(s.MaxSizeArray)
		arr.AddDependency(s.MinSizeArray)

		arr.AddDependency(s.AspectRatioArray)
		arr.AddDependency(s.PivotPointArray)

		arr.AddDependency(s.Hierarchy().Component())
		arr.AddDependency(s.InheritMaskArray)
		arr.AddDependency(s.ParentPivotPointArray)
	}

	for _, array := range arrays {
		array.AddDirtySet(s.DirtySet)
		array.BeforeGet(s.BeforeGet)
	}
}
