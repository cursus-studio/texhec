package ecs

import (
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) events.Builder {
		return events.NewBuilder()
	})
	ioc.Register(b, func(c ioc.Dic) events.Events {
		return ioc.Get[events.Builder](c).Build()
	})
	ioc.Register(b, func(c ioc.Dic) World {
		return NewWorld()
	})
})
