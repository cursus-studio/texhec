package frames

import (
	"engine/services/clock"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type config struct {
	tps,
	fps int
}

func NewConfig(tps, fps int) config {
	return config{
		tps: tps,
		fps: fps,
	}
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, config config) {
	ioc.Register(b, func(c ioc.Dic) Builder {
		return NewBuilder(config.tps, config.fps)
	})

	ioc.Register(b, func(c ioc.Dic) Frames {
		return ioc.Get[Builder](c).Build(ioc.Get[events.Events](c), ioc.Get[clock.Clock](c))
	})
})
