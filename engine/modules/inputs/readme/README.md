# inputs
## Architecture
integrates inputs into the engine

defines mouse and keyboard high-level abstractions.

### mouse
directly parses mouse inputs to specific entity click events

### keyboard
takes keyboard inputs and spreads them using signals which can be captured by entities in focused element hierarchy

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              12            177             27            823
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            13            177             27            824
-------------------------------------------------------------------------------
```
## TODO
Implement a proper input cursor and improve focusing and unfocusing on input

## Dependencies
`engine`:
  - `engine.Camera`
  - `engine.Collider`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Inputs`
  - `engine.Logger`
  - `engine.Text`
  - `engine.Window`
  - `engine.World`

`engine/modules/collider`:
  - `engine/modules/collider.Distance`
  - `engine/modules/collider.Entity`
  - `engine/modules/collider.Hit`
  - `engine/modules/collider.ObjectRayCollision`
  - `engine/modules/collider.RaycastAll`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSet`
  - `engine/modules/datastructures.SetReader`

`engine/modules/ecs`:
  - `engine/modules/ecs.ApplyEntity`
  - `engine/modules/ecs.ApplyEntityEvent`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.RegisterSystems`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/focus/pkg`:
  - `engine/modules/focus/pkg.BubblePkgT`

`engine/modules/inputs`:
  - `engine/modules/inputs.ApplyDrag`
  - `engine/modules/inputs.ApplyDragEvent`
  - `engine/modules/inputs.Camera`
  - `engine/modules/inputs.CursorCol`
  - `engine/modules/inputs.CursorRow`
  - `engine/modules/inputs.DoubleLeftClick`
  - `engine/modules/inputs.DoubleLeftClickComponent`
  - `engine/modules/inputs.DoubleRightClick`
  - `engine/modules/inputs.DoubleRightClickComponent`
  - `engine/modules/inputs.Drag`
  - `engine/modules/inputs.DragComponent`
  - `engine/modules/inputs.DragEvent`
  - `engine/modules/inputs.DraggedComponent`
  - `engine/modules/inputs.Entity`
  - `engine/modules/inputs.Event`
  - `engine/modules/inputs.EventTargetSetter`
  - `engine/modules/inputs.Hover`
  - `engine/modules/inputs.HoverComponent`
  - `engine/modules/inputs.Hovered`
  - `engine/modules/inputs.HoveredComponent`
  - `engine/modules/inputs.KeepSelected`
  - `engine/modules/inputs.KeepSelectedComponent`
  - `engine/modules/inputs.KeyboardEvent`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.LeftClickComponent`
  - `engine/modules/inputs.Lines`
  - `engine/modules/inputs.MouseEnter`
  - `engine/modules/inputs.MouseEnterComponent`
  - `engine/modules/inputs.MouseLeave`
  - `engine/modules/inputs.MouseLeaveComponent`
  - `engine/modules/inputs.NewKeyboardEvent`
  - `engine/modules/inputs.OnUpdate`
  - `engine/modules/inputs.RightClick`
  - `engine/modules/inputs.RightClickComponent`
  - `engine/modules/inputs.Service`
  - `engine/modules/inputs.SetTarget`
  - `engine/modules/inputs.Stack`
  - `engine/modules/inputs.StackComponent`
  - `engine/modules/inputs.Stacked`
  - `engine/modules/inputs.StackedComponent`
  - `engine/modules/inputs.StackedData`
  - `engine/modules/inputs.SynchronizePositionEvent`
  - `engine/modules/inputs.Target`
  - `engine/modules/inputs.Text`
  - `engine/modules/inputs.TextInput`
  - `engine/modules/inputs.TextInputComponent`
  - `engine/modules/inputs.TextInputEvent`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.NewStopEvent`

`engine/modules/text`:
  - `engine/modules/text.Content`
  - `engine/modules/text.NewText`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/window`:
  - `engine/modules/window.GetMousePos`
  - `engine/modules/window.MousePos`
  - `engine/modules/window.NewMousePos`
  - `engine/modules/window.X`
  - `engine/modules/window.Y`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/sdl`