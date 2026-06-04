package internal

import (
	"engine"
	"engine/modules/interactions"
	"reflect"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type featureService[Event any] struct {
	engine.EngineWorld `inject:""`
	GUI                Service `inject:""`
	name               interactions.Name
	interactions       []interactions.AnyInteractionService
}

func NewFeatureService[Event any](
	c ioc.Dic,
	name interactions.Name,
	featureInteractions []reflect.Type,
) interactions.FeatureService[Event] {
	s := ioc.GetServices[*featureService[Event]](c)
	s.name = name

	s.interactions = make([]interactions.AnyInteractionService, 0, len(featureInteractions))
	for _, interaction := range featureInteractions {
		s.interactions = append(s.interactions, s.GUI.Interactions()[interaction])
	}

	events.Listen(s.EventsBuilder(), s.OnFeature)
	return s
}

func (s *featureService[Event]) Name() interactions.Name { return s.name }
func (s *featureService[Event]) EventType() reflect.Type { return reflect.TypeFor[Event]() }
func (s *featureService[Event]) Interactions() []interactions.AnyInteractionService {
	return s.interactions
}

func (s *featureService[Event]) OnFeature(event interactions.FeatureEvent[Event]) {
	// clear
	for _, entity := range s.GUI.Feature().GetEntities() {
		s.World().RemoveEntity(entity)
	}

	// store feature
	featureEntity := s.World().NewEntity()
	s.GUI.Feature().Set(featureEntity, interactions.FeatureComponent{})
}
