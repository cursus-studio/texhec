package clockpkg

import (
	"engine/modules/clock"
	"engine/modules/clock/internal"
	"time"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) clock.Service {
		return internal.NewService(c, clock.DateFormat(time.RFC3339Nano))
	})
})
