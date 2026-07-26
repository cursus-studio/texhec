package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/interactions"
	"errors"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

var (
	ErrInteractionIsMissing error = errors.New("missing interaction state")
)

type AnyStepService interface {
	Key() interactions.StepKey
	Interaction() AnyInteractionService

	FillValue(ecs.EntityID, reflect.Value)
	// does rule apply to entity
	EntityRule(ecs.EntityID) error
}
type StepService[StepT interactions.Step[State], State any] interface {
	AnyStepService
	Rule(State) error
}

type stepService[StepT interactions.Step[State], State any] struct {
	engine.EngineWorld `inject:""`
	Interactions       Service                   `inject:""`
	InteractionService InteractionService[State] `inject:""`
	rule               func(State) error
}

func NewStepService[StepT interactions.Step[State], State any](
	c ioc.Dic,
	rule func(State) error,
) StepService[StepT, State] {
	s := ioc.GetServices[*stepService[StepT, State]](c)
	s.rule = rule

	return s
}

func (s *stepService[StepT, State]) Key() interactions.StepKey          { return reflect.TypeFor[StepT]() }
func (s *stepService[StepT, State]) Interaction() AnyInteractionService { return s.InteractionService }

func (s *stepService[StepT, State]) FillValue(entity ecs.EntityID, value reflect.Value) {
	state, _ := s.InteractionService.State().Get(entity)
	step := interactions.NewStepT[StepT](state.State)
	value.Set(reflect.ValueOf(step))
}

// does rule apply to entity
func (s *stepService[StepT, State]) EntityRule(entity ecs.EntityID) error {
	state, ok := s.InteractionService.State().Get(entity)
	if !ok {
		return ErrInteractionIsMissing
	}
	return s.rule(state.State)
}
func (s *stepService[StepT, State]) Rule(state State) error { return s.rule(state) }
