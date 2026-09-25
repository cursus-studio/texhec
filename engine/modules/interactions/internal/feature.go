package internal

import (
	"engine"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/interactions"
	"reflect"
	"slices"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

func OffsetFieldIndex[Type any](fieldOffset uintptr) int {
	structType := reflect.TypeFor[Type]()
	for i := range structType.NumField() {
		field := structType.Field(i)
		if field.Offset == fieldOffset {
			return i
		}
	}
	return -1
}

//

type RawRelation struct {
	// refers to fields OffsetOf
	Src, Tgt uintptr
	Set      func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID)
}

func NewRawRelation(
	src, tgt uintptr,
	set func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID),
) RawRelation {
	return RawRelation{src, tgt, set}
}

//

type Relation struct {
	// src and tgt field index
	Src, Tgt int
	Set      func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID)
}

func NewRelation[Type any](c ioc.Dic, raw RawRelation) (Relation, bool) {
	relation := Relation{
		Src: OffsetFieldIndex[Type](raw.Src), Tgt: OffsetFieldIndex[Type](raw.Tgt),
		Set: raw.Set(c)}
	return relation, relation.Src < relation.Tgt && relation.Src != -1 && relation.Tgt != -1
}

//

type AnyFeatureService interface {
	Key() interactions.FeatureKey
	Steps() []AnyStepService

	Init()

	// sets step interaction to missing or emits event
	Progress()
}
type FeatureService[Feature any] interface {
	AnyFeatureService
}

//

type featureService[Feature any] struct {
	C                  ioc.Dic
	engine.EngineWorld `inject:""`
	Interactions       Service                    `inject:""`
	ContextSetter      interactions.ContextSetter `inject:""`

	rawRelations       []RawRelation
	relationByTgtField datastructures.SparseArray[int, []Relation]
	steps              datastructures.SparseArray[int, AnyStepService]
}

func NewFeatureService[Feature any](c ioc.Dic, relations []RawRelation) FeatureService[Feature] {
	s := ioc.GetServices[*featureService[Feature]](c)
	s.C = c

	s.rawRelations = relations
	s.relationByTgtField = datastructures.NewSparseArray[int, []Relation]()
	s.steps = datastructures.NewSparseArray[int, AnyStepService]()

	return s
}

func (s *featureService[Feature]) Key() interactions.FeatureKey { return reflect.TypeFor[Feature]() }
func (s *featureService[Feature]) Steps() []AnyStepService {
	return s.steps.GetValues()
}

func (s *featureService[Feature]) Init() {
	event := reflect.TypeFor[Feature]()
	fieldsCount := event.NumField()
	for i := range fieldsCount {
		fieldType := event.Field(i).Type
		step, ok := s.Interactions.StepByKey(fieldType)
		if !ok {
			continue
		}

		s.steps.Set(i, step)
	}

	for _, rawRelation := range s.rawRelations {
		relation, ok := NewRelation[Feature](s.C, rawRelation)
		if !ok {
			panic("compiled code passed invalid relation")
		}
		relations, _ := s.relationByTgtField.Get(relation.Tgt)
		relations = append(relations, relation)
		s.relationByTgtField.Set(relation.Tgt, relations)
	}
}

func (s *featureService[Feature]) Progress() {
	featureEntity := s.Interactions.FeatureEntity()
	interactionEntities := slices.Clone(s.Hierarchy().Children(featureEntity).GetIndices())
	interactions := datastructures.NewSparseArray[int, ecs.EntityID]()
	for _, i := range s.steps.GetIndices() {
		step, ok := s.steps.Get(i)
		if !ok {
			continue
		}
		if len(interactionEntities) == 0 {
			interactionEntity := s.World().NewEntity()
			propertiesEntity := s.World().NewEntity()
			s.Hierarchy().SetParent(interactionEntity, featureEntity)
			relations, _ := s.relationByTgtField.Get(i)
			for _, relation := range relations {
				interactionEntity, _ := interactions.Get(relation.Src)
				relation.Set(interactionEntity, propertiesEntity)
			}
			step.Interaction().MarkMissing(propertiesEntity, interactionEntity)
			return
		}
		interactions.Set(i, interactionEntities[0])
		interactionEntities = interactionEntities[1:]
		interactionEntity, _ := interactions.Get(i)
		err := step.EntityRule(interactionEntity)
		if err == nil {
			continue
		}
		if err != ErrInteractionIsMissing {
			propertiesEntity := s.World().NewEntity()
			relations, _ := s.relationByTgtField.Get(i)
			for _, relation := range relations {
				interactionEntity, _ := interactions.Get(relation.Src)
				relation.Set(interactionEntity, propertiesEntity)
			}
			step.Interaction().MarkMissing(propertiesEntity, interactionEntity)
			s.Logger().Warn(err)
		}
		return
	}
	var feature Feature
	value := reflect.ValueOf(&feature).Elem()
	for _, i := range s.steps.GetIndices() {
		step, _ := s.steps.Get(i)
		interactionEntity, _ := interactions.Get(i)
		step.FillValue(interactionEntity, value.Field(i))
	}
	s.Interactions.ResetFeatureEntity()
	events.EmitAny(s.Events(), s.ContextSetter.SetContext(feature))
}
