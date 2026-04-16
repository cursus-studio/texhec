package dragpkg

import (
	"engine/modules/drag"
	"engine/modules/drag/internal"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// events
			Register(drag.DraggableEvent{})
	})

	ioc.Register(b, func(c ioc.Dic) drag.System {
		return internal.NewSystem(c)
	})
}
