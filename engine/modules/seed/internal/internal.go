package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/seed"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	seed               ecs.ComponentArray[seed.SeedComponent]
}

func NewService(c ioc.Dic) seed.Service {
	s := ioc.GetServices[*service](c)
	s.seed = ecs.GetComponentArray[seed.SeedComponent](s.World())
	return s
}

func (s *service) Seed() ecs.ComponentArray[seed.SeedComponent] {
	return s.seed
}

// #nosec G115
func (s *service) GenerateSeed() seed.Seed { return seed.Seed(s.Clock().Now().UnixNano()) }
