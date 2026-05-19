package gitpkg

import (
	"cicd/modules/git"
	"cicd/modules/git/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) git.Service {
		return internal.NewService(c)
	})
})
