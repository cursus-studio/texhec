# actions
## Architecture
Stores all game domain interactions and steps

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7             86             17            557
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             8             86             17            558
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/actions.Service`

#### method Service Blueprint
Type: `func() engine/modules/feature.FieldService[core/modules/actions.Blueprint]`

#### method Service CanAttack
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanAttackRule]`

#### method Service CanBeAttacked
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanBeAttackedRule]`

#### method Service CanDeploy
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanDeployRule]`

#### method Service CanDeployAt
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanDeployAtRule]`

#### method Service CanDeployBlueprint
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanDeployBlueprintRule]`

#### method Service CanMove
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanMoveRule]`

#### method Service CanMoveTo
Type: `func() engine/modules/feature.RuleService[core/modules/actions.CanMoveToRule]`

#### method Service Coords
Type: `func() engine/modules/feature.FieldService[core/modules/actions.Coords]`

#### method Service Entity
Type: `func() engine/modules/feature.FieldService[core/modules/actions.Entity]`

#### method Service IsEnemy
Type: `func() engine/modules/feature.RuleService[core/modules/actions.IsEnemyRule]`

#### method Service Own
Type: `func() engine/modules/feature.RuleService[core/modules/actions.OwnRule]`
rules

#### method Service PlayerContext
Type: `func() engine/modules/feature.FieldService[core/modules/actions.PlayerContext]`
fields

### type PlayerContextGetter
Type: `core/modules/actions.PlayerContextGetter`

#### method PlayerContextGetter Context
Type: `func() core/modules/actions.PlayerContext`

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

### type OwnRule
Type: `core/modules/actions.OwnRule`

#### property OwnRule PlayerContext
Type: `engine/modules/feature.FieldGetter[core/modules/actions.PlayerContext]`

#### property OwnRule Entity
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`

### type IsEnemyRule
Type: `core/modules/actions.IsEnemyRule`

#### property IsEnemyRule PlayerContext
Type: `engine/modules/feature.FieldGetter[core/modules/actions.PlayerContext]`

#### property IsEnemyRule Entity
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`

### type CanDeployRule
Type: `core/modules/actions.CanDeployRule`
build

#### property CanDeployRule Entity
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`

### type CanDeployBlueprintRule
Type: `core/modules/actions.CanDeployBlueprintRule`

#### property CanDeployBlueprintRule Entity
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`

#### property CanDeployBlueprintRule Blueprint
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Blueprint]`

### type CanDeployAtRule
Type: `core/modules/actions.CanDeployAtRule`

#### property CanDeployAtRule Blueprint
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Blueprint]`

#### property CanDeployAtRule Coords
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Coords]`

### type CanMoveRule
Type: `core/modules/actions.CanMoveRule`
move

#### property CanMoveRule Entity
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`

### type CanMoveToRule
Type: `core/modules/actions.CanMoveToRule`

#### property CanMoveToRule Entity
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`

#### property CanMoveToRule Coords
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Coords]`

### type CanAttackRule
Type: `core/modules/actions.CanAttackRule`
attack

#### property CanAttackRule Entity
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`

### type CanBeAttackedRule
Type: `core/modules/actions.CanBeAttackedRule`

#### property CanBeAttackedRule Attacker
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`

#### property CanBeAttackedRule Attacked
Type: `engine/modules/feature.FieldGetter[core/modules/actions.Entity]`


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
  - `core/modules/actions.Attacked`
  - `core/modules/actions.Attacker`
  - `core/modules/actions.Blueprint`
  - `core/modules/actions.CanAttackRule`
  - `core/modules/actions.CanBeAttackedRule`
  - `core/modules/actions.CanDeployAtRule`
  - `core/modules/actions.CanDeployBlueprintRule`
  - `core/modules/actions.CanDeployRule`
  - `core/modules/actions.CanMoveRule`
  - `core/modules/actions.CanMoveToRule`
  - `core/modules/actions.Coords`
  - `core/modules/actions.Entity`
  - `core/modules/actions.IsEnemyRule`
  - `core/modules/actions.OwnRule`
  - `core/modules/actions.PlayerContext`
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
  - `core/modules/definitions.RangePlaceholderLayer`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.TilePlaceholderLayer`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.Obstruction`

`core/modules/tile`:
  - `core/modules/tile.Blueprint`
  - `core/modules/tile.ClickBlueprintEvent`
  - `core/modules/tile.ClickEntityEvent`
  - `core/modules/tile.Entity`
  - `core/modules/tile.HoverEntityEvent`
  - `core/modules/tile.Layer`
  - `core/modules/tile.Name`
  - `core/modules/tile.NewClickBlueprintEvent`
  - `core/modules/tile.NewLayer`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.Pos`
  - `core/modules/tile.Size`

`engine/modules/ecs`:
  - `engine/modules/ecs.EntityID`

`engine/modules/feature`:
  - `engine/modules/feature.AllFields`
  - `engine/modules/feature.Error`
  - `engine/modules/feature.FieldGetter`
  - `engine/modules/feature.FieldRenderInteraction`
  - `engine/modules/feature.FieldService`
  - `engine/modules/feature.Fields`
  - `engine/modules/feature.From`
  - `engine/modules/feature.Interactions`
  - `engine/modules/feature.LockedFeature`
  - `engine/modules/feature.PreviewInteraction`
  - `engine/modules/feature.Render`
  - `engine/modules/feature.Rule`
  - `engine/modules/feature.RuleRender`
  - `engine/modules/feature.RuleRenderCompletePreview`
  - `engine/modules/feature.RuleRenderPreviewRulesFail`
  - `engine/modules/feature.RuleRenderPreviewRulesPass`
  - `engine/modules/feature.RuleService`
  - `engine/modules/feature.SubmitInteraction`
  - `engine/modules/feature.Value`

`engine/modules/feature/pkg`:
  - `engine/modules/feature/pkg.FieldPkg`
  - `engine/modules/feature/pkg.RulePkg`

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

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.SetUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`