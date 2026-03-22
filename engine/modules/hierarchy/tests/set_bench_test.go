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

func BenchmarkRemoveNChildren(b *testing.B) {
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

func BenchmarkRemoveParentWith100Children(b *testing.B) {
	setup := NewSetup()
	parents := make([]ecs.EntityID, b.N)
	for i := range b.N {
		parents[i] = setup.World.NewEntity()
		setup.Service.SetParent(setup.World.NewEntity(), parents[i])
		for range 10 {
			child := setup.World.NewEntity()
			setup.Service.SetParent(child, parents[i])
		}
	}

	b.ResetTimer()
	for _, parent := range parents {
		setup.World.RemoveEntity(parent)
	}
}
