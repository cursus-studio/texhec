# prototype
## Architecture
allows us to copy entity with copyable components

its to create copies of entities. its equivalent of unity prefabs (unity semantics)

## Benchmarks
```
$ go test ./... -bench=.
goos: linux
goarch: amd64
pkg: engine/modules/prototype/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkClone1-8         	  579832	      2063 ns/op
BenchmarkClone2-8         	  478113	      2703 ns/op
BenchmarkManual1Clone-8   	52015118	        26.10 ns/op
BenchmarkManual2Clone-8   	39357132	        34.08 ns/op
PASS
ok  	engine/modules/prototype/test	5.234s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7             46              4            194
-------------------------------------------------------------------------------
SUM:                             7             46              4            194
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/prototype.Service`

#### method Service Clone
Type: `func(cloned engine/services/ecs.EntityID) engine/services/ecs.EntityID`

#### method Service CloneTo
Type: `func(cloned engine/services/ecs.EntityID, clone engine/services/ecs.EntityID)`

### type CloneEvent
Type: `engine/modules/prototype.CloneEvent`

#### property CloneEvent Cloned
Type: `engine/services/ecs.EntityID`

#### property CloneEvent Clone
Type: `engine/services/ecs.EntityID`

## Functions
### func NewCloneEvent
Type: `func(cloned engine/services/ecs.EntityID, clone engine/services/ecs.EntityID) engine/modules/prototype.CloneEvent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.World`

`engine/modules/prototype`:
  - `engine/modules/prototype.NewCloneEvent`
  - `engine/modules/prototype.Service`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

`engine/services/ecs`:
  - `engine/services/ecs.AnyComponentArray`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.GetAny`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.SetAny`
  - `engine/services/ecs.World`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`