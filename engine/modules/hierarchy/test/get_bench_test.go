package test

import (
	"testing"
)

func BenchmarkChildren_1(b *testing.B) {
	setup := NewSetup()
	parent := setup.World.NewEntity()
	child := setup.World.NewEntity()

	setup.Service.SetParent(child, parent)

	b.ResetTimer()
	for b.Loop() {
		setup.Service.Children(parent)
	}
}

func BenchmarkChildren_10(b *testing.B) {
	setup := NewSetup()
	parent := setup.World.NewEntity()

	for range 10 {
		child := setup.World.NewEntity()
		setup.Service.SetParent(child, parent)
	}

	b.ResetTimer()
	for b.Loop() {
		setup.Service.Children(parent)
	}
}

func BenchmarkChildren_100(b *testing.B) {
	setup := NewSetup()
	parent := setup.World.NewEntity()

	for range 100 {
		child := setup.World.NewEntity()
		setup.Service.SetParent(child, parent)
	}

	b.ResetTimer()
	for b.Loop() {
		setup.Service.Children(parent)
	}
}

func BenchmarkFlatChildren_1_1(b *testing.B) {
	setup := NewSetup()
	parent := setup.World.NewEntity()
	child := setup.World.NewEntity()
	grandChild := setup.World.NewEntity()

	setup.Service.SetParent(child, parent)
	setup.Service.SetParent(grandChild, child)

	b.ResetTimer()
	for b.Loop() {
		setup.Service.FlatChildren(parent)
	}
}

func BenchmarkFlatChildren_10_10(b *testing.B) {
	setup := NewSetup()
	parent := setup.World.NewEntity()

	for range 10 {
		child := setup.World.NewEntity()
		setup.Service.SetParent(child, parent)

		for range 10 {
			grandChild := setup.World.NewEntity()
			setup.Service.SetParent(grandChild, child)
		}
	}

	for b.Loop() {
		setup.Service.FlatChildren(parent)
	}
}
