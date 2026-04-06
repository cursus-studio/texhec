package test

import "testing"

func BenchmarkClone1(b *testing.B) {
	s := NewSetup()

	cloned := s.World.NewEntity()
	cloned1Comp := Cloned1Component{1}
	s.Cloned1.Set(cloned, cloned1Comp)

	b.ResetTimer()
	for b.Loop() {
		s.Prototype.Clone(cloned)
	}
}

func BenchmarkClone2(b *testing.B) {
	s := NewSetup()

	cloned := s.World.NewEntity()
	cloned1Comp := Cloned1Component{1}
	cloned2Comp := Cloned2Component{1}
	s.Cloned1.Set(cloned, cloned1Comp)
	s.Cloned2.Set(cloned, cloned2Comp)

	b.ResetTimer()
	for b.Loop() {
		s.Prototype.Clone(cloned)
	}
}

func BenchmarkManual1Clone(b *testing.B) {
	s := NewSetup()

	cloned := s.World.NewEntity()
	cloned1Comp := Cloned1Component{1}
	s.Cloned1.Set(cloned, cloned1Comp)

	b.ResetTimer()
	for b.Loop() {
		clone := s.World.NewEntity()
		if clonedComp, ok := s.Cloned1.Get(clone); ok {
			s.Cloned1.Set(clone, clonedComp)
		}
	}
}

func BenchmarkManual2Clone(b *testing.B) {
	s := NewSetup()

	cloned := s.World.NewEntity()
	cloned1Comp := Cloned1Component{1}
	cloned2Comp := Cloned2Component{1}
	s.Cloned1.Set(cloned, cloned1Comp)
	s.Cloned2.Set(cloned, cloned2Comp)

	b.ResetTimer()
	for b.Loop() {
		clone := s.World.NewEntity()
		if clonedComp, ok := s.Cloned1.Get(clone); ok {
			s.Cloned1.Set(clone, clonedComp)
		}
		if clonedComp, ok := s.Cloned2.Get(clone); ok {
			s.Cloned2.Set(clone, clonedComp)
		}
	}
}
