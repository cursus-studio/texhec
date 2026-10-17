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
BenchmarkEntityRecording-8                  	  306679	      3700 ns/op
BenchmarkCreateNEntitiesEntityRecording-8   	 8346109	       132.7 ns/op
BenchmarkEntityApply1Entities-8             	 8308940	       143.1 ns/op
BenchmarkEntityApply10Entities-8            	 3907071	       308.6 ns/op
PASS
ok  	engine/modules/record/test	5.663s
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
## Types
### type Service
Type: `engine/modules/record.Service`

#### method Service Apply
Type: `func(engine/modules/record.Config, ...engine/modules/record.Recording)`

#### method Service ApplyUUIDState
Type: `func(engine/modules/record.Config, engine/modules/record.UUIDState)`
applies state and purges all entities which have any of config components
it is used to apply whole world state therefor it ensures it is synchronized

#### method Service Peek
Type: `func(engine/modules/record.RecordingID) (r engine/modules/record.Recording, ok bool)`
returns recording without closing it

#### method Service StartBackwardsRecording
Type: `func(engine/modules/record.Config) engine/modules/record.RecordingID`
starts opened recording (opened recording is recorded until stopped)
applying it rewinds state.

#### method Service StartRecording
Type: `func(engine/modules/record.Config) engine/modules/record.RecordingID`
starts opened recording (opened recording is recorded until stopped)
applying it on previous state will create current state

#### method Service Stop
Type: `func(engine/modules/record.RecordingID) (r engine/modules/record.Recording, ok bool)`
finishes recording if open (false is returned if recording isn't started)

#### method Service UUIDState
Type: `func(engine/modules/record.Recording) engine/modules/record.UUIDState`
doesn't work for removing entities
is it meant to be used for recordings with service lifetime

### type Config
Type: `engine/modules/record.Config`

#### property Config ComponentsOrder
Type: `*[]reflect.Type`

#### property Config ComponentsIndices
Type: `map[reflect.Type]int`

#### property Config RecordedComponents
Type: `map[reflect.Type]func(engine/modules/ecs.World) engine/modules/ecs.AnyComponentArray`

#### property Config InheritZero
Type: `map[reflect.Type]func(engine/modules/ecs.World)`

### type ComponentGetter
Type: `engine/modules/record.ComponentGetter[Component any]`

### type RecordingID
Type: `engine/modules/record.RecordingID`

### type Recording
Type: `engine/modules/record.Recording`
recording cannot be encoded

#### property Recording Entities
Type: `engine/modules/datastructures.SparseArray[engine/modules/ecs.EntityID, engine/modules/record.EntityComponents]`
[componentArrayLayoutID]any component
nil for removed entity

### type UUIDState
Type: `engine/modules/record.UUIDState`

#### property UUIDState Keys
Type: `[]engine/modules/uuid.UUID`

#### property UUIDState Values
Type: `[]engine/modules/record.EntityComponents`

## Functions
### func NewConfig
Type: `func() engine/modules/record.Config`

### func AddToConfig
Type: `func[Component any](config engine/modules/record.Config) engine/modules/record.ComponentGetter[Component]`


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
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.CreateGetEntity`
  - `engine/modules/uuid.CreateGetUUID`
  - `engine/modules/uuid.Service`
  - `engine/modules/uuid.UUID`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/ioc/v2`