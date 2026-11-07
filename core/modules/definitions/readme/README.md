# definitions
## Architecture
contains all game specific objects

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             43              9            336
Markdown                         1              1              0              3
-------------------------------------------------------------------------------
SUM:                             4             44              9            339
-------------------------------------------------------------------------------
```
## TODO
Move.
Clean up `/assets` and start using it instead of local copy.

Create automatic export from `.pxo` to `.png` and/or `.git`

## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.GameWorld`

`core/modules/definitions`:
  - `core/modules/definitions.AirspaceObstruction`
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.Blank`
  - `core/modules/definitions.Border`
  - `core/modules/definitions.Btn`
  - `core/modules/definitions.EaseOutElastic`
  - `core/modules/definitions.Farm`
  - `core/modules/definitions.HouseT1`
  - `core/modules/definitions.HouseT2`
  - `core/modules/definitions.HouseT3`
  - `core/modules/definitions.HouseT4`
  - `core/modules/definitions.Hud`
  - `core/modules/definitions.Input`
  - `core/modules/definitions.Linear`
  - `core/modules/definitions.LowlandObstruction`
  - `core/modules/definitions.MyEasing`
  - `core/modules/definitions.Objects`
  - `core/modules/definitions.Service`
  - `core/modules/definitions.SquareCollider`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.Tank`
  - `core/modules/definitions.Text`
  - `core/modules/definitions.Tiles`
  - `core/modules/definitions.Transitions`
  - `core/modules/definitions.WaterObstruction`

`core/modules/deploy`:
  - `core/modules/deploy.DeployComp`
  - `core/modules/deploy.NewDeploy`

`core/modules/obstruction`:
  - `core/modules/obstruction.Obstruction`
  - `core/modules/obstruction.Obstructions`
  - `core/modules/obstruction.Service`

`core/modules/tile`:
  - `core/modules/tile.Coord`

`engine/modules/assets`:
  - `engine/modules/assets.Cache`
  - `engine/modules/assets.GetAsset`
  - `engine/modules/assets.NewCache`

`engine/modules/audio`:
  - `engine/modules/audio.Channel`

`engine/modules/collider`:
  - `engine/modules/collider.AABB`
  - `engine/modules/collider.Collider`
  - `engine/modules/collider.Leaf`
  - `engine/modules/collider.NewAABB`
  - `engine/modules/collider.NewCollider`
  - `engine/modules/collider.NewColliderAsset`
  - `engine/modules/collider.NewPolygon`
  - `engine/modules/collider.NewRange`
  - `engine/modules/collider.Polygon`
  - `engine/modules/collider.Range`

`engine/modules/ecs`:
  - `engine/modules/ecs.EntityID`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.GetRegistry`

`engine/modules/focus`:
  - `engine/modules/focus.Bubbling`
  - `engine/modules/focus.FocusEvent`
  - `engine/modules/focus.NewBubbling`

`engine/modules/graphics`:
  - `engine/modules/graphics.Index`

`engine/modules/groups`:
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Group`
  - `engine/modules/groups.Groups`
  - `engine/modules/groups.InheritGroups`

`engine/modules/inputs`:
  - `engine/modules/inputs.KeepSelected`
  - `engine/modules/inputs.KeepSelectedComponent`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`
  - `engine/modules/inputs.NewTextInputEvent`

`engine/modules/render`:
  - `engine/modules/render.AspectRatio`
  - `engine/modules/render.Color`
  - `engine/modules/render.Mesh`
  - `engine/modules/render.NewColor`
  - `engine/modules/render.NewMesh`
  - `engine/modules/render.NewMeshAsset`
  - `engine/modules/render.NewTexture`
  - `engine/modules/render.NewTextureAsset`
  - `engine/modules/render.Texture`
  - `engine/modules/render.TextureAsset`
  - `engine/modules/render.Vertex`

`engine/modules/scene`:
  - `engine/modules/scene.NewSceneId`

`engine/modules/text`:
  - `engine/modules/text.Align`
  - `engine/modules/text.FontSize`
  - `engine/modules/text.NewAlign`
  - `engine/modules/text.NewFontSize`

`engine/modules/transform`:
  - `engine/modules/transform.AspectRatio`
  - `engine/modules/transform.Inherit`
  - `engine/modules/transform.MaxSize`
  - `engine/modules/transform.NewAspectRatio`
  - `engine/modules/transform.NewInherit`
  - `engine/modules/transform.NewMaxSize`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.PrimaryAxisX`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeSizeX`
  - `engine/modules/transform.Size`

`engine/modules/transition`:
  - `engine/modules/transition.EasingFunction`
  - `engine/modules/transition.NewEasingFunction`
  - `engine/modules/transition.Progress`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/ioc/v2`