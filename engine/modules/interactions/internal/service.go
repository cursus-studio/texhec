package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/interactions"
	"reflect"
	"slices"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

// marks entity storing all interactions instance
type FeatureEntityComponent struct{}
type InteractionKeyComponent struct{ Key interactions.InteractionKey }
type PreviewComponent struct{ PreviewEntity ecs.EntityID }

func NewInteractionKey[State any]() InteractionKeyComponent {
	return InteractionKeyComponent{reflect.TypeFor[State]()}
}
func NewPreview(previewEntity ecs.EntityID) PreviewComponent { return PreviewComponent{previewEntity} }

//

type Service interface {
	interactions.Service
	InteractionKey() ecs.ComponentArray[InteractionKeyComponent]
	Preview() ecs.ComponentArray[PreviewComponent]

	RegisterFeature(AnyFeatureService)
	RegisterStep(AnyStepService)
	RegisterInteraction(AnyInteractionService)
	StepByKey(interactions.StepKey) (AnyStepService, bool)

	Init()

	// if only one feature is available it progresses it
	Progress()
	OnChangeProgress(ecs.EntityID)
	FeatureEntity() ecs.EntityID
	ResetFeatureEntity()
}

type service struct {
	engine.EngineWorld `inject:""`
	availableFeatures  ecs.ComponentArray[interactions.AvailableFeaturesComponent]
	featureEntity      ecs.ComponentArray[FeatureEntityComponent]
	interactionKey     ecs.ComponentArray[InteractionKeyComponent]
	preview            ecs.ComponentArray[PreviewComponent]

	features                  []AnyFeatureService
	featuresKeys              []interactions.FeatureKey
	featuresByKey             map[interactions.FeatureKey]AnyFeatureService
	featureKeysByFirstStepKey map[interactions.StepKey][]interactions.FeatureKey

	steps                 []AnyStepService
	stepsByKey            map[interactions.StepKey]AnyStepService
	stepsByInteractionKey map[interactions.InteractionKey][]AnyStepService

	interactions []AnyInteractionService

	loaded bool
}

func NewService(c ioc.Dic) Service {
	s := ioc.GetServices[*service](c)
	s.availableFeatures = ecs.GetComponentArray[interactions.AvailableFeaturesComponent](s.World())
	s.featureEntity = ecs.GetComponentArray[FeatureEntityComponent](s.World())
	s.interactionKey = ecs.GetComponentArray[InteractionKeyComponent](s.World())
	s.preview = ecs.GetComponentArray[PreviewComponent](s.World())

	s.featuresByKey = make(map[interactions.FeatureKey]AnyFeatureService)
	s.featureKeysByFirstStepKey = make(map[interactions.StepKey][]interactions.FeatureKey)

	s.stepsByKey = make(map[interactions.StepKey]AnyStepService)
	s.stepsByInteractionKey = make(map[interactions.InteractionKey][]AnyStepService)

	s.interactionKey.OnMod(s.OnChangeProgress)
	events.Listen(s.EventsBuilder(), s.onSelectFeat)

	return s
}

func (s *service) InteractionKey() ecs.ComponentArray[InteractionKeyComponent] {
	return s.interactionKey
}
func (s *service) Preview() ecs.ComponentArray[PreviewComponent] { return s.preview }

func (s *service) RegisterFeature(feat AnyFeatureService) {
	s.features = append(s.features, feat)
	s.featuresKeys = append(s.featuresKeys, feat.Key())
	s.featuresByKey[feat.Key()] = feat
}
func (s *service) RegisterStep(step AnyStepService) {
	s.steps = append(s.steps, step)
	s.stepsByKey[step.Key()] = step

	interactionKey := step.Interaction().Key()
	steps := s.stepsByInteractionKey[interactionKey]
	steps = append(steps, step)
	s.stepsByInteractionKey[interactionKey] = steps
}
func (s *service) RegisterInteraction(int AnyInteractionService) {
	s.interactions = append(s.interactions, int)
}

func (s *service) Init() {
	if s.loaded {
		return
	}
	s.loaded = true
	for _, feat := range s.features {
		feat.Init()
		if len(feat.Steps()) == 0 {
			continue
		}
		stepKey := feat.Steps()[0].Key()
		feats := s.featureKeysByFirstStepKey[stepKey]
		feats = append(feats, feat.Key())
		s.featureKeysByFirstStepKey[stepKey] = feats
	}
}

func (s *service) StepByKey(key interactions.StepKey) (AnyStepService, bool) {
	step, ok := s.stepsByKey[key]
	return step, ok
}

func (s *service) Progress() {
	if features, _ := s.AvailableFeatures().Get(s.FeatureEntity()); features.Selected {
		featKey := features.Features[0]
		feat := s.featuresByKey[featKey]
		feat.Progress()
		return
	}

	featureEntity := s.FeatureEntity()
	children := s.Hierarchy().Children(featureEntity).GetIndices()
	if len(children) == 0 {
		s.ResetFeatureEntity()
		return
	}

	child := children[0]
	interactionKey, ok := s.interactionKey.Get(child)
	if !ok {
		s.ResetFeatureEntity()
		return
	}
	var availableFeatures interactions.AvailableFeaturesComponent

	steps := s.stepsByInteractionKey[interactionKey.Key]
	for _, step := range steps {
		if err := step.EntityRule(child); err != nil {
			continue
		}

		feats := s.featureKeysByFirstStepKey[step.Key()]
		availableFeatures.Features = append(availableFeatures.Features, feats...)
	}

	if len(availableFeatures.Features) == 0 {
		s.ResetFeatureEntity()
		return
	}
	s.AvailableFeatures().Set(featureEntity, availableFeatures)
}

func (s *service) FeatureEntity() ecs.EntityID {
	entities := s.featureEntity.GetEntities()
	if len(entities) == 0 {
		entity := s.World().NewEntity()
		entities = []ecs.EntityID{entity}
		s.featureEntity.Set(entity, FeatureEntityComponent{})
		s.AvailableFeatures().Set(entity, interactions.NewAvailableFeatures())
	}

	for _, entity := range entities[1:] {
		s.World().RemoveEntity(entity)
	}
	return entities[0]
}

func (s *service) ResetFeatureEntity() {
	entity := s.FeatureEntity()
	s.AvailableFeatures().Set(entity, interactions.NewAvailableFeatures())
	for _, interactionEntity := range s.Hierarchy().Children(entity).GetIndices() {
		if previewComp, ok := s.Preview().Get(interactionEntity); ok {
			s.World().RemoveEntity(previewComp.PreviewEntity)
		}
		s.World().RemoveEntity(interactionEntity)
	}
}

func (s *service) Features() []interactions.FeatureKey { return slices.Clone(s.featuresKeys) }
func (s *service) AvailableFeatures() ecs.ComponentArray[interactions.AvailableFeaturesComponent] {
	return s.availableFeatures
}

func (s *service) OnChangeProgress(ecs.EntityID) { s.Progress() }
func (s *service) onSelectFeat(event interactions.SelectFeatureEvent) {
	feat, ok := s.featuresByKey[event.FeatureKey]
	if !ok {
		s.ResetFeatureEntity()
		return
	}
	entity := s.FeatureEntity()
	s.AvailableFeatures().Set(entity, interactions.NewSelectedFeature(event.FeatureKey))
	feat.Progress()
}
