package recordimpl

import (
	"engine"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/record"
	"engine/modules/uuid"
	"reflect"
	"sync"

	"github.com/ogiusek/ioc/v2"
)

type entityArray struct {
	dirtySet ecs.DirtySet

	// backwards dependencies
	dependencies datastructures.Set[*BackwardRecording]

	ecs.AnyComponentArray
}

type service struct {
	engine.EngineWorld `inject:""`
	worldArrays        map[string]entityArray

	worldCopy       ecs.World
	worldCopyUUID   ecs.ComponentArray[uuid.Component]
	worldCopyArrays map[string]ecs.AnyComponentArray

	mutex *sync.Mutex
	*entityKeyedRecorder
}

func NewService(c ioc.Dic) record.Service {
	s := ioc.GetServices[*service](c)
	s.worldArrays = make(map[string]entityArray)

	s.worldCopy = ecs.NewWorld()
	s.worldCopyUUID = ecs.GetComponentArray[uuid.Component](s.worldCopy)
	s.worldCopyArrays = make(map[string]ecs.AnyComponentArray)

	s.mutex = &sync.Mutex{}

	s.entityKeyedRecorder = newEntityKeyedRecorder(s)

	return s
}

func (s *service) UUIDState(recording record.Recording) record.UUIDState {
	state := record.UUIDState{
		Keys:   make([]uuid.UUID, 0, len(recording.Entities.GetIndices())),
		Values: make([]record.EntityComponents, 0, len(recording.Entities.GetIndices())),
	}

	for i, entity := range recording.Entities.GetIndices() {
		// we do not record removed entities
		if len(recording.Entities.GetValues()[i]) == 0 {
			continue
		}
		state.Keys = append(state.Keys, s.UUID().CreateGetUUID(entity))
		state.Values = append(state.Values, recording.Entities.GetValues()[i])
	}

	return state
}
func (s *service) ApplyUUIDState(config record.Config, state record.UUIDState) {
	arrays := make([]ecs.AnyComponentArray, 0, len(*config.ComponentsOrder))
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		array.PrepareBulk()
		arrays = append(arrays, array)
	}

	overflowEntities := datastructures.NewSparseSet[ecs.EntityID]()
	for _, array := range arrays {
		for _, entity := range array.GetEntities() {
			overflowEntities.Add(entity)
		}
	}

	for i, uuid := range state.Keys {
		entity := s.UUID().CreateGetEntity(uuid)
		overflowEntities.Remove(entity)
		components := state.Values[i]
		for i, component := range components {
			array := arrays[i]
			if component == nil {
				array.Remove(entity)
				continue
			}
			array.SetAny(entity, component)
		}
	}

	for _, overflowEntity := range overflowEntities.GetIndices() {
		s.World().RemoveEntity(overflowEntity)
	}

	for _, array := range arrays {
		array.CommitBulk()
	}
}

//

func (s *service) SyncBackwardsRecordingState() {
	for arrayType, array := range s.worldArrays {
		s.synchronizeArrayState(
			array,
			s.worldCopyArrays[arrayType],
		)
	}
}

func (s *service) synchronizeArrayState(
	worldArray entityArray,
	worldCopyArray ecs.AnyComponentArray,
) {
	entities := worldArray.dirtySet.Get()
	if len(entities) == 0 {
		return
	}

	// apply in Entity arrays
	for _, recording := range s.backwardsRecordings.GetValues() {
		for _, entity := range entities {
			if _, ok := recording.Entities.Get(entity); ok {
				continue
			}
			var components []any
			if !s.worldCopy.EntityExists(entity) {
				goto saveEntity
			}
			components = make([]any, 0, len(recording.WorldCopyArrays))
			for _, array := range recording.WorldCopyArrays {
				component, ok := array.GetAny(entity)
				if !ok {
					component = nil
				}
				components = append(components, component)
			}
		saveEntity:
			recording.Entities.Set(entity, components)
		}
	}

	// apply in world
	for _, entity := range entities {
		if component, ok := worldArray.GetAny(entity); ok {
			s.worldCopy.EnsureExists(entity)
			worldCopyArray.SetAny(entity, component)
			continue
		}
		if s.World().EntityExists(entity) {
			worldCopyArray.Remove(entity)
			continue
		}
		s.worldCopy.RemoveEntity(entity)
	}
}

func (s *service) GetWorldArray(arrayType reflect.Type, config record.Config) entityArray {
	arrayKey := arrayType.String()
	if array, ok := s.worldArrays[arrayKey]; ok {
		return array
	}
	worldArray, _ := s.createWorldArrays(arrayType, config)
	return worldArray
}

func (s *service) GetWorldCopyArray(arrayType reflect.Type, config record.Config) ecs.AnyComponentArray {
	arrayKey := arrayType.String()
	if array, ok := s.worldCopyArrays[arrayKey]; ok {
		return array
	}
	_, worldCopyArray := s.createWorldArrays(arrayType, config)
	return worldCopyArray
}

func (s *service) createWorldArrays(arrayType reflect.Type, config record.Config) (entityArray, ecs.AnyComponentArray) {
	arrayKey := arrayType.String()
	arrayCtor := config.RecordedComponents[arrayType]
	inheritCtor := config.InheritZero[arrayType]
	entityArray := entityArray{
		dirtySet:          ecs.NewDirtySet(),
		dependencies:      datastructures.NewSet[*BackwardRecording](),
		AnyComponentArray: arrayCtor(s.World()),
	}
	entityArray.AddDirtySet(entityArray.dirtySet)
	s.worldArrays[arrayKey] = entityArray
	entityArray.dirtySet.Clear()

	inheritCtor(s.World())
	array := arrayCtor(s.worldCopy)
	s.worldCopyArrays[arrayKey] = array

	for _, entity := range entityArray.GetEntities() {
		component, _ := entityArray.GetAny(entity)
		array.SetAny(entity, component)
	}

	return entityArray, array
}
