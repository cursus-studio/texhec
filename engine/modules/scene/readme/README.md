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
## Types
### type Service
Type: `engine/modules/scene.Service`

#### method Service Scene
Type: `func() engine/modules/ecs.EntityID`

#### method Service SetScene
Type: `func(engine/modules/scene.ID, engine/modules/scene.Scene)`

### type System
Type: `engine/modules/scene.System`
change scene should happen after rendering
because on scene change everything is cleaned up

#### method System Register
Type: `func() error`

### type ChangeSceneEvent
Type: `engine/modules/scene.ChangeSceneEvent`

#### property ChangeSceneEvent ID
Type: `engine/modules/scene.ID`

### type ID
Type: `engine/modules/scene.ID`

#### property ID ID
Type: `string`

### type Scene
Type: `engine/modules/scene.Scene`

## Functions
### func NewChangeSceneEvent
Type: `func(ID engine/modules/scene.ID) engine/modules/scene.ChangeSceneEvent`

### func NewSceneId
Type: `func(id string) engine/modules/scene.ID`


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