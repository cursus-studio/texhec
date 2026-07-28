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
	t := ioc.GetServices[*service](c)
	t.uuidArray = ecs.GetComponentArray[uuid.Component](t.World())
	return t
}

func (t *service) Component() ecs.ComponentArray[uuid.Component] { return t.uuidArray }

func (t *service) Entity(uuid uuid.UUID) (ecs.EntityID, bool) {
	return t.Get(uuid)
}
