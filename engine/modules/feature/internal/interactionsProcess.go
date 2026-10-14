package internal

import (
	"engine/modules/ecs"
	"engine/modules/feature"
	"reflect"
	"slices"

	"github.com/ogiusek/events"
)

func (s *interactionsService) getFieldFeatures(field any) []feature.FeatureKey {
	fieldKey := reflect.TypeOf(field)
	fieldIndexConfig, ok := s.Config.featureFieldIndex[fieldKey]
	if !ok {
		return nil
	}
	features := slices.Clone(fieldIndexConfig.Features)

	features = slices.DeleteFunc(features, func(featureKey feature.FeatureKey) bool {
		featureConfig := s.Config.featureConfigs[featureKey]
		fields := []any{}
		for _, fieldConfig := range featureConfig.Fields {
			if fieldConfig.Provide != nil {
				fields = append(fields, fieldConfig.Provide())
				continue
			}
			fields = append(fields, field)
			for _, rule := range fieldConfig.Rules {
				ruleService := s.Service().GetRuleService(rule)
				if err := ruleService.AppliesAny(rule.Rule, fields); err != nil {
					return true
				}
			}
			break
		}
		return false
	})
	return features
}

func (s *interactionsService) RenderState(entity ecs.EntityID) {
	interactionsComp, ok := s.interactions.Get(entity)
	if !ok {
		return
	}
	for _, entity := range slices.Clone(s.render.GetEntities()) {
		s.World.World().RemoveEntity(entity)
	}

	// for Interaction in [feature.InteractionsComponent.Interactions] use:
	// [feature.FieldRenderComponent] with [feature.FieldRenderInteraction]
	for _, interaction := range interactionsComp.Interactions {
		fieldKey := reflect.TypeOf(interaction)
		fieldService := s.Service().GetFieldService(fieldKey)
		comp := fieldService.NewAnyRenderComponent(interaction, feature.FieldRenderInteraction)

		interactionEntity := s.World.World().NewEntity()
		fieldService.RenderAny().SetAny(interactionEntity, comp)
	}

	if !interactionsComp.LockedFeature {
		return
	}
	featureKey := interactionsComp.Features[0]
	featureConfig := s.Config.featureConfigs[featureKey]

	allFields := s.AllFields(interactionsComp)
	previewFieldIndex := len(allFields)
	if interactionsComp.PreviewInteraction != nil {
		previewFieldIndex--
	}

	var previewFieldKey feature.FieldKey
	var previewFieldConfig FeatureFieldConfig
	if previewFieldIndex < featureKey.NumField() {
		previewFieldKey = featureKey.Field(previewFieldIndex).Type
		previewFieldConfig = featureConfig.Fields[previewFieldIndex]
	}

	previewFieldService := s.Service().GetFieldService(previewFieldKey)

	// do not render a preview if there is no feature selected
	if len(interactionsComp.Interactions) == 0 {
		interactionsComp.PreviewInteraction = nil
	}

	// for Interaction in [feature.InteractionsComponent.Interactions] render:
	// [feature.RuleRenderComponent] with [feature.RuleRenderCompletePreview]
	for i := range previewFieldIndex {
		fieldConfig := featureConfig.Fields[i]
		for _, rule := range fieldConfig.Rules {
			ruleService := s.Service().GetRuleService(rule)
			renderComp := ruleService.NewAnyRenderComponent(rule.Rule, allFields, feature.RuleRenderCompletePreview)

			previewRuleEntity := s.World.World().NewEntity()
			ruleService.RenderAny().SetAny(previewRuleEntity, renderComp)
		}
	}

	rulesPass := interactionsComp.PreviewInteraction != nil
	for _, rule := range previewFieldConfig.Rules {
		if !rulesPass {
			break
		}
		ruleService := s.Service().GetRuleService(rule)
		rulesPass = rulesPass && ruleService.AppliesAny(rule.Rule, allFields) == nil
	}

	// for Rule in [feature.InteractionsComponent.PreviewInteraction] use:
	// [feature.RuleRenderComponent] with [feature.RuleRenderPreviewRulesFail] and [feature.RuleRenderPreviewRulesPass]
	for _, rule := range previewFieldConfig.Rules {
		ruleService := s.Service().GetRuleService(rule)

		ruleRender := feature.RuleRenderPreviewRulesPass
		if !rulesPass {
			ruleRender = feature.RuleRenderPreviewRulesFail
		}
		renderComp := ruleService.NewAnyRenderComponent(rule.Rule, allFields, ruleRender)

		previewRuleEntity := s.World.World().NewEntity()
		ruleService.RenderAny().SetAny(previewRuleEntity, renderComp)
	}

	// for [feature.InteractionsComponent.PreviewInteraction] use:
	// [feature.FieldRenderComponent] with [feature.FieldRenderPreviewRulesPass] and [feature.FieldRenderPreviewRulesFail]
	var comp any
	if interactionsComp.PreviewInteraction == nil {
		comp = previewFieldService.NewAnyRenderComponent(
			interactionsComp.PreviewInteraction, feature.FieldRenderMissing)
	} else if rulesPass {
		comp = previewFieldService.NewAnyRenderComponent(
			interactionsComp.PreviewInteraction, feature.FieldRenderPreviewRulesPass)
	} else if !rulesPass {
		comp = previewFieldService.NewAnyRenderComponent(
			interactionsComp.PreviewInteraction, feature.FieldRenderPreviewRulesFail)
	}
	previewInteractionEntity := s.World.World().NewEntity()
	previewFieldService.RenderAny().SetAny(previewInteractionEntity, comp)
}

