package test

import (
	"engine/services/ecs"
	"testing"
)

func BenchmarkEntityRecording(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	b.ResetTimer()
	for b.Loop() {
		recordingID := s.Record.Entity().StartRecording(s.Config)
		s.Record.Entity().Stop(recordingID)
	}
}

func BenchmarkCreateNEntitiesEntityRecording(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	recordingID := s.Record.Entity().StartRecording(s.Config)
	for i := range b.N {
		s.ComponentArray.Set(ecs.EntityID(i), Component{Counter: i})
	}
	b.ResetTimer()
	s.Record.Entity().Stop(recordingID)
}

func BenchmarkEntityApply1Entities(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	recording := s.Record.Entity().GetState(s.Config)
	b.ResetTimer()
	for b.Loop() {
		s.Record.Entity().Apply(s.Config, recording)
	}
}
func BenchmarkEntityApply10Entities(b *testing.B) {
	s := NewSetup()

	for range 10 {
		entity := s.World.NewEntity()
		s.ComponentArray.Set(entity, Component{Counter: 6})
	}

	recording := s.Record.Entity().GetState(s.Config)
	b.ResetTimer()
	for b.Loop() {
		s.Record.Entity().Apply(s.Config, recording)
	}
}
