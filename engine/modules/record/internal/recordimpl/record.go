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
	dependencies     datastructures.Set[*BackwardRecording]
	uuidDependencies datastructures.Set[*UUIDBackwardRecording]

	ecs.AnyComponentArray
}

type service struct {
	engine.EngineWorld `inject:""`
	worldArrays        map[string]entityArray

	worldCopy       ecs.World
	worldCopyUUID   ecs.ComponentArray[uuid.Component]
	worldCopyArrays map[string]ecs.AnyComponentArray

	mutex  *sync.Mutex
	entity *entityKeyedRecorder
	uuid   *uuidKeyedRecorder
}

func NewService(c ioc.Dic) record.Service {
	s := ioc.GetServices[*service](c)
	s.worldArrays = make(map[string]entityArray)

	s.worldCopy = ecs.NewWorld()
	s.worldCopyUUID = ecs.GetComponentArray[uuid.Component](s.worldCopy)
	s.worldCopyArrays = make(map[string]ecs.AnyComponentArray)

	s.mutex = &sync.Mutex{}

	s.entity = newEntityKeyedRecorder(s)
	s.uuid = newUUIDKeyedRecorder(s)

	return s
}

//

func (s *service) Entity() record.EntityKeyedRecorder {
	return s.entity
}
func (s *service) UUID() record.UUIDKeyedRecorder {
	return s.uuid
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
	for _, recording := range s.entity.backwardsRecordings.GetValues() {
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

	// apply in UUID arrays
	for _, recording := range s.uuid.backwardsRecordings.GetValues() {
		for _, entity := range entities {
			uuid, ok := s.worldCopyUUID.Get(entity)
			if !ok {
				uuid, ok = s.EngineWorld.UUID().Component().Get(entity)
				if !ok {
					uuid.ID = s.EngineWorld.UUID().NewUUID()
					s.EngineWorld.UUID().Component().Set(entity, uuid)
				}
				s.worldCopyUUID.Set(entity, uuid)
			}
			if _, ok := recording.Entities[uuid.ID]; ok {
				continue
			}
			var components []any
			if !s.worldCopy.EntityExists(entity) {
				goto saveUUID
			}
			components = make([]any, 0, len(recording.WorldCopyArrays))
			for _, array := range recording.WorldCopyArrays {
				component, ok := array.GetAny(entity)
				if !ok {
					component = nil
				}
				components = append(components, component)
			}
		saveUUID:
			recording.EntitiesOrder = append(recording.EntitiesOrder, uuid.ID)
			recording.Entities[uuid.ID] = components
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
	arrayCtor := config.RecordedComponents[arrayType]
	inheritCtor := config.InheritZero[arrayType]
	entityArray := entityArray{
		dirtySet:          ecs.NewDirtySet(),
		dependencies:      datastructures.NewSet[*BackwardRecording](),
		uuidDependencies:  datastructures.NewSet[*UUIDBackwardRecording](),
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

	return entityArray
}

func (s *service) GetWorldCopyArray(arrayType reflect.Type, config record.Config) ecs.AnyComponentArray {
	arrayKey := arrayType.String()
	if array, ok := s.worldCopyArrays[arrayKey]; ok {
		return array
	}
	arrayCtor := config.RecordedComponents[arrayType]
	inheritCtor := config.InheritZero[arrayType]
	entityArray := entityArray{
		dirtySet:          ecs.NewDirtySet(),
		dependencies:      datastructures.NewSet[*BackwardRecording](),
		uuidDependencies:  datastructures.NewSet[*UUIDBackwardRecording](),
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

	return array
}
