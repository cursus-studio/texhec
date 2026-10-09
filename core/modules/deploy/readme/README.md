# deploy
## Architecture
defines GUI for deploying objects

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             51             16            347
-------------------------------------------------------------------------------
SUM:                             3             51             16            347
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/deploy.Service`

#### method Service CanDeploy
Type: `func(core/modules/actions.Entity) error`

#### method Service CanDeployAt
Type: `func(core/modules/actions.Entity, core/modules/actions.Blueprint, core/modules/actions.Coords) error`

#### method Service CanDeployBlueprint
Type: `func(core/modules/actions.Entity, core/modules/actions.Blueprint) error`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[core/modules/deploy.Component]`

#### method Service Deploy
Type: `func(blueprint engine/modules/ecs.EntityID, owner engine/modules/ecs.EntityID, coords engine/modules/grid.Coords) (engine/modules/ecs.EntityID, error)`
deploy differs from execute event by who deploys.
execute adds costs and everything where deploy just deploys without any costs (its deployed by system)

#### method Service DeployEvent
Type: `func(core/modules/deploy.DeployEvent)`

#### method Service DestroyEvent
Type: `func(core/modules/deploy.DestroyEvent)`

#### method Service Reach
Type: `func() core/modules/reach.ServiceT[core/modules/deploy.Component]`

#### method Service Register
Type: `func() error`

### type Component
Type: `core/modules/deploy.Component`

#### property Component Deployable
Type: `[]engine/modules/ecs.EntityID`

### type DeployEvent
Type: `core/modules/deploy.DeployEvent`

#### property DeployEvent PlayerContext
Type: `core/modules/actions.PlayerContext`

#### property DeployEvent By
Type: `core/modules/actions.Entity`

#### property DeployEvent Blueprint
Type: `core/modules/actions.Blueprint`

#### property DeployEvent Coords
Type: `core/modules/actions.Coords`

### type DestroyEvent
Type: `core/modules/deploy.DestroyEvent`

#### property DestroyEvent PlayerContext
Type: `core/modules/actions.PlayerContext`

#### property DestroyEvent Entity
Type: `core/modules/actions.Entity`

## Functions
### func NewDeploy
Type: `func(deployable ...engine/modules/ecs.EntityID) core/modules/deploy.Component`

### func NewDeployEvent
Type: `func(ctx core/modules/actions.PlayerContext, by engine/modules/uuid.UUID, blueprint engine/modules/uuid.UUID, coords engine/modules/grid.Coords) core/modules/deploy.DeployEvent`

### func NewDestroyEvent
Type: `func(ctx core/modules/actions.PlayerContext, uuid engine/modules/uuid.UUID) core/modules/deploy.DestroyEvent`


## Dependencies
`core/game`:
  - `core/game.Actions`
  - `core/game.Deploy`
  - `core/game.Economy`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Player`
  - `core/game.Reach`
  - `core/game.Tile`

`core/modules/actions`:
  - `core/modules/actions.Blueprint`
  - `core/modules/actions.CanDeploy`
  - `core/modules/actions.CanDeployAt`
  - `core/modules/actions.CanDeployAtRule`
  - `core/modules/actions.CanDeployBlueprint`
  - `core/modules/actions.CanDeployBlueprintRule`
  - `core/modules/actions.CanDeployRule`
  - `core/modules/actions.Coords`
  - `core/modules/actions.Entity`
  - `core/modules/actions.Own`
  - `core/modules/actions.OwnRule`
  - `core/modules/actions.PlayerContext`

`core/modules/deploy`:
  - `core/modules/deploy.Blueprint`
  - `core/modules/deploy.By`
  - `core/modules/deploy.Component`
  - `core/modules/deploy.Coords`
  - `core/modules/deploy.DeployEvent`
  - `core/modules/deploy.Deployable`
  - `core/modules/deploy.DestroyEvent`
  - `core/modules/deploy.Entity`
  - `core/modules/deploy.PlayerContext`
  - `core/modules/deploy.Reach`
  - `core/modules/deploy.Service`

`core/modules/economy`:
  - `core/modules/economy.Cost`
  - `core/modules/economy.ErrToExpensive`
  - `core/modules/economy.Money`
  - `core/modules/economy.Pay`
  - `core/modules/economy.Wallet`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Deployed`
  - `core/modules/obstruction.ErrPositionIsOccupied`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.NewDeployed`
  - `core/modules/obstruction.Obstruction`

`core/modules/reach`:
  - `core/modules/reach.Component`
  - `core/modules/reach.Distance`
  - `core/modules/reach.ErrOutsideOfReach`
  - `core/modules/reach.NewReach`
  - `core/modules/reach.Reach`
  - `core/modules/reach.ReachesCoords`
  - `core/modules/reach.ServiceT`

`core/modules/reach/pkg`:
  - `core/modules/reach/pkg.PkgT`

`core/modules/tile`:
  - `core/modules/tile.Blueprint`
  - `core/modules/tile.Coord`
  - `core/modules/tile.ErrBlueprintIsMissingUUID`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.Pos`
  - `core/modules/tile.Size`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/feature`:
  - `engine/modules/feature.Error`
  - `engine/modules/feature.NewFieldGetter`

`engine/modules/feature/pkg`:
  - `engine/modules/feature/pkg.Config`
  - `engine/modules/feature/pkg.RegisterFeature`
  - `engine/modules/feature/pkg.RegisterRuleToField`

`engine/modules/grid`:
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`

`engine/modules/loop`:
  - `engine/modules/loop.TickEvent`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.AddEventValidation`
  - `engine/modules/netsync/pkg.Config`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.NewUUIDFromAny`
  - `engine/modules/uuid.SetUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`