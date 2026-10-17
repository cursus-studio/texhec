package netsyncpkg

import (
	"engine/modules/netsync"
	"engine/modules/netsync/internal/client"
	"engine/modules/netsync/internal/clienttypes"
	"engine/modules/netsync/internal/config"
	"engine/modules/netsync/internal/prediction"
	"engine/modules/netsync/internal/server"
	"engine/modules/netsync/internal/servertypes"
	"engine/modules/netsync/internal/service"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[clienttypes.ReconcileDTO],
		typeregistrypkg.PkgT[clienttypes.PredictedEventDTO],

		typeregistrypkg.PkgT[servertypes.ReconcileDTO],
		typeregistrypkg.PkgT[servertypes.EventDTO],
		typeregistrypkg.PkgT[servertypes.PredictionResponseDTO],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) Config { return newConfig() })
	ioc.Register(b, func(c ioc.Dic) config.Config {
		return ioc.Get[Config](c).config
	})

	ioc.Register(b, func(c ioc.Dic) netsync.Service {
		return service.NewService(c)
	})

	ioc.Register(b, func(c ioc.Dic) prediction.Service {
		return prediction.NewService(c)
	})
	ioc.Register(b, func(c ioc.Dic) prediction.IDService { return prediction.NewEventIDService(c) })
	ioc.Register(b, func(c ioc.Dic) *server.Service { return server.NewService(c) })
	ioc.Register(b, func(c ioc.Dic) *client.Service { return client.NewService(c) })
})
