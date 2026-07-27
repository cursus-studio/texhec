package ecspkg

import (
	"engine/modules/ecs"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) ecs.World {
		return ecs.NewWorld()
	})
})
