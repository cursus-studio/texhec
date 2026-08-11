# ecs
## Architecture
This **ECS** framework tries to follow only 2 rules:
- **DOD** (data oriented design) to ensure highest performance
- **Simplicity** to ensure developer productivity, scalability and performance

### What is DOD
**DOD** is focusing on data layout so its performant to access and modify.\
In short writing code so its efficient.\
This simple goal (write efficient code) has massive consequences in how code is written and how we think about code.\
I won't talk about all of them here but the core idea comes from this goal.

### Data structures
We use sparse structures to store data. This enables us to efficiently access data.

### Lazy listeners
`BeforeGet` allows us to lazily modify data before getting it.\
Lazy listener should be a default choice for modifying data.\
Lazily modifying data allows us to do this in batches which is highly efficient

### Active listeners
`OnUpsert` and `OnRemove` are active listeners.\
They allow to instantly act on data modification but:
- they are heavy and called for every single entity
- using them can cause calling them too much in dependency loops

### Why golang GC (garbage collector) isn't a problem
We follow **DOD** there for GC isn't laden with managing pointers because there are little pointers to manage.\
This makes golang a perfect candidate for this project because of high developer efficiency and low performance overhead.

### Architecture changes to revise in the future
Change architecture to:
- add entity mechanism to wait until entity is released in all systems
- call before get on all components

This would **simplify** codebase and would make it follow **DOD** more.
This would depracate `EnsureExists`.

## Usage examples
### World
```go
type World interface {
	entitiesInterface
	componentsInterface
}
```

### Entities
This is one of interfaces from which `ecs.World` is composed.
```go
type entitiesInterface interface {
	GetEntities() []EntityID
	EntityExists(EntityID) bool

	NewEntity() EntityID
	EnsureExists(EntityID)
	RemoveEntity(EntityID)
}
```

#### `GetEntities`
Returns all entities.\
It returns original slice so if you want to perform operations on slice or\
you want to perform write action on `entitiesInterface` then copy this slice.

#### `EntityExists`
Returns `true` if entity exists.

#### `NewEntity`
Creates new entity and returns its id.

#### `EnsureExists`
Its very niche method. It ensures that entity with specific id exists.\
It isn't recommended for most use cases.\
It ensures that entity with specific id exists by creating it if it doesn't exist.

#### `RemoveEntity`
Removes entity with specific id.

### Access to components array
#### Interface
```go
func GetComponentsArray[Component any](world World) ComponentsArray[Component]
```

#### Example usage
```go
func _(world ecs.World) {
    arr := ecs.GetComponentsArray(world)
    // do something with components array
}
```

### Components array
```go
type AnyComponentArray interface {
	GetAny(entity EntityID) (any, bool)
	GetEntities() []EntityID

	// when type doesn't match error is returned
	SetAny(EntityID, any) error
	Remove(EntityID)

	// configuration
	// on dependency change its also applied here
	AddDependency(AnyComponentArray)
	AddDirtySet(DirtySet)
	BeforeGet(BeforeGet)

	OnUpsert(OnMod)
	OnRemove(OnMod)
}

type ComponentsArray[Component any] interface {
	AnyComponentArray
	Get(entity EntityID) (Component, bool)

	Set(EntityID, Component)

	// configuration
	SetEmpty(Component)
}
```

#### `GetAny`
Returns component as any.\
Its for generic applications

#### `GetEntities`
Returns all entities which have this component.\
It returns original slice so if you want to perform operations on slice or\
you want to perform write actions on `ComponentsArray` then copy this slice.

#### `SetAny`
Sets component for specific entity.\
Its for generic applications

#### `Remove`
It removes component from specific entity.

#### `AddDependency`
It adds other components array as a dependency.\
If component array is dependent from other component array then\
when one component array dirty sets are marked dirty then second dirty sets are marked dirty to automatically.

#### `AddDirtySet`
It adds dirty set to mark dirty upon any component modification.

#### `BeforeGet`
It gets called each time on get.\
It can be used to update data from which component is dependent.\
It should use dirty set on start to ensure we escape it instantly if there is nothing to do.

Calling `BeforeGet` might sound not efficient but it is efficient because checking\
dirty set takes 2ns.
```
BenchmarkDirtySetGet-8                 	612629619	         1.974 ns/op	       0 B/op	       0 allocs/op
```

