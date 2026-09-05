package recordimpl

import (
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/record"
	"engine/modules/uuid"
)

type UUIDForwardRecording struct {
	Config   record.Config
	DirtySet ecs.DirtySet
}
type UUIDBackwardRecording struct {
	Config          record.Config
	WorldCopyArrays []ecs.AnyComponentArray
	EntitiesOrder   []uuid.UUID
	Entities        map[uuid.UUID][]any
}

type uuidKeyedRecorder struct {
	*service

	i     record.UUIDRecordingID
	holes datastructures.SparseSet[record.UUIDRecordingID]

	forwardRecordings   datastructures.SparseArray[record.UUIDRecordingID, *UUIDForwardRecording]
	backwardsRecordings datastructures.SparseArray[record.UUIDRecordingID, *UUIDBackwardRecording]
}

func newUUIDKeyedRecorder(
	s *service,
) *uuidKeyedRecorder {
	uuidKeyedRecorder := &uuidKeyedRecorder{
		s,

		1,
		datastructures.NewSparseSet[record.UUIDRecordingID](),

		datastructures.NewSparseArray[record.UUIDRecordingID, *UUIDForwardRecording](),
		datastructures.NewSparseArray[record.UUIDRecordingID, *UUIDBackwardRecording](),
	}

	return uuidKeyedRecorder
}

func (s *uuidKeyedRecorder) GetState(config record.Config) record.UUIDRecording {
	entities := datastructures.NewSparseSet[ecs.EntityID]()
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		for _, entity := range array.GetEntities() {
			entities.Add(entity)
		}
	}
	return s.getStateFor(config, entities.GetIndices())
}

func (s *uuidKeyedRecorder) StartBackwardsRecording(config record.Config) record.UUIDRecordingID {
	s.WarmUp().WarmUp()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.SyncBackwardsRecordingState()

	id := s.getID()
	recording := &UUIDBackwardRecording{
		Config:          config,
		WorldCopyArrays: make([]ecs.AnyComponentArray, 0, len(*config.ComponentsOrder)),
		Entities:        map[uuid.UUID][]any{},
	}
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		array.uuidDependencies.Add(recording)
		worldCopyArray := s.GetWorldCopyArray(arrayType, config)
		recording.WorldCopyArrays = append(recording.WorldCopyArrays, worldCopyArray)
	}
	s.backwardsRecordings.Set(id, recording)

	return id
}
func (s *uuidKeyedRecorder) StartRecording(config record.Config) record.UUIDRecordingID {
	s.WarmUp().WarmUp()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.SyncBackwardsRecordingState()

	id := s.getID()
	recording := &UUIDForwardRecording{
		Config:   config,
		DirtySet: ecs.NewDirtySet(),
	}
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		array.AddDirtySet(recording.DirtySet)
	}
	recording.DirtySet.Clear()
	s.forwardRecordings.Set(id, recording)

	return id
}
func (s *uuidKeyedRecorder) Stop(id record.UUIDRecordingID) (record.UUIDRecording, bool) {
	s.WarmUp().WarmUp()
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if recording, ok := s.forwardRecordings.Get(id); ok {
		entities := recording.DirtySet.Get()
		recording.DirtySet.Release()

		s.forwardRecordings.Remove(id)
		s.holes.Add(id)

		return s.getStateFor(recording.Config, entities), true
	}
	if recording, ok := s.backwardsRecordings.Get(id); ok {
		s.SyncBackwardsRecordingState()
		for _, arrayType := range *recording.Config.ComponentsOrder {
			array := s.GetWorldArray(arrayType, recording.Config)
			array.uuidDependencies.RemoveElements(recording)
		}

		s.backwardsRecordings.Remove(id)
		s.holes.Add(id)

		return record.UUIDRecording{
			EntitiesOrder: recording.EntitiesOrder,
			Entities:      recording.Entities,
		}, true
	}
	return record.UUIDRecording{}, false
}
func (s *uuidKeyedRecorder) Apply(config record.Config, recordings ...record.UUIDRecording) {
	arrays := make([]ecs.AnyComponentArray, 0, len(*config.ComponentsOrder))
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		arrays = append(arrays, array)
	}

	for _, recording := range recordings {
		for _, uuidValue := range recording.EntitiesOrder {
			components := recording.Entities[uuidValue]
			entity, ok := s.EngineWorld.UUID().Entity(uuidValue)
			if !ok && components != nil {
				entity = s.World().NewEntity()
				s.EngineWorld.UUID().Component().Set(entity, uuid.New(uuidValue))
			}
			if components == nil {
				s.World().RemoveEntity(entity)
				continue
			}

			for i, component := range components {
				array := arrays[i]
				if component == nil {
					array.Remove(entity)
					continue
				}
				array.SetAny(entity, component)
			}
		}
	}
}

func (s *uuidKeyedRecorder) getStateFor(config record.Config, entities []ecs.EntityID) record.UUIDRecording {
	arrays := make([]ecs.AnyComponentArray, 0, len(*config.ComponentsOrder))
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		arrays = append(arrays, array)
	}

	recording := record.UUIDRecording{
		Entities: make(map[uuid.UUID][]any, len(entities)),
	}

	for _, entity := range entities {
		uuidComponent, ok := s.EngineWorld.UUID().Component().Get(entity)
		if !ok {
			uuidComponent.ID = s.EngineWorld.UUID().NewUUID()
			s.EngineWorld.UUID().Component().Set(entity, uuidComponent)
		}
		components := make([]any, 0, len(arrays))
		for _, array := range arrays {
			v, ok := array.GetAny(entity)
			if !ok {
				v = nil
			}
			components = append(components, v)
		}
		recording.EntitiesOrder = append(recording.EntitiesOrder, uuidComponent.ID)
		recording.Entities[uuidComponent.ID] = components
	}

	return recording
}

func (s *uuidKeyedRecorder) getID() record.UUIDRecordingID {
	if holes := s.holes.GetIndices(); len(holes) != 0 {
		hole := holes[0]
		s.holes.Remove(hole)
		return hole
	}
	i := s.i
	s.i++
	return i
}
