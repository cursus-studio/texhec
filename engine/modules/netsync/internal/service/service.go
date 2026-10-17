package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/netsync"
	"engine/modules/netsync/internal/client"
	"engine/modules/netsync/internal/clienttypes"
	"engine/modules/netsync/internal/config"
	"engine/modules/netsync/internal/prediction"
	"engine/modules/netsync/internal/server"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	PredictionService  ioc.Lazy[prediction.Service] `inject:""`
	ClientService      ioc.Lazy[*client.Service]    `inject:""`
	ServerService      ioc.Lazy[*server.Service]    `inject:""`
	Config             ioc.Lazy[config.Config]      `inject:""`
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
		for _, config := range s.Config().EventConfigs {
			config.Builder(s.EventsBuilder(), func(a any) {
				if len(s.Server().GetEntities()) != 0 {
					err := s.PredictionService().StartPrediction(a)
					if err == prediction.ErrPredictionMachineFailed {
						conn := s.ClientService().GetConnection()
						if conn != nil {
							err := conn.Send(clienttypes.ReconcileDTO{})
							s.Logger().Warn(err)
						}
					} else {
						s.Logger().Warn(err)
					}
				}
				if len(s.Clients().GetEntities()) != 0 {
					s.ServerService().BeforeEvent(a)
				}
			})
		}
		return nil
	})
}
func (s *service) Stop() ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s.ServerService().AddAfterListeners()
		for _, config := range s.Config().EventConfigs {
			config.Builder(s.EventsBuilder(), func(a any) {
				if len(s.Server().GetEntities()) != 0 {
					s.Logger().Warn(s.PredictionService().StopPrediction())
				}
			})
		}
		return nil
	})
}

func (s *service) Server() ecs.ComponentArray[netsync.ServerComponent]   { return s.server }
func (s *service) Clients() ecs.ComponentArray[netsync.ClientsComponent] { return s.clients }
func (s *service) Client() ecs.ComponentArray[netsync.ClientComponent]   { return s.client }
