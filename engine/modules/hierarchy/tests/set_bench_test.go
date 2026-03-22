package test

import (
	"engine/services/ecs"
	"testing"
)

func BenchmarkAddNChildrenWithParent(b *testing.B) {
	setup := NewSetup()
	grandParent := setup.World.NewEntity()
	parent := grandParent
	parentCount := 0
	for range parentCount {
		child := setup.World.NewEntity()
		setup.Service.SetParent(child, parent)
		parent = child
	}

	b.ResetTimer()
	for b.Loop() {
		child := setup.World.NewEntity()
		setup.Service.SetParent(child, parent)
	}
}

func BenchmarkAddNChildrenWith5Parents(b *testing.B) {
	setup := NewSetup()
	grandParent := setup.World.NewEntity()
	parent := grandParent
	parentCount := 5
	for range parentCount {
		child := setup.World.NewEntity()
		setup.Service.SetParent(child, parent)
		parent = child
	}

	b.ResetTimer()
	for b.Loop() {
		child := setup.World.NewEntity()
		setup.Service.SetParent(child, parent)
	}
}

func BenchmarkRemoveChild(b *testing.B) {
	setup := NewSetup()
	parent := setup.World.NewEntity()

	children := make([]ecs.EntityID, b.N)
	for i := range b.N {
		children[i] = setup.World.NewEntity()
		setup.Service.SetParent(children[i], parent)
	}

	b.ResetTimer()
	for i := range b.N {
		setup.World.RemoveEntity(children[i])
	}
}

func BenchmarkRemoveParentWith1Children(b *testing.B) {
	setup := NewSetup()
	b.ResetTimer()
	for b.Loop() {
		b.StopTimer()
		parent := setup.World.NewEntity()
		setup.Service.SetParent(setup.World.NewEntity(), parent)
		for range 1 {
			child := setup.World.NewEntity()
			setup.Service.SetParent(child, parent)
		}
		b.StartTimer()
		setup.World.RemoveEntity(parent)
	}
}

func BenchmarkRemoveParentWith100Children(b *testing.B) {
	setup := NewSetup()
	b.ResetTimer()
	for b.Loop() {
		b.StopTimer()
		parent := setup.World.NewEntity()
		setup.Service.SetParent(setup.World.NewEntity(), parent)
		for range 100 {
			child := setup.World.NewEntity()
			setup.Service.SetParent(child, parent)
		}
		b.StartTimer()
		setup.World.RemoveEntity(parent)
	}
}
