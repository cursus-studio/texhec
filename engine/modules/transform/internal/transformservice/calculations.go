package transformservice

import (
	"engine/modules/ecs"
	"engine/modules/transform"

	"github.com/go-gl/mathgl/mgl32"
)

func (s *service) GetRelativeParentPos(entity ecs.EntityID) mgl32.Vec3 {
	parent, ok := s.Hierarchy().Component().Get(entity)
	if !ok {
		return mgl32.Vec3{}
	}
	parentMask, _ := s.InheritMaskArray.Get(entity)
	if parentMask.RelativeMask&transform.RelativePos == 0 {
		return mgl32.Vec3{}
	}
	parentPos, _ := s.AbsolutePosArray.Get(parent.Parent)
	parentSize, _ := s.AbsoluteSizeArray.Get(parent.Parent)
	parentPivot, _ := s.ParentPivotPointArray.Get(entity)
	parentPivot.Point = parentPivot.Point.Sub(s.defaultParentPivot.Point)
	return parentPos.Pos.Add(mgl32.Vec3{
		parentSize.Size[0] * parentPivot.Point[0],
		parentSize.Size[1] * parentPivot.Point[1],
		parentSize.Size[2] * parentPivot.Point[2],
	})
}

func (s *service) GetPivotPos(entity ecs.EntityID, size transform.AbsoluteSizeComponent) mgl32.Vec3 {
	pivot, _ := s.PivotPointArray.Get(entity)
	if pivot == s.defaultPivot {
		return mgl32.Vec3{}
	}
	pivot.Point = pivot.Point.Sub(s.defaultPivot.Point)
	return mgl32.Vec3{
		size.Size[0] * (-pivot.Point[0]),
		size.Size[1] * (-pivot.Point[1]),
		size.Size[2] * (-pivot.Point[2]),
	}
}

//

func (s *service) GetRelativeParentRotation(entity ecs.EntityID) mgl32.Quat {
	parent, ok := s.Hierarchy().Component().Get(entity)
	if !ok {
		return mgl32.QuatIdent()
	}
	parentMask, _ := s.InheritMaskArray.Get(entity)
	if parentMask.RelativeMask&transform.RelativeRotation == 0 {
		return mgl32.QuatIdent()
	}
	parentRot, _ := s.AbsoluteRotationArray.Get(parent.Parent)
	return parentRot.Rotation
}

//

func (s *service) GetRelativeParentSize(entity ecs.EntityID) mgl32.Vec3 {
	size := s.defaultSize.Size
	parent, ok := s.Hierarchy().Component().Get(entity)
	if !ok {
		return size
	}
	parentMask, _ := s.InheritMaskArray.Get(entity)
	parentSize, _ := s.AbsoluteSizeArray.Get(parent.Parent)
	if parentMask.RelativeMask&transform.RelativeSizeX != 0 {
		size[0] = parentSize.Size[0]
	}
	if parentMask.RelativeMask&transform.RelativeSizeY != 0 {
		size[1] = parentSize.Size[1]
	}
	if parentMask.RelativeMask&transform.RelativeSizeZ != 0 {
		size[2] = parentSize.Size[2]
	}
	return size
}

func (s *service) ApplyMinMaxSize(entity ecs.EntityID, size *transform.SizeComponent) {
	if maxSize, ok := s.MaxSizeArray.Get(entity); ok {
		if maxSize.Size[0] != 0 && size.Size[0] > maxSize.Size[0] {
			size.Size[0] = maxSize.Size[0]
		}
		if maxSize.Size[1] != 0 && size.Size[1] > maxSize.Size[1] {
			size.Size[1] = maxSize.Size[1]
		}
		if maxSize.Size[2] != 0 && size.Size[2] > maxSize.Size[2] {
			size.Size[2] = maxSize.Size[2]
		}
	}
	if minSize, ok := s.MinSizeArray.Get(entity); ok {
		if minSize.Size[0] != 0 && size.Size[0] < minSize.Size[0] {
			size.Size[0] = minSize.Size[0]
		}
		if minSize.Size[1] != 0 && size.Size[1] < minSize.Size[1] {
			size.Size[1] = minSize.Size[1]
		}
		if minSize.Size[2] != 0 && size.Size[2] < minSize.Size[2] {
			size.Size[2] = minSize.Size[2]
		}
	}
}

// its used by full ApplyAspectRatio method.
// ignores min and max size.
// also concludes that aspect ratio is verified.
func (s *service) getAspectRatio(size transform.SizeComponent, ratio transform.AspectRatioComponent) transform.SizeComponent {
	var base float32
	switch ratio.PrimaryAxis {
	case transform.PrimaryAxisX:
		base = size.Size[0] / ratio.AspectRatio[0]
	case transform.PrimaryAxisY:
		base = size.Size[1] / ratio.AspectRatio[1]
	case transform.PrimaryAxisZ:
		base = size.Size[2] / ratio.AspectRatio[2]
	default:
		return size
	}
	if ratio.AspectRatio[0] != 0 {
		size.Size[0] = base * ratio.AspectRatio[0]
	}
	if ratio.AspectRatio[1] != 0 {
		size.Size[1] = base * ratio.AspectRatio[1]
	}
	if ratio.AspectRatio[2] != 0 {
		size.Size[2] = base * ratio.AspectRatio[2]
	}
	return size
}

