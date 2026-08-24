# definitions
## Architecture
contains all game specific objects

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             42              8            336
Markdown                         1              1              0              3
-------------------------------------------------------------------------------
SUM:                             4             43              8            339
-------------------------------------------------------------------------------
```
## TODO
Move.
Clean up `/assets` and start using it instead of local copy.

Create automatic export from `.pxo` to `.png` and/or `.git`

## Types
### type Service
Type: `core/modules/definitions.Service`
In DI container
Definitions have more dependencies

#### method Service Assets
Type: `func() core/modules/definitions.Assets`

#### method Service Hud
Type: `func() core/modules/definitions.Hud`

#### method Service Load
Type: `func()`

#### method Service Objects
Type: `func() core/modules/definitions.Objects`

#### method Service Tiles
Type: `func() core/modules/definitions.Tiles`

#### method Service Transitions
Type: `func() core/modules/definitions.Transitions`

### type Assets
Type: `core/modules/definitions.Assets`
In DI container
Assets have fewer dependencies

#### property Assets ExampleAudio
Type: `engine/modules/ecs.EntityID`

#### property Assets Blank
Type: `engine/modules/ecs.EntityID`

#### property Assets Border
Type: `engine/modules/ecs.EntityID`

#### property Assets SquareMesh
Type: `engine/modules/ecs.EntityID`

#### property Assets SquareCollider
Type: `engine/modules/ecs.EntityID`

#### property Assets FontAsset
Type: `engine/modules/ecs.EntityID`

### type Hud
Type: `core/modules/definitions.Hud`

#### property Hud Btn
Type: `engine/modules/ecs.EntityID`

#### property Hud Text
Type: `engine/modules/ecs.EntityID`

#### property Hud Input
Type: `engine/modules/ecs.EntityID`

#### property Hud Cursor
Type: `engine/modules/ecs.EntityID`

#### property Hud Settings
Type: `engine/modules/ecs.EntityID`

#### property Hud Background1
Type: `engine/modules/ecs.EntityID`

#### property Hud Background2
Type: `engine/modules/ecs.EntityID`

#### property Hud Selected
Type: `engine/modules/ecs.EntityID`

#### property Hud Target
Type: `engine/modules/ecs.EntityID`

#### property Hud Can
Type: `engine/modules/ecs.EntityID`

#### property Hud Cannot
Type: `engine/modules/ecs.EntityID`

### type Transitions
Type: `core/modules/definitions.Transitions`

#### property Transitions Linear
Type: `engine/modules/ecs.EntityID`

#### property Transitions MyEasing
Type: `engine/modules/ecs.EntityID`

#### property Transitions EaseOutElastic
Type: `engine/modules/ecs.EntityID`

### type Tiles
Type: `core/modules/definitions.Tiles`
generation configs should be in registry or in destined path and dispatched instantly on initialization

#### property Tiles Water
Type: `engine/modules/ecs.EntityID`

#### property Tiles Sand
Type: `engine/modules/ecs.EntityID`

#### property Tiles Texhec
Type: `engine/modules/ecs.EntityID`

#### property Tiles Grass
Type: `engine/modules/ecs.EntityID`

#### property Tiles Mountain
Type: `engine/modules/ecs.EntityID`

### type Objects
Type: `core/modules/definitions.Objects`

#### property Objects Farm
Type: `engine/modules/ecs.EntityID`

#### property Objects HouseT1
Type: `engine/modules/ecs.EntityID`

#### property Objects HouseT2
Type: `engine/modules/ecs.EntityID`

#### property Objects HouseT3
Type: `engine/modules/ecs.EntityID`

#### property Objects HouseT4
Type: `engine/modules/ecs.EntityID`

#### property Objects Tank
Type: `engine/modules/ecs.EntityID`

## Variables
### var MenuID
Type: `engine/modules/scene.ID`

### var GameID
Type: `engine/modules/scene.ID`

### var GameServerID
Type: `engine/modules/scene.ID`

### var GameClientID
Type: `engine/modules/scene.ID`

### var SettingsID
Type: `engine/modules/scene.ID`

### var CreditsID
Type: `engine/modules/scene.ID`

### var EffectChannel
Type: `engine/modules/audio.Channel`

### var MusicChannel
Type: `engine/modules/audio.Channel`

### var UiGroup
Type: `engine/modules/groups.Group`

### var GameGroup
Type: `engine/modules/groups.Group`

### var BgGroup
Type: `engine/modules/groups.Group`

### var TileLayer
Type: `core/modules/tile.Coord`

### var ConstructLayer
Type: `core/modules/tile.Coord`

### var UnitLayer
Type: `core/modules/tile.Coord`

### var TilePlaceholderLayer
Type: `core/modules/tile.Coord`
PathLayer

### var RangePlaceholderLayer
Type: `core/modules/tile.Coord`

### var ObjectSelectionPlaceholderLayer
Type: `core/modules/tile.Coord`

### var ObjectPlaceholderLayer
Type: `core/modules/tile.Coord`

### var AirspaceObstruction
Type: `core/modules/obstruction.Obstruction`

### var WaterObstruction
Type: `core/modules/obstruction.Obstruction`

### var LowlandObstruction
Type: `core/modules/obstruction.Obstruction`


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
  - `core/modules/deploy.Component`
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
  - `engine/modules/collider.Component`
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
  - `engine/modules/groups.Component`
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Group`
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