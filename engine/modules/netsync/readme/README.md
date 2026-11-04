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
Go                              14            198             80           1000
Markdown                         1              2              0              5
-------------------------------------------------------------------------------
SUM:                            15            200             80           1005
-------------------------------------------------------------------------------
```
## TODO
This module needs clean up for transparent events and ticks.
Currently we hope event won't land on tick instead of handling it properly.
Transparent events aren't tracked to reverse back in state.
Client also hopes he doesn't receive event at a bad time because if event from server lands in a bad time
transparent events are going to be messed.



## Dependencies
`engine`:
  - `engine.Connection`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Logger`
  - `engine.Loop`
  - `engine.NetSync`
  - `engine.Record`
  - `engine.World`

`engine/modules/connection`:
  - `engine/modules/connection.Close`
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.ConnEntity`
  - `engine/modules/connection.Connection`
  - `engine/modules/connection.MsgCtx`
  - `engine/modules/connection.Send`

`engine/modules/connection/pkg`:
  - `engine/modules/connection/pkg.ParentPkg`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewStableSparseArray`
  - `engine/modules/datastructures.SparseArray`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.LastTick`
  - `engine/modules/loop.SyncToUnixNano`
  - `engine/modules/loop.TickEvent`
  - `engine/modules/loop.UnixNano`

`engine/modules/netsync`:
  - `engine/modules/netsync.Client`
  - `engine/modules/netsync.ClientComponent`
  - `engine/modules/netsync.ClientsComponent`
  - `engine/modules/netsync.Server`
  - `engine/modules/netsync.ServerComponent`
  - `engine/modules/netsync.Service`
  - `engine/modules/netsync.SetConnection`
  - `engine/modules/netsync.ValidatedEvent`

`engine/modules/record`:
  - `engine/modules/record.AddToConfig`
  - `engine/modules/record.Apply`
  - `engine/modules/record.ApplyUUIDState`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entities`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.Peek`
  - `engine/modules/record.Recording`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.StartBackwardsRecording`
  - `engine/modules/record.StartRecording`
  - `engine/modules/record.Stop`
  - `engine/modules/record.UUIDState`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/ogiusek/null`