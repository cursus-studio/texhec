package test

import "testing"

func BenchmarkSpatialIndexingGetEmpty(b *testing.B) {
	setup := NewSetup()
	component := Component{Index: 69}

	b.StartTimer()
	for b.Loop() {
		setup.Service.Get(component.Index)
	}
}

func BenchmarkSpatialIndexingGet(b *testing.B) {
	setup := NewSetup()
	component := Component{Index: 69}
	entity := setup.W.NewEntity()
	setup.Array.Set(entity, component)

	b.StartTimer()
	for b.Loop() {
		setup.Service.Get(component.Index)
	}
}

func BenchmarkSpatialIndexingSave(b *testing.B) {
	setup := NewSetup()
	component := Component{Index: 69}
	entity := setup.W.NewEntity()

	b.StartTimer()
	for b.Loop() {
		setup.Array.Remove(entity)
		setup.Array.Set(entity, component)
	}
	setup.Service.Get(0)
}
