# reach
## Architecture
is responsible for managing object features range

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	core/modules/reach/test	0.007s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7             59             28            334
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             8             59             28            335
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/reach.Service`

#### method Service Distance
Type: `func(from core/modules/tile.PosComponent, fromSize core/modules/tile.SizeComponent, to core/modules/tile.PosComponent, toSize core/modules/tile.SizeComponent) core/modules/tile.Coord`
returns rounded up distance between nearest coordinates

### type ServiceT
Type: `core/modules/reach.ServiceT[FeatureComponent any]`

#### method ServiceT Component
Type: `func() engine/services/ecs.ComponentsArray[core/modules/reach.Component[FeatureComponent]]`

#### method ServiceT Reaches
Type: `func(fromEntity engine/services/ecs.EntityID, toEntity engine/services/ecs.EntityID) bool`

#### method ServiceT TilesWithinReach
Type: `func(entity engine/services/ecs.EntityID) []engine/modules/grid.Coords`

### type Component
Type: `core/modules/reach.Component[FeatureComponent any]`
stores reach distance squared (squared to avoid Sqrt)

#### property Component Reach
Type: `core/modules/tile.Coord`

## Functions
### func NewReach
Type: `func[FeatureComponent any](reach core/modules/tile.Coord) core/modules/reach.Component[FeatureComponent]`


## Dependencies
`core/game`:
  - `core/game.GameWorld`
  - `core/game.Reach`
  - `core/game.Tile`

`core/modules/reach`:
  - `core/modules/reach.Component`
  - `core/modules/reach.Distance`
  - `core/modules/reach.Reach`
  - `core/modules/reach.Service`
  - `core/modules/reach.ServiceT`

`core/modules/reach/pkg`:
  - `core/modules/reach/pkg.PkgT`

`core/modules/tile`:
  - `core/modules/tile.Coord`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`

`core/pkg`:
  - `core/pkg.Pkg`

`engine/modules/grid`:
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`

### Third Party
- `github.com/ogiusek/ioc/v2`