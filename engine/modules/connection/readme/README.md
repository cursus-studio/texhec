# connection
## Architecture
defines connection and stores it in component

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/connection/test	0.033s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8            117             27            522
-------------------------------------------------------------------------------
SUM:                             8            117             27            522
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/connection.Service`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/connection.ConnectionComponent]`

#### method Service Connect
Type: `func(entity engine/modules/ecs.EntityID, addr string) error`

#### method Service Host
Type: `func(entity engine/modules/ecs.EntityID, addr string) error`

#### method Service Listener
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/connection.ListenerComponent]`

#### method Service Register
Type: `func() error`

#### method Service TransferConnection
Type: `func(fromEntity engine/modules/ecs.EntityID, toEntity engine/modules/ecs.EntityID) error`

### type MsgCtxSetter
Type: `engine/modules/connection.MsgCtxSetter`

#### method MsgCtxSetter SetCtx
Type: `func(tgtMsg engine/modules/connection.MsgCtx)`

### type Conn
Type: `engine/modules/connection.Conn`
singular connection interface

#### method Conn Close
Type: `func()`

#### method Conn Send
Type: `func(message any) error`
send has block behavior

### type MsgCtx
Type: `engine/modules/connection.MsgCtx`

#### property MsgCtx ConnEntity
Type: `engine/modules/ecs.EntityID`

#### method MsgCtx SetCtx
Type: `func(tgtMsg engine/modules/connection.MsgCtx)`

### type ListenerComponent
Type: `engine/modules/connection.ListenerComponent`

#### method ListenerComponent Listener
Type: `func() net.Listener`

### type ConnectionComponent
Type: `engine/modules/connection.ConnectionComponent`

#### method ConnectionComponent Conn
Type: `func() engine/modules/connection.Conn`

## Functions
### func NewMsgCtx
Type: `func(connEntity engine/modules/ecs.EntityID) engine/modules/connection.MsgCtx`

### func NewListener
Type: `func(listener net.Listener) engine/modules/connection.ListenerComponent`

### func NewConnection
Type: `func(conn engine/modules/connection.Conn) engine/modules/connection.ConnectionComponent`


## Dependencies
`engine`:
  - `engine.Codec`
  - `engine.Connection`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Hierarchy`
  - `engine.Logger`
  - `engine.World`

`engine/modules/connection`:
  - `engine/modules/connection.Close`
  - `engine/modules/connection.Component`
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.ConnectionComponent`
  - `engine/modules/connection.Listener`
  - `engine/modules/connection.ListenerComponent`
  - `engine/modules/connection.MsgCtx`
  - `engine/modules/connection.MsgCtxSetter`
  - `engine/modules/connection.NewConnection`
  - `engine/modules/connection.NewListener`
  - `engine/modules/connection.NewMsgCtx`
  - `engine/modules/connection.Service`
  - `engine/modules/connection.SetCtx`

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

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`