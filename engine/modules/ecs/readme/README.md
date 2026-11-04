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
BenchmarkDirtySetDirty-8               	160960870	         7.432 ns/op
BenchmarkDirtySetDirtyInversed-8       	169464208	         6.810 ns/op
BenchmarkDirtySetGet-8                 	473988548	         2.533 ns/op
BenchmarkDirtySetDirtyAndGet-8         	174952015	         6.852 ns/op
BenchmarkDirtySetDirtyAnd1Get-8        	147201364	         8.143 ns/op
Benchmark4SavesWith7Systems-8          	28698441	        37.96 ns/op
Benchmark16SavesWith7Systems-8         	 7674924	       158.3 ns/op
Benchmark256SavesWith7Systems-8        	  494852	      2375 ns/op
Benchmark4096SavesWith7Systems-8       	   31630	     37944 ns/op
Benchmark16384SavesWith7Systems-8      	    7890	    151916 ns/op
Benchmark65536SavesWith7Systems-8      	    1976	    610258 ns/op
Benchmark262144SavesWith7Systems-8     	     493	   2443931 ns/op
BenchmarkGetComponent-8                	82264460	        14.59 ns/op
BenchmarkCreateComponents-8            	45491002	        26.08 ns/op
BenchmarkUpdateComponents-8            	100000000	        10.08 ns/op
BenchmarkRemoveComponent-8             	81012231	        15.05 ns/op
BenchmarkRemoveEntityWithComponent-8   	36389577	        33.11 ns/op
BenchmarkRemoveEntity-8                	60956052	        17.37 ns/op
PASS
ok  	engine/modules/ecs/test	36.592s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              13            226             46            960
Markdown                         1             44              0            157
-------------------------------------------------------------------------------
SUM:                            14            270             46           1117
-------------------------------------------------------------------------------
```
## Dependencies
`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSet`
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewSparseSet`
  - `engine/modules/datastructures.Set`
  - `engine/modules/datastructures.SparseArray`
  - `engine/modules/datastructures.SparseSet`

`engine/modules/ecs`:
  - `engine/modules/ecs.NewWorld`
  - `engine/modules/ecs.World`

### Third Party
- `github.com/ogiusek/ioc/v2`