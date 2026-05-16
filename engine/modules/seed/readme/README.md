# seed
## Architecture
defines seed data structure used for generating pseudo random numbers

## Types
### type Seed
Type: `engine/modules/seed.Seed`

#### method Seed Source
Type: `func() uint64`

#### method Seed Value
Type: `func() uint64`

#### method Seed SeededRand
Type: `func(s2 engine/modules/seed.Seed) *math/rand/v2.Rand`

## Functions
### func New
Type: `func[Number golang.org/x/exp/constraints.Integer](s Number) engine/modules/seed.Seed`


## Dependencies
### Third Party
`golang.org/x/exp/constraints`