# seed
## Architecture
defines seed data structure used for generating pseudo random numbers

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             21              4             75
-------------------------------------------------------------------------------
SUM:                             3             21              4             75
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/seed.Service`

#### method Service GenerateSeed
Type: `func() engine/modules/seed.Seed`

#### method Service Seed
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/seed.SeedComponent]`

### type Seed
Type: `engine/modules/seed.Seed`

#### method Seed Source
Type: `func() uint64`

#### method Seed Value
Type: `func() uint64`

#### method Seed SeededRand
Type: `func(s2 engine/modules/seed.Seed) *math/rand/v2.Rand`

### type SeedComponent
Type: `engine/modules/seed.SeedComponent`

#### property SeedComponent Seed
Type: `engine/modules/seed.Seed`

## Variables
### var ErrWorldShouldHaveOneSeed
Type: `error`

## Functions
### func New
Type: `func[Number golang.org/x/exp/constraints.Integer](s Number) engine/modules/seed.Seed`

### func NewSeed
Type: `func[Number golang.org/x/exp/constraints.Integer](s Number) engine/modules/seed.SeedComponent`


## Dependencies
`engine`:
  - `engine.Clock`
  - `engine.EngineWorld`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.ComponentPkg`

`engine/modules/seed`:
  - `engine/modules/seed.Seed`
  - `engine/modules/seed.SeedComponent`
  - `engine/modules/seed.Service`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`