# scene
## Architecture
defines scenes and decouples scene constructor from its id

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             24              7             83
-------------------------------------------------------------------------------
SUM:                             3             24              7             83
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.EventsBuilder`
  - `engine.Logger`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/scene`:
  - `engine/modules/scene.ChangeSceneEvent`
  - `engine/modules/scene.ID`
  - `engine/modules/scene.Scene`
  - `engine/modules/scene.Service`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`