package internal

import (
	"core/game"
	"core/modules/player"
	"engine/modules/interactions"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type contextSetter struct {
	game.GameWorld `inject:""`
}

func NewContextSetter(c ioc.Dic) interactions.ContextSetter {
	return ioc.GetServices[*contextSetter](c)
}

func (s *contextSetter) getContext() player.PlayerContext {
	entities := s.Player().ActingPlayer().GetEntities()
	if len(entities) != 1 {
		return player.PlayerContext{}
	}
	entity := entities[0]
	uuid, ok := s.Player().PlayerUUID().Get(entity)
	if !ok {
		return player.PlayerContext{}
	}
	return player.PlayerContext{PlayerUUID: uuid.UUID}
}

func (s *contextSetter) SetContext(feat interactions.Feature) interactions.Feature {
	v := reflect.ValueOf(feat)
	ptrCopy := reflect.New(v.Type())
	structCopy := ptrCopy.Elem()
	structCopy.Set(v)

	fieldVal := structCopy.FieldByName(reflect.TypeFor[player.PlayerContext]().Name())
	if fieldVal.IsValid() && fieldVal.CanSet() {
		fieldVal.Set(reflect.ValueOf(s.getContext()))
	}

	return structCopy.Interface()
}
