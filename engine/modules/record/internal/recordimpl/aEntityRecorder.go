package recordimpl

import (
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/record"
)

type FowardRecording struct {
	Config   record.Config
	DirtySet ecs.DirtySet
}

type BackwardRecording struct {
	Config          record.Config
	WorldCopyArrays []ecs.AnyComponentArray
	Entities        datastructures.SparseArray[ecs.EntityID, []any]
}

type entityKeyedRecorder struct {
	*service

	i     record.RecordingID
	holes datastructures.SparseSet[record.RecordingID]

	forwardRecordings   datastructures.SparseArray[record.RecordingID, *FowardRecording]
	backwardsRecordings datastructures.SparseArray[record.RecordingID, *BackwardRecording]
}

func newEntityKeyedRecorder(
	t *service,
) *entityKeyedRecorder {
	entityKeyedRecorder := &entityKeyedRecorder{
		t,

		1,
		datastructures.NewSparseSet[record.RecordingID](),

		datastructures.NewSparseArray[record.RecordingID, *FowardRecording](),
		datastructures.NewSparseArray[record.RecordingID, *BackwardRecording](),
	}

	return entityKeyedRecorder
}

func (s *entityKeyedRecorder) GetState(config record.Config) record.Recording {
	entities := datastructures.NewSparseSet[ecs.EntityID]()
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		for _, entity := range array.GetEntities() {
			entities.Add(entity)
		}
	}
	return s.getStateFor(config, entities.GetIndices())
}
func (s *entityKeyedRecorder) StartBackwardsRecording(config record.Config) record.RecordingID {
	s.WarmUp().WarmUp()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.SyncBackwardsRecordingState()

	id := s.getID()
	recording := &BackwardRecording{
		Config:          config,
		WorldCopyArrays: make([]ecs.AnyComponentArray, 0, len(*config.ComponentsOrder)),
		Entities:        datastructures.NewSparseArray[ecs.EntityID, []any](),
	}
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		array.dependencies.Add(recording)
		worldCopyArray := s.GetWorldCopyArray(arrayType, config)
		recording.WorldCopyArrays = append(recording.WorldCopyArrays, worldCopyArray)
	}
	s.backwardsRecordings.Set(id, recording)

	return id
}
func (s *entityKeyedRecorder) StartRecording(config record.Config) record.RecordingID {
	s.WarmUp().WarmUp()
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := s.getID()
	recording := &FowardRecording{
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
func (s *entityKeyedRecorder) Stop(id record.RecordingID) (record.Recording, bool) {
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
			array.dependencies.RemoveElements(recording)
		}

		s.backwardsRecordings.Remove(id)
		s.holes.Add(id)

		return record.Recording{Entities: recording.Entities}, true
	}
	return record.Recording{}, false
}
func (s *entityKeyedRecorder) Apply(config record.Config, recordings ...record.Recording) {
	arrays := make([]ecs.AnyComponentArray, 0, len(*config.ComponentsOrder))
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		arrays = append(arrays, array)
	}

	for _, recording := range recordings {
		for _, entity := range recording.Entities.GetIndices() {
			components, ok := recording.Entities.Get(entity)
			if !ok {
				continue
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

func (s *entityKeyedRecorder) getStateFor(config record.Config, entities []ecs.EntityID) record.Recording {
	arrays := make([]ecs.AnyComponentArray, 0, len(*config.ComponentsOrder))
	for _, arrayType := range *config.ComponentsOrder {
		array := s.GetWorldArray(arrayType, config)
		arrays = append(arrays, array)
	}

	recording := record.Recording{
		Entities: datastructures.NewSparseArray[ecs.EntityID, []any](),
	}

	for _, entity := range entities {
		components := make([]any, 0, len(arrays))
		for _, array := range arrays {
			v, ok := array.GetAny(entity)
			if !ok {
				v = nil
			}
			components = append(components, v)
		}
		recording.Entities.Set(entity, components)
	}

	return recording
}

//

func (s *entityKeyedRecorder) getID() record.RecordingID {
	if holes := s.holes.GetIndices(); len(holes) != 0 {
		hole := holes[0]
		s.holes.Remove(hole)
		return hole
	}
	i := s.i
	s.i++
	return i
}
