# actions
## Architecture
Stores and renders all game fields and rules.
Currently all rendering is in one module to ensure consistency but
certain rules could be spreaded across multiple modules and if this module grows to much it should.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7             83             17            590
Markdown                         1              0              0              3
-------------------------------------------------------------------------------
SUM:                             8             83             17            593
-------------------------------------------------------------------------------
```
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
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.Obstrucion`
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
  - `engine/modules/feature.ErrInvalidType`
  - `engine/modules/feature.Error`
  - `engine/modules/feature.FieldGetter`
  - `engine/modules/feature.FieldRenderInteraction`
  - `engine/modules/feature.FieldService`
  - `engine/modules/feature.Fields`
  - `engine/modules/feature.From`
  - `engine/modules/feature.Index`
  - `engine/modules/feature.Interactions`
  - `engine/modules/feature.LockedFeature`
  - `engine/modules/feature.NewFieldGetter`
  - `engine/modules/feature.PreviewInteraction`
  - `engine/modules/feature.Render`
  - `engine/modules/feature.Rule`
  - `engine/modules/feature.RuleRender`
  - `engine/modules/feature.RuleRenderCompletePreview`
  - `engine/modules/feature.RuleRenderPreviewRulesFail`
  - `engine/modules/feature.RuleRenderPreviewRulesPass`
  - `engine/modules/feature.RuleService`
  - `engine/modules/feature.SubmitInteraction`
  - `engine/modules/feature.Validate`
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
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`
  - `engine/modules/groups.Groups`
  - `engine/modules/groups.InheritGroups`

`engine/modules/inputs`:
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.AddGenericEventValidation`
  - `engine/modules/netsync/pkg.Config`

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
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.SetUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`