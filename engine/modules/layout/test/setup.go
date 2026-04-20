package test

import (
	"engine"
	"engine/mock"
	"engine/services/ecs"
	"testing"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	engine.EngineWorld `inject:""`
	T                  *testing.T
}

func NewSetup(t *testing.T) Setup {
	c := ioc.NewContainer(
		mock.Pkg,
	)
	setup := ioc.GetServices[Setup](c)
	setup.T = t
	return setup
}

func (s Setup) Expect(entity ecs.EntityID, x, y float32) {
	s.T.Helper()
	expected := mgl32.Vec3{x, y, 1}
	pos, _ := s.Transform().AbsolutePos().Get(entity)
	if pos.Pos != expected {
		pivot, _ := s.Transform().PivotPoint().Get(entity)
		parentPivot, _ := s.Transform().ParentPivotPoint().Get(entity)
		size, _ := s.Transform().AbsoluteSize().Get(entity)

		parent, _ := s.Hierarchy().Parent(entity)
		pSize, _ := s.Transform().AbsoluteSize().Get(parent)
		s.T.Errorf(
			"expected %v and got %v (pivot %v, parent %v, size %v, pSize %v)",
			expected,
			pos,
			pivot,
			parentPivot,
			size,
			pSize,
		)
	}
}
