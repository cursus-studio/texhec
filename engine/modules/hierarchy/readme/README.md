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
BenchmarkChildren_1-8                          	215221099	         5.477 ns/op
BenchmarkChildren_10-8                         	206531972	         5.733 ns/op
BenchmarkChildren_100-8                        	221496177	         5.477 ns/op
BenchmarkFlatChildren_1_1-8                    	178174026	         6.668 ns/op
BenchmarkFlatChildren_10_10-8                  	172217704	         6.895 ns/op
BenchmarkAddChildToParentWithGrandParent-8     	 7234621	       170.8 ns/op
BenchmarkAddChildToParentWith5GrandParents-8   	 6127701	       164.6 ns/op
BenchmarkRemoveChild-8                         	 1536834	       842.3 ns/op
BenchmarkRemoveParentWith1Children-8           	   82946	     14288 ns/op
BenchmarkRemoveParentWith100Children-8         	   10000	    116548 ns/op
PASS
ok  	engine/modules/hierarchy/test	15.862s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7            103             21            483
-------------------------------------------------------------------------------
SUM:                             7            103             21            483
-------------------------------------------------------------------------------
```
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

### type Component
Type: `engine/modules/hierarchy.Component`

#### property Component Parent
Type: `engine/services/ecs.EntityID`

## Variables
### var ErrParentCycle
Type: `error`

## Functions
### func NewParent
Type: `func(parent engine/services/ecs.EntityID) engine/modules/hierarchy.Component`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.World`

`engine/modules/hierarchy`:
  - `engine/modules/hierarchy.Component`
  - `engine/modules/hierarchy.NewParent`
  - `engine/modules/hierarchy.Parent`
  - `engine/modules/hierarchy.Service`

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
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.OnRemove`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.RemoveEntity`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.World`

### Third Party
- `github.com/ogiusek/ioc/v2`