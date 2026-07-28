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
Go                              12            177             27            821
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            13            177             27            822
-------------------------------------------------------------------------------
```
## TODO
Implement a proper input cursor and improve focusing and unfocusing on input

## Types
### type Service
Type: `engine/modules/inputs.Service`

#### method Service DoubleLeftClick
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.DoubleLeftClickComponent]`

#### method Service DoubleRightClick
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.DoubleRightClickComponent]`

#### method Service Drag
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.DragComponent]`

#### method Service Dragged
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.DraggedComponent]`

#### method Service Hover
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.HoverComponent]`

#### method Service Hovered
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.HoveredComponent]`

#### method Service KeepSelected
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.KeepSelectedComponent]`

#### method Service LeftClick
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.LeftClickComponent]`

#### method Service MouseEnter
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.MouseEnterComponent]`

#### method Service MouseLeave
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.MouseLeaveComponent]`

#### method Service Register
Type: `func() error`

#### method Service RightClick
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.RightClickComponent]`

#### method Service Stack
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.StackComponent]`

#### method Service Stacked
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.StackedComponent]`

#### method Service StackedData
Type: `func() []engine/modules/inputs.Target`
returns ordered targets with additional data

#### method Service TextInput
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/inputs.TextInputComponent]`

### type ApplyDragEvent
Type: `engine/modules/inputs.ApplyDragEvent`
interfaces which can be implemented by events {

#### method ApplyDragEvent ApplyDrag
Type: `func(engine/modules/inputs.DragEvent) (event any)`

### type EventTargetSetter
Type: `engine/modules/inputs.EventTargetSetter`

#### method EventTargetSetter SetTarget
Type: `func(engine/modules/inputs.Target) engine/modules/inputs.EventTargetSetter`

### type KeyboardEvent
Type: `engine/modules/inputs.KeyboardEvent`

#### property KeyboardEvent KeyboardEvent
Type: `github.com/veandco/go-sdl2/sdl.KeyboardEvent`

### type TextInputComponent
Type: `engine/modules/inputs.TextInputComponent`

#### property TextInputComponent Lines
Type: `[]string`
split in lines

#### property TextInputComponent CursorRow
Type: `int`

#### property TextInputComponent CursorCol
Type: `int`

#### method TextInputComponent Text
Type: `func() string`

### type TextInputEvent
Type: `engine/modules/inputs.TextInputEvent`
handles inputs and saves change in component

#### property TextInputEvent Entity
Type: `engine/modules/ecs.EntityID`

#### property TextInputEvent KeyboardEvent
Type: `engine/modules/inputs.KeyboardEvent`

#### method TextInputEvent CapturesEvents
Type: `func() engine/modules/datastructures.SetReader[reflect.Type]`

#### method TextInputEvent Capture
Type: `func(event any) any`

#### method TextInputEvent ApplyEntity
Type: `func(entity engine/modules/ecs.EntityID) any`

#### method TextInputEvent Fallthrough
Type: `func() bool`

### type HoveredComponent
Type: `engine/modules/inputs.HoveredComponent`
many elements can be hovered at once

#### property HoveredComponent Camera
Type: `engine/modules/ecs.EntityID`

### type DraggedComponent
Type: `engine/modules/inputs.DraggedComponent`

#### property DraggedComponent Camera
Type: `engine/modules/ecs.EntityID`

### type KeepSelectedComponent
Type: `engine/modules/inputs.KeepSelectedComponent`
keeps element selected even if user drags outside

### type StackComponent
Type: `engine/modules/inputs.StackComponent`
this is special component stating that on click it enables clicking elements below

### type StackedComponent
Type: `engine/modules/inputs.StackedComponent`
this means that element got pressed and if not next that isn't going to be pressed

### type LeftClickComponent
Type: `engine/modules/inputs.LeftClickComponent`

#### property LeftClickComponent Event
Type: `any`

### type DoubleLeftClickComponent
Type: `engine/modules/inputs.DoubleLeftClickComponent`

#### property DoubleLeftClickComponent Event
Type: `any`

### type RightClickComponent
Type: `engine/modules/inputs.RightClickComponent`

#### property RightClickComponent Event
Type: `any`

### type DoubleRightClickComponent
Type: `engine/modules/inputs.DoubleRightClickComponent`

#### property DoubleRightClickComponent Event
Type: `any`

### type MouseEnterComponent
Type: `engine/modules/inputs.MouseEnterComponent`

#### property MouseEnterComponent Event
Type: `any`

### type MouseLeaveComponent
Type: `engine/modules/inputs.MouseLeaveComponent`

#### property MouseLeaveComponent Event
Type: `any`

### type HoverComponent
Type: `engine/modules/inputs.HoverComponent`

#### property HoverComponent Event
Type: `any`

### type DragComponent
Type: `engine/modules/inputs.DragComponent`

#### property DragComponent Event
Type: `any`

### type DragEvent
Type: `engine/modules/inputs.DragEvent`
this event is called when nothing is dragged

#### property DragEvent Camera
Type: `engine/modules/ecs.EntityID`

#### property DragEvent From
Type: `engine/modules/window.MousePos`
from and to is normalized

#### property DragEvent To
Type: `engine/modules/window.MousePos`
from and to is normalized

### type SynchronizePositionEvent
Type: `engine/modules/inputs.SynchronizePositionEvent`

#### property SynchronizePositionEvent Camera
Type: `engine/modules/ecs.EntityID`

#### property SynchronizePositionEvent From
Type: `engine/modules/window.MousePos`
from and to is normalized

#### property SynchronizePositionEvent To
Type: `engine/modules/window.MousePos`
from and to is normalized

#### method SynchronizePositionEvent ApplyDrag
Type: `func(dragEvent engine/modules/inputs.DragEvent) any`

### type Target
Type: `engine/modules/inputs.Target`

#### property Target ObjectRayCollision
Type: `engine/modules/collider.ObjectRayCollision`

#### property Target Camera
Type: `engine/modules/ecs.EntityID`

## Functions
### func NewKeyboardEvent
Type: `func(e github.com/veandco/go-sdl2/sdl.KeyboardEvent) engine/modules/inputs.KeyboardEvent`

### func NewTextInputEvent
Type: `func() engine/modules/inputs.TextInputEvent`

### func NewHovered
Type: `func(camera engine/modules/ecs.EntityID) engine/modules/inputs.HoveredComponent`

### func NewDragged
Type: `func(camera engine/modules/ecs.EntityID) engine/modules/inputs.DraggedComponent`

### func NewLeftClick
Type: `func(e any) engine/modules/inputs.LeftClickComponent`

### func NewDoubleLeftClick
Type: `func(e any) engine/modules/inputs.DoubleLeftClickComponent`

### func NewRightClick
Type: `func(e any) engine/modules/inputs.RightClickComponent`

### func NewDoubleRightClick
Type: `func(e any) engine/modules/inputs.DoubleRightClickComponent`

### func NewMouseEnterComponent
Type: `func(event any) engine/modules/inputs.MouseEnterComponent`

### func NewMouseLeaveComponent
Type: `func(event any) engine/modules/inputs.MouseLeaveComponent`

### func NewHoverComponent
Type: `func(event any) engine/modules/inputs.HoverComponent`

### func NewDragComponent
Type: `func(event any) engine/modules/inputs.DragComponent`


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