func (s *interactionsService) ProcessState(
	entity ecs.EntityID,
	interactionsComp feature.InteractionsComponent,
) {
	defer s.RenderState(entity)
	if !interactionsComp.LockedFeature {
		if len(interactionsComp.Interactions) == 0 {
			interactionsComp = feature.NewInteractions(
				nil, interactionsComp.PreviewInteraction, nil, false)
		} else {
			if len(interactionsComp.Interactions) > 1 {
				s.interactions.Remove(entity)
			}
			field := interactionsComp.Interactions[len(interactionsComp.Interactions)-1]
			interactionsComp = feature.NewInteractions(
				[]any{field}, interactionsComp.PreviewInteraction, s.getFieldFeatures(field), false)
		}
		s.interactions.Set(entity, interactionsComp)
		return
	}

	featureKey := interactionsComp.Features[0]
	featureConfig := s.Config.featureConfigs[featureKey]
	interactions := slices.Clone(interactionsComp.Interactions)
	lastPassingInteractionIndex := 0
	allFields := []any{}
fieldsLoop:
	for numField := range featureKey.NumField() {
		fieldType := featureKey.Field(numField).Type
		fieldConfig := featureConfig.Fields[numField]

		var field any
		if fieldConfig.Provide != nil {
			field = fieldConfig.Provide()
		} else if len(interactions) != 0 {
			field = interactions[0]
			interactions = interactions[1:]
			lastPassingInteractionIndex++
		} else {
			break
		}
		allFields = append(allFields, field)

		if reflect.TypeOf(field) != fieldType {
			lastPassingInteractionIndex--
			allFields = allFields[:len(allFields)-1]
			break
		}
		for _, rule := range fieldConfig.Rules {
			ruleService := s.Service().GetRuleService(rule)
			if err := ruleService.AppliesAny(rule.Rule, allFields); err != nil {
				lastPassingInteractionIndex--
				allFields = allFields[:len(allFields)-1]
				s.World.Logger().Warn(err)
				break fieldsLoop
			}
		}
	}

	if len(allFields) < featureKey.NumField() {
		interactionsComp.Interactions = interactionsComp.Interactions[:lastPassingInteractionIndex]
		s.interactions.Set(entity, interactionsComp)
		return
	}

	featureValue := reflect.New(interactionsComp.Features[0]).Elem()
	for numField := range featureKey.NumField() {
		featureValue.Field(numField).Set(reflect.ValueOf(allFields[numField]))
	}
	s.interactions.Remove(entity)
	s.interactions.Set(entity, feature.NewInteractions(nil, nil, nil, false))
	for _, entity := range slices.Clone(s.render.GetEntities()) {
		s.World.World().RemoveEntity(entity)
	}
	events.EmitAny(s.World.Events(), featureValue.Interface())
}
