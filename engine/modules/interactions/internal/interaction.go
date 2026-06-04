package internal

import (
	"engine"
	"engine/modules/interactions"
	"engine/services/ecs"
	"fmt"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type interactionService[State any] struct {
	engine.EngineWorld `inject:""`
	missingInteraction ecs.ComponentsArray[interactions.MissingInteractionComponent[State]]
	interaction        ecs.ComponentsArray[interactions.InteractionComponent[State]]
	name               interactions.Name
}

func NewInteractionService[State any](
	c ioc.Dic,
	name interactions.Name,
) interactions.InteractionService[State] {
	s := ioc.GetServices[*interactionService[State]](c)
	s.missingInteraction = ecs.GetComponentsArray[interactions.MissingInteractionComponent[State]](s.World())
	s.interaction = ecs.GetComponentsArray[interactions.InteractionComponent[State]](s.World())
	s.name = name
	s.interaction.OnUpsert(s.Interactions().Proceed)
	return s
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

func (s *interactionService[State]) FinishMeasurement(state State) error {
	entity := s.Interactions().FeatureEntity()
	if _, ok := s.interaction.Get(entity); ok {
		return fmt.Errorf("measurement is already done")
	}
	s.missingInteraction.Remove(entity)
	s.interaction.Set(entity, interactions.NewInteraction(state))
	return nil
}
