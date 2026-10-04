package internal

import (
	"engine"
	"engine/modules/feature"
	"slices"

	"github.com/ogiusek/ioc/v2"
)

type Service interface {
	feature.Service

	RegisterRuleService(feature.RuleKey, feature.AnyRuleService)
	GetRuleService(feature.RuleKey) feature.AnyRuleService
	RegisterFieldService(feature.FieldKey, feature.AnyFieldService)
	GetFieldService(feature.FieldKey) feature.AnyFieldService
}

type service struct {
	World  engine.EngineWorld `inject:""`
	Config Config             `inject:""`
	feature.InteractionsService

	rules  map[feature.RuleKey]feature.AnyRuleService
	fields map[feature.FieldKey]feature.AnyFieldService
}

func NewService(c ioc.Dic) Service {
	s := ioc.GetServices[*service](c)
	s.InteractionsService = NewInteractionsService(c)

	s.rules = make(map[feature.RuleKey]feature.AnyRuleService)
	s.fields = make(map[feature.FieldKey]feature.AnyFieldService)
	return s
}

//

func (s *service) RegisterFieldService(key feature.FieldKey, anyFieldService feature.AnyFieldService) {
	s.fields[key] = anyFieldService
}
func (s *service) GetFieldService(key feature.FieldKey) feature.AnyFieldService {
	return s.fields[key]
}
func (s *service) RegisterRuleService(key feature.RuleKey, anyRuleService feature.AnyRuleService) {
	s.rules[key] = anyRuleService
}
func (s *service) GetRuleService(key feature.RuleKey) feature.AnyRuleService {
	return s.rules[key]
}

//

func (s *service) Register() error {
	for _, s := range s.Config.systems {
		if err := s.Register(); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Features() []feature.FeatureKey {
	return slices.Clone(s.Config.features)
}
