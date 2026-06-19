package transformservice

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/transform"
	"slices"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	DirtySet           ecs.DirtySet

	AbsolutePosArray      ecs.ComponentArray[transform.AbsolutePosComponent]
	AbsoluteSizeArray     ecs.ComponentArray[transform.AbsoluteSizeComponent]
	AbsoluteRotationArray ecs.ComponentArray[transform.AbsoluteRotationComponent]

	AbsolutePosWrapper      ecs.ComponentArray[transform.AbsolutePosComponent]
	AbsoluteSizeWrapper     ecs.ComponentArray[transform.AbsoluteSizeComponent]
	AbsoluteRotationWrapper ecs.ComponentArray[transform.AbsoluteRotationComponent]

	PosArray      ecs.ComponentArray[transform.PosComponent]
	RotationArray ecs.ComponentArray[transform.RotationComponent]
	SizeArray     ecs.ComponentArray[transform.SizeComponent]

	MaxSizeArray     ecs.ComponentArray[transform.MaxSizeComponent]
	MinSizeArray     ecs.ComponentArray[transform.MinSizeComponent]
	AspectRatioArray ecs.ComponentArray[transform.AspectRatioComponent]

	PivotPointArray       ecs.ComponentArray[transform.PivotPointComponent]
	ParentMaskArray       ecs.ComponentArray[transform.ParentComponent]
	ParentPivotPointArray ecs.ComponentArray[transform.ParentPivotPointComponent]

	defaultRot         transform.RotationComponent
	defaultSize        transform.SizeComponent
	defaultPivot       transform.PivotPointComponent
	defaultParentPivot transform.ParentPivotPointComponent
}

func NewService(
	c ioc.Dic,
	defaultRot transform.RotationComponent,
	defaultSize transform.SizeComponent,
	defaultPivot transform.PivotPointComponent,
	defaultParentPivot transform.ParentPivotPointComponent,
) transform.Service {
	s := ioc.GetServices[*service](c)

	s.DirtySet = ecs.NewDirtySet()

	s.AbsolutePosArray = ecs.GetComponentArray[transform.AbsolutePosComponent](s.World())
	s.AbsoluteSizeArray = ecs.GetComponentArray[transform.AbsoluteSizeComponent](s.World())
	s.AbsoluteRotationArray = ecs.GetComponentArray[transform.AbsoluteRotationComponent](s.World())

	s.AbsolutePosWrapper = &absolutePosArray{s, s.AbsolutePosArray}
	s.AbsoluteSizeWrapper = &absoluteSizeArray{s, s.AbsoluteSizeArray}
	s.AbsoluteRotationWrapper = &absoluteRotationArray{s, s.AbsoluteRotationArray}

	s.PosArray = ecs.GetComponentArray[transform.PosComponent](s.World())
	s.SizeArray = ecs.GetComponentArray[transform.SizeComponent](s.World())
	s.RotationArray = ecs.GetComponentArray[transform.RotationComponent](s.World())

	s.MaxSizeArray = ecs.GetComponentArray[transform.MaxSizeComponent](s.World())
	s.MinSizeArray = ecs.GetComponentArray[transform.MinSizeComponent](s.World())
	s.AspectRatioArray = ecs.GetComponentArray[transform.AspectRatioComponent](s.World())

	s.PivotPointArray = ecs.GetComponentArray[transform.PivotPointComponent](s.World())
	s.ParentMaskArray = ecs.GetComponentArray[transform.ParentComponent](s.World())
	s.ParentPivotPointArray = ecs.GetComponentArray[transform.ParentPivotPointComponent](s.World())

	s.defaultRot = transform.NewRotation(mgl32.QuatIdent())
	s.defaultSize = transform.NewSize(1, 1, 1)
	s.defaultPivot = transform.NewPivotPoint(.5, .5, .5)
	s.defaultParentPivot = transform.NewParentPivotPoint(.5, .5, .5)

	s.Init()
	return s

}

