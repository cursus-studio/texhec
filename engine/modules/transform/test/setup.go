package test

import (
	"engine/modules/hierarchy"
	"engine/modules/transform"
	"engine/pkg"
	"engine/services/ecs"
	"testing"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	ecs.World
	hierarchy hierarchy.Service
	transform transform.Service
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
	)
	return Setup{
		ioc.Get[ecs.World](c),
		ioc.Get[hierarchy.Service](c),
		ioc.Get[transform.Service](c),
	}
}

func (setup Setup) expectAbsolutePos(t *testing.T, entity ecs.EntityID, expectedPos transform.PosComponent) {
	t.Helper()
	pos, _ := setup.transform.AbsolutePos().Get(entity)
	if pos.Pos != expectedPos.Pos {
		t.Errorf("expected pos %v but has %v", expectedPos, pos)
	}
}

func (setup Setup) expectAbsoluteSize(t *testing.T, entity ecs.EntityID, expectedSize transform.SizeComponent) {
	t.Helper()
	size, _ := setup.transform.AbsoluteSize().Get(entity)
	if size.Size != expectedSize.Size {
		t.Errorf("expected size %v but has %v", expectedSize, size)
	}
}
