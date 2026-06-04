package internal

import (
	"engine"
	"engine/modules/interactions"
	"engine/services/ecs"
	"fmt"
	"reflect"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type Service interface {
	interactions.Service
	RegisterFeature(interactions.AnyFeatureService)
	RegisterInteraction(interactions.AnyInteractionService)

	Features() map[reflect.Type]interactions.AnyFeatureService
	Interactions() map[reflect.Type]interactions.AnyInteractionService
}

type service struct {
	engine.EngineWorld `inject:""`
	feature            ecs.ComponentsArray[interactions.FeatureComponent]
	featureEvent       ecs.ComponentsArray[interactions.FeatureEventComponent]

	features     map[reflect.Type]interactions.AnyFeatureService
	interactions map[reflect.Type]interactions.AnyInteractionService
}

func NewService(c ioc.Dic) Service {
	s := ioc.GetServices[*service](c)
	s.feature = ecs.GetComponentsArray[interactions.FeatureComponent](s.World())
	s.featureEvent = ecs.GetComponentsArray[interactions.FeatureEventComponent](s.World())

	s.features = make(map[reflect.Type]interactions.AnyFeatureService)
	s.interactions = make(map[reflect.Type]interactions.AnyInteractionService)

	return s
}

// interanal {
func (s *service) RegisterFeature(feature interactions.AnyFeatureService) {
	s.features[feature.EventType()] = feature
}
func (s *service) RegisterInteraction(interaction interactions.AnyInteractionService) {
	s.interactions[interaction.StateType()] = interaction
}
func (s *service) Features() map[reflect.Type]interactions.AnyFeatureService { return s.features }
func (s *service) Interactions() map[reflect.Type]interactions.AnyInteractionService {
	return s.interactions
}

// }

func (s *service) Feature() ecs.ComponentsArray[interactions.FeatureComponent] { return s.feature }
func (s *service) FeatureEvent() ecs.ComponentsArray[interactions.FeatureEventComponent] {
	return s.featureEvent
}

func (s *service) FeatureEntity() ecs.EntityID {
	entities := s.feature.GetEntities()
	if len(entities) == 0 {
		entity := s.World().NewEntity()
		s.feature.Set(entity, s.feature.GetEmpty())
		return entity
	}
	if len(entities) > 1 {
		s.Logger().Warn(fmt.Errorf("GUI cannot handle many features at once"))
		for _, entity := range entities[1:] {
			s.World().RemoveEntity(entity)
		}
	}
	entity := entities[0]
	return entity
}

// on AnyFeatureComponent or FinishedInteractionComponent upsert
func (s *service) Proceed(ecs.EntityID) {
	entity := s.FeatureEntity()
	featureEvent, _ := s.featureEvent.Get(entity)
	featureKey := reflect.TypeOf(featureEvent.Event)
	feature, ok := s.features[featureKey]
	if !ok {
		return
	}
	interactions := feature.Interactions()
	// if there are measurements left start new measurement
	for _, interaction := range interactions {
		if _, ok := interaction.MissingInteractionAny().GetAny(entity); ok {
			return
		}
		if _, ok := interaction.InteractionAny().GetAny(entity); !ok {
			interaction.MissingInteractionAny().SetAny(entity, nil)
			return
		}
	}

	// if all measurements are done emit feature event and remove all features and measurements
	events.EmitAny(s.Events(), featureEvent.Event)
	s.World().RemoveEntity(entity)
}
