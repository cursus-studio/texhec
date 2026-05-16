# netsync
## Architecture
synchronizes client and server worlds while accounting for permissions

## Types
### type Service
Type: `engine/modules/netsync.Service`

#### method Service Client
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/netsync.ClientComponent]`

#### method Service Server
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/netsync.ServerComponent]`

#### method Service Start
Type: `func() engine/services/ecs.SystemRegister`

#### method Service Stop
Type: `func() engine/services/ecs.SystemRegister`

### type AuthorizedEvent
Type: `engine/modules/netsync.AuthorizedEvent`
event pointer should implement it

#### method AuthorizedEvent SetConnection
Type: `func(engine/services/ecs.EntityID)`

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

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSparseSet`
  - `engine/services/datastructures.Remove`
  - `engine/services/datastructures.SparseSet`

`engine/services/ecs`:
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.RemoveEntity`
  - `engine/services/ecs.SystemRegister`

### Third Party
`github.com/ogiusek/events`
`github.com/ogiusek/ioc/v2`