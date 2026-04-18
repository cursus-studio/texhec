package internal

import (
	"engine"
	"engine/modules/relation"
	"engine/modules/uuid"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld          `inject:""`
	relation.Service[uuid.UUID] `inject:""`
	uuid.Factory                `inject:""`

	uuidArray ecs.ComponentsArray[uuid.Component]
}

func NewService(c ioc.Dic) uuid.Service {
	t := ioc.GetServices[*service](c)
	t.uuidArray = ecs.GetComponentsArray[uuid.Component](t.World())
	return t
}

func (t *service) Component() ecs.ComponentsArray[uuid.Component] { return t.uuidArray }

func (t *service) Entity(uuid uuid.UUID) (ecs.EntityID, bool) {
	return t.Get(uuid)
}
