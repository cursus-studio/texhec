package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/loop"
	"reflect"
	"slices"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type interactionsService struct {
	World   engine.EngineWorld `inject:""`
	Service ioc.Lazy[Service]  `inject:""`
	Config  Config             `inject:""`

	interactions ecs.ComponentArray[feature.InteractionsComponent]
}

func NewInteractionsService(c ioc.Dic) feature.InteractionsService {
	s := ioc.GetServices[*interactionsService](c)

	s.interactions = ecs.GetComponentArray[feature.InteractionsComponent](s.World.World())

	events.Listen(s.World.EventsBuilder(), func(event feature.LockFeatureEvent) {
		entity := s.InteractionsEntity()
		comp, _ := s.interactions.Get(entity)
		comp.Features = []feature.FeatureKey{event.Feature}
		comp.LockedFeature = true
		s.ProcessState(entity, comp)
	})
	events.Listen(s.World.EventsBuilder(), func(loop.TickEvent) {
		entity := s.InteractionsEntity()
		comp, _ := s.interactions.Get(entity)
		s.ProcessState(entity, comp)
	})
	return s
}

func (s *interactionsService) getFieldFeatures(field any) []feature.FeatureKey {
	fieldKey := reflect.TypeOf(field)
	fieldIndexConfig := s.Config.featureFieldIndex[fieldKey]
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
				ruleService := s.Service().GetRuleService(rule.RuleType)
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
	for _, child := range slices.Clone(s.World.Hierarchy().Children(entity).GetIndices()) {
		s.World.World().RemoveEntity(child)
	}

	// for Interaction in [feature.InteractionsComponent.Interactions] use:
	// [feature.FieldRenderComponent] with [feature.Interaction]
	for _, interaction := range interactionsComp.Interactions {
		fieldKey := reflect.TypeOf(interaction)
		fieldService := s.Service().GetFieldService(fieldKey)
		comp := fieldService.NewAnyRenderComponent(interaction, feature.Interaction)

		interactionEntity := s.World.World().NewEntity()
		s.World.Hierarchy().SetParent(interactionEntity, entity)
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
	{
		fieldsLeft := slices.Clone(interactionsComp.Interactions)
		for numField := range featureKey.NumField() {
			fieldConfig := featureConfig.Fields[numField]
			if fieldConfig.Provide != nil {
				allFields = append(allFields, fieldConfig.Provide())
				continue
			}
			fieldKey := featureKey.Field(numField).Type
			if len(fieldsLeft) == 0 {
				previewFieldKey = fieldKey
				previewFieldConfig = fieldConfig
				break
			}
			allFields = append(allFields, fieldsLeft[0])
			fieldsLeft = fieldsLeft[1:]
		}
	}
	previewFieldService := s.Service().GetFieldService(previewFieldKey)

	// for Rule in [feature.InteractionsComponent.PreviewInteraction] use:
	// [feature.RuleRenderComponent]
	rulesPass := true
	for _, rule := range previewFieldConfig.Rules {
		ruleService := s.Service().GetRuleService(rule.RuleType)
		dataIn := allFields[rule.DataInFieldIndex]
		appliedTo := interactionsComp.PreviewInteraction

		renderComp := ruleService.NewAnyRenderComponent(dataIn, appliedTo)

		previewRuleEntity := s.World.World().NewEntity()
		s.World.Hierarchy().SetParent(previewRuleEntity, entity)
		ruleService.RenderAny().SetAny(previewRuleEntity, renderComp)

		rulesPass = rulesPass && ruleService.AppliesToAny(dataIn, appliedTo) == nil
	}

	// for [feature.InteractionsComponent.PreviewInteraction] use:
	// [feature.FieldRenderComponent] with ([feature.PreviewRulesPass] or [feature.PreviewRulesFail])
	var comp any
	if interactionsComp.PreviewInteraction == nil {
		comp = previewFieldService.NewAnyRenderComponent(
			interactionsComp.PreviewInteraction, feature.MissingField)
	} else if rulesPass {
		comp = previewFieldService.NewAnyRenderComponent(
			interactionsComp.PreviewInteraction, feature.PreviewRulesPass)
	} else if !rulesPass {
		comp = previewFieldService.NewAnyRenderComponent(
			interactionsComp.PreviewInteraction, feature.PreviewRulesFail)
	}
	previewInteractionEntity := s.World.World().NewEntity()
	s.World.Hierarchy().SetParent(previewInteractionEntity, entity)
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
			ruleService := s.Service().GetRuleService(rule.RuleType)
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
	s.Interactions()
	events.EmitAny(s.World.Events(), featureValue.Interface())
}

func (s *interactionsService) InteractionsEntity() ecs.EntityID {
	entities := slices.Clone(s.interactions.GetEntities())
	if len(entities) == 0 {
		entity := s.World.World().NewEntity()
		comp := feature.NewInteractions(
			nil, nil, nil, false)
		s.interactions.Set(entity, comp)
		return entity
	}
	for _, entity := range entities[1:] {
		s.interactions.Remove(entity)
	}
	return entities[0]
}
func (s *interactionsService) Interactions() feature.InteractionsComponent {
	entity := s.InteractionsEntity()
	comp, _ := s.interactions.Get(entity)
	return comp
}
func (s *interactionsService) OnInteractionsMod(listener func(ecs.EntityID)) {
	s.interactions.OnUpsert(listener)
}

func (s *interactionsService) PreviewInteraction(field any) error {
	entity := s.InteractionsEntity()
	comp, _ := s.interactions.Get(entity)
	comp.PreviewInteraction = field
	s.ProcessState(entity, comp)
	return nil
}

func (s *interactionsService) SubmitPreviewInteraction(key feature.FieldKey) {
	entity := s.InteractionsEntity()
	comp, _ := s.interactions.Get(entity)
	if reflect.TypeOf(comp.PreviewInteraction) != key {
		return
	}
	comp.Interactions = append(comp.Interactions, comp.PreviewInteraction)
	comp.PreviewInteraction = nil
	s.ProcessState(entity, comp)
}
