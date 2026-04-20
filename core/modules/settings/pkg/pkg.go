package settingspkg

import (
	"core/modules/settings"
	"core/modules/settings/internal"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// events
			Register(settings.EnterSettingsEvent{})
	})

	ioc.Register(b, func(c ioc.Dic) settings.System {
		return internal.NewSystem(c)
	})
})