#### `OnUpsert`
It gets called instantly on component modification or addition removal.\
Its less efficient then `BeforeGet`

#### `OnRemove`
It gets called instantly on component removal.

#### `Get`
It returns entity component if entity has any.
Else it returns default value and `false` (`false` standing for !ok).

#### `Set`
It sets component for entity.

#### `SetEmpty`
It sets default value of component.

### Dirty set
```go
type DirtySet interface {
	// get also clears
	Get() []EntityID
	Dirty(EntityID)
	Clear()

	Ok() bool
	Release()
}
```

#### `Get`
Returns all dirty entities and marks them clear.

#### `Dirty`
Marks entity as dirty (modified).

#### `Clear`
Clears all dirty entities.

#### `Ok`
Returns `false` after being released.

#### `Release`
Releases dirty set and allows framework to release it properly

## Benchmarks
```
$ go test ./... -bench=.
goos: linux
goarch: amd64
pkg: engine/modules/ecs/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkDirtySetDirty-8               	137151241	         8.231 ns/op
BenchmarkDirtySetDirtyInversed-8       	144657516	         7.394 ns/op
BenchmarkDirtySetGet-8                 	590244398	         2.090 ns/op
BenchmarkDirtySetDirtyAndGet-8         	169425655	         7.102 ns/op
BenchmarkDirtySetDirtyAnd1Get-8        	127674126	         8.851 ns/op
Benchmark4SavesWith7Systems-8          	26503406	        42.23 ns/op
Benchmark16SavesWith7Systems-8         	 6938511	       165.1 ns/op
Benchmark256SavesWith7Systems-8        	  488564	      2509 ns/op
Benchmark4096SavesWith7Systems-8       	   29869	     40203 ns/op
Benchmark16384SavesWith7Systems-8      	    7695	    159021 ns/op
Benchmark65536SavesWith7Systems-8      	    1840	    643265 ns/op
Benchmark262144SavesWith7Systems-8     	     460	   2573859 ns/op
BenchmarkGetComponent-8                	81757089	        14.96 ns/op
BenchmarkCreateComponents-8            	39956193	        31.43 ns/op
BenchmarkUpdateComponents-8            	100000000	        11.46 ns/op
BenchmarkRemoveComponent-8             	77082870	        15.48 ns/op
BenchmarkRemoveEntityWithComponent-8   	35099274	        35.41 ns/op
BenchmarkRemoveEntity-8                	65063078	        18.66 ns/op
PASS
ok  	engine/modules/ecs/test	37.578s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              13            218             43            920
Markdown                         1             44              0            157
-------------------------------------------------------------------------------
SUM:                            14            262             43           1077
-------------------------------------------------------------------------------
```
## Types
### type ApplyEntityEvent
Type: `engine/modules/ecs.ApplyEntityEvent`
event wrappers

#### method ApplyEntityEvent ApplyEntity
Type: `func(entityEmitting engine/modules/ecs.EntityID) (event any)`

### type SystemRegister
Type: `engine/modules/ecs.SystemRegister`
systems

#### method SystemRegister Register
Type: `func() error`

### type DirtySet
Type: `engine/modules/ecs.DirtySet`
dirty set

#### method DirtySet Clear
Type: `func()`

#### method DirtySet Dirty
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID)`

#### method DirtySet Get
Type: `func() []engine/modules/ecs/internal/ecstypes.EntityID`

#### method DirtySet Ok
Type: `func() bool`

#### method DirtySet Release
Type: `func()`

### type Component
Type: `engine/modules/ecs.Component`

### type World
Type: `engine/modules/ecs.World`

#### method World EnsureExists
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID)`

#### method World EntityExists
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID) bool`

#### method World GetArrByComp
Type: `func(engine/modules/ecs/internal/ecstypes.Component) (engine/modules/ecs/internal/ecstypes.AnyComponentArray, bool)`

#### method World GetEntities
Type: `func() []engine/modules/ecs/internal/ecstypes.EntityID`

#### method World NewEntity
Type: `func() engine/modules/ecs/internal/ecstypes.EntityID`

#### method World OnArrayInitialization
Type: `func(func(engine/modules/ecs/internal/ecstypes.AnyComponentArray))`

#### method World RemoveEntity
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID)`

#### method World WarmUp
Type: `func()`

### type AnyComponentArray
Type: `engine/modules/ecs.AnyComponentArray`
component array

