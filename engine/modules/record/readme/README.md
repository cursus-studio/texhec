# record
## Architecture
This module splits into two separate parts.\
One to record by `EntityID` which should be used to perform things localy.\
Second to record by `UUID` which should be used to record things for external machines.

Backwards recording returns state on recording start.\
Backwards recording can be used to record state before to smoothen changes.

Forwards recording returns state on recording end.\
Forwards recording can be used to record changes to send them somewhere else to replicate them.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              12            204             44           1013
-------------------------------------------------------------------------------
SUM:                            12            204             44           1013
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/record.Service`

#### method Service Entity
Type: `func() engine/modules/record.EntityKeyedRecorder`

#### method Service UUID
Type: `func() engine/modules/record.UUIDKeyedRecorder`

### type EntityKeyedRecorder
Type: `engine/modules/record.EntityKeyedRecorder`

#### method EntityKeyedRecorder Apply
Type: `func(engine/modules/record.Config, ...engine/modules/record.Recording)`

#### method EntityKeyedRecorder GetState
Type: `func(engine/modules/record.Config) engine/modules/record.Recording`
gets state as finished recording

#### method EntityKeyedRecorder StartBackwardsRecording
Type: `func(engine/modules/record.Config) engine/modules/record.RecordingID`
starts opened recording (opened recording is recorded until stopped)
applying it rewinds state.

#### method EntityKeyedRecorder StartRecording
Type: `func(engine/modules/record.Config) engine/modules/record.RecordingID`
starts opened recording (opened recording is recorded until stopped)
applying it on previous state will create current state

#### method EntityKeyedRecorder Stop
Type: `func(engine/modules/record.RecordingID) (r engine/modules/record.Recording, ok bool)`
finishes recording if open (false is returned if recording isn't started)

### type UUIDKeyedRecorder
Type: `engine/modules/record.UUIDKeyedRecorder`

#### method UUIDKeyedRecorder Apply
Type: `func(engine/modules/record.Config, ...engine/modules/record.UUIDRecording)`

#### method UUIDKeyedRecorder GetState
Type: `func(engine/modules/record.Config) engine/modules/record.UUIDRecording`
gets state as finished recording

#### method UUIDKeyedRecorder StartBackwardsRecording
Type: `func(engine/modules/record.Config) engine/modules/record.UUIDRecordingID`
starts opened recording (opened recording is recorded until stopped)
applying it rewinds state.

#### method UUIDKeyedRecorder StartRecording
Type: `func(engine/modules/record.Config) engine/modules/record.UUIDRecordingID`
starts opened recording (opened recording is recorded until stopped)
applying it on previous state will create current state

#### method UUIDKeyedRecorder Stop
Type: `func(engine/modules/record.UUIDRecordingID) (r engine/modules/record.UUIDRecording, ok bool)`
finishes recording if open (false is returned if recording isn't started)

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
Type: `engine/modules/datastructures.SparseArray[engine/modules/ecs.EntityID, []any]`
map[componentUUID][componentArrayLayoutID]any component
map[componentUUID]nil is when entity is removed

### type UUIDRecordingID
Type: `engine/modules/record.UUIDRecordingID`

### type UUIDRecording
Type: `engine/modules/record.UUIDRecording`

#### property UUIDRecording Entities
Type: `map[engine/modules/uuid.UUID][]any`
map[componentUUID][componentArrayLayoutID]any component
map[componentUUID]nil is when entity is removed

## Functions
### func NewConfig
Type: `func() engine/modules/record.Config`

### func AddToConfig
Type: `func[Component any](config engine/modules/record.Config) engine/modules/record.ComponentGetter[Component]`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.UUID`
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
  - `engine/modules/record.EntityKeyedRecorder`
  - `engine/modules/record.InheritZero`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.RecordedComponents`
  - `engine/modules/record.Recording`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.Service`
  - `engine/modules/record.UUIDKeyedRecorder`
  - `engine/modules/record.UUIDRecording`
  - `engine/modules/record.UUIDRecordingID`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.New`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.Service`
  - `engine/modules/uuid.UUID`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/ioc/v2`