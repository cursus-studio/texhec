package test

import (
	"core/modules/reach"
	"core/modules/tile"
	"testing"
)

func BenchmarkDist(b *testing.B) {
	s := NewSetup()
	b.ResetTimer()
	for b.Loop() {
		s.Reach().Distance(
			tile.NewPos(1, 1), tile.NewSize(1, 1),
			tile.NewPos(1, 0), tile.NewSize(1, 1),
		)
	}
}

func Benchmark4TilesWithinReach(b *testing.B) {
	s := NewSetup()
	entity := s.World().NewEntity()
	s.Tile().Pos().Set(entity, tile.NewPos(0, 0))
	s.Tile().Size().Set(entity, tile.NewSize(1, 1))
	s.ReachT.Component().Set(entity, reach.NewReach[FeatureComponent](1))
	b.ResetTimer()
	for b.Loop() {
		_ = s.ReachT.TilesWithinReach(entity)
	}
}
func Benchmark12TilesWithinReach(b *testing.B) {
	s := NewSetup()
	entity := s.World().NewEntity()
	s.Tile().Pos().Set(entity, tile.NewPos(2, 2))
	s.Tile().Size().Set(entity, tile.NewSize(1, 1))
	s.ReachT.Component().Set(entity, reach.NewReach[FeatureComponent](2))
	b.ResetTimer()
	for b.Loop() {
		_ = s.ReachT.TilesWithinReach(entity)
	}
}
