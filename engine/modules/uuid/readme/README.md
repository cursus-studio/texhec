# uuid
## Architecture
integrates uuid

uuid allows us to create universaly unique entity

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               4             25              6             96
-------------------------------------------------------------------------------
SUM:                             4             25              6             96
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/uuid.Service`

#### method Service Component
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/uuid.Component]`

#### method Service Entity
Type: `func(engine/modules/uuid.UUID) (engine/services/ecs.EntityID, bool)`

#### method Service NewUUID
Type: `func() engine/modules/uuid.UUID`

### type Factory
Type: `engine/modules/uuid.Factory`

#### method Factory NewUUID
Type: `func() engine/modules/uuid.UUID`

### type Component
Type: `engine/modules/uuid.Component`

#### property Component ID
Type: `engine/modules/uuid.UUID`

### type UUID
Type: `engine/modules/uuid.UUID`

#### method UUID String
Type: `func() string`

#### method UUID Bytes
Type: `func() []byte`

## Functions
### func New
Type: `func(id engine/modules/uuid.UUID) engine/modules/uuid.Component`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.World`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.MapRelationPkg`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Factory`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.Service`
  - `engine/modules/uuid.UUID`

`engine/services/ecs`:
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.World`

### Third Party
- `github.com/google/uuid`
- `github.com/ogiusek/ioc/v2`