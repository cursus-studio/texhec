package batcherpkg

import (
	"engine/modules/batcher"
	"engine/modules/batcher/internal"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type config struct {
	workers            int
	frameLoadingBudget time.Duration
}

func NewConfig(
	workers int,
	frameLoadingBudget time.Duration,
) config {
	return config{
		workers,
		frameLoadingBudget,
	}
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, config config) {
	ioc.Register(b, func(c ioc.Dic) *internal.Service {
		return internal.NewService(
			c,
			config.workers,
			config.frameLoadingBudget,
		)
	})

	ioc.Register(b, func(c ioc.Dic) batcher.Service {
		return ioc.Get[*internal.Service](c)
	})

	ioc.Register(b, func(c ioc.Dic) batcher.System {
		return ioc.Get[*internal.Service](c).System()
	})
})
