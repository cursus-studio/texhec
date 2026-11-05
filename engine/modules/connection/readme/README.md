# connection
## Architecture
defines connection and stores it in component

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/connection/test	0.044s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8            124             29            565
-------------------------------------------------------------------------------
SUM:                             8            124             29            565
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.Codec`
  - `engine.Connection`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Hierarchy`
  - `engine.Logger`
  - `engine.UUID`
  - `engine.World`

`engine/modules/connection`:
  - `engine/modules/connection.Close`
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.ConnEntity`
  - `engine/modules/connection.Connection`
  - `engine/modules/connection.ConnectionComponent`
  - `engine/modules/connection.Listener`
  - `engine/modules/connection.ListenerComponent`
  - `engine/modules/connection.MsgCtx`
  - `engine/modules/connection.MsgCtxSetter`
  - `engine/modules/connection.NewConnection`
  - `engine/modules/connection.NewListener`
  - `engine/modules/connection.NewMsgCtx`
  - `engine/modules/connection.NewSetConnUUIDDTO`
  - `engine/modules/connection.Send`
  - `engine/modules/connection.Service`
  - `engine/modules/connection.SetConnUUIDDTO`
  - `engine/modules/connection.SetCtx`
  - `engine/modules/connection.UUID`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSet`
  - `engine/modules/datastructures.Set`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.RegisterSystems`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`