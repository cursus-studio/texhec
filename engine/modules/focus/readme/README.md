# focus
## Architecture
Creates all focus related components and events and allows event bubbling from focused or any other entities

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               5             45             22            199
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             6             45             22            200
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Focus`
  - `engine.Hierarchy`
  - `engine.Scene`
  - `engine.World`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.SetReader`

`engine/modules/ecs`:
  - `engine/modules/ecs.ApplyEntity`
  - `engine/modules/ecs.ApplyEntityEvent`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/focus`:
  - `engine/modules/focus.BubbleEvent`
  - `engine/modules/focus.BubblingComponent`
  - `engine/modules/focus.Capture`
  - `engine/modules/focus.CapturesEvents`
  - `engine/modules/focus.DefaultFocusEvent`
  - `engine/modules/focus.DefaultFocused`
  - `engine/modules/focus.DefaultFocusedComponent`
  - `engine/modules/focus.Emit`
  - `engine/modules/focus.Entity`
  - `engine/modules/focus.Event`
  - `engine/modules/focus.EventType`
  - `engine/modules/focus.Fallthrough`
  - `engine/modules/focus.FocusEvent`
  - `engine/modules/focus.Focused`
  - `engine/modules/focus.FocusedComponent`
  - `engine/modules/focus.NewBubbleEvent`
  - `engine/modules/focus.NewDefaultFocusEvent`
  - `engine/modules/focus.NewDefaultFocused`
  - `engine/modules/focus.NewFocused`
  - `engine/modules/focus.NewFocusedBubbleEvent`
  - `engine/modules/focus.Service`
  - `engine/modules/focus.UnfocusEvent`

`engine/modules/scene`:
  - `engine/modules/scene.ChangeSceneEvent`
  - `engine/modules/scene.Scene`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`