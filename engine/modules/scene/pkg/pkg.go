package scenepkg

import (
	"engine/modules/scene"
	"engine/modules/scene/internal"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// events
			Register(scene.ChangeSceneEvent{})
	})

	ioc.Register(b, func(c ioc.Dic) scene.Service {
		return internal.NewService(c)
	})
})
