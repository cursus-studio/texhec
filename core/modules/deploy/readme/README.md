# deploy
## Architecture
defines GUI for deploying objects

## Types
### type Service
Type: `core/modules/deploy.Service`

#### method Service Component
Type: `func() engine/services/ecs.ComponentsArray[core/modules/deploy.Component]`

#### method Service Deploy
Type: `func(blueprint engine/services/ecs.EntityID, owner engine/services/ecs.EntityID, coords engine/modules/grid.Coords) (engine/services/ecs.EntityID, error)`
deploy differs from execute event by who deploys.
execute adds costs and everything where deploy just deploys without any costs (its deployed by system)

#### method Service Execute
Type: `func(core/modules/deploy.ExecuteEvent)`

#### method Service Preview
Type: `func(core/modules/deploy.PreviewEvent)`

#### method Service Select
Type: `func(core/modules/deploy.SelectEvent)`

### type Component
Type: `core/modules/deploy.Component`

#### property Component Deployable
Type: `[]engine/services/ecs.EntityID`

### type SelectEvent
Type: `core/modules/deploy.SelectEvent`
Select unit.
Add in gui some indicator.
Change on click event.

#### property SelectEvent By
Type: `engine/services/ecs.EntityID`

#### property SelectEvent Blueprint
Type: `engine/services/ecs.EntityID`

### type PreviewEvent
Type: `core/modules/deploy.PreviewEvent`
Select unit.
Add in gui some indicator.
Perform all checks and costs

#### property PreviewEvent By
Type: `engine/services/ecs.EntityID`

#### property PreviewEvent Blueprint
Type: `engine/services/ecs.EntityID`

#### property PreviewEvent Coords
Type: `engine/modules/grid.Coords`

#### method PreviewEvent ApplyCoords
Type: `func(coords engine/modules/grid.Coords) any`

### type ExecuteEvent
Type: `core/modules/deploy.ExecuteEvent`
Deploys on coords something if it doesn't collide

#### property ExecuteEvent By
Type: `engine/services/ecs.EntityID`

#### property ExecuteEvent Blueprint
Type: `engine/services/ecs.EntityID`

#### property ExecuteEvent Coords
Type: `engine/modules/grid.Coords`

#### method ExecuteEvent ApplyCoords
Type: `func(coords engine/modules/grid.Coords) any`

## Functions
### func NewDeploy
Type: `func(deployable ...engine/services/ecs.EntityID) core/modules/deploy.Component`

### func NewSelectEvent
Type: `func(by engine/services/ecs.EntityID, blueprint engine/services/ecs.EntityID) core/modules/deploy.SelectEvent`

### func NewPreviewEvent
Type: `func(by engine/services/ecs.EntityID, blueprint engine/services/ecs.EntityID) core/modules/deploy.PreviewEvent`

### func NewExecuteEvent
Type: `func(by engine/services/ecs.EntityID, blueprint engine/services/ecs.EntityID) core/modules/deploy.ExecuteEvent`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Player`
  - `core/game.Tile`
  - `core/game.Ui`

`core/modules/definitions`:
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.Blank`
  - `core/modules/definitions.GameGroup`
  - `core/modules/definitions.ObjectPlaceholderLayer`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.TilePlaceholderLayer`

`core/modules/deploy`:
  - `core/modules/deploy.ApplyCoords`
  - `core/modules/deploy.Blueprint`
  - `core/modules/deploy.By`
  - `core/modules/deploy.Component`
  - `core/modules/deploy.Coords`
  - `core/modules/deploy.ExecuteEvent`
  - `core/modules/deploy.NewExecuteEvent`
  - `core/modules/deploy.NewPreviewEvent`
  - `core/modules/deploy.PreviewEvent`
  - `core/modules/deploy.SelectEvent`
  - `core/modules/deploy.Service`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Deployed`
  - `core/modules/obstruction.ErrPositionIsOccupied`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.NewDeployed`
  - `core/modules/obstruction.Obstruction`
  - `core/modules/obstruction.Tiles`

`core/modules/player`:
  - `core/modules/player.NewOwner`
  - `core/modules/player.Owner`

`core/modules/tile`:
  - `core/modules/tile.Layer`
  - `core/modules/tile.NewClickEntityEvent`
  - `core/modules/tile.NewLayer`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.NewSelectEvent`
  - `core/modules/tile.Pos`
  - `core/modules/tile.Size`

`core/modules/ui`:
  - `core/modules/ui.ActionComponent`
  - `core/modules/ui.Actions`
  - `core/modules/ui.NewUnselect`
  - `core/modules/ui.ObjectComponent`

`engine/modules/grid`:
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.GetIndex`
  - `engine/modules/grid.GetTile`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`
  - `engine/modules/groups.Ptr`
  - `engine/modules/groups.Val`

`engine/modules/inputs`:
  - `engine/modules/inputs.KeepSelected`
  - `engine/modules/inputs.KeepSelectedComponent`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

`engine/modules/render`:
  - `engine/modules/render.Color`
  - `engine/modules/render.Mesh`
  - `engine/modules/render.NewColor`
  - `engine/modules/render.NewMesh`
  - `engine/modules/render.NewTexture`
  - `engine/modules/render.Texture`

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.Set`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`