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
Go                               5             38              7            168
-------------------------------------------------------------------------------
SUM:                             5             38              7            168
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/uuid.Service`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/uuid.Component]`

#### method Service Entity
Type: `func(engine/modules/uuid.UUID) (engine/modules/ecs.EntityID, bool)`

#### method Service NewUUID
Type: `func() engine/modules/uuid.UUID`

#### method Service NewUUIDFromString
Type: `func(string) engine/modules/uuid.UUID`

### type LinkService
Type: `engine/modules/uuid.LinkService[Wrapped any]`

#### method LinkService Cache
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/uuid.LinkCacheComponent[Wrapped]]`

#### method LinkService Get
Type: `func(linkSrc engine/modules/ecs.EntityID) (linkDst engine/modules/ecs.EntityID, ok bool)`

#### method LinkService SetUUID
Type: `func(engine/modules/ecs.EntityID, engine/modules/uuid.UUID)`

#### method LinkService UUID
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/uuid.LinkUUIDComponent[Wrapped]]`

### type Factory
Type: `engine/modules/uuid.Factory`

#### method Factory NewUUID
Type: `func() engine/modules/uuid.UUID`

#### method Factory NewUUIDFromString
Type: `func(string) engine/modules/uuid.UUID`

### type Component
Type: `engine/modules/uuid.Component`

#### property Component ID
Type: `engine/modules/uuid.UUID`

### type LinkUUIDComponent
Type: `engine/modules/uuid.LinkUUIDComponent[Wrappd any]`

#### property LinkUUIDComponent UUID
Type: `engine/modules/uuid.UUID`

### type LinkCacheComponent
Type: `engine/modules/uuid.LinkCacheComponent[Wrappd any]`

#### property LinkCacheComponent Entity
Type: `engine/modules/ecs.EntityID`

### type UUID
Type: `engine/modules/uuid.UUID`

#### method UUID String
Type: `func() string`

#### method UUID Bytes
Type: `func() []byte`

## Functions
### func New
Type: `func(id engine/modules/uuid.UUID) engine/modules/uuid.Component`

### func NewLinkUUID
Type: `func[Wrapped any](uuid engine/modules/uuid.UUID) engine/modules/uuid.LinkUUIDComponent[Wrapped]`

### func NewLinkCache
Type: `func[Wrapped any](entity engine/modules/ecs.EntityID) engine/modules/uuid.LinkCacheComponent[Wrapped]`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.UUID`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.World`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.MapRelationPkg`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Factory`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.LinkCacheComponent`
  - `engine/modules/uuid.LinkService`
  - `engine/modules/uuid.LinkUUIDComponent`
  - `engine/modules/uuid.NewLinkCache`
  - `engine/modules/uuid.NewLinkUUID`
  - `engine/modules/uuid.Service`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/google/uuid`
- `github.com/ogiusek/ioc/v2`