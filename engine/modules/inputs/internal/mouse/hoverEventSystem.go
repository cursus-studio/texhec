package mouse

import (
	"engine"
	"engine/modules/inputs"
	"engine/services/ecs"
	"engine/services/frames"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type hoverEventSystem struct {
	engine.EngineWorld `inject:""`
}

func NewHoverEventsSystem(c ioc.Dic) inputs.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*hoverEventSystem](c)
		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (s *hoverEventSystem) Listen(event frames.FrameEvent) {
	for _, entity := range s.Inputs().Hovered().GetEntities() {
		eventsComponent, ok := s.Inputs().Hover().Get(entity)
		if !ok {
			continue
		}

		if e, ok := eventsComponent.Event.(inputs.ApplyEntityEvent); ok {
			eventsComponent.Event = e.ApplyEntity(entity)
		}
		if setter, ok := eventsComponent.Event.(inputs.EventTargetSetter); ok {
			for _, data := range s.Inputs().StackedData() {
				if data.Entity != entity {
					continue
				}
				eventsComponent.Event = setter.SetTarget(data)
				break
			}

		}
		events.EmitAny(s.Events(), eventsComponent.Event)
	}
}
