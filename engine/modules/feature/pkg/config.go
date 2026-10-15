package featurepkg

import (
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/feature/internal"
)

type Config = internal.Config

func RegisterFieldProvider[FieldType any](
	config Config, provide feature.ProvideFieldService[FieldType]) {
	internal.RegisterFieldProvider(config, provide)
}
func RegisterRuleToField[FieldRule feature.Rule, Feature any](
	config Config, _ Feature, rule FieldRule) {
	internal.RegisterRuleToField[Feature](config, rule)
}
func RegisterFeature[Feature any](config Config) (zero Feature) {
	return internal.RegisterFeature[Feature](config)
}
func RegisterSystems(config Config, systems ...ecs.SystemRegister) {
	internal.RegisterSystems(config, systems...)
}
