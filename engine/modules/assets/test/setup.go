package test

import (
	"engine"
	enginepkg "engine/pkg"

	"github.com/ogiusek/ioc/v2"
)

type setup struct {
	engine.EngineWorld `inject:""`
}

func NewSetup() setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
	)
	return ioc.GetServices[setup](c)
}
