# grid
## Architecture
defines generic slice based data structure to store grids

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               5             48              9            196
-------------------------------------------------------------------------------
SUM:                             5             48              9            196
-------------------------------------------------------------------------------

```
## Types
### type Service
Type: `engine/modules/grid.Service[Tile engine/modules/grid.TileConstraint]`

#### method Service Component
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/grid.SquareGridComponent[Tile]]`

### type TileConstraint
Type: `engine/modules/grid.TileConstraint`

### type Coord
Type: `engine/modules/grid.Coord`

### type Coords
Type: `engine/modules/grid.Coords`

#### property Coords X
Type: `engine/modules/grid.Coord`

#### property Coords Y
Type: `engine/modules/grid.Coord`

#### method Coords Coords
Type: `func() (X engine/modules/grid.Coord, Y engine/modules/grid.Coord)`

### type Index
Type: `engine/modules/grid.Index`

### type SquareGridComponent
Type: `engine/modules/grid.SquareGridComponent[Tile engine/modules/grid.TileConstraint]`

#### method SquareGridComponent Size
Type: `func() (engine/modules/grid.Coord, engine/modules/grid.Coord)`
getters for consts

#### method SquareGridComponent Width
Type: `func() engine/modules/grid.Coord`

#### method SquareGridComponent Height
Type: `func() engine/modules/grid.Coord`

#### method SquareGridComponent GetIndex
Type: `func(x engine/modules/grid.Coord, y engine/modules/grid.Coord) (engine/modules/grid.Index, bool)`
index and coord getters

#### method SquareGridComponent GetCoords
Type: `func(index engine/modules/grid.Index) engine/modules/grid.Coords`

#### method SquareGridComponent GetTiles
Type: `func() []Tile`

#### method SquareGridComponent GetLastIndex
Type: `func() engine/modules/grid.Index`

#### method SquareGridComponent GetTile
Type: `func(index engine/modules/grid.Index) Tile`
getters and setters tiles

#### method SquareGridComponent SetTile
Type: `func(index engine/modules/grid.Index, tile Tile)`

## Functions
### func NewCoords
Type: `func[Number golang.org/x/exp/constraints.Integer](x Number, y Number) engine/modules/grid.Coords`

### func NewSquareGrid
Type: `func[Tile engine/modules/grid.TileConstraint](w engine/modules/grid.Coord, h engine/modules/grid.Coord) engine/modules/grid.SquareGridComponent[Tile]`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Inputs`
  - `engine.World`

`engine/modules/collider`:
  - `engine/modules/collider.AddRayFallThroughPolicy`
  - `engine/modules/collider.Entity`
  - `engine/modules/collider.FallTroughPolicy`
  - `engine/modules/collider.Hit`
  - `engine/modules/collider.ObjectRayCollision`
  - `engine/modules/collider.Point`
  - `engine/modules/collider.Service`

`engine/modules/grid`:
  - `engine/modules/grid.Component`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.GetIndex`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.Height`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.Service`
  - `engine/modules/grid.SquareGridComponent`
  - `engine/modules/grid.TileConstraint`
  - `engine/modules/grid.Width`

`engine/modules/inputs`:
  - `engine/modules/inputs.EventTargetSetter`
  - `engine/modules/inputs.Hover`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewHoverComponent`
  - `engine/modules/inputs.NewLeftClick`
  - `engine/modules/inputs.ObjectRayCollision`
  - `engine/modules/inputs.Target`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.Set`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`