package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/netsync"
	"engine/modules/netsync/internal/client"
	"engine/modules/netsync/internal/config"
	"engine/modules/netsync/internal/server"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	ClientService      ioc.Lazy[*client.Service] `inject:""`
	ServerService      ioc.Lazy[*server.Service] `inject:""`
	Config             config.InjectedConfig     `inject:""`
	server             ecs.ComponentArray[netsync.ServerComponent]
	clients            ecs.ComponentArray[netsync.ClientsComponent]
	client             ecs.ComponentArray[netsync.ClientComponent]
}

func NewService(c ioc.Dic) netsync.Service {
	s := ioc.GetServices[*service](c)
	s.server = ecs.GetComponentArray[netsync.ServerComponent](s.World())
	s.clients = ecs.GetComponentArray[netsync.ClientsComponent](s.World())
	s.client = ecs.GetComponentArray[netsync.ClientComponent](s.World())
	return s
}

func (s *service) Start() ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s.ServerService().AddBeforeListeners()
		for _, listen := range s.Config().ListenToEvents {
			listen(s.EventsBuilder(), s.ClientService().BeforeEvent)
			listen(s.EventsBuilder(), s.ServerService().BeforeEvent)
		}
		for _, listen := range s.Config().ListenToSimulatedEvents {
			listen(s.EventsBuilder(), s.ClientService().BeforeEventRecord)
			listen(s.EventsBuilder(), s.ServerService().BeforeEvent)
		}
		for _, listen := range s.Config().ListenToTransparentEvents {
			listen(s.EventsBuilder(), s.ClientService().OnTransparentEvent)
			listen(s.EventsBuilder(), s.ServerService().OnTransparentEvent)
		}
		for _, listen := range s.Config().ListenToVerifyHappenEvents {
			listen(s.EventsBuilder(), s.ClientService().OnVerifyEventHappen)
			listen(s.EventsBuilder(), s.ServerService().OnVerifyEventHappen)
		}
		return nil
	})
}
func (s *service) Stop() ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		for _, listen := range s.Config().ListenToEvents {
			listen(s.EventsBuilder(), s.ClientService().AfterEvent)
			listen(s.EventsBuilder(), s.ServerService().AfterEvent)
		}
		for _, listen := range s.Config().ListenToSimulatedEvents {
			listen(s.EventsBuilder(), s.ClientService().AfterEvent)
			listen(s.EventsBuilder(), s.ServerService().AfterEvent)
		}
		return nil
	})
}

func (s *service) Server() ecs.ComponentArray[netsync.ServerComponent]   { return s.server }
func (s *service) Clients() ecs.ComponentArray[netsync.ClientsComponent] { return s.clients }
func (s *service) Client() ecs.ComponentArray[netsync.ClientComponent]   { return s.client }
