package internal

import (
	"engine"
	"engine/modules/warmup"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
}

func NewService(c ioc.Dic) warmup.Service {
	return ioc.GetServices[*service](c)
}

func (s *service) WarmUp() {
	events.Emit(s.Events(), warmup.Event{})
}
