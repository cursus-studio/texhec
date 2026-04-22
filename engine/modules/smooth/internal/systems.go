package internal

import (
	"engine"
	"engine/modules/loop"
	"engine/modules/record"
	"engine/modules/transition"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type Service[Component transition.LerpConstraint[Component]] struct {
	engine.EngineWorld `inject:""`
	recordingID        record.RecordingID
	config             record.Config

	componentArray ecs.ComponentsArray[Component]
	lerpArray      ecs.ComponentsArray[transition.TransitionComponent[Component]]
}

func NewService[Component transition.LerpConstraint[Component]](c ioc.Dic) *Service[Component] {
	config := record.NewConfig()
	record.AddToConfig[Component](config)

	s := ioc.GetServices[*Service[Component]](c)

	s.recordingID = 0
	s.config = config
	s.componentArray = ecs.GetComponentsArray[Component](s.World())
	s.lerpArray = ecs.GetComponentsArray[transition.TransitionComponent[Component]](s.World())
	return s
}

//

type system[Component transition.LerpConstraint[Component]] struct {
	engine.EngineWorld `inject:""`
	Service            *Service[Component] `inject:""`
}

type FirstEvent loop.TickEvent
type LastEvent loop.TickEvent

func NewSystems[Component transition.LerpConstraint[Component]](c ioc.Dic) {
	s := ioc.GetServices[*system[Component]](c)
	events.Listen(s.EventsBuilder(), func(FirstEvent) {
		for _, entity := range s.Service.lerpArray.GetEntities() {
			transitionComponent, ok := s.Service.lerpArray.Get(entity)
			if !ok {
				continue
			}
			s.Service.lerpArray.Remove(entity)
			s.Service.componentArray.Set(entity, transitionComponent.To)
		}

		s.Service.recordingID = s.Record().Entity().StartBackwardsRecording(s.Service.config)
	})

	events.Listen(s.EventsBuilder(), func(tick LastEvent) {
		r, ok := s.Record().Entity().Stop(s.Service.recordingID)
		if !ok {
			return
		}
		for _, entity := range r.Entities.GetIndices() {
			beforeComponents, ok := r.Entities.Get(entity)
			if !ok || beforeComponents == nil {
				continue
			}
			before, ok := beforeComponents[0].(Component)
			if !ok {
				continue
			}
			after, ok := s.Service.componentArray.Get(entity)
			if !ok {
				continue
			}
			lerpComponent := transition.NewTransition(before, after, tick.Delta)
			s.Service.lerpArray.Set(entity, lerpComponent)
		}
	})
}
