package logger

import (
	"engine/services/clock"

	"github.com/ogiusek/ioc/v2"
)

type config struct {
	panicOnWarn bool
	print       func(c ioc.Dic, message string)
}

func NewConfig(
	panicOnWarn bool,
	print func(c ioc.Dic, message string),
) config {
	return config{
		panicOnWarn: panicOnWarn,
		print:       print,
	}
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, config config) {
	ioc.Register(b, func(c ioc.Dic) Logger {
		return &logger{
			PanicOnError: config.panicOnWarn,
			Clock:        ioc.Get[clock.Clock](c),
			Print:        func(s string) { config.print(c, s) },
			Panic:        func(s string) { panic(s) },
		}
	})
})