#### method AnyComponentArray AddDependency
Type: `func(engine/modules/ecs/internal/ecstypes.AnyComponentArray)`

#### method AnyComponentArray AddDirtySet
Type: `func(engine/modules/ecs/internal/ecstypes.DirtySet)`

#### method AnyComponentArray BeforeGet
Type: `func(engine/modules/ecs/internal/ecstypes.BeforeGet)`

#### method AnyComponentArray GetAny
Type: `func(entity engine/modules/ecs/internal/ecstypes.EntityID) (engine/modules/ecs/internal/ecstypes.Component, bool)`

#### method AnyComponentArray GetEntities
Type: `func() []engine/modules/ecs/internal/ecstypes.EntityID`

#### method AnyComponentArray OnMod
Type: `func(engine/modules/ecs/internal/ecstypes.OnMod)`

#### method AnyComponentArray OnRemove
Type: `func(engine/modules/ecs/internal/ecstypes.OnMod)`

#### method AnyComponentArray OnUpsert
Type: `func(engine/modules/ecs/internal/ecstypes.OnMod)`

#### method AnyComponentArray Remove
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID)`

#### method AnyComponentArray SetAny
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID, engine/modules/ecs/internal/ecstypes.Component)`

### type ComponentArray
Type: `engine/modules/ecs.ComponentArray[Component any]`

#### method ComponentArray AddDependency
Type: `func(engine/modules/ecs/internal/ecstypes.AnyComponentArray)`

#### method ComponentArray AddDirtySet
Type: `func(engine/modules/ecs/internal/ecstypes.DirtySet)`

#### method ComponentArray BeforeGet
Type: `func(engine/modules/ecs/internal/ecstypes.BeforeGet)`

#### method ComponentArray Get
Type: `func(entity engine/modules/ecs/internal/ecstypes.EntityID) (Component, bool)`

#### method ComponentArray GetAny
Type: `func(entity engine/modules/ecs/internal/ecstypes.EntityID) (engine/modules/ecs/internal/ecstypes.Component, bool)`

#### method ComponentArray GetEmpty
Type: `func() Component`

#### method ComponentArray GetEntities
Type: `func() []engine/modules/ecs/internal/ecstypes.EntityID`

#### method ComponentArray OnMod
Type: `func(engine/modules/ecs/internal/ecstypes.OnMod)`

#### method ComponentArray OnRemove
Type: `func(engine/modules/ecs/internal/ecstypes.OnMod)`

#### method ComponentArray OnUpsert
Type: `func(engine/modules/ecs/internal/ecstypes.OnMod)`

#### method ComponentArray Remove
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID)`

#### method ComponentArray Set
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID, Component)`

#### method ComponentArray SetAny
Type: `func(engine/modules/ecs/internal/ecstypes.EntityID, engine/modules/ecs/internal/ecstypes.Component)`

#### method ComponentArray SetEmpty
Type: `func(Component)`

### type SetEvent
Type: `engine/modules/ecs.SetEvent`

#### property SetEvent Entity
Type: `engine/modules/ecs.EntityID`

#### property SetEvent Component
Type: `engine/modules/ecs.Component`

## Functions
### func NewSetEvent
Type: `func(entity engine/modules/ecs.EntityID, comp engine/modules/ecs.Component) engine/modules/ecs.SetEvent`

### func NewSystemRegister
Type: `func(l func() error) engine/modules/ecs.SystemRegister`

### func RegisterSystems
Type: `func(systems ...engine/modules/ecs.SystemRegister) []error`

### func ComponentComparator
Type: `func[Component any]() func(c1 Component, c2 Component) bool`

### func NewDirtySet
Type: `func() engine/modules/ecs.DirtySet`

### func NewWorld
Type: `func() engine/modules/ecs.World`

### func GetComponentArray
Type: `func[Component any](world engine/modules/ecs.World) engine/modules/ecs.ComponentArray[Component]`
component array getter


## Dependencies
`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSet`
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewSparseSet`
  - `engine/modules/datastructures.Set`
  - `engine/modules/datastructures.SparseArray`
  - `engine/modules/datastructures.SparseSet`

`engine/modules/ecs`:
  - `engine/modules/ecs.Component`
  - `engine/modules/ecs.Entity`
  - `engine/modules/ecs.NewWorld`
  - `engine/modules/ecs.SetEvent`
  - `engine/modules/ecs.World`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`