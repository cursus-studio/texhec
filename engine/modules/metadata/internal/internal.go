package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/metadata"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	name        ecs.ComponentArray[metadata.NameComponent]
	description ecs.ComponentArray[metadata.DescriptionComponent]
	link        ecs.ComponentArray[metadata.LinkComponent]
}

func NewService(c ioc.Dic) metadata.Service {
	s := ioc.GetServices[*service](c)
	s.name = ecs.GetComponentArray[metadata.NameComponent](s.World())
	s.description = ecs.GetComponentArray[metadata.DescriptionComponent](s.World())
	s.link = ecs.GetComponentArray[metadata.LinkComponent](s.World())
	return s
}

func (s *service) Name() ecs.ComponentArray[metadata.NameComponent] {
	return s.name
}
func (s *service) Description() ecs.ComponentArray[metadata.DescriptionComponent] {
	return s.description
}
func (s *service) Link() ecs.ComponentArray[metadata.LinkComponent] {
	return s.link
}
