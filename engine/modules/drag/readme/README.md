# drag
## Architecture
allows us to drag objects and/or camera

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

`engine/modules/ecs`:
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/inputs`:
  - `engine/modules/inputs.Camera`
  - `engine/modules/inputs.DragEvent`
  - `engine/modules/inputs.From`
  - `engine/modules/inputs.To`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`