package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/feature"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type ruleService[RuleType feature.Rule] struct {
	engine.EngineWorld `inject:""`
	Service            ioc.Lazy[Service] `inject:""`
	Config             Config            `inject:""`
	renderMark         ecs.ComponentArray[feature.RenderComponent]
	render             ecs.ComponentArray[feature.RuleRenderComponent[RuleType]]

	applies func(RuleType, feature.AllFields) error
	err     error
}

func NewRuleService[RuleType feature.Rule](
	c ioc.Dic,
	applies func(RuleType, feature.AllFields) error,
	err error,
) feature.RuleService[RuleType] {
	s := ioc.GetServices[*ruleService[RuleType]](c)
	s.renderMark = ecs.GetComponentArray[feature.RenderComponent](s.World())
	s.render = ecs.GetComponentArray[feature.RuleRenderComponent[RuleType]](s.World())

	s.applies = applies
	s.err = err

	s.render.OnUpsert(func(entity ecs.EntityID) {
		s.renderMark.Set(entity, feature.NewRender())
	})
	s.Service().RegisterRuleService(reflect.TypeFor[RuleType](), s)
	return s
}

func (s *ruleService[RuleType]) Render() ecs.ComponentArray[feature.RuleRenderComponent[RuleType]] {
	return s.render
}

func (s *ruleService[RuleType]) Applies(rule RuleType, allFields []any) error {
	return s.applies(rule, allFields)
}

func (s *ruleService[RuleType]) AppliesAny(ruleAny any, allFields []any) error {
	if rule, ok := ruleAny.(RuleType); ok {
		return s.applies(rule, allFields)
	}
	return s.Error()
}

//

func (s *ruleService[RuleType]) RenderAny() ecs.AnyComponentArray { return s.render }
func (s *ruleService[RuleType]) NewAnyRenderComponent(rule any, fields feature.AllFields, render feature.RuleRender) any {
	return feature.NewRuleRender(rule.(RuleType), fields, render)
}

func (s *ruleService[RuleType]) Error() error { return s.err }
