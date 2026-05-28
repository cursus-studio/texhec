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
BenchmarkChildren_1-8                          	223546735	         5.381 ns/op
BenchmarkChildren_10-8                         	217609297	         5.488 ns/op
BenchmarkChildren_100-8                        	223566892	         5.360 ns/op
BenchmarkFlatChildren_1_1-8                    	172945767	         6.968 ns/op
BenchmarkFlatChildren_10_10-8                  	179091206	         6.708 ns/op
BenchmarkAddChildToParentWithGrandParent-8     	 6948628	       161.4 ns/op
BenchmarkAddChildToParentWith5GrandParents-8   	 6925312	       162.8 ns/op
BenchmarkRemoveChild-8                         	 1547041	       760.1 ns/op
BenchmarkRemoveParentWith1Children-8           	   76017	     15405 ns/op
BenchmarkRemoveParentWith100Children-8         	   10000	    109442 ns/op
PASS
ok  	engine/modules/hierarchy/test	15.437s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               9            127             22            587
Markdown                         1              0              0              2
-------------------------------------------------------------------------------
SUM:                            10            127             22            589
-------------------------------------------------------------------------------
```
## TODO
abstract `InheritGroups` to inherit allowing to inherit any component.
Try to reduce calls during removal.

## Types
### type Service
Type: `engine/modules/hierarchy.Service`

#### method Service Children
Type: `func(parent engine/services/ecs.EntityID) engine/services/datastructures.SparseSetReader[engine/services/ecs.EntityID]`

#### method Service Component
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/hierarchy.Component]`

#### method Service FlatChildren
Type: `func(parent engine/services/ecs.EntityID) engine/services/datastructures.SparseSetReader[engine/services/ecs.EntityID]`
includes children of children

#### method Service GetOrderedParents
Type: `func(child engine/services/ecs.EntityID) []engine/services/ecs.EntityID`

#### method Service GetParents
Type: `func(child engine/services/ecs.EntityID) engine/services/datastructures.SparseSetReader[engine/services/ecs.EntityID]`
from closest to furthest

#### method Service IsChildOf
Type: `func(child engine/services/ecs.EntityID, parent engine/services/ecs.EntityID) bool`
returns true if is child of any parent doesn't matter the depth

#### method Service Parent
Type: `func(child engine/services/ecs.EntityID) (engine/services/ecs.EntityID, bool)`

#### method Service SetChildren
Type: `func(parent engine/services/ecs.EntityID, children ...engine/services/ecs.EntityID)`
maintains order of children and adds component to children
even if children doesn't exist

#### method Service SetParent
Type: `func(child engine/services/ecs.EntityID, parent engine/services/ecs.EntityID)`

### type ServiceT
Type: `engine/modules/hierarchy.ServiceT[Component any]`

#### method ServiceT Inherit
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/hierarchy.InheritComponent[Component]]`

### type Component
Type: `engine/modules/hierarchy.Component`

#### property Component Parent
Type: `engine/services/ecs.EntityID`

### type InheritComponent
Type: `engine/modules/hierarchy.InheritComponent[Component any]`

## Variables
### var ErrParentCycle
Type: `error`

## Functions
### func NewParent
Type: `func(parent engine/services/ecs.EntityID) engine/modules/hierarchy.Component`

### func NewInherit
Type: `func[Component any]() engine/modules/hierarchy.InheritComponent[Component]`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Hierarchy`
  - `engine.World`

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

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndices`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.NewSparseSetWithPaging`
  - `engine/services/datastructures.Remove`
  - `engine/services/datastructures.Set`
  - `engine/services/datastructures.SparseArray`
  - `engine/services/datastructures.SparseSet`
  - `engine/services/datastructures.SparseSetReader`

`engine/services/ecs`:
  - `engine/services/ecs.AddDependency`
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.BeforeGet`
  - `engine/services/ecs.Clear`
  - `engine/services/ecs.ComponentComparator`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEmpty`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.OnRemove`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.RemoveEntity`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.World`

### Third Party
- `github.com/ogiusek/ioc/v2`