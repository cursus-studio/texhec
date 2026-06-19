package test

import (
	"engine/modules/ecs"
	"engine/modules/transform"
	"testing"
)

func BenchmarkGetPos(b *testing.B) {
	setup := NewSetup()
	entity := setup.NewEntity()
	for b.Loop() {
		setup.transform.AbsolutePos().Get(entity)
	}
}

func BenchmarkRawGetPos(b *testing.B) {
	world := ecs.NewWorld()
	arr := ecs.GetComponentArray[transform.AbsolutePosComponent](world) // no wrappers
	entity := world.NewEntity()
	for b.Loop() {
		arr.Get(entity)
	}
}

func BenchmarkSetAbsolutePos(b *testing.B) {
	setup := NewSetup()

	entity := setup.NewEntity()
	for i := range b.N {
		pos := transform.NewPos(0, 0, float32(i))
		setup.transform.AbsolutePos().Set(entity, transform.AbsolutePosComponent(pos))
	}
}

func BenchmarkSetAndGetAbsolutePos(b *testing.B) {
	setup := NewSetup()

	entity := setup.NewEntity()
	for i := range b.N {
		pos := transform.NewPos(0, 0, float32(i))
		setup.transform.Pos().Set(entity, pos)
		for range 1 {
			setup.transform.AbsolutePos().Get(entity)
		}
	}
}
