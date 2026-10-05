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
		if event.Feature == nil {
			comp = feature.NewInteractions(nil, nil, nil, false)
		} else {
			comp.Features = []feature.FeatureKey{event.Feature}
			comp.LockedFeature = true
		}
		s.ProcessState(entity, comp)
	})
	events.Listen(s.World.EventsBuilder(), func(loop.TickEvent) {
		entity := s.InteractionsEntity()
		comp, _ := s.interactions.Get(entity)
		s.ProcessState(entity, comp)
	})
	return s
}

func (s *interactionsService) InteractionsEntity() ecs.EntityID {
	entities := slices.Clone(s.interactions.GetEntities())
	if len(entities) == 0 {
		entity := s.World.World().NewEntity()
		comp := feature.NewInteractions(nil, nil, nil, false)
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

func (s *interactionsService) canUseFieldKey(fieldKey reflect.Type) bool {
	comp := s.Interactions()
	if !comp.LockedFeature {
		return true
	}
	feat := comp.Features[0]
	featConfig := s.Config.featureConfigs[feat]
	var wantedFieldKey reflect.Type
	userInteractionIndex := 0

	for _, fieldConfig := range featConfig.Fields {
		if fieldConfig.Provide != nil {
			continue
		}
		if userInteractionIndex == len(comp.Interactions) {
			wantedFieldKey = fieldConfig.Key
			break
		}
		userInteractionIndex++
	}
	canUse := wantedFieldKey != nil && wantedFieldKey == fieldKey
	return canUse
}

func (s *interactionsService) SubmitInteraction(field any) {
	if !s.canUseFieldKey(reflect.TypeOf(field)) {
		return
	}
	entity := s.InteractionsEntity()
	comp, _ := s.interactions.Get(entity)
	comp.Interactions = append(comp.Interactions, field)
	comp.PreviewInteraction = nil
	s.ProcessState(entity, comp)
}
func (s *interactionsService) PreviewInteraction(field any) {
	if !s.canUseFieldKey(reflect.TypeOf(field)) {
		return
	}
	entity := s.InteractionsEntity()
	comp, _ := s.interactions.Get(entity)
	comp.PreviewInteraction = field
	s.ProcessState(entity, comp)
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
