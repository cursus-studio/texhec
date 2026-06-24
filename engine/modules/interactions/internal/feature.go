package internal

import (
	"engine"
	"engine/modules/interactions"
	"reflect"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type populateFeatureEvent[Event any] struct{}

type featureService[Event any] struct {
	engine.EngineWorld  `inject:""`
	InteractionsService Service `inject:""`
	name                interactions.Name
	featureInteractions []reflect.Type
	interactions        []interactions.AnyInteractionService
	populate            func() Event
}

func NewFeatureService[Event any](
	c ioc.Dic,
	name interactions.Name,
	featureInteractions []reflect.Type,
	populate func() Event,
) interactions.FeatureService[Event] {
	s := ioc.GetServices[*featureService[Event]](c)
	s.name = name
	s.featureInteractions = featureInteractions
	s.populate = populate

	s.InteractionsService.Feature().OnUpsert(s.EngineWorld.Interactions().Proceed)

	events.Listen(s.EventsBuilder(), s.OnPopulate)
	events.Listen(s.EventsBuilder(), s.OnFeature)
	return s
}

func (s *featureService[Event]) Name() interactions.Name { return s.name }
func (s *featureService[Event]) EventType() reflect.Type {
	return reflect.TypeFor[populateFeatureEvent[Event]]()
}
func (s *featureService[Event]) Interactions() []interactions.AnyInteractionService {
	if len(s.interactions) != 0 {
		return s.interactions
	}

	s.interactions = make([]interactions.AnyInteractionService, 0, len(s.featureInteractions))
	for _, featureInteraction := range s.featureInteractions {
		s.interactions = append(s.interactions, s.InteractionsService.Interactions()[featureInteraction])
	}
	return s.interactions
}

func (s *featureService[Event]) OnFeature(event interactions.FeatureEvent[Event]) {
	featureEntity := s.InteractionsService.FeatureEntity()
	s.InteractionsService.Feature().Set(featureEntity, interactions.NewFeature(populateFeatureEvent[Event]{}))
}

func (s *featureService[Event]) OnPopulate(event populateFeatureEvent[Event]) {
	events.Emit(s.Events(), s.populate())
}
