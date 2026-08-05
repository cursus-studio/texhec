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
BenchmarkClone1-8         	  567273	      2123 ns/op
BenchmarkClone2-8         	  553878	      2056 ns/op
BenchmarkManual1Clone-8   	48528190	        24.75 ns/op
BenchmarkManual2Clone-8   	41206909	        29.87 ns/op
PASS
ok  	engine/modules/prototype/test	4.795s
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
Type: `func(cloned engine/modules/ecs.EntityID) engine/modules/ecs.EntityID`

#### method Service CloneTo
Type: `func(cloned engine/modules/ecs.EntityID, clone engine/modules/ecs.EntityID)`

### type CloneEvent
Type: `engine/modules/prototype.CloneEvent`

#### property CloneEvent Cloned
Type: `engine/modules/ecs.EntityID`

#### property CloneEvent Clone
Type: `engine/modules/ecs.EntityID`

## Functions
### func NewCloneEvent
Type: `func(cloned engine/modules/ecs.EntityID, clone engine/modules/ecs.EntityID) engine/modules/prototype.CloneEvent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.AnyComponentArray`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.World`

`engine/modules/prototype`:
  - `engine/modules/prototype.NewCloneEvent`
  - `engine/modules/prototype.Service`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`