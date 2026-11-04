# loop
## Architecture
defines game loop

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             56             21            212
-------------------------------------------------------------------------------
SUM:                             3             56             21            212
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.Clock`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/loop`:
  - `engine/modules/loop.ConfigureEvent`
  - `engine/modules/loop.EmitOnFrameComponent`
  - `engine/modules/loop.EmitOnFrameEvent`
  - `engine/modules/loop.EmitOnTickComponent`
  - `engine/modules/loop.EmitOnTickEvent`
  - `engine/modules/loop.Event`
  - `engine/modules/loop.FPS`
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.LastTickComponent`
  - `engine/modules/loop.NewLastTick`
  - `engine/modules/loop.Service`
  - `engine/modules/loop.Stats`
  - `engine/modules/loop.StopEvent`
  - `engine/modules/loop.TPS`
  - `engine/modules/loop.TickEvent`
  - `engine/modules/loop.UnixNano`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.AddTickEvent`
  - `engine/modules/netsync/pkg.Config`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`