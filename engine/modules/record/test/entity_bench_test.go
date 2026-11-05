package test

import (
	"engine/modules/ecs"
	"testing"
)

func BenchmarkEntityRecording(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	b.ResetTimer()
	for b.Loop() {
		recordingID := s.Record.StartRecording(s.Config)
		s.Record.Stop(recordingID)
	}
}

func BenchmarkCreateNEntitiesEntityRecording(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	recordingID := s.Record.StartRecording(s.Config)
	for i := range b.N {
		s.ComponentArray.Set(ecs.EntityID(i), Component{Counter: i})
	}
	b.ResetTimer()
	s.Record.Stop(recordingID)
}

func BenchmarkEntityApply1Entities(b *testing.B) {
	s := NewSetup()

	recordingID := s.Record.StartRecording(s.Config)

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	recording, _ := s.Record.Stop(recordingID)
	b.ResetTimer()
	for b.Loop() {
		s.Record.Apply(s.Config, recording)
	}
}
func BenchmarkEntityApply10Entities(b *testing.B) {
	s := NewSetup()

	recordingID := s.Record.StartRecording(s.Config)

	for range 10 {
		entity := s.World.NewEntity()
		s.ComponentArray.Set(entity, Component{Counter: 6})
	}

	recording, _ := s.Record.Stop(recordingID)
	b.ResetTimer()
	for b.Loop() {
		s.Record.Apply(s.Config, recording)
	}
}
