package test

import (
	"engine/modules/hierarchy"
	hierarchypkg "engine/modules/hierarchy/pkg"
	"engine/modules/transform"
	transformpkg "engine/modules/transform/pkg"
	"engine/services/clock"
	"engine/services/ecs"
	"engine/services/logger"
	"testing"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	ecs.World
	hierarchy hierarchy.Service
	transform transform.Service
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		logger.Pkg(logger.NewConfig(true, func(c ioc.Dic, message string) { print(message) })),
		clock.Pkg(time.RFC3339Nano),
		ecs.Pkg,
		hierarchypkg.Pkg,
		transformpkg.Pkg,
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
