# netsync
## Architecture
synchronizes client and server worlds while accounting for permissions

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              11            112             91            660
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            12            112             91            661
-------------------------------------------------------------------------------
```
## TODO
Create more features to allow more specific features to allow more specific calls

## Types
### type Service
Type: `engine/modules/netsync.Service`

#### method Service Client
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/netsync.ClientComponent]`

#### method Service Server
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/netsync.ServerComponent]`

#### method Service Start
Type: `func() engine/modules/ecs.SystemRegister`

#### method Service Stop
Type: `func() engine/modules/ecs.SystemRegister`

### type AuthorizedEvent
Type: `engine/modules/netsync.AuthorizedEvent`
event pointer should implement it

#### method AuthorizedEvent SetConnection
Type: `func(engine/modules/ecs.EntityID)`

### type ServerComponent
Type: `engine/modules/netsync.ServerComponent`
entity with this component and with connection component will be one with which we'll synchronize

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
  - `engine.NetSync`
  - `engine.Record`
  - `engine.UUID`
  - `engine.World`

`engine/modules/connection`:
  - `engine/modules/connection.Close`
  - `engine/modules/connection.Component`
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.Messages`
  - `engine/modules/connection.Send`

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

`engine/modules/netsync`:
  - `engine/modules/netsync.AuthorizedEvent`
  - `engine/modules/netsync.Client`
  - `engine/modules/netsync.ClientComponent`
  - `engine/modules/netsync.Server`
  - `engine/modules/netsync.ServerComponent`
  - `engine/modules/netsync.Service`
  - `engine/modules/netsync.SetConnection`

`engine/modules/record`:
  - `engine/modules/record.Apply`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entities`
  - `engine/modules/record.GetState`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.StartBackwardsRecording`
  - `engine/modules/record.StartRecording`
  - `engine/modules/record.Stop`
  - `engine/modules/record.UUID`
  - `engine/modules/record.UUIDRecording`
  - `engine/modules/record.UUIDRecordingID`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`