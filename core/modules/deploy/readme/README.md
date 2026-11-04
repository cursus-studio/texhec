# deploy
## Architecture
defines GUI for deploying objects

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             57             17            357
-------------------------------------------------------------------------------
SUM:                             3             57             17            357
-------------------------------------------------------------------------------
```
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
  - `core/modules/actions.CanDeployBlueprint`
  - `core/modules/actions.Coords`
  - `core/modules/actions.Entity`
  - `core/modules/actions.NewCanDeployAtRule`
  - `core/modules/actions.NewCanDeployBlueprintRule`
  - `core/modules/actions.NewCanDeployRule`
  - `core/modules/actions.NewOwnRule`
  - `core/modules/actions.Own`
  - `core/modules/actions.PlayerContext`

`core/modules/deploy`:
  - `core/modules/deploy.Blueprint`
  - `core/modules/deploy.By`
  - `core/modules/deploy.Coords`
  - `core/modules/deploy.DeployComponent`
  - `core/modules/deploy.DeployCounterComponent`
  - `core/modules/deploy.DeployEvent`
  - `core/modules/deploy.Deployable`
  - `core/modules/deploy.DestroyEvent`
  - `core/modules/deploy.Entity`
  - `core/modules/deploy.I`
  - `core/modules/deploy.Inc`
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
  - `core/modules/obstruction.Deployed`
  - `core/modules/obstruction.ErrPositionIsOccupied`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.NewDeployed`
  - `core/modules/obstruction.Obstrucion`
  - `core/modules/obstruction.Obstruction`

`core/modules/reach`:
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
  - `engine/modules/netsync/pkg.ComponentPkg`
  - `engine/modules/netsync/pkg.EventPkg`

`engine/modules/seed`:
  - `engine/modules/seed.ErrWorldShouldHaveOneSeed`
  - `engine/modules/seed.Seed`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.NewDeterministicUUID`
  - `engine/modules/uuid.NewUUIDFromAny`
  - `engine/modules/uuid.SetUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`