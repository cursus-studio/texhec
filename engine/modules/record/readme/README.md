# record
## Architecture
This module splits into two separate parts.\
One to record by `EntityID` which should be used to perform things localy.\
Second to record by `UUID` which should be used to record things for external machines.

Backwards recording returns state on recording start.\
Backwards recording can be used to record state before to smoothen changes.

Forwards recording returns state on recording end.\
Forwards recording can be used to record changes to send them somewhere else to replicate them.

## Benchmarks
```
$ go test ./... -bench=.
goos: linux
goarch: amd64
pkg: engine/modules/record/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkEntityRecording-8                  	  333686	      3678 ns/op
BenchmarkCreateNEntitiesEntityRecording-8   	10982412	       128.6 ns/op
BenchmarkEntityApply1Entities-8             	 7506349	       155.2 ns/op
BenchmarkEntityApply10Entities-8            	 3837843	       310.2 ns/op
PASS
ok  	engine/modules/record/test	6.100s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8            126             34            631
-------------------------------------------------------------------------------
SUM:                             8            126             34            631
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.UUID`
  - `engine.WarmUp`
  - `engine.World`

`engine/modules/codec`:
  - `engine/modules/codec.Service`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSet`
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewSparseSet`
  - `engine/modules/datastructures.Set`
  - `engine/modules/datastructures.SparseArray`
  - `engine/modules/datastructures.SparseSet`

`engine/modules/ecs`:
  - `engine/modules/ecs.AnyComponentArray`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.NewWorld`
  - `engine/modules/ecs.World`

`engine/modules/record`:
  - `engine/modules/record.AddToConfig`
  - `engine/modules/record.ComponentsOrder`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entities`
  - `engine/modules/record.EntityComponents`
  - `engine/modules/record.InheritZero`
  - `engine/modules/record.Keys`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.RecordedComponents`
  - `engine/modules/record.Recording`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.Service`
  - `engine/modules/record.UUIDState`
  - `engine/modules/record.Values`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.CreateGetEntity`
  - `engine/modules/uuid.CreateGetUUID`
  - `engine/modules/uuid.Service`
  - `engine/modules/uuid.UUID`
  - `engine/modules/uuid.UUIDComponent`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/ioc/v2`