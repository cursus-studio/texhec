package projectfspkg

import (
	"cicd/modules/projectfs"
	"cicd/modules/projectfs/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) projectfs.Service {
		return internal.NewService(c)
	})
})
