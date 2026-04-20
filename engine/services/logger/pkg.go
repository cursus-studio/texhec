package logger

import (
	"engine/services/clock"

	"github.com/ogiusek/ioc/v2"
)

type config struct {
	panicOnWarn bool
	flush       func(c ioc.Dic) func(message string)
}

func NewConfig(
	panicOnWarn bool,
	flush func(c ioc.Dic) func(message string),
) config {
	return config{
		panicOnWarn: panicOnWarn,
		flush:       flush,
	}
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, config config) {
	ioc.Register(b, func(c ioc.Dic) Logger {
		return &logger{
			PanicOnError: config.panicOnWarn,
			Clock:        ioc.Get[clock.Clock](c),
			Flush:        config.flush(c),
			Panic:        func(s string) { panic(s) },
		}
	})
})
