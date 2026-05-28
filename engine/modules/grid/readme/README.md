# grid
## Architecture
defines generic slice based data structure to store grids
it implements unified chunk size

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               5             71             29            338
-------------------------------------------------------------------------------
SUM:                             5             71             29            338
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/grid.Service`

#### method Service AbsoluteCoords
Type: `func(engine/modules/grid.ChunkCoordsComponent, engine/modules/grid.Coords) engine/modules/grid.Coords`
calculate chunk coords

#### method Service ChunkSize
Type: `func() engine/modules/grid.Coord`
getters within chunk

#### method Service Coords
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/grid.ChunkCoordsComponent]`
arrays

#### method Service CoordsIndex
Type: `func(engine/modules/grid.Coords) (engine/modules/grid.Index, bool)`

#### method Service GetChunk
Type: `func(engine/modules/grid.ChunkCoordsComponent) (engine/services/ecs.EntityID, bool)`

#### method Service GetLastIndex
Type: `func() engine/modules/grid.Index`

#### method Service IndexCoords
Type: `func(index engine/modules/grid.Index) engine/modules/grid.Coords`

#### method Service RelativeCoords
Type: `func(engine/modules/grid.Coords) (engine/modules/grid.ChunkCoordsComponent, engine/modules/grid.Coords)`

### type TileConstraint
Type: `engine/modules/grid.TileConstraint`

### type ServiceT
Type: `engine/modules/grid.ServiceT[Tile engine/modules/grid.TileConstraint]`

#### method ServiceT Chunk
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/grid.ChunkComponent[Tile]]`
arrays

#### method ServiceT CoordsData
Type: `func(engine/modules/grid.Coords) (engine/modules/grid.CoordsData[Tile], bool)`
calculate chunk coords

#### method ServiceT NewChunk
Type: `func() engine/modules/grid.ChunkComponent[Tile]`
ctors

### type ChunkSize
Type: `engine/modules/grid.ChunkSize`

#### method ChunkSize Val
Type: `func() engine/modules/grid.Coord`

### type Coord
Type: `engine/modules/grid.Coord`

### type Coords
Type: `engine/modules/grid.Coords`

#### property Coords X
Type: `engine/modules/grid.Coord`

#### property Coords Y
Type: `engine/modules/grid.Coord`

#### method Coords Size
Type: `func() (engine/modules/grid.Coord, engine/modules/grid.Coord)`

#### method Coords Coords
Type: `func() (X engine/modules/grid.Coord, Y engine/modules/grid.Coord)`

### type Index
Type: `engine/modules/grid.Index`

### type ChunkComponent
Type: `engine/modules/grid.ChunkComponent[Tile engine/modules/grid.TileConstraint]`

#### method ChunkComponent GetTiles
Type: `func() []Tile`

#### method ChunkComponent GetTile
Type: `func(index engine/modules/grid.Index) Tile`

#### method ChunkComponent SetTile
Type: `func(index engine/modules/grid.Index, tile Tile)`

### type ChunkCoordsComponent
Type: `engine/modules/grid.ChunkCoordsComponent`

#### property ChunkCoordsComponent X
Type: `engine/modules/grid.Coord`

#### property ChunkCoordsComponent Y
Type: `engine/modules/grid.Coord`

### type CoordsData
Type: `engine/modules/grid.CoordsData[Tile engine/modules/grid.TileConstraint]`
stores coords chunk data

#### property CoordsData Entity
Type: `engine/services/ecs.EntityID`

#### property CoordsData Component
Type: `engine/modules/grid.ChunkComponent[Tile]`

#### property CoordsData Index
Type: `engine/modules/grid.Index`

## Functions
### func NewChunkSize
Type: `func(size uint8) engine/modules/grid.ChunkSize`
cannot use number bigger then 31 because this would overflow Coord (uint32)

### func NewCoord
Type: `func[Num golang.org/x/exp/constraints.Integer](n Num) engine/modules/grid.Coord`
allows to create negative coords

### func NewCoords
Type: `func[Number golang.org/x/exp/constraints.Integer](x Number, y Number) engine/modules/grid.Coords`

### func NewChunk
Type: `func[Tile engine/modules/grid.TileConstraint](s engine/modules/grid.Coord) engine/modules/grid.ChunkComponent[Tile]`

### func NewChunkCoords
Type: `func(x engine/modules/grid.Coord, y engine/modules/grid.Coord) engine/modules/grid.ChunkCoordsComponent`

### func NewCoordsData
Type: `func[Tile engine/modules/grid.TileConstraint](entity engine/services/ecs.EntityID, component engine/modules/grid.ChunkComponent[Tile], index engine/modules/grid.Index) engine/modules/grid.CoordsData[Tile]`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Grid`
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
  - `engine/modules/grid.Chunk`
  - `engine/modules/grid.ChunkComponent`
  - `engine/modules/grid.ChunkCoordsComponent`
  - `engine/modules/grid.ChunkSize`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.CoordsData`
  - `engine/modules/grid.CoordsIndex`
  - `engine/modules/grid.GetChunk`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.NewChunk`
  - `engine/modules/grid.NewChunkCoords`
  - `engine/modules/grid.NewChunkSize`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.NewCoordsData`
  - `engine/modules/grid.RelativeCoords`
  - `engine/modules/grid.Service`
  - `engine/modules/grid.ServiceT`
  - `engine/modules/grid.TileConstraint`
  - `engine/modules/grid.Val`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/inputs`:
  - `engine/modules/inputs.EventTargetSetter`
  - `engine/modules/inputs.Hover`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewHoverComponent`
  - `engine/modules/inputs.NewLeftClick`
  - `engine/modules/inputs.ObjectRayCollision`
  - `engine/modules/inputs.Target`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.MapRelationPkg`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/ecs`:
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.World`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`