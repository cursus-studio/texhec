# deploy
## Architecture
defines GUI for deploying objects

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             36             14            200
-------------------------------------------------------------------------------
SUM:                             3             36             14            200
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/deploy.Service`

#### method Service Component
Type: `func() engine/services/ecs.ComponentsArray[core/modules/deploy.Component]`

#### method Service Deploy
Type: `func(blueprint engine/services/ecs.EntityID, owner engine/services/ecs.EntityID, coords engine/modules/grid.Coords) (engine/services/ecs.EntityID, error)`
deploy differs from execute event by who deploys.
execute adds costs and everything where deploy just deploys without any costs (its deployed by system)

#### method Service DeployEvent
Type: `func(core/modules/deploy.DeployEvent)`

#### method Service Reach
Type: `func() core/modules/reach.ServiceT[core/modules/deploy.Component]`

### type Component
Type: `core/modules/deploy.Component`

#### property Component Deployable
Type: `[]engine/services/ecs.EntityID`

### type DeployEvent
Type: `core/modules/deploy.DeployEvent`

#### property DeployEvent By
Type: `engine/services/ecs.EntityID`

#### property DeployEvent Blueprint
Type: `engine/services/ecs.EntityID`

#### property DeployEvent Coords
Type: `engine/modules/grid.Coords`

## Functions
### func NewDeploy
Type: `func(deployable ...engine/services/ecs.EntityID) core/modules/deploy.Component`

### func NewFeatureDeployEvent
Type: `func() engine/modules/interactions.FeatureEvent[core/modules/deploy.DeployEvent]`

### func NewDeployEvent
Type: `func(by engine/services/ecs.EntityID, blueprint engine/services/ecs.EntityID, coords engine/modules/grid.Coords) core/modules/deploy.DeployEvent`


## Dependencies
`core/game`:
  - `core/game.Deploy`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Player`
  - `core/game.Reach`
  - `core/game.Tile`

`core/modules/deploy`:
  - `core/modules/deploy.Blueprint`
  - `core/modules/deploy.By`
  - `core/modules/deploy.Component`
  - `core/modules/deploy.Coords`
  - `core/modules/deploy.DeployEvent`
  - `core/modules/deploy.Reach`
  - `core/modules/deploy.Service`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Deployed`
  - `core/modules/obstruction.ErrPositionIsOccupied`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.NewDeployed`
  - `core/modules/obstruction.Obstruction`

`core/modules/player`:
  - `core/modules/player.NewOwner`
  - `core/modules/player.Owner`

`core/modules/reach`:
  - `core/modules/reach.Component`
  - `core/modules/reach.Distance`
  - `core/modules/reach.ErrOutsideOfReach`
  - `core/modules/reach.NewReach`
  - `core/modules/reach.Reach`
  - `core/modules/reach.ServiceT`

`core/modules/reach/pkg`:
  - `core/modules/reach/pkg.PkgT`

`core/modules/tile`:
  - `core/modules/tile.Coord`
  - `core/modules/tile.Coords`
  - `core/modules/tile.CoordsInteraction`
  - `core/modules/tile.Entity`
  - `core/modules/tile.NewClickEntityEvent`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.ObjectInteraction`
  - `core/modules/tile.Pos`
  - `core/modules/tile.Size`
  - `core/modules/tile.SourceObjectInteraction`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/grid`:
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`

`engine/modules/inputs`:
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

`engine/modules/interactions`:
  - `engine/modules/interactions.FeatureEntity`
  - `engine/modules/interactions.FeatureEvent`
  - `engine/modules/interactions.Interaction`
  - `engine/modules/interactions.NewFeatureEvent`
  - `engine/modules/interactions.State`

`engine/modules/interactions/pkg`:
  - `engine/modules/interactions/pkg.FeaturePkg`

`engine/modules/seed`:
  - `engine/modules/seed.ErrWorldCanHaveOneSeed`
  - `engine/modules/seed.WorldSeed`

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SetEmpty`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`