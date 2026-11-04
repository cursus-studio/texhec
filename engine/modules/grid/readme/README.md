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
Go                               5             66             31            332
-------------------------------------------------------------------------------
SUM:                             5             66             31            332
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Grid`
  - `engine.Inputs`
  - `engine.Transform`
  - `engine.UUID`
  - `engine.World`

`engine/modules/collider`:
  - `engine/modules/collider.Entity`
  - `engine/modules/collider.Hit`
  - `engine/modules/collider.ObjectRayCollision`
  - `engine/modules/collider.Point`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.World`

`engine/modules/grid`:
  - `engine/modules/grid.ChunkComponent`
  - `engine/modules/grid.ChunkCoordsComponent`
  - `engine/modules/grid.ChunkSize`
  - `engine/modules/grid.ClickEvent`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.CoordsData`
  - `engine/modules/grid.GetChunk`
  - `engine/modules/grid.HoverEvent`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.NewChunk`
  - `engine/modules/grid.NewChunkCoords`
  - `engine/modules/grid.NewChunkSize`
  - `engine/modules/grid.NewClickEvent`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.NewCoordsData`
  - `engine/modules/grid.NewHoverEvent`
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
  - `engine/modules/inputs.Stack`
  - `engine/modules/inputs.StackComponent`
  - `engine/modules/inputs.Target`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.MapRelationPkg`

`engine/modules/transform`:
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.Size`
  - `engine/modules/transform.SizeComponent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.NewDeterministicUUID`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`