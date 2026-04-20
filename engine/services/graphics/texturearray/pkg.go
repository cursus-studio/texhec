package texturearray

import (
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Factory {
		return &factory{
			make([]func(TextureArray), 0),
		}
	})
})
