# hierarchy
## Architecture
defines child-parent relationship.
this is one of core modules on which relies most of the engine.

service stores separate relation cache and updates it on changes to the hierarchy.
this allows us to have O(1) access time to children and parents

## Benchmarks
```
$ go test ./... -bench=.
goos: linux
goarch: amd64
pkg: engine/modules/hierarchy/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkChildren_1-8                          	223990726	         5.403 ns/op
BenchmarkChildren_10-8                         	220043870	         5.389 ns/op
BenchmarkChildren_100-8                        	224304896	         5.352 ns/op
BenchmarkFlatChildren_1_1-8                    	178237136	         6.736 ns/op
BenchmarkFlatChildren_10_10-8                  	185739192	         6.464 ns/op
BenchmarkAddChildToParentWithGrandParent-8     	 6815106	       162.5 ns/op
BenchmarkAddChildToParentWith5GrandParents-8   	 7389226	       149.9 ns/op
BenchmarkRemoveChild-8                         	 1165906	      1030 ns/op
BenchmarkRemoveParentWith1Children-8           	   64238	     17186 ns/op
BenchmarkRemoveParentWith100Children-8         	    9790	    136048 ns/op
PASS
ok  	engine/modules/hierarchy/test	15.219s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               9            127             22            587
-------------------------------------------------------------------------------
SUM:                             9            127             22            587
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/hierarchy.Service`

#### method Service Children
Type: `func(parent engine/modules/ecs.EntityID) engine/modules/datastructures.SparseSetReader[engine/modules/ecs.EntityID]`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/hierarchy.Component]`

#### method Service FlatChildren
Type: `func(parent engine/modules/ecs.EntityID) engine/modules/datastructures.SparseSetReader[engine/modules/ecs.EntityID]`
includes children of children

#### method Service GetOrderedParents
Type: `func(child engine/modules/ecs.EntityID) []engine/modules/ecs.EntityID`

#### method Service GetParents
Type: `func(child engine/modules/ecs.EntityID) engine/modules/datastructures.SparseSetReader[engine/modules/ecs.EntityID]`
from closest to furthest

#### method Service IsChildOf
Type: `func(child engine/modules/ecs.EntityID, parent engine/modules/ecs.EntityID) bool`
returns true if is child of any parent doesn't matter the depth

#### method Service Parent
Type: `func(child engine/modules/ecs.EntityID) (engine/modules/ecs.EntityID, bool)`

#### method Service SetChildren
Type: `func(parent engine/modules/ecs.EntityID, children ...engine/modules/ecs.EntityID)`
maintains order of children and adds component to children
even if children doesn't exist

#### method Service SetParent
Type: `func(child engine/modules/ecs.EntityID, parent engine/modules/ecs.EntityID)`

### type ServiceT
Type: `engine/modules/hierarchy.ServiceT[Component any]`

#### method ServiceT Inherit
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/hierarchy.InheritComponent[Component]]`

### type Component
Type: `engine/modules/hierarchy.Component`

#### property Component Parent
Type: `engine/modules/ecs.EntityID`

### type InheritComponent
Type: `engine/modules/hierarchy.InheritComponent[Component any]`

## Variables
### var ErrParentCycle
Type: `error`

## Functions
### func NewParent
Type: `func(parent engine/modules/ecs.EntityID) engine/modules/hierarchy.Component`

### func NewInherit
Type: `func[Component any]() engine/modules/hierarchy.InheritComponent[Component]`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Hierarchy`
  - `engine.World`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewSparseSetWithPaging`
  - `engine/modules/datastructures.SparseArray`
  - `engine/modules/datastructures.SparseSet`
  - `engine/modules/datastructures.SparseSetReader`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.ComponentComparator`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.World`

`engine/modules/hierarchy`:
  - `engine/modules/hierarchy.Children`
  - `engine/modules/hierarchy.Component`
  - `engine/modules/hierarchy.InheritComponent`
  - `engine/modules/hierarchy.NewParent`
  - `engine/modules/hierarchy.Parent`
  - `engine/modules/hierarchy.Service`
  - `engine/modules/hierarchy.ServiceT`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/ioc/v2`