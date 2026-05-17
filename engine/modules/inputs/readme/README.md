# inputs
## Architecture
integrates inputs into the engine

defines mouse and keyboard high-level abstractions.

### mouse
directly parses mouse inputs to specific entity click events

### keyboard
takes keyboard inputs and spreads them using signals which can be captured by entities in focused element hierarchy

## Types
### type Service
Type: `engine/modules/inputs.Service`

#### method Service CaptureKeyboard
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.CaptureKeyboardComponent]`

#### method Service DefaultFocused
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.DefaultFocusedComponent]`

#### method Service DoubleLeftClick
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.DoubleLeftClickComponent]`

#### method Service DoubleRightClick
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.DoubleRightClickComponent]`

#### method Service Drag
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.DragComponent]`

#### method Service Dragged
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.DraggedComponent]`

#### method Service Focused
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.FocusedComponent]`

#### method Service Hover
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.HoverComponent]`

#### method Service Hovered
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.HoveredComponent]`

#### method Service IsCaptured
Type: `func(engine/services/ecs.EntityID) bool`
checks is entity event captured
if it is then it shouldn't listen to keyboard events

#### method Service KeepSelected
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.KeepSelectedComponent]`

#### method Service LeftClick
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.LeftClickComponent]`

#### method Service MouseEnter
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.MouseEnterComponent]`

#### method Service MouseLeave
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.MouseLeaveComponent]`

#### method Service Register
Type: `func() error`

#### method Service RightClick
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.RightClickComponent]`

#### method Service Stack
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.StackComponent]`

#### method Service Stacked
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.StackedComponent]`

#### method Service StackedData
Type: `func() []engine/modules/inputs.Target`
returns ordered targets with additional data

