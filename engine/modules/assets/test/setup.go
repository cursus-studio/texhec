package test

import (
	"engine"
	assetspkg "engine/modules/assets/pkg"
	registrypkg "engine/modules/registry/pkg"
	uuidpkg "engine/modules/uuid/pkg"
	"engine/services/clock"
	"engine/services/ecs"
	"engine/services/logger"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type setup struct {
	engine.EngineWorld `inject:""`
}

func NewSetup() setup {
	c := ioc.NewContainer(
		clock.Pkg(time.RFC3339Nano),
		logger.Pkg(logger.NewConfig(true, func(c ioc.Dic, message string) { print(message) })),
		registrypkg.Pkg,
		ecs.Pkg,
		uuidpkg.Pkg,
		assetspkg.Pkg(""),
	)
	return ioc.GetServices[setup](c)
}
