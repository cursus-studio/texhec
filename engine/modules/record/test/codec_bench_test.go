package test

import (
	"testing"
)

func BenchmarkEntityCodecEncode(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	originalRecording := s.Record.Entity().GetState(s.Config)

	for b.Loop() {
		_, _ = s.Codec.Encode(originalRecording)
	}
}
func BenchmarkEntityCodecDecode(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	originalRecording := s.Record.Entity().GetState(s.Config)

	encodedRecording, err := s.Codec.Encode(originalRecording)
	if err != nil {
		b.Error(err)
		return
	}

	for b.Loop() {
		_, _ = s.Codec.Decode(encodedRecording)
	}
}
func BenchmarkUUIDCodecEncode(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	originalRecording := s.Record.UUID().GetState(s.Config)

	for b.Loop() {
		_, _ = s.Codec.Encode(originalRecording)
	}
}
func BenchmarkUUIDCodecDecode(b *testing.B) {
	s := NewSetup()

	entity := s.World.NewEntity()
	s.ComponentArray.Set(entity, Component{Counter: 6})

	originalRecording := s.Record.UUID().GetState(s.Config)

	encodedRecording, err := s.Codec.Encode(originalRecording)
	if err != nil {
		b.Error(err)
		return
	}

	for b.Loop() {
		_, _ = s.Codec.Decode(encodedRecording)
	}
}
