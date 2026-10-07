package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/feature"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type ruleService[RuleType feature.Rule[DataIn, AppliedTo], DataIn any, AppliedTo any] struct {
	engine.EngineWorld `inject:""`
	Service            ioc.Lazy[Service] `inject:""`
	Config             Config            `inject:""`
	renderMark         ecs.ComponentArray[feature.RenderComponent]
	render             ecs.ComponentArray[feature.RuleRenderComponent[RuleType, DataIn, AppliedTo]]

	applies func(DataIn, AppliedTo) error
	err     error
}

func NewRuleService[RuleType feature.Rule[DataIn, AppliedTo], DataIn any, AppliedTo any](
	c ioc.Dic,
	applies func(DataIn, AppliedTo) error,
	err error,
) feature.RuleService[RuleType, DataIn, AppliedTo] {
	s := ioc.GetServices[*ruleService[RuleType, DataIn, AppliedTo]](c)
	s.renderMark = ecs.GetComponentArray[feature.RenderComponent](s.World())
	s.render = ecs.GetComponentArray[feature.RuleRenderComponent[RuleType, DataIn, AppliedTo]](s.World())

	s.applies = applies
	s.err = err

	s.render.OnUpsert(func(entity ecs.EntityID) {
		s.renderMark.Set(entity, feature.NewRender())
	})
	s.Service().RegisterRuleService(reflect.TypeFor[RuleType](), s)
	return s
}

func (s *ruleService[RuleType, DataIn, AppliedTo]) Render() ecs.ComponentArray[feature.RuleRenderComponent[RuleType, DataIn, AppliedTo]] {
	return s.render
}

func (s *ruleService[RuleType, DataIn, AppliedTo]) Applies(dataIn DataIn, appliedTo AppliedTo) error {
	return s.applies(dataIn, appliedTo)
}

//

func (s *ruleService[RuleType, DataIn, AppliedTo]) RenderAny() ecs.AnyComponentArray { return s.render }
func (s *ruleService[RuleType, DataIn, AppliedTo]) NewAnyRenderComponent(anyDataIn, anyAppliedTo any, render feature.RuleRender) any {
	dataIn, _ := anyDataIn.(DataIn)
	appliedTo, _ := anyAppliedTo.(AppliedTo)
	return feature.NewRuleRender[RuleType](dataIn, appliedTo, render)
}

func (s *ruleService[RuleType, DataIn, AppliedTo]) AppliesToAny(anyDataIn, anyAppliedTo any) error {
	dataIn, _ := anyDataIn.(DataIn)
	appliedTo, _ := anyAppliedTo.(AppliedTo)
	return s.Applies(dataIn, appliedTo)
}

func (s *ruleService[RuleType, DataIn, AppliedTo]) Error() error { return s.err }
