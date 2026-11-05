# reach
## Architecture
is responsible for managing object features range

## Benchmarks
```
$ go test ./... -bench=.
goos: linux
goarch: amd64
pkg: core/modules/reach/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkDist-8                 	120405490	         9.943 ns/op
Benchmark4TilesWithinReach-8    	12106434	        98.63 ns/op
Benchmark12TilesWithinReach-8   	 8124784	       142.1 ns/op
PASS
ok  	core/modules/reach/test	3.575s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8             60             27            382
Markdown                         2              0              0              2
-------------------------------------------------------------------------------
SUM:                            10             60             27            384
-------------------------------------------------------------------------------
```
## TODO
Define circle shader for reach

## Dependencies
`core/game`:
  - `core/game.GameWorld`
  - `core/game.Reach`
  - `core/game.Tile`

`core/modules/reach`:
  - `core/modules/reach.Distance`
  - `core/modules/reach.NewReach`
  - `core/modules/reach.Reach`
  - `core/modules/reach.ReachComponent`
  - `core/modules/reach.Service`
  - `core/modules/reach.ServiceT`
  - `core/modules/reach.TilesWithinReach`

`core/modules/reach/pkg`:
  - `core/modules/reach/pkg.PkgT`

`core/modules/tile`:
  - `core/modules/tile.Aligned`
  - `core/modules/tile.Blueprint`
  - `core/modules/tile.Coord`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`

`core/pkg`:
  - `core/pkg.Pkg`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/grid`:
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/ioc/v2`