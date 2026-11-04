# pathfind
## Architecture
finds path on a grid and moves objects to their target according to their speed

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7            110             34            800
Markdown                         1              1              0              5
-------------------------------------------------------------------------------
SUM:                             8            111             34            805
-------------------------------------------------------------------------------
```
## TODO
Modify speed to allow moving many tiles per tick.
Modify pathfinding to do not use shortest route and instead use optimal path (optimize paths per chunks).

Research materials:
- `supreme commander`
- `multi agent pathfinding`

## Dependencies
`core/game`:
  - `core/game.Actions`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Pathfind`
  - `core/game.Tile`

`core/modules/actions`:
  - `core/modules/actions.CanMove`
  - `core/modules/actions.CanMoveTo`
  - `core/modules/actions.Coords`
  - `core/modules/actions.Entity`
  - `core/modules/actions.NewCanMoveRule`
  - `core/modules/actions.NewCanMoveToRule`
  - `core/modules/actions.NewOwnRule`
  - `core/modules/actions.PlayerContext`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.Obstrucion`
  - `core/modules/obstruction.Obstruction`
  - `core/modules/obstruction.ObstructionComponent`
  - `core/modules/obstruction.Obstructions`

`core/modules/pathfind`:
  - `core/modules/pathfind.CanStep`
  - `core/modules/pathfind.Coords`
  - `core/modules/pathfind.Entity`
  - `core/modules/pathfind.ErrInvalidPath`
  - `core/modules/pathfind.FindPathEvent`
  - `core/modules/pathfind.InvSpeed`
  - `core/modules/pathfind.NewSpeed`
  - `core/modules/pathfind.NewStep`
  - `core/modules/pathfind.NewTarget`
  - `core/modules/pathfind.PlayerContext`
  - `core/modules/pathfind.Region`
  - `core/modules/pathfind.Service`
  - `core/modules/pathfind.Speed`
  - `core/modules/pathfind.SpeedComponent`
  - `core/modules/pathfind.Step`
  - `core/modules/pathfind.StepComponent`
  - `core/modules/pathfind.Target`
  - `core/modules/pathfind.TargetComponent`

`core/modules/tile`:
  - `core/modules/tile.Aligned`
  - `core/modules/tile.Blueprint`
  - `core/modules/tile.Coord`
  - `core/modules/tile.Coords`
  - `core/modules/tile.ErrInvalidPosition`
  - `core/modules/tile.ErrInvalidStep`
  - `core/modules/tile.ErrPositionAndSpeedIsRequiredToStep`
  - `core/modules/tile.GetTile`
  - `core/modules/tile.Grid`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.NewRot`
  - `core/modules/tile.NewSize`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Rot`
  - `core/modules/tile.RotComponent`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.UnloadChunkEvent`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewSparseSet`
  - `engine/modules/datastructures.SparseArray`
  - `engine/modules/datastructures.SparseSet`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.SystemRegister`
  - `engine/modules/ecs.World`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/feature/pkg`:
  - `engine/modules/feature/pkg.Config`
  - `engine/modules/feature/pkg.RegisterFeature`
  - `engine/modules/feature/pkg.RegisterRuleToField`

`engine/modules/grid`:
  - `engine/modules/grid.AbsoluteCoords`
  - `engine/modules/grid.Chunk`
  - `engine/modules/grid.ChunkCoordsComponent`
  - `engine/modules/grid.Component`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.CoordsData`
  - `engine/modules/grid.CoordsIndex`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.GetTiles`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.IndexCoords`
  - `engine/modules/grid.NewChunk`
  - `engine/modules/grid.NewCoord`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.RelativeCoords`
  - `engine/modules/grid.ServiceT`
  - `engine/modules/grid.SetTile`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/grid/pkg`:
  - `engine/modules/grid/pkg.PkgT`

`engine/modules/loop`:
  - `engine/modules/loop.TickEvent`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.ComponentPkg`
  - `engine/modules/netsync/pkg.EventPkg`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.MapRelationPkg`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`