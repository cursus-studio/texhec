# drag
## Architecture
allows us to drag objects and/or camera

## Types
### type Service
Type: `engine/modules/drag.Service`

#### method Service Register
Type: `func() error`

### type DraggableEvent
Type: `engine/modules/drag.DraggableEvent`

#### property DraggableEvent Entity
Type: `engine/services/ecs.EntityID`

#### property DraggableEvent Drag
Type: `engine/modules/inputs.DragEvent`

#### method DraggableEvent ApplyDrag
Type: `func(dragEvent engine/modules/inputs.DragEvent) any`

## Functions
### func NewDraggable
Type: `func(entity engine/services/ecs.EntityID) engine/modules/drag.DraggableEvent`


## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             18              1             71
-------------------------------------------------------------------------------
SUM:                             3             18              1             71
-------------------------------------------------------------------------------

```
## Dependencies
`engine`:
  - `engine.Camera`
  - `engine.EngineWorld`
  - `engine.EventsBuilder`
  - `engine.Transform`

`engine/modules/drag`:
  - `engine/modules/drag.Drag`
  - `engine/modules/drag.DraggableEvent`
  - `engine/modules/drag.Entity`
  - `engine/modules/drag.Service`

`engine/modules/inputs`:
  - `engine/modules/inputs.Camera`
  - `engine/modules/inputs.DragEvent`
  - `engine/modules/inputs.From`
  - `engine/modules/inputs.To`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/ecs`:
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`