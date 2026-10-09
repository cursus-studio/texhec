package internal

import (
	"engine/modules/ecs"
	"engine/modules/feature"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type FieldIndex = int

type FieldConfig = *fieldConfig
type FeatureFieldConfig = *featureFieldConfig
type RuleConfig = *ruleConfig
type FeatureConfig = *featureConfig
type FeatureFieldIndexConfig = *featureIndexConfig

type fieldConfig struct {
	Key     feature.FieldKey
	Provide func() any
}
type featureFieldConfig struct {
	FieldConfig
	Rules []RuleConfig
}
type ruleConfig struct {
	RuleType feature.RuleKey
	Rule     any
	Service  feature.AnyRuleService
}
type featureConfig struct {
	Fields []FeatureFieldConfig
}
type featureIndexConfig struct {
	Features []feature.FeatureKey
}

func NewFieldConfig(key feature.FieldKey) FieldConfig {
	return &fieldConfig{
		Key:     key,
		Provide: nil,
	}
}
func NewFeatureFieldConfig(fieldConfig FieldConfig) FeatureFieldConfig {
	return &featureFieldConfig{
		FieldConfig: fieldConfig,
		Rules:       make([]RuleConfig, 0),
	}
}
func NewRuleConfig[RuleType feature.Rule](rule RuleType) RuleConfig {
	config := &ruleConfig{
		RuleType: reflect.TypeFor[RuleType](),
		Rule:     rule,
	}
	return config
}
func NewFeatureConfig[FeatureType any](
	fieldsConfigs map[feature.FieldKey]FieldConfig,
) FeatureConfig {
	featureKey := reflect.TypeFor[FeatureType]()
	fields := featureKey.NumField()
	config := &featureConfig{
		Fields: make([]FeatureFieldConfig, fields),
	}
	for i := range fields {
		fieldKey := featureKey.Field(i).Type
		fieldConfig, ok := fieldsConfigs[fieldKey]
		if !ok {
			fieldConfig = NewFieldConfig(fieldKey)
			fieldsConfigs[fieldKey] = fieldConfig
		}
		config.Fields[i] = NewFeatureFieldConfig(fieldConfig)
	}
	return config
}
func NewFeatureFieldIndexConfig(
	featureKey feature.FeatureKey) FeatureFieldIndexConfig {
	return &featureIndexConfig{
		Features: make([]feature.FeatureKey, 0),
	}
}

//

type config struct {
	systems []ecs.SystemRegister

	featureConfigs    map[feature.FeatureKey]FeatureConfig
	features          []feature.FeatureKey
	featureFieldIndex map[feature.FieldKey]FeatureFieldIndexConfig

	fields map[feature.FieldKey]FieldConfig
}
type Config = *config

func NewConfig(c ioc.Dic) Config {
	return &config{
		systems: make([]ecs.SystemRegister, 0),

		featureConfigs:    make(map[feature.FeatureKey]FeatureConfig),
		features:          make([]feature.FeatureKey, 0),
		featureFieldIndex: make(map[feature.FieldKey]FeatureFieldIndexConfig),

		fields: make(map[feature.FieldKey]FieldConfig),
	}
}

// getter methods for tests

// SHOULD ONLY BE USED IN TESTS
func (c *config) GetFeatureConfig(key feature.FeatureKey) FeatureConfig {
	return c.featureConfigs[key]
}

// register methods

// returns first field which isn't provided field
func featureFieldIndex(config Config, featureKey feature.FeatureKey) feature.FieldKey {
	featureConfig := config.featureConfigs[featureKey]
	for _, field := range featureConfig.Fields {
		if field.Provide == nil {
			return field.Key
		}
	}
	panic("feature has to have not provided fields")
}
func addFeatureFieldIndex(config Config, featureKey feature.FeatureKey) {
	fieldKey := featureFieldIndex(config, featureKey)
	featureFieldIndex, ok := config.featureFieldIndex[fieldKey]
	if !ok {
		featureFieldIndex = NewFeatureFieldIndexConfig(featureKey)
		config.featureFieldIndex[fieldKey] = featureFieldIndex
	}
	featureFieldIndex.Features = append(featureFieldIndex.Features, featureKey)
}
func RegisterFieldProvider[FieldType any](
	config Config, provide feature.ProvideFieldService[FieldType]) {
	fieldKey := reflect.TypeFor[FieldType]()
	fieldConfig, ok := config.fields[fieldKey]
	if !ok {
		fieldConfig = NewFieldConfig(fieldKey)
		config.fields[fieldKey] = fieldConfig
	}
	fieldConfig.Provide = func() any { return provide() }

	fieldIndexConfig, ok := config.featureFieldIndex[fieldKey]
	if !ok {
		return
	}
	delete(config.featureFieldIndex, fieldKey)
	for _, featureKey := range fieldIndexConfig.Features {
		addFeatureFieldIndex(config, featureKey)
	}
}
func RegisterRuleToField[FieldRule feature.Rule, Feature any, AppliedTo any](
	config Config, zero *Feature, appliedTo *AppliedTo, rule FieldRule) {
	featureKey := reflect.TypeFor[Feature]()
	appliedToIndex := feature.GetFieldIndex(zero, appliedTo)
	feat, ok := config.featureConfigs[featureKey]
	if !ok {
		RegisterFeature[Feature](config)
		feat = config.featureConfigs[featureKey]
	}
	ruleConfig := feat.Fields[appliedToIndex]
	ruleConfig.Rules = append(ruleConfig.Rules, NewRuleConfig(rule))
}
func RegisterFeature[Feature any](config Config) (zero Feature) {
	featureKey := reflect.TypeFor[Feature]()
	if _, ok := config.featureConfigs[featureKey]; ok {
		return zero
	}
	config.featureConfigs[featureKey] = NewFeatureConfig[Feature](config.fields)
	config.features = append(config.features, featureKey)
	addFeatureFieldIndex(config, featureKey)

	return zero
}
func RegisterSystems(config Config, systems ...ecs.SystemRegister) {
	config.systems = append(config.systems, systems...)
}
