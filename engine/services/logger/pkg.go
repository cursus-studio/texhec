package logger

import (
	"engine/services/clock"

	"github.com/ogiusek/ioc/v2"
)

type Config interface {
	PanicOnWarn(bool)
	Flush(func(message string))
}

type config struct {
	panicOnWarn bool
	flush       func(message string)
}

func newConfig() Config {
	return &config{
		panicOnWarn: false,
		flush:       func(message string) { print(message) },
	}
}

func (c *config) PanicOnWarn(panicOnWarn bool)     { c.panicOnWarn = panicOnWarn }
func (c *config) Flush(flush func(message string)) { c.flush = flush }

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Config { return newConfig() })
	ioc.Register(b, func(c ioc.Dic) Logger {
		config := ioc.Get[Config](c).(*config)
		return &logger{
			PanicOnError: config.panicOnWarn,
			Clock:        ioc.Get[clock.Clock](c),
			Flush:        config.flush,
			Panic:        func(s string) { panic(s) },
		}
	})
})
