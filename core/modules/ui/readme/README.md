# ui
## Architecture
this module is responsible for create foundations for creating in game GUI

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7             83             13            418
Markdown                         1              8              0             27
-------------------------------------------------------------------------------
SUM:                             8             91             13            445
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/ui.Service`

#### method Service AnimatedBackground
Type: `func() engine/modules/ecs.ComponentArray[core/modules/ui.AnimatedBackgroundComponent]`

#### method Service CursorCamera
Type: `func() engine/modules/ecs.ComponentArray[core/modules/ui.CursorCameraComponent]`

#### method Service HideMenu
Type: `func()`
removes all children

#### method Service Register
Type: `func() error`

#### method Service ShowMenu
Type: `func() (parents []engine/modules/ecs.EntityID)`
returns parent to attach ui elements
potentially with enter animation

#### method Service UiCamera
Type: `func() engine/modules/ecs.ComponentArray[core/modules/ui.UiCameraComponent]`

### type UiCameraComponent
Type: `core/modules/ui.UiCameraComponent`
marker which says module relative to which element to position

### type AnimatedBackgroundComponent
Type: `core/modules/ui.AnimatedBackgroundComponent`

### type CursorCameraComponent
Type: `core/modules/ui.CursorCameraComponent`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.GameWorld`
  - `core/game.Ui`

`core/modules/ui`:
  - `core/modules/ui.AnimatedBackground`
  - `core/modules/ui.AnimatedBackgroundComponent`
  - `core/modules/ui.CursorCamera`
  - `core/modules/ui.CursorCameraComponent`
  - `core/modules/ui.Service`
  - `core/modules/ui.ShowMenu`
  - `core/modules/ui.UiCameraComponent`

`engine/modules/assets`:
  - `engine/modules/assets.GetAsset`

`engine/modules/collider`:
  - `engine/modules/collider.Component`
  - `engine/modules/collider.Direction`
  - `engine/modules/collider.NewCollider`
  - `engine/modules/collider.Pos`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.RegisterSystems`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/inputs`:
  - `engine/modules/inputs.KeepSelected`
  - `engine/modules/inputs.KeepSelectedComponent`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

`engine/modules/interactions`:
  - `engine/modules/interactions.AvailableFeatures`
  - `engine/modules/interactions.Features`
  - `engine/modules/interactions.NewDeselectFeatureEvent`
  - `engine/modules/interactions.NewSelectFeatureEvent`
  - `engine/modules/interactions.SelectFeatureEvent`

`engine/modules/layout`:
  - `engine/modules/layout.Align`
  - `engine/modules/layout.Gap`
  - `engine/modules/layout.NewAlign`
  - `engine/modules/layout.NewGap`
  - `engine/modules/layout.NewOrder`
  - `engine/modules/layout.Order`
  - `engine/modules/layout.OrderVectical`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`

`engine/modules/render`:
  - `engine/modules/render.Color`
  - `engine/modules/render.Images`
  - `engine/modules/render.Mesh`
  - `engine/modules/render.NewColor`
  - `engine/modules/render.NewMesh`
  - `engine/modules/render.NewTexture`
  - `engine/modules/render.NewTextureFrame`
  - `engine/modules/render.Texture`
  - `engine/modules/render.TextureAsset`
  - `engine/modules/render.TextureFrameComponent`

`engine/modules/text`:
  - `engine/modules/text.Align`
  - `engine/modules/text.Content`
  - `engine/modules/text.FontSize`
  - `engine/modules/text.NewAlign`
  - `engine/modules/text.NewFontSize`
  - `engine/modules/text.NewText`

`engine/modules/transform`:
  - `engine/modules/transform.Absolute`
  - `engine/modules/transform.Inherit`
  - `engine/modules/transform.NewInherit`
  - `engine/modules/transform.NewParentPivotPoint`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.ParentPivotPoint`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeSizeXY`
  - `engine/modules/transform.Size`

`engine/modules/transition`:
  - `engine/modules/transition.NewDelayedEvent`
  - `engine/modules/transition.NewTransition`
  - `engine/modules/transition.NewTransitionEvent`
  - `engine/modules/transition.TransitionComponent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/sdl`