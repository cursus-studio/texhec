# economy
## Architecture
economy module is responsible for managing players resources

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             30              2            178
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             4             30              2            179
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/economy.Service`

#### method Service Cost
Type: `func() engine/modules/ecs.ComponentArray[core/modules/economy.CostComponent]`

#### method Service Factory
Type: `func() engine/modules/ecs.ComponentArray[core/modules/economy.FactoryComponent]`

#### method Service Register
Type: `func() error`

#### method Service Wallet
Type: `func() engine/modules/ecs.ComponentArray[core/modules/economy.WalletComponent]`

### type Money
Type: `core/modules/economy.Money`

### type WalletComponent
Type: `core/modules/economy.WalletComponent`

#### property WalletComponent Money
Type: `core/modules/economy.Money`

#### method WalletComponent Pay
Type: `func(cost core/modules/economy.CostComponent) core/modules/economy.WalletComponent`

#### method WalletComponent Smooth
Type: `func()`

#### method WalletComponent Lerp
Type: `func(c2 core/modules/economy.WalletComponent, mix32 float32) core/modules/economy.WalletComponent`

### type FactoryComponent
Type: `core/modules/economy.FactoryComponent`

#### property FactoryComponent MoneyPerSecond
Type: `core/modules/economy.Money`

### type CostComponent
Type: `core/modules/economy.CostComponent`

#### property CostComponent Cost
Type: `core/modules/economy.Money`

## Variables
### var ErrToExpensive
Type: `error`

## Functions
### func NewWallet
Type: `func(money core/modules/economy.Money) core/modules/economy.WalletComponent`

### func NewFactory
Type: `func(money core/modules/economy.Money) core/modules/economy.FactoryComponent`

### func NewCost
Type: `func(money core/modules/economy.Money) core/modules/economy.CostComponent`


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
  - `engine/modules/groups.Component`
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`

`engine/modules/loop`:
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.TickEvent`

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