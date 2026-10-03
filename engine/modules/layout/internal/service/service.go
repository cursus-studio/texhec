package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/layout"
	"engine/modules/transform"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	align         ecs.ComponentArray[layout.AlignComponent]
	order         ecs.ComponentArray[layout.OrderComponent]
	gap           ecs.ComponentArray[layout.GapComponent]
	dirtyParents  ecs.DirtySet
	dirtyChildren ecs.DirtySet
}

func NewLayoutService(c ioc.Dic) layout.Service {
	s := ioc.GetServices[*service](c)
	s.align = ecs.GetComponentArray[layout.AlignComponent](s.World())
	s.order = ecs.GetComponentArray[layout.OrderComponent](s.World())
	s.gap = ecs.GetComponentArray[layout.GapComponent](s.World())
	s.dirtyParents = ecs.NewDirtySet()
	s.dirtyChildren = ecs.NewDirtySet()
	s.Init()
	return s
}

func (s *service) Align() ecs.ComponentArray[layout.AlignComponent] { return s.align }
func (s *service) Order() ecs.ComponentArray[layout.OrderComponent] { return s.order }
func (s *service) Gap() ecs.ComponentArray[layout.GapComponent]     { return s.gap }

//

func (s *service) Init() {
	// t.order.SetEmpty(layout.NewOrder(layout.OrderHorizontal))
	s.align.SetEmpty(layout.NewAlign(.5, .5))
	s.gap.SetEmpty(layout.NewGap(0))

	s.Transform().AbsolutePos().AddDependency(s.align)
	s.Transform().AbsolutePos().AddDependency(s.order)
	s.Transform().AbsolutePos().AddDependency(s.gap)

	s.align.AddDirtySet(s.dirtyParents)
	s.order.AddDirtySet(s.dirtyParents)
	s.gap.AddDirtySet(s.dirtyParents)
	s.Transform().AddDirtySet(s.dirtyParents)

	s.Transform().AddDirtySet(s.dirtyChildren)
	s.Hierarchy().Component().AddDirtySet(s.dirtyChildren)

	// before get trigger
	s.Transform().AbsolutePos().BeforeGet(s.BeforeGet)
	s.Transform().AbsoluteSize().BeforeGet(s.BeforeGet)
}

type save struct {
	entity      ecs.EntityID
	pos         transform.PosComponent
	pivot       transform.PivotPointComponent
	parentPivot transform.ParentPivotPointComponent
}

func (s *service) BeforeGet() {
	for _, child := range s.dirtyChildren.Get() {
		if parent, ok := s.Hierarchy().Parent(child); ok {
			s.dirtyParents.Dirty(parent)
		}
	}
	parents := s.dirtyParents.Get()
	if len(parents) == 0 {
		return
	}
	defer s.dirtyChildren.Clear()
	defer s.dirtyParents.Clear()

	saves := []save{}

	for _, parent := range parents {
		parentSaves := s.handleParentChildren(parent)
		saves = append(saves, parentSaves...)
	}

	for _, save := range saves {
		s.Transform().Pos().Set(save.entity, save.pos)
		s.Transform().PivotPoint().Set(save.entity, save.pivot)
		s.Transform().ParentPivotPoint().Set(save.entity, save.parentPivot)
	}
}

func (s *service) handleParentChildren(parent ecs.EntityID) []save {
	children := s.Hierarchy().Children(parent).GetIndices()
	if len(children) == 0 {
		return nil
	}
	order, ok := s.order.Get(parent)
	if !ok {
		return nil
	}
	saves := make([]save, 0, len(children))
	align, _ := s.align.Get(parent)
	gap, _ := s.gap.Get(parent)

	// including gaps
	var totalSize float32 = 0
	for _, child := range children {
		size, _ := s.Transform().AbsoluteSize().Get(child)
		totalSize += size.Size[order.Order] + gap.Gap
	}
	totalSize -= gap.Gap

	size, _ := s.Transform().AbsoluteSize().Get(parent)
	progress := totalSize - size.Size[order.Primary()]
	progress *= align.Primary

	for _, child := range children {
		// pos
		pos := transform.NewPos(0, 0, 1)
		pos.Pos[order.Primary()] = progress

		// pivot point
		pivot := transform.NewPivotPoint(.5, .5, .5)
		pivot.Point[order.Primary()] = 1
		pivot.Point[order.Secondary()] = align.Secondary

		// parent pivot
		parentPivot := transform.NewParentPivotPoint(.5, .5, .5)
		parentPivot.Point[order.Primary()] = 1
		parentPivot.Point[order.Secondary()] = align.Secondary

		save := save{
			child,
			pos,
			pivot,
			parentPivot,
		}
		saves = append(saves, save)

		// update progress
		size, _ := s.Transform().AbsoluteSize().Get(child)
		progress -= size.Size[order.Primary()] + gap.Gap
	}

	return saves
}
