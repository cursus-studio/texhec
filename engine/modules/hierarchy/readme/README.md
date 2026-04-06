# Hierarchy
## Architecture
How module is built and general flow of data and why this way in case of controversial choices.

This module is built mainly with 2 main components.\
Public `hierarchy.Component` pointing to parent and internal `ParentComponent`.\
Relations are stored internaly and to use them we use `hierarhy.Service`.\
On `hierarchy.Component` modifications we modify internal state instantly.\
We should do this instantly to ensure children are always removed on parent removal and lazy reaction could do not add child before it was removed.\
Instant reaction complicates flow a lot and this has space for improvement.

## Benchmarks

```sh
$ go test . -bench=. -benchmem
goos: linux
goarch: amd64
pkg: engine/modules/hierarchy/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkChildren_1-8                          	211353626	         5.708 ns/op	       0 B/op	       0 allocs/op
BenchmarkChildren_10-8                         	204640497	         5.825 ns/op	       0 B/op	       0 allocs/op
BenchmarkChildren_100-8                        	211510869	         5.654 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlatChildren_1_1-8                    	180521926	         6.592 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlatChildren_10_10-8                  	186883720	         6.427 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddChildToParentWithGrandParent-8     	11669298	        95.11 ns/op	     198 B/op	       0 allocs/op
BenchmarkAddChildToParentWith5GrandParents-8   	10917055	       101.0 ns/op	     211 B/op	       0 allocs/op
BenchmarkRemoveChild-8                         	 6898168	       162.1 ns/op	      50 B/op	       1 allocs/op
BenchmarkRemoveParentWith1Children-8           	  188970	      6325 ns/op	   16488 B/op	       5 allocs/op
BenchmarkRemoveParentWith100Children-8         	   38294	     30320 ns/op	   17496 B/op	      11 allocs/op
```

## Usage examples
### Service interface
Service methods are self explainatory.
```go
type Service interface {
	Component() ecs.ComponentsArray[Component]

	// returns true if is child of any parent doesn't matter the depth
	IsChildOf(child ecs.EntityID, parent ecs.EntityID) bool
	SetParent(child ecs.EntityID, parent ecs.EntityID)
	Parent(child ecs.EntityID) (ecs.EntityID, bool)

	// from closest to furthest
	GetParents(child ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID]
	GetOrderedParents(child ecs.EntityID) []ecs.EntityID

	// maintains order of children and adds component to children
	// even if children doesn't exist
	SetChildren(parent ecs.EntityID, children ...ecs.EntityID)

	Children(parent ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID]
	// includes children of children
	FlatChildren(parent ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID]
}
```

### `IsChildOf`
Returns `true` if child is child of parent.
```go
func _(hierarchy hierarchy.Service, child, parent ecs.EntityID) {
    if hierarchy.IsChildOf(child, parent) {
        // do something
    }
}
```

### `SetParent`
Adds `hierarchy.Component` to a child and doing so sets child-parent relation
```go
func _(hierarchy hierarchy.Service, child, parent ecs.EntityID) {
    hierarchy.SetParent(child, parent)
}
```

### `Parent`
Reads `hierarchy.Component` from a child and returns parent and ok
```go
func _(hierarchy hierarchy.Service, child ecs.EntityID) {
    parent, ok := hierarchy.Parent(child)
}
```

### `GetParents`
Returns parents in `SparseSetReader`.
```go
func _(hierarchy hierarchy.Service, child, parent ecs.EntityID) {
    parents, ok := hierarchy.GetParents(child)
    isRelated := parents.Get(parent)
}
```

### `GetOrderedParents`
Returns parents in `slice` in order.\
Its separate from `GetParents` because `SparseSetReader` doesn't maintain order.
```go
func _(hierarchy hierarchy.Service, child ecs.EntityID) {
    for _, parent := range hierarchy.GetOrderedParents(child) {
    }
}
```

### `SetChildren`
Sets children for each child parent as parent
```go
func _(hierarchy hierarchy.Service, parent, child1, child2 ecs.EntityID) {
    hierarchy.SetChildren(parent, child1, child2)
}
```

### `Children`
Reads parent children.
```go
func _(hierarchy hierarchy.Service, parent ecs.EntityID) {
    children := hierarchy.Children(parent)
}
```

### `FlatChildren`
Reads parent children and grand children.
```go
func _(hierarchy hierarchy.Service, parent ecs.EntityID) {
    flatChildren := hierarchy.FlatChildren(parent)
}
```

## Dependencies
- [codec](/engine/services/codec/readme/README.md)
- [datastructures](/engine/services/datastructures/readme/README.md)
- [ecs](/engine/services/ecs/readme/README.md)
- [logger](/engine/services/logger/readme/README.md)
