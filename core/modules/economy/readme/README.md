# economy
## Architecture
economy module is responsible for managing players resources

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             31              2            180
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             4             31              2            181
-------------------------------------------------------------------------------
```
## Dependencies
`core/game`:
  - `core/game.Economy`
  - `core/game.GameWorld`
  - `core/game.Player`
  - `core/game.Ui`

`core/modules/definitions`:
  - `core/modules/definitions.UiGroup`

`core/modules/economy`:
  - `core/modules/economy.Cost`
  - `core/modules/economy.CostComponent`
  - `core/modules/economy.Factory`
  - `core/modules/economy.FactoryComponent`
  - `core/modules/economy.Money`
  - `core/modules/economy.MoneyPerSecond`
  - `core/modules/economy.NewCost`
  - `core/modules/economy.NewFactory`
  - `core/modules/economy.Service`
  - `core/modules/economy.Wallet`
  - `core/modules/economy.WalletComponent`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/groups`:
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`
  - `engine/modules/groups.Groups`

`engine/modules/loop`:
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.TickEvent`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.ComponentPkg`

`engine/modules/text`:
  - `engine/modules/text.Align`
  - `engine/modules/text.Break`
  - `engine/modules/text.BreakNone`
  - `engine/modules/text.Content`
  - `engine/modules/text.FontSize`
  - `engine/modules/text.NewAlign`
  - `engine/modules/text.NewBreak`
  - `engine/modules/text.NewFontSize`
  - `engine/modules/text.NewText`

`engine/modules/transform`:
  - `engine/modules/transform.Inherit`
  - `engine/modules/transform.NewInherit`
  - `engine/modules/transform.NewParentPivotPoint`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.ParentPivotPoint`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.RelativePos`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`