func (t *service) BeforeGet() {
	entities := t.DirtySet.Get()
	if len(entities) == 0 {
		return
	}
	children := []ecs.EntityID{}
	entities = slices.DeleteFunc(entities, func(entity ecs.EntityID) bool {
		return !t.World().EntityExists(entity)
	})

	saves := []save{}

	for len(entities) != 0 || len(children) != 0 {
		if len(entities) == 0 {
			for _, save := range saves {
				t.AbsolutePosArray.Set(save.entity, save.pos)
				t.AbsoluteRotationArray.Set(save.entity, save.rot)
				t.AbsoluteSizeArray.Set(save.entity, save.size)
			}
			t.DirtySet.Clear()

			entities = children
			children = nil
			saves = nil
		}
		entity := entities[0]
		entities = entities[1:]

		pos, rot, size := t.CalculateAbsolute(entity)
		save := save{
			entity: entity,
			pos:    pos, rot: rot, size: size,
		}

		saves = append(saves, save)

		for _, child := range t.Hierarchy().Children(entity).GetIndices() {
			comparedMask := transform.RelativePos | transform.RelativeRotation | transform.RelativeSizeXYZ
			mask, _ := t.ParentMaskArray.Get(child)
			if mask.RelativeMask&comparedMask == 0 {
				continue
			}
			children = append(children, child)
		}
	}

	for _, save := range saves {
		t.AbsolutePosArray.Set(save.entity, save.pos)
		t.AbsoluteRotationArray.Set(save.entity, save.rot)
		t.AbsoluteSizeArray.Set(save.entity, save.size)
	}
	t.DirtySet.Clear()
}

func (t *service) AbsolutePos() ecs.ComponentArray[transform.AbsolutePosComponent] {
	return t.AbsolutePosWrapper
}
func (t *service) AbsoluteRotation() ecs.ComponentArray[transform.AbsoluteRotationComponent] {
	return t.AbsoluteRotationWrapper
}
func (t *service) AbsoluteSize() ecs.ComponentArray[transform.AbsoluteSizeComponent] {
	return t.AbsoluteSizeWrapper
}
func (t *service) Pos() ecs.ComponentArray[transform.PosComponent] {
	return t.PosArray
}
func (t *service) Rotation() ecs.ComponentArray[transform.RotationComponent] {
	return t.RotationArray
}
func (t *service) Size() ecs.ComponentArray[transform.SizeComponent] {
	return t.SizeArray
}
func (t *service) MaxSize() ecs.ComponentArray[transform.MaxSizeComponent] {
	return t.MaxSizeArray
}
func (t *service) MinSize() ecs.ComponentArray[transform.MinSizeComponent] {
	return t.MinSizeArray
}
func (t *service) AspectRatio() ecs.ComponentArray[transform.AspectRatioComponent] {
	return t.AspectRatioArray
}
func (t *service) PivotPoint() ecs.ComponentArray[transform.PivotPointComponent] {
	return t.PivotPointArray
}
func (t *service) Parent() ecs.ComponentArray[transform.ParentComponent] {
	return t.ParentMaskArray
}
func (t *service) ParentPivotPoint() ecs.ComponentArray[transform.ParentPivotPointComponent] {
	return t.ParentPivotPointArray
}

func (t *service) Mat4(entity ecs.EntityID) mgl32.Mat4 {
	pos, ok := t.AbsolutePosArray.Get(entity)
	if !ok {
		pos.Pos = mgl32.Vec3{0, 0, 0}
	}
	rot, ok := t.AbsoluteRotationArray.Get(entity)
	if !ok {
		rot.Rotation = mgl32.QuatIdent()
	}
	size, ok := t.AbsoluteSizeArray.Get(entity)
	if !ok {
		size.Size = mgl32.Vec3{1, 1, 1}
	}

	translation := mgl32.Translate3D(pos.Pos.X(), pos.Pos.Y(), pos.Pos.Z())
	rotation := rot.Rotation.Mat4()
	scale := mgl32.Scale3D(size.Size.X()/2, size.Size.Y()/2, size.Size.Z()/2)
	return translation.Mul4(rotation).Mul4(scale)
}

func (t *service) AddDirtySet(set ecs.DirtySet) {
	t.AbsolutePosArray.AddDirtySet(set)
	t.AbsoluteRotationArray.AddDirtySet(set)
	t.AbsoluteSizeArray.AddDirtySet(set)
}
