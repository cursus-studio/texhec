package internal

import (
	"engine/modules/datastructures/internal/types"
	"sort"
	"sync"
)

type trackingArray[Stored comparable] struct {
	Data       []Stored
	ChangesMap map[int]*Stored
}

func NewTrackingArray[Stored comparable]() types.TrackingArray[Stored] {
	return &trackingArray[Stored]{
		ChangesMap: map[int]*Stored{},
	}
}

func (s *trackingArray[Stored]) Get() []Stored { return s.Data }

func (s *trackingArray[Stored]) Add(elements ...Stored) {
	for i := range len(elements) {
		s.ChangesMap[i+len(s.Data)] = nil
	}
	s.Data = append(s.Data, elements...)
}

func (s *trackingArray[Stored]) Set(index int, e Stored) {
	if diffSize := index - len(s.Data) + 1; diffSize > 0 {
		diff := make([]Stored, diffSize)
		s.Add(diff...)
	}
	original := s.Data[index]
	if s.Data[index] == e {
		return
	}
	if _, ok := s.ChangesMap[index]; !ok {
		s.ChangesMap[index] = &original
	}
	s.Data[index] = e
}

func (s *trackingArray[Stored]) Remove(indices ...int) {
	sort.Slice(indices, func(i, j int) bool { return indices[i] > indices[j] })
	for _, index := range indices {
		if index >= len(s.Data) {
			continue
		}
		if _, ok := s.ChangesMap[index]; !ok {
			e := s.Data[index]
			s.ChangesMap[index] = &e
		}
		if _, ok := s.ChangesMap[len(s.Data)-1]; !ok {
			e := s.Data[len(s.Data)-1]
			s.ChangesMap[len(s.Data)-1] = &e
		}

		s.Data[index] = s.Data[len(s.Data)-1]
		s.Data = s.Data[:len(s.Data)-1]
	}
}

func (s *trackingArray[Stored]) Changes() []types.Change[Stored] {
	changes := make([]types.Change[Stored], 0, len(s.ChangesMap))
	for index, from := range s.ChangesMap {
		changes = append(changes, types.NewChange(index, from))
	}
	sort.Slice(changes, func(i, j int) bool { return changes[i].Index > changes[j].Index })
	return changes
}

func (s *trackingArray[Stored]) ClearChanges() { s.ChangesMap = map[int]*Stored{} }

//

type threadSafeArr[Stored comparable] struct {
	mutex sync.Locker
	types.TrackingArray[Stored]
}

func NewThreadSafeTrackingArray[Stored comparable](mutex sync.Locker) types.TrackingArray[Stored] {
	return &threadSafeArr[Stored]{
		mutex:         mutex,
		TrackingArray: NewTrackingArray[Stored](),
	}
}

func (arr *threadSafeArr[Stored]) Add(elements ...Stored) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()
	arr.TrackingArray.Add(elements...)
}
func (arr *threadSafeArr[Stored]) Set(index int, e Stored) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()
	arr.TrackingArray.Set(index, e)
}
func (arr *threadSafeArr[Stored]) Remove(indices ...int) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()
	arr.TrackingArray.Remove(indices...)
}

func (arr *threadSafeArr[Stored]) Changes() []types.Change[Stored] {
	arr.mutex.Lock()
	return arr.TrackingArray.Changes()
}

func (arr *threadSafeArr[Stored]) ClearChanges() {
	arr.TrackingArray.ClearChanges()
	arr.mutex.Unlock()
}
