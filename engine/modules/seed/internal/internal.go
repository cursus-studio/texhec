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
func (s *service) WorldSeed() (ecs.EntityID, bool) {
	seedEntities := s.seed.GetEntities()
	if len(seedEntities) == 0 {
		return 0, false
	}
	if len(seedEntities) > 1 {
		s.Logger().Warn(seed.ErrWorldCanHaveOneSeed)
	}
	for _, seedEntity := range seedEntities[1:] {
		s.World().RemoveEntity(seedEntity)

	}
	return seedEntities[0], true
}
