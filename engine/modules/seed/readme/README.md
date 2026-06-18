# seed
## Architecture
defines seed data structure used for generating pseudo random numbers

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             19              3             78
-------------------------------------------------------------------------------
SUM:                             3             19              3             78
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/seed.Service`

#### method Service Seed
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/seed.SeedComponent]`

#### method Service WorldSeed
Type: `func() (engine/services/ecs.EntityID, bool)`

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
### var ErrWorldCanHaveOneSeed
Type: `error`

## Functions
### func New
Type: `func[Number golang.org/x/exp/constraints.Integer](s Number) engine/modules/seed.Seed`

### func NewSeed
Type: `func[Number golang.org/x/exp/constraints.Integer](s Number) engine/modules/seed.SeedComponent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Logger`
  - `engine.World`

`engine/modules/seed`:
  - `engine/modules/seed.ErrWorldCanHaveOneSeed`
  - `engine/modules/seed.SeedComponent`
  - `engine/modules/seed.Service`

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.RemoveEntity`

### Third Party
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`