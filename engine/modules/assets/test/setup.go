package test

import (
	"engine/modules/assets"
	assetspkg "engine/modules/assets/pkg"
	"engine/modules/registry"
	registrypkg "engine/modules/registry/pkg"
	uuidpkg "engine/modules/uuid/pkg"
	"engine/services/clock"
	"engine/services/ecs"
	"engine/services/logger"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type setup struct {
	World  ecs.World      `inject:"1"`
	Assets assets.Service `inject:"1"`

	Registry registry.Service `inject:"1"`
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
