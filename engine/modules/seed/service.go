// defines seed data structure used for generating pseudo random numbers
package seed

import (
	"engine/services/ecs"
	"errors"
	"math/rand/v2"

	"golang.org/x/exp/constraints"
)

type Seed uint64

func New[Number constraints.Integer](s Number) Seed {
	return Seed(s)
}

func (s *Seed) Source() uint64 { return uint64(*s) }
func (s *Seed) Value() uint64 {
	seed := uint64(*s)
	// SplitMix64 variant: A fast, high-quality 64-bit mixer
	seed = (seed ^ (seed >> 30)) * 0xbf58476d1ce4e5b9
	seed = (seed ^ (seed >> 27)) * 0x94d049bb133111eb
	seed = seed ^ (seed >> 31)
	return seed
}

func (s1 Seed) SeededRand(s2 Seed) *rand.Rand {
	// #nosec G404
	return rand.New(rand.NewPCG(s1.Value(), s2.Value()))
}

var (
	ErrWorldCanHaveOneSeed error = errors.New("world can have only one seed")
)

type SeedComponent struct {
	Seed Seed
}

func NewSeed[Number constraints.Integer](s Number) SeedComponent {
	return SeedComponent{New(s)}
}

type Service interface {
	Seed() ecs.ComponentsArray[SeedComponent]
	WorldSeed() (ecs.EntityID, bool)
}
