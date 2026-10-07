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
				dataIn := fields[rule.DataInFieldIndex]
				if err := ruleService.AppliesToAny(dataIn, field); err != nil {
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

	// get preview type
	var previewFieldKey feature.FieldKey
	var previewFieldConfig FeatureFieldConfig
	allFields := make([]any, 0, featureKey.NumField())
	allFieldsConfigs := make([]FeatureFieldConfig, 0, featureKey.NumField())
	{
		fieldsLeft := slices.Clone(interactionsComp.Interactions)
		for numField := range featureKey.NumField() {
			fieldConfig := featureConfig.Fields[numField]
			if fieldConfig.Provide != nil {
				allFields = append(allFields, fieldConfig.Provide())
				allFieldsConfigs = append(allFieldsConfigs, fieldConfig)
				continue
			}
			fieldKey := featureKey.Field(numField).Type
			if len(fieldsLeft) == 0 {
				previewFieldKey = fieldKey
				previewFieldConfig = fieldConfig
				break
			}
			allFields = append(allFields, fieldsLeft[0])
			allFieldsConfigs = append(allFieldsConfigs, fieldConfig)
			fieldsLeft = fieldsLeft[1:]
		}
	}
	previewFieldService := s.Service().GetFieldService(previewFieldKey)

	// do not render a preview if there is no feature selected
	if len(interactionsComp.Interactions) == 0 {
		interactionsComp.PreviewInteraction = nil
	}

	// for Interaction in [feature.InteractionsComponent.Interactions] render:
	// [feature.RuleRenderComponent] with [feature.RuleRenderCompletePreview]
	for i, fieldConfig := range allFieldsConfigs {
		for _, rule := range fieldConfig.Rules {
			ruleService := s.Service().GetRuleService(rule)
			dataIn := allFields[rule.DataInFieldIndex]
			appliedTo := allFields[i]

			renderComp := ruleService.NewAnyRenderComponent(dataIn, appliedTo, feature.RuleRenderCompletePreview)

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
		dataIn := allFields[rule.DataInFieldIndex]
		appliedTo := interactionsComp.PreviewInteraction

		rulesPass = rulesPass && ruleService.AppliesToAny(dataIn, appliedTo) == nil
	}

	// for Rule in [feature.InteractionsComponent.PreviewInteraction] use:
	// [feature.RuleRenderComponent] with [feature.RuleRenderPreviewRulesFail] and [feature.RuleRenderPreviewRulesPass]
	for _, rule := range previewFieldConfig.Rules {
		ruleService := s.Service().GetRuleService(rule)
		dataIn := allFields[rule.DataInFieldIndex]
		appliedTo := interactionsComp.PreviewInteraction

		ruleRender := feature.RuleRenderPreviewRulesPass
		if !rulesPass {
			ruleRender = feature.RuleRenderPreviewRulesFail
		}
		renderComp := ruleService.NewAnyRenderComponent(dataIn, appliedTo, ruleRender)

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
			dataIn := allFields[rule.DataInFieldIndex]
			if err := ruleService.AppliesToAny(dataIn, field); err != nil {
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
	s.World.World().RemoveEntity(entity)
	for _, entity := range slices.Clone(s.render.GetEntities()) {
		s.World.World().RemoveEntity(entity)
	}
	s.Interactions()
	events.EmitAny(s.World.Events(), featureValue.Interface())
}
