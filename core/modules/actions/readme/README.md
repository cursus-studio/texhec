# actions
## Architecture
Stores all game domain interactions and steps

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             57            380            267
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             7             57            380            268
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/actions.Service`

#### method Service Blueprint
Type: `func() engine/modules/feature.FieldService[core/modules/actions.Blueprint]`

#### method Service CanAttack
Type: `func() engine/modules/feature.NoInputRuleService[core/modules/actions.CanAttackRule, core/modules/actions.Entity]`

#### method Service CanBeAttacked
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanBeAttackedRule, core/modules/actions.Entity, core/modules/actions.Entity]`

#### method Service CanBuild
Type: `func() engine/modules/feature.NoInputRuleService[core/modules/actions.CanBuildRule, core/modules/actions.Entity]`

#### method Service CanBuildBlueprint
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanBuildBlueprintRule, core/modules/actions.Entity, core/modules/actions.Blueprint]`

#### method Service CanMove
Type: `func() engine/modules/feature.NoInputRuleService[core/modules/actions.CanMoveRule, core/modules/actions.Entity]`

#### method Service CanMoveTo
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanMoveToRule, core/modules/actions.Entity, core/modules/actions.Coords]`

#### method Service CanPlace
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanPlaceRule, core/modules/actions.Blueprint, core/modules/actions.Coords]`

#### method Service Coords
Type: `func() engine/modules/feature.FieldService[core/modules/actions.Coords]`

#### method Service Entity
Type: `func() engine/modules/feature.FieldService[core/modules/actions.Entity]`

#### method Service IsEnemy
Type: `func() engine/modules/feature.RuleService[core/modules/actions.IsEnemyRule, core/modules/actions.PlayerContext, core/modules/actions.Entity]`

#### method Service Own
Type: `func() engine/modules/feature.RuleService[core/modules/actions.OwnRule, core/modules/actions.PlayerContext, core/modules/actions.Entity]`
rules

#### method Service PlayerContext
Type: `func() engine/modules/feature.FieldService[core/modules/actions.PlayerContext]`
fields

### type PlayerContextGetter
Type: `core/modules/actions.PlayerContextGetter`

#### method PlayerContextGetter Context
Type: `func() core/modules/actions.PlayerContext`

### type OwnRule
Type: `core/modules/actions.OwnRule`

### type IsEnemyRule
Type: `core/modules/actions.IsEnemyRule`

### type CanBuildRule
Type: `core/modules/actions.CanBuildRule`
build

### type CanBuildBlueprintRule
Type: `core/modules/actions.CanBuildBlueprintRule`

### type CanPlaceRule
Type: `core/modules/actions.CanPlaceRule`

### type CanMoveRule
Type: `core/modules/actions.CanMoveRule`
move

### type CanMoveToRule
Type: `core/modules/actions.CanMoveToRule`

### type CanAttackRule
Type: `core/modules/actions.CanAttackRule`
attack

### type CanBeAttackedRule
Type: `core/modules/actions.CanBeAttackedRule`

### type PlayerContext
Type: `core/modules/actions.PlayerContext`

#### method PlayerContext Context
Type: `func() core/modules/actions.PlayerContext`

### type Coords
Type: `core/modules/actions.Coords`

#### property Coords X
Type: `engine/modules/grid.Coord`

#### property Coords Y
Type: `engine/modules/grid.Coord`

### type Entity
Type: `core/modules/actions.Entity`

### type Blueprint
Type: `core/modules/actions.Blueprint`


## Dependencies
`core/game`:
  - `core/game.Attack`
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.EngineWorld`
  - `core/game.GameWorld`
  - `core/game.Pathfind`
  - `core/game.Player`
  - `core/game.Tile`
  - `core/game.Ui`

`core/modules/actions`:
  - `core/modules/actions.Blueprint`
  - `core/modules/actions.CanAttackRule`
  - `core/modules/actions.CanBeAttackedRule`
  - `core/modules/actions.CanBuildBlueprintRule`
  - `core/modules/actions.CanBuildRule`
  - `core/modules/actions.CanMoveRule`
  - `core/modules/actions.CanMoveToRule`
  - `core/modules/actions.CanPlaceRule`
  - `core/modules/actions.Coords`
  - `core/modules/actions.Entity`
  - `core/modules/actions.IsEnemyRule`
  - `core/modules/actions.OwnRule`
  - `core/modules/actions.PlayerContext`
  - `core/modules/actions.Service`

`core/modules/tile`:
  - `core/modules/tile.Blueprint`
  - `core/modules/tile.ClickBlueprintEvent`
  - `core/modules/tile.ClickEntityEvent`
  - `core/modules/tile.Entity`
  - `core/modules/tile.HoverEntityEvent`
  - `core/modules/tile.Name`
  - `core/modules/tile.NewClickBlueprintEvent`

`engine/modules/ecs`:
  - `engine/modules/ecs.EntityID`

`engine/modules/feature`:
  - `engine/modules/feature.Error`
  - `engine/modules/feature.FieldService`
  - `engine/modules/feature.In`
  - `engine/modules/feature.NoInputRule`
  - `engine/modules/feature.NoInputRuleService`
  - `engine/modules/feature.PreviewInteraction`
  - `engine/modules/feature.Render`
  - `engine/modules/feature.Rule`
  - `engine/modules/feature.RuleService`
  - `engine/modules/feature.SubmitInteraction`

`engine/modules/feature/pkg`:
  - `engine/modules/feature/pkg.FieldPkg`
  - `engine/modules/feature/pkg.NoInputRulePkg`
  - `engine/modules/feature/pkg.RulePkg`

`engine/modules/grid`:
  - `engine/modules/grid.AbsoluteCoords`
  - `engine/modules/grid.Chunk`
  - `engine/modules/grid.ClickEvent`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.HoverEvent`

`engine/modules/inputs`:
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

`engine/modules/text`:
  - `engine/modules/text.Content`
  - `engine/modules/text.NewText`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`