package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/interactions"
	"reflect"

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
type FeatureService[Feature interactions.Feature] interface {
	AnyFeatureService
}

//

type featureService[Feature interactions.Feature] struct {
	C                  ioc.Dic
	engine.EngineWorld `inject:""`
	Interactions       Service `inject:""`
	rawRelations       []RawRelation
	relationByTgtField [][]Relation
	steps              []AnyStepService
}

func NewFeatureService[Feature interactions.Feature](c ioc.Dic, relations []RawRelation) FeatureService[Feature] {
	s := ioc.GetServices[*featureService[Feature]](c)
	s.C = c

	s.rawRelations = relations

	return s
}

func (s *featureService[Feature]) Key() interactions.FeatureKey { return reflect.TypeFor[Feature]() }
func (s *featureService[Feature]) Steps() []AnyStepService {
	return s.steps
}

func (s *featureService[Feature]) Init() {
	event := reflect.TypeFor[Feature]()
	fieldsCount := event.NumField()
	s.steps = make([]AnyStepService, 0, fieldsCount)
	for i := range fieldsCount {
		fieldType := event.Field(i).Type
		step, ok := s.Interactions.StepByKey(fieldType)
		if !ok {
			panic("feature cannot use not registered step")
		}

		s.steps = append(s.steps, step)
	}

	s.relationByTgtField = make([][]Relation, len(s.steps))
	for _, rawRelation := range s.rawRelations {
		relation, ok := NewRelation[Feature](s.C, rawRelation)
		if !ok {
			panic("compiled code passed invalid relation")
		}
		relations := s.relationByTgtField[relation.Tgt]
		relations = append(relations, relation)
		s.relationByTgtField[relation.Tgt] = relations
	}
}

func (s *featureService[Feature]) Progress() {
	featureEntity := s.Interactions.FeatureEntity()
	interactionEntities := s.Hierarchy().Children(featureEntity).GetIndices()
	for i, step := range s.steps {
		if len(interactionEntities) <= i {
			interactionEntity := s.World().NewEntity()
			propertiesEntity := s.World().NewEntity()
			s.Hierarchy().SetParent(interactionEntity, featureEntity)
			relations := s.relationByTgtField[i]
			for _, relation := range relations {
				relation.Set(interactionEntities[relation.Src], propertiesEntity)
			}
			step.Interaction().MarkMissing(propertiesEntity, interactionEntity)
			return
		}
		interactionEntity := interactionEntities[i]
		err := step.EntityRule(interactionEntity)
		if err == nil {
			continue
		}
		if err != ErrInteractionIsMissing {
			propertiesEntity := s.World().NewEntity()
			relations := s.relationByTgtField[i]
			for _, relation := range relations {
				relation.Set(interactionEntities[relation.Src], propertiesEntity)
			}
			step.Interaction().MarkMissing(propertiesEntity, interactionEntity)
			s.Logger().Warn(err)
		}
		return
	}
	var feature Feature
	value := reflect.ValueOf(&feature).Elem()
	for i, step := range s.steps {
		step.FillValue(interactionEntities[i], value.Field(i))
	}
	s.Interactions.ResetFeatureEntity()
	events.EmitAny(s.Events(), feature.Event())
}
