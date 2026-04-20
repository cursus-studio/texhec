package test

import (
	"engine"
	"engine/mock"

	"github.com/ogiusek/ioc/v2"
)

type setup struct {
	engine.EngineWorld `inject:""`
}

func NewSetup() setup {
	c := ioc.NewContainer(
		mock.Pkg,
	)
	return ioc.GetServices[setup](c)
}
