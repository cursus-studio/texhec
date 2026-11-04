# ui
## Architecture
this module is responsible for create foundations for creating in game GUI

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               9            120             15            598
Markdown                         1              8              0             27
-------------------------------------------------------------------------------
SUM:                            10            128             15            625
-------------------------------------------------------------------------------
```
## Dependencies
`core/game`:
  - `core/game.Attack`
  - `core/game.Definitions`
  - `core/game.GameWorld`
  - `core/game.Tile`
  - `core/game.Ui`

`core/modules/definitions`:
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.Background1`
  - `core/modules/definitions.Background2`
  - `core/modules/definitions.BgGroup`
  - `core/modules/definitions.Blank`
  - `core/modules/definitions.Btn`
  - `core/modules/definitions.Cursor`
  - `core/modules/definitions.GameGroup`
  - `core/modules/definitions.Hud`
  - `core/modules/definitions.Settings`
  - `core/modules/definitions.SquareCollider`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.Text`
  - `core/modules/definitions.UiGroup`

`core/modules/settings`:
  - `core/modules/settings.EnterSettingsEvent`

`core/modules/ui`:
  - `core/modules/ui.AnimatedBackground`
  - `core/modules/ui.AnimatedBackgroundComponent`
  - `core/modules/ui.CursorCamera`
  - `core/modules/ui.CursorCameraComponent`
  - `core/modules/ui.GameCameraComponent`
  - `core/modules/ui.MaxChunksZoomOut`
  - `core/modules/ui.Service`
  - `core/modules/ui.ShowMenu`
  - `core/modules/ui.UiCamera`
  - `core/modules/ui.UiCameraComponent`

`engine/modules/assets`:
  - `engine/modules/assets.GetAsset`

`engine/modules/camera`:
  - `engine/modules/camera.GetViewport`
  - `engine/modules/camera.Limits`
  - `engine/modules/camera.Mobile`
  - `engine/modules/camera.NewCameraLimits`
  - `engine/modules/camera.NewMobileCamera`
  - `engine/modules/camera.NewOrtho`
  - `engine/modules/camera.NewPriority`
  - `engine/modules/camera.Ortho`
  - `engine/modules/camera.Priority`
  - `engine/modules/camera.ShootRay`

`engine/modules/collider`:
  - `engine/modules/collider.Collider`
  - `engine/modules/collider.Direction`
  - `engine/modules/collider.NewCollider`
  - `engine/modules/collider.Pos`

`engine/modules/delay`:
  - `engine/modules/delay.NewDelayedEvent`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.RegisterSystems`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/feature`:
  - `engine/modules/feature.Features`
  - `engine/modules/feature.Interactions`
  - `engine/modules/feature.InteractionsListeners`
  - `engine/modules/feature.LockFeatureEvent`
  - `engine/modules/feature.NewLockFeatureEvent`
  - `engine/modules/feature.NewUnlockFeatureEvent`

`engine/modules/groups`:
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`
  - `engine/modules/groups.Groups`
  - `engine/modules/groups.InheritGroups`

`engine/modules/inputs`:
  - `engine/modules/inputs.KeepSelected`
  - `engine/modules/inputs.KeepSelectedComponent`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

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

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.ComponentPkg`

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
  - `engine/modules/transform.RelativeSizeX`
  - `engine/modules/transform.RelativeSizeXY`
  - `engine/modules/transform.Size`

`engine/modules/transition`:
  - `engine/modules/transition.NewTransition`
  - `engine/modules/transition.NewTransitionEvent`
  - `engine/modules/transition.TransitionComponent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/sdl`