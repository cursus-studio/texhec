# connection
## Architecture
defines connection and stores it in component

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/connection/test	0.051s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7             89             17            411
-------------------------------------------------------------------------------
SUM:                             7             89             17            411
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/connection.Service`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/connection.ConnectionComponent]`

#### method Service Connect
Type: `func(addr string) error`

#### method Service Host
Type: `func(addr string) error`

#### method Service Listener
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/connection.ListenerComponent]`

#### method Service Register
Type: `func() error`

#### method Service TransferConnection
Type: `func(fromEntity engine/modules/ecs.EntityID, toEntity engine/modules/ecs.EntityID) error`

### type Conn
Type: `engine/modules/connection.Conn`
singular connection interface

#### method Conn Close
Type: `func() error`

#### method Conn Messages
Type: `func() chan any`
closed channel can be returned if connection is closed

#### method Conn Send
Type: `func(message any) error`
send has block behavior

### type ListenerComponent
Type: `engine/modules/connection.ListenerComponent`

#### method ListenerComponent Listener
Type: `func() net.Listener`

### type ConnectionComponent
Type: `engine/modules/connection.ConnectionComponent`

#### method ConnectionComponent Conn
Type: `func() engine/modules/connection.Conn`

## Functions
### func NewListener
Type: `func(listener net.Listener) engine/modules/connection.ListenerComponent`

### func NewConnection
Type: `func(conn engine/modules/connection.Conn) engine/modules/connection.ConnectionComponent`


## Dependencies
`engine`:
  - `engine.Codec`
  - `engine.EngineWorld`
  - `engine.Hierarchy`
  - `engine.Logger`
  - `engine.World`

`engine/modules/connection`:
  - `engine/modules/connection.Close`
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.ConnectionComponent`
  - `engine/modules/connection.Listener`
  - `engine/modules/connection.ListenerComponent`
  - `engine/modules/connection.NewConnection`
  - `engine/modules/connection.NewListener`
  - `engine/modules/connection.Service`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndex`
  - `engine/services/datastructures.NewSet`
  - `engine/services/datastructures.RemoveElements`
  - `engine/services/datastructures.Set`

### Third Party
- `github.com/ogiusek/ioc/v2`