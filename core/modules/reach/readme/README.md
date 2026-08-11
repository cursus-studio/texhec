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
BenchmarkDist-8                 	100000000	        12.41 ns/op
Benchmark4TilesWithinReach-8    	11624863	        97.99 ns/op
Benchmark12TilesWithinReach-8   	 9029640	       131.6 ns/op
PASS
ok  	core/modules/reach/test	3.589s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8             58             27            373
Markdown                         2              0              0              2
-------------------------------------------------------------------------------
SUM:                            10             58             27            375
-------------------------------------------------------------------------------
```
## TODO
Define circle shader for reach

## Types
### type Service
Type: `core/modules/reach.Service`

#### method Service Distance
Type: `func(from core/modules/tile.PosComponent, fromSize core/modules/tile.SizeComponent, to core/modules/tile.PosComponent, toSize core/modules/tile.SizeComponent) core/modules/tile.Coord`
returns rounded up distance between nearest coordinates

### type ServiceT
Type: `core/modules/reach.ServiceT[FeatureComponent any]`

#### method ServiceT Component
Type: `func() engine/modules/ecs.ComponentArray[core/modules/reach.Component[FeatureComponent]]`

#### method ServiceT Reaches
Type: `func(fromEntity engine/modules/ecs.EntityID, toEntity engine/modules/ecs.EntityID) bool`

#### method ServiceT TilesFrom
Type: `func(core/modules/tile.PosComponent, core/modules/tile.SizeComponent, core/modules/reach.Component[FeatureComponent]) []engine/modules/grid.Coords`

#### method ServiceT TilesWithinReach
Type: `func(entity engine/modules/ecs.EntityID) []engine/modules/grid.Coords`

### type Component
Type: `core/modules/reach.Component[FeatureComponent any]`
stores reach distance squared (squared to avoid Sqrt)

#### property Component Reach
Type: `engine/modules/grid.Coord`

## Variables
### var ErrOutsideOfReach
Type: `error`

## Functions
### func NewReach
Type: `func[FeatureComponent any](reach engine/modules/grid.Coord) core/modules/reach.Component[FeatureComponent]`
takes square of distnace


## Dependencies
`core/game`:
  - `core/game.GameWorld`
  - `core/game.Reach`
  - `core/game.Tile`

`core/modules/reach`:
  - `core/modules/reach.Component`
  - `core/modules/reach.Distance`
  - `core/modules/reach.NewReach`
  - `core/modules/reach.Reach`
  - `core/modules/reach.Service`
  - `core/modules/reach.ServiceT`
  - `core/modules/reach.TilesWithinReach`

`core/modules/reach/pkg`:
  - `core/modules/reach/pkg.PkgT`

`core/modules/tile`:
  - `core/modules/tile.Aligned`
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