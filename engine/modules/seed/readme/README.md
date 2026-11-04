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