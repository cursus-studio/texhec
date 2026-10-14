# netsync
## Architecture
synchronizes client and server worlds while accounting for permissions

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/netsync/test/prediction_test	0.014s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              14            198             78            937
Markdown                         1              2              0              5
-------------------------------------------------------------------------------
SUM:                            15            200             78            942
-------------------------------------------------------------------------------
```
## TODO
This module needs clean up for transparent events and ticks.
Currently we hope event won't land on tick instead of handling it properly.
Transparent events aren't tracked to reverse back in state.
Client also hopes he doesn't receive event at a bad time because if event from server lands in a bad time
transparent events are going to be messed.



## Types
### type Service
Type: `engine/modules/netsync.Service`

#### method Service Client
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/netsync.ClientComponent]`

#### method Service Clients
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/netsync.ClientsComponent]`

#### method Service Server
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/netsync.ServerComponent]`

#### method Service Start
Type: `func() engine/modules/ecs.SystemRegister`

#### method Service Stop
Type: `func() engine/modules/ecs.SystemRegister`

### type ValidatedEvent
Type: `engine/modules/netsync.ValidatedEvent`
event pointer should implement it

#### method ValidatedEvent SetConnection
Type: `func(engine/modules/ecs.EntityID)`

### type ServerComponent
Type: `engine/modules/netsync.ServerComponent`
entity with this component and with connection component will be one with which we'll synchronize

### type ClientsComponent
Type: `engine/modules/netsync.ClientsComponent`

### type ClientComponent
Type: `engine/modules/netsync.ClientComponent`
entity with this component and connection will get notifications about changes


## Dependencies
`engine`:
  - `engine.Connection`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Hierarchy`
  - `engine.Logger`
  - `engine.Loop`
  - `engine.NetSync`
  - `engine.Record`
  - `engine.World`

`engine/modules/connection`:
  - `engine/modules/connection.Close`
  - `engine/modules/connection.Component`
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.ConnEntity`
  - `engine/modules/connection.MsgCtx`
  - `engine/modules/connection.Send`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewStableSparseArray`
  - `engine/modules/datastructures.SparseArray`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.LastTick`
  - `engine/modules/loop.NewEmitOnTickEvent`
  - `engine/modules/loop.SyncToUnixNano`
  - `engine/modules/loop.TickEvent`
  - `engine/modules/loop.UnixNano`

`engine/modules/netsync`:
  - `engine/modules/netsync.Client`
  - `engine/modules/netsync.ClientComponent`
  - `engine/modules/netsync.Clients`
  - `engine/modules/netsync.ClientsComponent`
  - `engine/modules/netsync.Server`
  - `engine/modules/netsync.ServerComponent`
  - `engine/modules/netsync.Service`
  - `engine/modules/netsync.SetConnection`
  - `engine/modules/netsync.ValidatedEvent`

`engine/modules/record`:
  - `engine/modules/record.Apply`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entity`
  - `engine/modules/record.GetState`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.Recording`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.StartBackwardsRecording`
  - `engine/modules/record.Stop`
  - `engine/modules/record.UUID`
  - `engine/modules/record.UUIDRecording`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/ogiusek/null`