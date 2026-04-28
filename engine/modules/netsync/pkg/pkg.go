package netsyncpkg

import (
	"engine/modules/netsync"
	"engine/modules/netsync/internal/client"
	"engine/modules/netsync/internal/clienttypes"
	"engine/modules/netsync/internal/server"
	"engine/modules/netsync/internal/servertypes"
	"engine/modules/netsync/internal/service"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/services/ecs"

	"github.com/ogiusek/events"
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
	ioc.Register(b, func(c ioc.Dic) netsync.StartSystem {
		clientService := ioc.Get[*client.Service](c)
		serverService := ioc.Get[*server.Service](c)
		eventsBuilder := ioc.Get[events.Builder](c)
		return ecs.NewSystemRegister(func() error {
			for _, listen := range clientService.ListenToEvents {
				listen(eventsBuilder, clientService.BeforeEvent)
			}
			for _, listen := range clientService.ListenToSimulatedEvents {
				listen(eventsBuilder, clientService.BeforeEventRecord)
			}
			for _, listen := range clientService.ListenToTransparentEvents {
				listen(eventsBuilder, clientService.OnTransparentEvent)
			}

			for _, listen := range serverService.ListenToEvents {
				listen(eventsBuilder, serverService.BeforeEvent)
			}
			for _, listen := range clientService.ListenToSimulatedEvents {
				listen(eventsBuilder, serverService.BeforeEvent)
			}
			for _, listen := range serverService.ListenToTransparentEvents {
				listen(eventsBuilder, serverService.OnTransparentEvent)
			}
			return nil
		})
	})
	ioc.Register(b, func(c ioc.Dic) netsync.StopSystem {
		clientService := ioc.Get[*client.Service](c)
		serverService := ioc.Get[*server.Service](c)
		eventsBuilder := ioc.Get[events.Builder](c)
		return ecs.NewSystemRegister(func() error {
			for _, listen := range clientService.ListenToEvents {
				listen(eventsBuilder, clientService.AfterEvent)
			}
			for _, listen := range clientService.ListenToSimulatedEvents {
				listen(eventsBuilder, clientService.AfterEvent)
			}

			for _, listen := range serverService.ListenToEvents {
				listen(eventsBuilder, serverService.AfterEvent)
			}
			for _, listen := range serverService.ListenToSimulatedEvents {
				listen(eventsBuilder, serverService.AfterEvent)
			}
			return nil
		})
	})
})
