package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/netsync"
	"engine/modules/netsync/internal/client"
	"engine/modules/netsync/internal/server"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	ClientService      ioc.Lazy[*client.Service] `inject:""`
	ServerService      ioc.Lazy[*server.Service] `inject:""`
	server             ecs.ComponentArray[netsync.ServerComponent]
	client             ecs.ComponentArray[netsync.ClientComponent]
}

func NewService(c ioc.Dic) netsync.Service {
	s := ioc.GetServices[*service](c)
	s.server = ecs.GetComponentArray[netsync.ServerComponent](s.World())
	s.client = ecs.GetComponentArray[netsync.ClientComponent](s.World())
	return s
}

func (s *service) Start() ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		for _, listen := range s.ClientService().ListenToEvents {
			listen(s.EventsBuilder(), s.ClientService().BeforeEvent)
		}
		for _, listen := range s.ClientService().ListenToSimulatedEvents {
			listen(s.EventsBuilder(), s.ClientService().BeforeEventRecord)
		}
		for _, listen := range s.ClientService().ListenToTransparentEvents {
			listen(s.EventsBuilder(), s.ClientService().OnTransparentEvent)
		}

		for _, listen := range s.ServerService().ListenToEvents {
			listen(s.EventsBuilder(), s.ServerService().BeforeEvent)
		}
		for _, listen := range s.ClientService().ListenToSimulatedEvents {
			listen(s.EventsBuilder(), s.ServerService().BeforeEvent)
		}
		for _, listen := range s.ServerService().ListenToTransparentEvents {
			listen(s.EventsBuilder(), s.ServerService().OnTransparentEvent)
		}
		return nil
	})
}
func (s *service) Stop() ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		for _, listen := range s.ClientService().ListenToEvents {
			listen(s.EventsBuilder(), s.ClientService().AfterEvent)
		}
		for _, listen := range s.ClientService().ListenToSimulatedEvents {
			listen(s.EventsBuilder(), s.ClientService().AfterEvent)
		}

		for _, listen := range s.ServerService().ListenToEvents {
			listen(s.EventsBuilder(), s.ServerService().AfterEvent)
		}
		for _, listen := range s.ServerService().ListenToSimulatedEvents {
			listen(s.EventsBuilder(), s.ServerService().AfterEvent)
		}
		return nil
	})
}

func (s *service) Server() ecs.ComponentArray[netsync.ServerComponent] { return s.server }
func (s *service) Client() ecs.ComponentArray[netsync.ClientComponent] { return s.client }
