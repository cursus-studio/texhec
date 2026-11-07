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
Go                               4             39              8            196
-------------------------------------------------------------------------------
SUM:                             4             39              8            196
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Logger`
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
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.LinkCacheComponent`
  - `engine/modules/uuid.LinkService`
  - `engine/modules/uuid.LinkUUIDComponent`
  - `engine/modules/uuid.NewLinkCache`
  - `engine/modules/uuid.NewLinkUUID`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.Service`
  - `engine/modules/uuid.UUID`
  - `engine/modules/uuid.UUIDComponent`

### Third Party
- `github.com/google/uuid`
- `github.com/ogiusek/ioc/v2`