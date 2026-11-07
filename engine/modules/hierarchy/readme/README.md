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
BenchmarkChildren_1-8                          	222155427	         5.443 ns/op
BenchmarkChildren_10-8                         	219609798	         5.464 ns/op
BenchmarkChildren_100-8                        	222874813	         5.380 ns/op
BenchmarkFlatChildren_1_1-8                    	182304591	         6.585 ns/op
BenchmarkFlatChildren_10_10-8                  	183836434	         6.518 ns/op
BenchmarkAddChildToParentWithGrandParent-8     	 7330680	       160.6 ns/op
BenchmarkAddChildToParentWith5GrandParents-8   	 7226611	       164.7 ns/op
BenchmarkRemoveChild-8                         	  976084	      1141 ns/op
BenchmarkRemoveParentWith1Children-8           	   63080	     17757 ns/op
BenchmarkRemoveParentWith100Children-8         	    8671	    139983 ns/op
PASS
ok  	engine/modules/hierarchy/test	14.038s
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
  - `engine/modules/hierarchy.GetParent`
  - `engine/modules/hierarchy.InheritComponent`
  - `engine/modules/hierarchy.NewParent`
  - `engine/modules/hierarchy.Parent`
  - `engine/modules/hierarchy.ParentComponent`
  - `engine/modules/hierarchy.Service`
  - `engine/modules/hierarchy.ServiceT`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/ioc/v2`