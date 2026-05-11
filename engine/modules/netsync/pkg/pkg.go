package netsyncpkg

import (
	"engine/modules/netsync"
	"engine/modules/netsync/internal/client"
	"engine/modules/netsync/internal/clienttypes"
	"engine/modules/netsync/internal/server"
	"engine/modules/netsync/internal/servertypes"
	"engine/modules/netsync/internal/service"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[clienttypes.PredictedEvent],
		typeregistrypkg.PkgT[clienttypes.FetchStateDTO],
		typeregistrypkg.PkgT[clienttypes.EmitEventDTO],
		typeregistrypkg.PkgT[clienttypes.TransparentEventDTO],

		typeregistrypkg.PkgT[servertypes.SendStateDTO],
		typeregistrypkg.PkgT[servertypes.SendChangeDTO],
		typeregistrypkg.PkgT[servertypes.TransparentEventDTO],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) Config { return newConfig() })

	ioc.Register(b, func(c ioc.Dic) netsync.Service {
		return service.NewService(c)
	})

	ioc.Register(b, func(c ioc.Dic) *server.Service {
		return server.NewService(c, *ioc.Get[Config](c).config)
	})
	ioc.Register(b, func(c ioc.Dic) *client.Service {
		return client.NewService(c, *ioc.Get[Config](c).config)
	})
})
