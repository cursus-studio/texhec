package test

import (
	enginepkg "engine/pkg"
	"engine/services/codec"

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
		enginepkg.Pkg,
		ioc.NewPkg(func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
				b.Register(Type{})
			})
		}),
	)
	return setup{ioc.Get[codec.Codec](c)}
}
