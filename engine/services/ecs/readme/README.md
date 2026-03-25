# ECS
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

Active listeners might be discarded in the future.

### Why golang GC (garbage collector) isn't a problem
We follow **DOD** there for GC isn't laden with managing pointers because there are little pointers to manage.\
This makes golang a perfect candidate for this project because of high developer efficiency and low performance overhead.

### Architecture changes to revise in the future
Change architecture to:
- remove active listeners
- add entity mechanism to wait until entity is released in all systems
- call before get on all components

This would **simplify** codebase and would make it follow **DOD** more.
This would depracate `EnsureExists`, `OnUpsert`, `OnRemove`.

## Benchmarks
```sh
$ go test ./... -bench=. -benchmem
goos: linux
goarch: amd64
pkg: engine/services/ecs/tests
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkDirtySetDirty-8               	162088540	         7.127 ns/op	      23 B/op	       0 allocs/op
BenchmarkDirtySetDirtyInversed-8       	173312954	         6.332 ns/op	      21 B/op	       0 allocs/op
BenchmarkDirtySetGet-8                 	599886150	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkDirtySetDirtyAndGet-8         	242420068	         4.986 ns/op	       0 B/op	       0 allocs/op
BenchmarkDirtySetDirtyAnd1Get-8        	157544990	         8.248 ns/op	      23 B/op	       0 allocs/op
Benchmark4SavesWith7Systems-8          	26450721	        39.41 ns/op	       0 B/op	       0 allocs/op
Benchmark16SavesWith7Systems-8         	 7368591	       161.1 ns/op	       0 B/op	       0 allocs/op
Benchmark256SavesWith7Systems-8        	  487632	      2466 ns/op	       0 B/op	       0 allocs/op
Benchmark4096SavesWith7Systems-8       	   30714	     39386 ns/op	       0 B/op	       0 allocs/op
Benchmark16384SavesWith7Systems-8      	    7208	    155196 ns/op	       0 B/op	       0 allocs/op
Benchmark65536SavesWith7Systems-8      	    1920	    619240 ns/op	       0 B/op	       0 allocs/op
Benchmark262144SavesWith7Systems-8     	     480	   2435807 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetComponent-8                	75503276	        15.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkCreateComponents-8            	36927805	        30.03 ns/op	      82 B/op	       0 allocs/op
BenchmarkUpdateComponents-8            	100000000	        10.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveComponent-8             	76977810	        14.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveEntityWithComponent-8   	15177441	        77.81 ns/op	      41 B/op	       0 allocs/op
BenchmarkRemoveEntity-8                	48522978	        23.43 ns/op	      49 B/op	       0 allocs/op
```

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

## Dependencies
- [datastructures](/engine/services/datastructures/readme/README.md)
