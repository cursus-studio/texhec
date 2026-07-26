package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/interactions"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type StateComponent[State any] struct{ State State }
type MissingComponent[State any] struct{}

func NewState[State any](state State) StateComponent[State] { return StateComponent[State]{state} }
func NewMissing[State any]() MissingComponent[State]        { return MissingComponent[State]{} }

type AnyInteractionService interface {
	Key() interactions.InteractionKey
	MarkMissing(propertiesEntity, interactionEntity ecs.EntityID)
}
type InteractionService[State any] interface {
	AnyInteractionService
	interactions.InteractionService[State]

	State() ecs.ComponentArray[StateComponent[State]]
	Missing() ecs.ComponentArray[MissingComponent[State]]
}

type interactionService[State any] struct {
	engine.EngineWorld `inject:""`
	Interactions       Service `inject:""`
	state              ecs.ComponentArray[StateComponent[State]]
	missing            ecs.ComponentArray[MissingComponent[State]]

	statePreview   ecs.ComponentArray[interactions.StatePreviewComponent[State]]
	missingPreview ecs.ComponentArray[interactions.MissingPreviewComponent[State]]
}

func NewInteractionService[State any](c ioc.Dic) InteractionService[State] {
	s := ioc.GetServices[*interactionService[State]](c)
	s.state = ecs.GetComponentArray[StateComponent[State]](s.World())
	s.missing = ecs.GetComponentArray[MissingComponent[State]](s.World())

	s.statePreview = ecs.GetComponentArray[interactions.StatePreviewComponent[State]](s.World())
	s.missingPreview = ecs.GetComponentArray[interactions.MissingPreviewComponent[State]](s.World())

	s.state.OnUpsert(s.addInteractionKey)
	s.missing.OnUpsert(s.addInteractionKey)

	s.state.OnUpsert(s.Interactions.OnChangeProgress)

	return s
}

func (s *interactionService[State]) Key() interactions.InteractionKey {
	return reflect.TypeFor[State]()
}
func (s *interactionService[State]) MarkMissing(propertiesEntity, interactionEntity ecs.EntityID) {
	defer s.World().RemoveEntity(propertiesEntity)
	s.state.Remove(interactionEntity)
	s.missing.Set(interactionEntity, NewMissing[State]())
	if previewComp, ok := s.Interactions.Preview().Get(interactionEntity); ok {
		s.World().RemoveEntity(previewComp.PreviewEntity)
	}

	previewEntity := s.World().NewEntity()
	s.Prototype().CloneTo(propertiesEntity, previewEntity)
	s.Interactions.Preview().Set(interactionEntity, NewPreview(previewEntity))
	s.MissingPreview().Set(previewEntity, interactions.NewMissingPreview[State]())
}

func (s *interactionService[State]) State() ecs.ComponentArray[StateComponent[State]] {
	return s.state
}
func (s *interactionService[State]) Missing() ecs.ComponentArray[MissingComponent[State]] {
	return s.missing
}

func (s *interactionService[State]) StatePreview() ecs.ComponentArray[interactions.StatePreviewComponent[State]] {
	return s.statePreview
}
func (s *interactionService[State]) MissingPreview() ecs.ComponentArray[interactions.MissingPreviewComponent[State]] {
	return s.missingPreview
}

func (s *interactionService[State]) Save(propertiesEntity ecs.EntityID, state State) {
	defer s.World().RemoveEntity(propertiesEntity)
	if entities := s.Missing().GetEntities(); len(entities) != 0 {
		interactionEntity := entities[0]
		s.Prototype().CloneTo(propertiesEntity, interactionEntity)
		s.Missing().Remove(interactionEntity)
		if previewComp, ok := s.Interactions.Preview().Get(interactionEntity); ok {
			s.World().RemoveEntity(previewComp.PreviewEntity)
		}
		previewEntity := s.World().NewEntity()
		s.Interactions.Preview().Set(interactionEntity, NewPreview(previewEntity))
		s.StatePreview().Set(previewEntity, interactions.NewStatePreview(state))
		s.State().Set(interactionEntity, NewState(state))
		return
	}
	s.Interactions.ResetFeatureEntity()
	featureEntity := s.Interactions.FeatureEntity()

	interactionEntity := s.World().NewEntity()
	s.Prototype().CloneTo(propertiesEntity, interactionEntity)

	previewEntity := s.World().NewEntity()
	s.Hierarchy().SetParent(interactionEntity, featureEntity)
	s.Interactions.Preview().Set(interactionEntity, NewPreview(previewEntity))
	s.StatePreview().Set(previewEntity, interactions.NewStatePreview(state))
	s.State().Set(interactionEntity, NewState(state))
}

func (s *interactionService[State]) addInteractionKey(entity ecs.EntityID) {
	s.Interactions.InteractionKey().Set(entity, NewInteractionKey[State]())
}
