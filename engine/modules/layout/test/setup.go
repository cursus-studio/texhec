package test

import (
	"engine"
	hierarchypkg "engine/modules/hierarchy/pkg"
	layoutpkg "engine/modules/layout/pkg"
	transformpkg "engine/modules/transform/pkg"
	"engine/services/clock"
	"engine/services/ecs"
	"engine/services/logger"
	"testing"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	engine.EngineWorld `inject:""`
	T                  *testing.T
}

func NewSetup(t *testing.T) Setup {
	c := ioc.NewContainer(
		logger.Pkg(logger.NewConfig(true, func(c ioc.Dic, message string) { print(message) })),
		clock.Pkg(time.RFC3339Nano),
		ecs.Pkg,
		hierarchypkg.Pkg,
		transformpkg.Pkg,
		layoutpkg.Pkg,
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
