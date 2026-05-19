# ui
## Architecture
this module is responsible for create foundations for creating in game GUI

## Types
### type Service
Type: `core/modules/ui.Service`

#### method Service Actions
Type: `func() core/modules/ui.SelectionGroup[core/modules/ui.ActionComponent]`

#### method Service AnimatedBackground
Type: `func() engine/services/ecs.ComponentsArray[core/modules/ui.AnimatedBackgroundComponent]`

#### method Service CursorCamera
Type: `func() engine/services/ecs.ComponentsArray[core/modules/ui.CursorCameraComponent]`

#### method Service HideMenu
Type: `func()`
removes all children

#### method Service Objects
Type: `func() core/modules/ui.SelectionGroup[core/modules/ui.ObjectComponent]`

#### method Service Register
Type: `func() error`

#### method Service ShowMenu
Type: `func() (parents []engine/services/ecs.EntityID)`
returns parent to attach ui elements
potentially with enter animation

#### method Service UiCamera
Type: `func() engine/services/ecs.ComponentsArray[core/modules/ui.UiCameraComponent]`

### type SelectionGroup
Type: `core/modules/ui.SelectionGroup[Component any]`
groups selected elements with component and allows to remove all of them at once
[SelectionGroup] differs from [ecs.ComponentsArray] that it listens to extra events

#### method SelectionGroup AddDependency
Type: `func(engine/services/ecs.AnyComponentArray)`

#### method SelectionGroup AddDirtySet
Type: `func(engine/services/ecs.DirtySet)`

#### method SelectionGroup BeforeGet
Type: `func(engine/services/ecs.BeforeGet)`

#### method SelectionGroup Get
Type: `func(entity engine/services/ecs.EntityID) (Component, bool)`

#### method SelectionGroup GetAny
Type: `func(entity engine/services/ecs.EntityID) (any, bool)`

#### method SelectionGroup GetEmpty
Type: `func() Component`

#### method SelectionGroup GetEntities
Type: `func() []engine/services/ecs.EntityID`

#### method SelectionGroup OnMod
Type: `func(engine/services/ecs.OnMod)`

#### method SelectionGroup OnRemove
Type: `func(engine/services/ecs.OnMod)`

#### method SelectionGroup OnUpsert
Type: `func(engine/services/ecs.OnMod)`

#### method SelectionGroup Remove
Type: `func(engine/services/ecs.EntityID)`

#### method SelectionGroup Set
Type: `func(engine/services/ecs.EntityID, Component)`

#### method SelectionGroup SetAny
Type: `func(engine/services/ecs.EntityID, any) error`

#### method SelectionGroup SetEmpty
Type: `func(Component)`

### type UiCameraComponent
Type: `core/modules/ui.UiCameraComponent`
marker which says module relative to which element to position

### type AnimatedBackgroundComponent
Type: `core/modules/ui.AnimatedBackgroundComponent`

### type CursorCameraComponent
Type: `core/modules/ui.CursorCameraComponent`

### type UnselectEvent
Type: `core/modules/ui.UnselectEvent[Component any]`
selection group events

### type SelectEvent
Type: `core/modules/ui.SelectEvent[Component any]`

#### property SelectEvent Entities
Type: `[]engine/services/ecs.EntityID`

### type SelectTickEvent
Type: `core/modules/ui.SelectTickEvent[Component any]`
each tick is emited with currently selected entity

#### property SelectTickEvent Tick
Type: `engine/modules/loop.TickEvent`

#### property SelectTickEvent Entities
Type: `[]engine/services/ecs.EntityID`

### type ObjectComponent
Type: `core/modules/ui.ObjectComponent`

### type ActionComponent
Type: `core/modules/ui.ActionComponent`

## Functions
### func NewUnselect
Type: `func[Component any]() core/modules/ui.UnselectEvent[Component]`

### func NewSelect
Type: `func[Component any](entities ...engine/services/ecs.EntityID) core/modules/ui.SelectEvent[Component]`

### func NewSelectTick
Type: `func[Component any](tick engine/modules/loop.TickEvent, entities []engine/services/ecs.EntityID) core/modules/ui.SelectTickEvent[Component]`


## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7             98             16            451
Markdown                         1              8              0             27
-------------------------------------------------------------------------------
SUM:                             8            106             16            478
-------------------------------------------------------------------------------

```
## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.GameWorld`
  - `core/game.Ui`

`core/modules/ui`:
  - `core/modules/ui.ActionComponent`
  - `core/modules/ui.AnimatedBackground`
  - `core/modules/ui.AnimatedBackgroundComponent`
  - `core/modules/ui.CursorCamera`
  - `core/modules/ui.CursorCameraComponent`
  - `core/modules/ui.Entities`
  - `core/modules/ui.NewSelectTick`
  - `core/modules/ui.NewUnselect`
  - `core/modules/ui.ObjectComponent`
  - `core/modules/ui.SelectEvent`
  - `core/modules/ui.SelectionGroup`
  - `core/modules/ui.Service`
  - `core/modules/ui.UiCameraComponent`
  - `core/modules/ui.UnselectEvent`

`engine/modules/assets`:
  - `engine/modules/assets.GetAsset`

`engine/modules/collider`:
  - `engine/modules/collider.Component`
  - `engine/modules/collider.Direction`
  - `engine/modules/collider.NewCollider`
  - `engine/modules/collider.Pos`

`engine/modules/groups`:
  - `engine/modules/groups.Inherit`
  - `engine/modules/groups.InheritGroupsComponent`

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
  - `engine/modules/loop.TickEvent`

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
  - `engine/modules/transform.NewParent`
  - `engine/modules/transform.NewParentPivotPoint`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.Parent`
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

`engine/services/ecs`:
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.RegisterSystems`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.RemoveEntity`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/sdl`