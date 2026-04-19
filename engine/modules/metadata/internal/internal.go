package internal

import (
	"engine"
	"engine/modules/metadata"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	name        ecs.ComponentsArray[metadata.NameComponent]
	description ecs.ComponentsArray[metadata.DescriptionComponent]
	link        ecs.ComponentsArray[metadata.LinkComponent]
}

func NewService(c ioc.Dic) metadata.Service {
	s := ioc.GetServices[*service](c)
	s.name = ecs.GetComponentsArray[metadata.NameComponent](s.World())
	s.description = ecs.GetComponentsArray[metadata.DescriptionComponent](s.World())
	s.link = ecs.GetComponentsArray[metadata.LinkComponent](s.World())
	return s
}

func (s *service) Name() ecs.ComponentsArray[metadata.NameComponent] {
	return s.name
}
func (s *service) Description() ecs.ComponentsArray[metadata.DescriptionComponent] {
	return s.description
}
func (s *service) Link() ecs.ComponentsArray[metadata.LinkComponent] {
	return s.link
}
