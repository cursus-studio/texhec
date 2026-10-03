package internal

import (
	"encoding/json"
	"engine"
	"engine/modules/ecs"
	"engine/modules/relation"
	"engine/modules/uuid"
	"reflect"

	uuidSource "github.com/google/uuid"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld          `inject:""`
	relation.Service[uuid.UUID] `inject:""`

	uuidArray ecs.ComponentArray[uuid.Component]
}

func NewService(c ioc.Dic) uuid.Service {
	s := ioc.GetServices[*service](c)
	s.uuidArray = ecs.GetComponentArray[uuid.Component](s.World())
	return s
}

func (s *service) NewUUID() uuid.UUID {
	return uuid.UUID(uuidSource.New())
}

func (s *service) NewUUIDFromString(seed string) uuid.UUID {
	return uuid.UUID(uuidSource.NewSHA1(uuidSource.NameSpaceURL, []byte(seed)))
}

type SeededUUID struct {
	Tick int64
	Seed any
}

func (s *service) NewUUIDFromAny(seed any) uuid.UUID {
	bytes, err := json.Marshal(seed)
	if err != nil {
		s.Logger().Fatal(err)
	}
	bytes = append([]byte(reflect.TypeOf(seed).String()), bytes...)
	return uuid.UUID(uuidSource.NewSHA1(uuidSource.NameSpaceURL, bytes))
}

//

func (s *service) Component() ecs.ComponentArray[uuid.Component] { return s.uuidArray }

func (s *service) Entity(uuid uuid.UUID) (ecs.EntityID, bool) {
	return s.Get(uuid)
}
