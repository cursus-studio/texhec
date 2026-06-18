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

type interactionService[State any] struct {
	engine.EngineWorld `inject:""`
	interactionGUI     ecs.ComponentsArray[interactions.InteractionGUIComponent[State]]
	missingInteraction ecs.ComponentsArray[interactions.MissingInteractionComponent[State]]
	interaction        ecs.ComponentsArray[interactions.InteractionComponent[State]]
	name               interactions.Name
}

func NewInteractionService[State any](
	c ioc.Dic,
	name interactions.Name,
) interactions.InteractionService[State] {
	s := ioc.GetServices[*interactionService[State]](c)
	s.interactionGUI = ecs.GetComponentsArray[interactions.InteractionGUIComponent[State]](s.World())
	s.missingInteraction = ecs.GetComponentsArray[interactions.MissingInteractionComponent[State]](s.World())
	s.interaction = ecs.GetComponentsArray[interactions.InteractionComponent[State]](s.World())
	s.name = name

	s.missingInteraction.OnRemove(s.Interactions().Proceed)
	s.interaction.OnUpsert(s.Interactions().Proceed)
	s.interaction.OnRemove(s.onInteractionRemove)
	events.Listen(s.EventsBuilder(), s.FinishMeasurement)
	return s
}

func (s *interactionService[State]) onInteractionRemove(ecs.EntityID) {
	entities := s.interactionGUI.GetEntities()
	for _, entity := range entities {
		s.World().RemoveEntity(entity)
	}
}

func (s *interactionService[State]) InteractionGUI() ecs.ComponentsArray[interactions.InteractionGUIComponent[State]] {
	return s.interactionGUI
}
func (s *interactionService[State]) MissingInteraction() ecs.ComponentsArray[interactions.MissingInteractionComponent[State]] {
	return s.missingInteraction
}
func (s *interactionService[State]) Interaction() ecs.ComponentsArray[interactions.InteractionComponent[State]] {
	return s.interaction
}
func (s *interactionService[State]) MissingInteractionAny() ecs.AnyComponentArray {
	return s.missingInteraction
}
func (s *interactionService[State]) InteractionAny() ecs.AnyComponentArray { return s.interaction }
func (s *interactionService[State]) Name() interactions.Name               { return s.name }
func (s *interactionService[State]) StateType() reflect.Type               { return reflect.TypeFor[State]() }
func (s *interactionService[State]) Measure() bool {
	entity := s.Interactions().FeatureEntity()
	if _, ok := s.interaction.Get(entity); ok {
		return true
	}
	s.missingInteraction.Set(entity, s.missingInteraction.GetEmpty())
	return false
}

func (s *interactionService[State]) FinishMeasurement(event interactions.FinishMeasurementEvent[State]) {
	entity := s.Interactions().FeatureEntity()
	if _, ok := s.interaction.Get(entity); ok {
		s.Logger().Warn(fmt.Errorf("measurement is already done"))
		return
	}
	s.interaction.Set(entity, interactions.NewInteraction(event.State))
	s.missingInteraction.Remove(entity)
}
