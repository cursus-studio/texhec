package test

import (
	"engine/services/clock"
	"engine/services/codec"
	"engine/services/logger"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type Type struct {
	Value int
}

type setup struct {
	codec codec.Codec
}

func NewSetup() setup {
	c := ioc.NewContainer(
		codec.Pkg,
		clock.Pkg(time.RFC3339Nano),
		logger.Pkg(logger.NewConfig(true, func(c ioc.Dic, message string) { print(message) })),
		ioc.NewPkg(func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
				b.Register(Type{})
			})
		}),
	)
	return setup{ioc.Get[codec.Codec](c)}
}