// integrates min and max size
func (s *service) ApplyAspectRatio(entity ecs.EntityID, size *transform.SizeComponent) {
	ratio, _ := s.AspectRatioArray.Get(entity)
	if ratio.PrimaryAxis == 0 || ratio.PrimaryAxis > 3 || ratio.AspectRatio[ratio.PrimaryAxis-1] == 0 {
		return
	}
	sizeRatio := s.getAspectRatio(*size, ratio)
	if maxSize, ok := s.MaxSizeArray.Get(entity); ok {
		for i := range 3 {
			if maxSize.Size[i] == 0 {
				maxSize.Size[i] = sizeRatio.Size[i]
			}
		}
		maxSizeRatio := ratio

		primaryAxisIndex := -1
		for i := range 3 {
			if maxSize.Size[i] != 0 && maxSizeRatio.AspectRatio[i] != 0 &&
				(primaryAxisIndex == -1 || maxSizeRatio.AspectRatio[primaryAxisIndex] > maxSizeRatio.AspectRatio[i]) {
				primaryAxisIndex = i
			}
		}
		if primaryAxisIndex != -1 {
			maxSizeRatio.PrimaryAxis = transform.PrimaryAxis(primaryAxisIndex + 1) // +1 because 0 isn't an axis
			maxSize = transform.MaxSizeComponent(
				s.getAspectRatio(transform.SizeComponent(maxSize), maxSizeRatio),
			)
			for i := range 3 {
				if maxSize.Size[i] >= sizeRatio.Size[i] {
					continue
				}
				sizeRatio = transform.SizeComponent(maxSize)
				break
			}
		}
	}
	if minSize, ok := s.MinSizeArray.Get(entity); ok {
		minSizeRatio := ratio

		primaryAxisIndex := -1
		for i := range 3 {
			if minSize.Size[i] != 0 && minSizeRatio.AspectRatio[i] != 0 &&
				(primaryAxisIndex == -1 || minSizeRatio.AspectRatio[primaryAxisIndex] < minSizeRatio.AspectRatio[i]) {
				primaryAxisIndex = i
			}
		}
		if primaryAxisIndex != -1 {
			minSizeRatio.PrimaryAxis = transform.PrimaryAxis(primaryAxisIndex + 1) // +1 because 0 isn't an axis
			minSize = transform.MinSizeComponent(
				s.getAspectRatio(transform.SizeComponent(minSize), minSizeRatio),
			)
			for i := range 3 {
				if minSize.Size[i] <= sizeRatio.Size[i] {
					continue
				}
				sizeRatio = transform.SizeComponent(minSize)
				break
			}
		}
	}
	*size = sizeRatio
}

//

func (s *service) CalculateAbsolutePos(entity ecs.EntityID, size transform.AbsoluteSizeComponent) transform.AbsolutePosComponent {
	pos, _ := s.PosArray.Get(entity)
	relativeToParentPos := s.GetRelativeParentPos(entity)
	pivotPos := s.GetPivotPos(entity, size)

	pos.Pos = pos.Pos.
		Add(relativeToParentPos).
		Add(pivotPos)

	return transform.AbsolutePosComponent(pos)
}
func (s *service) CalculateAbsoluteRot(entity ecs.EntityID) transform.AbsoluteRotationComponent {
	rot, _ := s.RotationArray.Get(entity)
	rot.Rotation = rot.Rotation.
		Mul(s.GetRelativeParentRotation(entity))

	return transform.AbsoluteRotationComponent(rot)
}
func (s *service) CalculateAbsoluteSize(entity ecs.EntityID) transform.AbsoluteSizeComponent {
	size, _ := s.SizeArray.Get(entity)

	relativeParentSize := s.GetRelativeParentSize(entity)
	size.Size = mgl32.Vec3{
		size.Size[0] * relativeParentSize[0],
		size.Size[1] * relativeParentSize[1],
		size.Size[2] * relativeParentSize[2],
	}
	s.ApplyAspectRatio(entity, &size)
	s.ApplyMinMaxSize(entity, &size)

	return transform.AbsoluteSizeComponent(size)
}

func (s *service) CalculateAbsolute(
	entity ecs.EntityID,
) (transform.AbsolutePosComponent, transform.AbsoluteRotationComponent, transform.AbsoluteSizeComponent) {
	rot := s.CalculateAbsoluteRot(entity)
	size := s.CalculateAbsoluteSize(entity)
	pos := s.CalculateAbsolutePos(entity, size)
	return pos, rot, size
}
