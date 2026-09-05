package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/relation"
	"engine/modules/uuid"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld          `inject:""`
	relation.Service[uuid.UUID] `inject:""`
	uuid.Factory                `inject:""`

	uuidArray ecs.ComponentArray[uuid.Component]
}

func NewService(c ioc.Dic) uuid.Service {
	s := ioc.GetServices[*service](c)
	s.uuidArray = ecs.GetComponentArray[uuid.Component](s.World())
	return s
}

func (s *service) Component() ecs.ComponentArray[uuid.Component] { return s.uuidArray }

func (s *service) Entity(uuid uuid.UUID) (ecs.EntityID, bool) {
	return s.Get(uuid)
}
