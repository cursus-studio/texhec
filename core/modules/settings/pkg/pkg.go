package settingspkg

import (
	"core/modules/settings"
	"core/modules/settings/internal"
	codecpkg "engine/modules/codec/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[settings.EnterSettingsEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) settings.System {
		return internal.NewSystem(c)
	})
})
