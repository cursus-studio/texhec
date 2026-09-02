# actions
## Architecture
Stores all game domain interactions and steps

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             86             13            467
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             7             86             13            468
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/actions.Service`

#### method Service Anchor
Type: `func() engine/modules/ecs.ComponentArray[core/modules/actions.AnchorComponent]`

#### method Service BlueprintInteraction
Type: `func() engine/modules/interactions.InteractionService[core/modules/actions.BlueprintInteraction]`

#### method Service CanDeploy
Type: `func() engine/modules/ecs.ComponentArray[core/modules/actions.CanDeployComponent]`

#### method Service CoordsCursor
Type: `func() engine/modules/ecs.ComponentArray[core/modules/actions.CoordsCursorComponent]`

#### method Service CoordsInteraction
Type: `func() engine/modules/interactions.InteractionService[core/modules/actions.CoordsInteraction]`

#### method Service EntityInteraction
Type: `func() engine/modules/interactions.InteractionService[core/modules/actions.EntityInteraction]`

#### method Service RegionAnchor
Type: `func() engine/modules/ecs.ComponentArray[core/modules/actions.RegionAnchorComponent]`

### type CoordsStep
Type: `core/modules/actions.CoordsStep`

#### method CoordsStep State
Type: `func() core/modules/actions.CoordsInteraction`

### type EntityStep
Type: `core/modules/actions.EntityStep`

#### method EntityStep State
Type: `func() core/modules/actions.EntityInteraction`

### type FriendlyEntityStep
Type: `core/modules/actions.FriendlyEntityStep`

#### method FriendlyEntityStep State
Type: `func() core/modules/actions.EntityInteraction`

### type FriendlyMobileEntityStep
Type: `core/modules/actions.FriendlyMobileEntityStep`

#### method FriendlyMobileEntityStep State
Type: `func() core/modules/actions.EntityInteraction`

### type FriendlyBuilderEntityStep
Type: `core/modules/actions.FriendlyBuilderEntityStep`

#### method FriendlyBuilderEntityStep State
Type: `func() core/modules/actions.EntityInteraction`

### type FriendlyOffensiveEntityStep
Type: `core/modules/actions.FriendlyOffensiveEntityStep`

#### method FriendlyOffensiveEntityStep State
Type: `func() core/modules/actions.EntityInteraction`

### type EnemyEntityStep
Type: `core/modules/actions.EnemyEntityStep`

#### method EnemyEntityStep State
Type: `func() core/modules/actions.EntityInteraction`

### type BlueprintStep
Type: `core/modules/actions.BlueprintStep`

#### method BlueprintStep State
Type: `func() core/modules/actions.BlueprintInteraction`

### type CanDeployComponent
Type: `core/modules/actions.CanDeployComponent`
components to configure interactions

#### property CanDeployComponent Entity
Type: `engine/modules/ecs.EntityID`

### type CoordsCursorComponent
Type: `core/modules/actions.CoordsCursorComponent`

#### property CoordsCursorComponent PropertiesEntity
Type: `engine/modules/ecs.EntityID`

#### property CoordsCursorComponent CustomImage
Type: `bool`
if true then entity is used as an image else default icon is used

### type AnchorComponent
Type: `core/modules/actions.AnchorComponent`

#### property AnchorComponent Entity
Type: `engine/modules/ecs.EntityID`

### type RegionAnchorComponent
Type: `core/modules/actions.RegionAnchorComponent`

#### property RegionAnchorComponent Region
Type: `core/modules/pathfind.Region`

### type CoordsInteraction
Type: `core/modules/actions.CoordsInteraction`

#### property CoordsInteraction Coords
Type: `engine/modules/grid.Coords`

### type EntityInteraction
Type: `core/modules/actions.EntityInteraction`

#### property EntityInteraction Entity
Type: `engine/modules/ecs.EntityID`

### type BlueprintInteraction
Type: `core/modules/actions.BlueprintInteraction`

#### property BlueprintInteraction Entity
Type: `engine/modules/ecs.EntityID`

## Variables
### var ErrRequiresSpeed
Type: `error`

### var ErrRequiresDeploy
Type: `error`

### var ErrRequiresAttack
Type: `error`

## Functions
### func NewCanDeploy
Type: `func(canDeploy engine/modules/ecs.EntityID) core/modules/actions.CanDeployComponent`

### func NewCoordsCursor
Type: `func(propertiesEntity engine/modules/ecs.EntityID, customImage bool) core/modules/actions.CoordsCursorComponent`

### func NewAnchor
Type: `func(entity engine/modules/ecs.EntityID) core/modules/actions.AnchorComponent`

### func NewRegionAnchor
Type: `func(region core/modules/pathfind.Region) core/modules/actions.RegionAnchorComponent`

### func NewCoordsInteraction
Type: `func(coords engine/modules/grid.Coords) core/modules/actions.CoordsInteraction`

### func NewEntityInteraction
Type: `func(entity engine/modules/ecs.EntityID) core/modules/actions.EntityInteraction`

### func NewBlueprintInteraction
Type: `func(entity engine/modules/ecs.EntityID) core/modules/actions.BlueprintInteraction`


## Dependencies
`core/game`:
  - `core/game.Attack`
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.EngineWorld`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Pathfind`
  - `core/game.Player`
  - `core/game.Tile`
  - `core/game.Ui`

`core/modules/actions`:
  - `core/modules/actions.AnchorComponent`
  - `core/modules/actions.BlueprintInteraction`
  - `core/modules/actions.BlueprintStep`
  - `core/modules/actions.CanDeployComponent`
  - `core/modules/actions.Coords`
  - `core/modules/actions.CoordsCursorComponent`
  - `core/modules/actions.CoordsInteraction`
  - `core/modules/actions.CoordsStep`
  - `core/modules/actions.CustomImage`
  - `core/modules/actions.EnemyEntityStep`
  - `core/modules/actions.Entity`
  - `core/modules/actions.EntityInteraction`
  - `core/modules/actions.EntityStep`
  - `core/modules/actions.ErrRequiresAttack`
  - `core/modules/actions.ErrRequiresDeploy`
  - `core/modules/actions.ErrRequiresSpeed`
  - `core/modules/actions.FriendlyBuilderEntityStep`
  - `core/modules/actions.FriendlyEntityStep`
  - `core/modules/actions.FriendlyMobileEntityStep`
  - `core/modules/actions.FriendlyOffensiveEntityStep`
  - `core/modules/actions.NewAnchor`
  - `core/modules/actions.NewBlueprintInteraction`
  - `core/modules/actions.NewCanDeploy`
  - `core/modules/actions.NewCoordsCursor`
  - `core/modules/actions.NewCoordsInteraction`
  - `core/modules/actions.NewEntityInteraction`
  - `core/modules/actions.NewRegionAnchor`
  - `core/modules/actions.PropertiesEntity`
  - `core/modules/actions.Region`
  - `core/modules/actions.RegionAnchorComponent`
  - `core/modules/actions.Service`

`core/modules/definitions`:
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.Blank`
  - `core/modules/definitions.Border`
  - `core/modules/definitions.Btn`
  - `core/modules/definitions.Can`
  - `core/modules/definitions.GameGroup`
  - `core/modules/definitions.Hud`
  - `core/modules/definitions.ObjectPlaceholderLayer`
  - `core/modules/definitions.ObjectSelectionPlaceholderLayer`
  - `core/modules/definitions.RangePlaceholderLayer`
  - `core/modules/definitions.Selected`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.TilePlaceholderLayer`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Deployed`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.Obstruction`

`core/modules/pathfind`:
  - `core/modules/pathfind.CoordsRegion`
  - `core/modules/pathfind.EntityRegion`
  - `core/modules/pathfind.Region`
  - `core/modules/pathfind.RegionObstruction`
  - `core/modules/pathfind.Speed`

`core/modules/player`:
  - `core/modules/player.ControlsObject`
  - `core/modules/player.ErrRequiresToBeEnemy`

`core/modules/tile`:
  - `core/modules/tile.Blueprint`
  - `core/modules/tile.ClickBlueprintEvent`
  - `core/modules/tile.ClickEntityEvent`
  - `core/modules/tile.Entity`
  - `core/modules/tile.ErrBlueprintIsMissingUUID`
  - `core/modules/tile.Layer`
  - `core/modules/tile.Name`
  - `core/modules/tile.NewClickBlueprintEvent`
  - `core/modules/tile.NewLayer`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.Pos`
  - `core/modules/tile.Size`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/grid`:
  - `engine/modules/grid.AbsoluteCoords`
  - `engine/modules/grid.Chunk`
  - `engine/modules/grid.ClickEvent`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.HoverEvent`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`
  - `engine/modules/groups.InheritGroups`

`engine/modules/inputs`:
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

`engine/modules/interactions`:
  - `engine/modules/interactions.InteractionService`
  - `engine/modules/interactions.MissingPreview`
  - `engine/modules/interactions.Save`
  - `engine/modules/interactions.State`
  - `engine/modules/interactions.StatePreview`
  - `engine/modules/interactions.Step`

`engine/modules/interactions/pkg`:
  - `engine/modules/interactions/pkg.InteractionPkg`
  - `engine/modules/interactions/pkg.StepPkg`

`engine/modules/render`:
  - `engine/modules/render.Color`
  - `engine/modules/render.Mesh`
  - `engine/modules/render.NewColor`
  - `engine/modules/render.NewMesh`
  - `engine/modules/render.NewTexture`
  - `engine/modules/render.Texture`

`engine/modules/text`:
  - `engine/modules/text.Content`
  - `engine/modules/text.NewText`

`engine/modules/transform`:
  - `engine/modules/transform.Absolute`
  - `engine/modules/transform.Inherit`
  - `engine/modules/transform.NewInherit`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeSizeXY`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`