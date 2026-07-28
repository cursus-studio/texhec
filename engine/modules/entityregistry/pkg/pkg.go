package entityregistrypkg

import (
	"engine/modules/entityregistry"
	"engine/modules/entityregistry/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) entityregistry.Service {
		return internal.NewService(c)
	})
})