#### method Service TextInput
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/inputs.TextInputComponent]`

### type CaptureKeyboardConstraint
Type: `engine/modules/inputs.CaptureKeyboardConstraint`
on keyboardEvent capture events are emitted
from child with [FocusedComponent]
to uppermost parent with [CaptureKeyboardComponent] with fallthrough == false
if none [CaptureKeyboardComponent] stops further emission then [KeyboardEvent] is emited

#### method CaptureKeyboardConstraint Capture
Type: `func(engine/modules/inputs.KeyboardEvent) any`

#### method CaptureKeyboardConstraint Fallthrough
Type: `func() bool`

### type ApplyDragEvent
Type: `engine/modules/inputs.ApplyDragEvent`
interfaces which can be implemented by events {

#### method ApplyDragEvent ApplyDrag
Type: `func(engine/modules/inputs.DragEvent) (event any)`

### type ApplyEntityEvent
Type: `engine/modules/inputs.ApplyEntityEvent`

#### method ApplyEntityEvent ApplyEntity
Type: `func(entityEmitting engine/services/ecs.EntityID) (event any)`

### type EventTargetSetter
Type: `engine/modules/inputs.EventTargetSetter`

#### method EventTargetSetter SetTarget
Type: `func(engine/modules/inputs.Target) engine/modules/inputs.EventTargetSetter`

### type KeyboardEvent
Type: `engine/modules/inputs.KeyboardEvent`

#### property KeyboardEvent KeyboardEvent
Type: `github.com/veandco/go-sdl2/sdl.KeyboardEvent`

### type UnfocusEvent
Type: `engine/modules/inputs.UnfocusEvent`
focuses default entity like scene or camera

### type FocusEvent
Type: `engine/modules/inputs.FocusEvent`
unfocuses all elements and only focuses specific one

#### property FocusEvent Entity
Type: `engine/services/ecs.EntityID`

#### method FocusEvent ApplyEntity
Type: `func(entity engine/services/ecs.EntityID) any`

### type DefaultFocusEvent
Type: `engine/modules/inputs.DefaultFocusEvent`

#### property DefaultFocusEvent Entity
Type: `engine/services/ecs.EntityID`

### type CaptureKeyboardComponent
Type: `engine/modules/inputs.CaptureKeyboardComponent`

#### property CaptureKeyboardComponent CaptureKeyboardConstraint
Type: `engine/modules/inputs.CaptureKeyboardConstraint`

### type DefaultFocusedComponent
Type: `engine/modules/inputs.DefaultFocusedComponent`

### type FocusedComponent
Type: `engine/modules/inputs.FocusedComponent`
element should be focused on click for example
on right click or escape element should get unfocused

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
handles inputs and saves change in description
this can become generic

#### property TextInputEvent Entity
Type: `engine/services/ecs.EntityID`

#### property TextInputEvent KeyboardEvent
Type: `engine/modules/inputs.KeyboardEvent`

#### method TextInputEvent Capture
Type: `func(event engine/modules/inputs.KeyboardEvent) any`

#### method TextInputEvent ApplyEntity
Type: `func(entity engine/services/ecs.EntityID) any`

#### method TextInputEvent Fallthrough
Type: `func() bool`

### type HoveredComponent
Type: `engine/modules/inputs.HoveredComponent`
many elements can be hovered at once

#### property HoveredComponent Camera
Type: `engine/services/ecs.EntityID`

### type DraggedComponent
Type: `engine/modules/inputs.DraggedComponent`

#### property DraggedComponent Camera
Type: `engine/services/ecs.EntityID`

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
Type: `engine/services/ecs.EntityID`

#### property DragEvent From
Type: `engine/modules/window.MousePos`
from and to is normalized

#### property DragEvent To
Type: `engine/modules/window.MousePos`
from and to is normalized

### type SynchronizePositionEvent
Type: `engine/modules/inputs.SynchronizePositionEvent`

#### property SynchronizePositionEvent Camera
Type: `engine/services/ecs.EntityID`

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
Type: `engine/services/ecs.EntityID`

## Functions
### func NewKeyboardEvent
Type: `func(e github.com/veandco/go-sdl2/sdl.KeyboardEvent) engine/modules/inputs.KeyboardEvent`

### func NewUnfocusEvent
Type: `func() engine/modules/inputs.UnfocusEvent`

### func NewFocusEvent
Type: `func(entity engine/services/ecs.EntityID) engine/modules/inputs.FocusEvent`

### func NewDefaultFocusEvent
Type: `func(entity engine/services/ecs.EntityID) engine/modules/inputs.DefaultFocusEvent`

### func NewCaptureKeyboard
Type: `func(event engine/modules/inputs.CaptureKeyboardConstraint) engine/modules/inputs.CaptureKeyboardComponent`

### func NewDefaultFocused
Type: `func() engine/modules/inputs.DefaultFocusedComponent`

### func NewFocused
Type: `func() engine/modules/inputs.FocusedComponent`

### func NewTextInputEvent
Type: `func() engine/modules/inputs.TextInputEvent`

### func NewHovered
Type: `func(camera engine/services/ecs.EntityID) engine/modules/inputs.HoveredComponent`

### func NewDragged
Type: `func(camera engine/services/ecs.EntityID) engine/modules/inputs.DraggedComponent`

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
  - `engine.Hierarchy`
  - `engine.Inputs`
  - `engine.Logger`
  - `engine.Scene`
  - `engine.Text`
  - `engine.Window`
  - `engine.World`

`engine/modules/collider`:
  - `engine/modules/collider.Distance`
  - `engine/modules/collider.Entity`
  - `engine/modules/collider.Hit`
  - `engine/modules/collider.ObjectRayCollision`
  - `engine/modules/collider.RaycastAll`

`engine/modules/inputs`:
  - `engine/modules/inputs.ApplyDrag`
  - `engine/modules/inputs.ApplyDragEvent`
  - `engine/modules/inputs.ApplyEntity`
  - `engine/modules/inputs.ApplyEntityEvent`
  - `engine/modules/inputs.Camera`
  - `engine/modules/inputs.Capture`
  - `engine/modules/inputs.CaptureKeyboard`
  - `engine/modules/inputs.CaptureKeyboardComponent`
  - `engine/modules/inputs.CursorCol`
  - `engine/modules/inputs.CursorRow`
  - `engine/modules/inputs.DefaultFocusEvent`
  - `engine/modules/inputs.DefaultFocused`
  - `engine/modules/inputs.DefaultFocusedComponent`
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
  - `engine/modules/inputs.Fallthrough`
  - `engine/modules/inputs.FocusEvent`
  - `engine/modules/inputs.Focused`
  - `engine/modules/inputs.FocusedComponent`
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
  - `engine/modules/inputs.NewDefaultFocusEvent`
  - `engine/modules/inputs.NewDefaultFocused`
  - `engine/modules/inputs.NewFocused`
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
  - `engine/modules/inputs.UnfocusEvent`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.NewStopEvent`

`engine/modules/scene`:
  - `engine/modules/scene.ChangeSceneEvent`
  - `engine/modules/scene.Scene`

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

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityExists`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.OnMod`
  - `engine/services/ecs.RegisterSystems`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/sdl`