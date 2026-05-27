# focus
## Architecture
Creates all focus related components and events and allows event bubbling from focused or any other entities

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               5             47             22            214
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             6             47             22            215
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/focus.Service`

#### method Service Bubbling
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/focus.BubblingComponent]`
bubbling

#### method Service DefaultFocused
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/focus.DefaultFocusedComponent]`
focus

#### method Service DryRun
Type: `func(engine/modules/focus.BubbleEvent) (bubbles []engine/services/ecs.EntityID, captured bool)`

#### method Service Emit
Type: `func(engine/modules/focus.BubbleEvent)`

#### method Service Focused
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/focus.FocusedComponent]`

#### method Service FocusedEntity
Type: `func() (engine/services/ecs.EntityID, bool)`

#### method Service NewFocusedBubbleEvent
Type: `func(event any) (engine/modules/focus.BubbleEvent, bool)`

### type BubblingConstraint
Type: `engine/modules/focus.BubblingConstraint`
captures events that are emitted
from child with [FocusedComponent]
to uppermost parent with [BubblingComponent] with fallthrough(T) == false
if none [BubblingComponent] stops further emission then [Event] is emited

#### method BubblingConstraint Capture
Type: `func(any) any`
should wrap event and return wrapping event.

#### method BubblingConstraint CapturesEvents
Type: `func() engine/services/datastructures.SetReader[reflect.Type]`
stores a list of events which can be passed to capture
this should be a global variable it never should be stored in component

#### method BubblingConstraint Fallthrough
Type: `func() bool`
should return a constant

### type BubblingComponent
Type: `engine/modules/focus.BubblingComponent`
This implements bubbling

#### property BubblingComponent BubblingConstraint
Type: `engine/modules/focus.BubblingConstraint`

### type BubbleEvent
Type: `engine/modules/focus.BubbleEvent`

#### property BubbleEvent Entity
Type: `engine/services/ecs.EntityID`

#### property BubbleEvent Event
Type: `any`
golang generics are to restrictive to use them.
this has to use any because propagating it everywhere would require to granural configuration everywhere

#### property BubbleEvent EventType
Type: `reflect.Type`

### type UnfocusEvent
Type: `engine/modules/focus.UnfocusEvent`
focuses default entity like scene or camera

### type FocusEvent
Type: `engine/modules/focus.FocusEvent`
unfocuses all elements and only focuses specific one

#### property FocusEvent Entity
Type: `engine/services/ecs.EntityID`

#### method FocusEvent ApplyEntity
Type: `func(entity engine/services/ecs.EntityID) any`

### type DefaultFocusEvent
Type: `engine/modules/focus.DefaultFocusEvent`

#### property DefaultFocusEvent Entity
Type: `engine/services/ecs.EntityID`

### type FocusedComponent
Type: `engine/modules/focus.FocusedComponent`
element should be focused on click for example
on right click or escape element should get unfocused

### type DefaultFocusedComponent
Type: `engine/modules/focus.DefaultFocusedComponent`

## Functions
### func NewBubbling
Type: `func(event engine/modules/focus.BubblingConstraint) engine/modules/focus.BubblingComponent`

### func NewBubbleEvent
Type: `func(entity engine/services/ecs.EntityID, event any) engine/modules/focus.BubbleEvent`

### func NewUnfocusEvent
Type: `func() engine/modules/focus.UnfocusEvent`

### func NewFocusEvent
Type: `func(entity engine/services/ecs.EntityID) engine/modules/focus.FocusEvent`

### func NewDefaultFocusEvent
Type: `func(entity engine/services/ecs.EntityID) engine/modules/focus.DefaultFocusEvent`

### func NewFocused
Type: `func() engine/modules/focus.FocusedComponent`

### func NewDefaultFocused
Type: `func() engine/modules/focus.DefaultFocusedComponent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Focus`
  - `engine.Hierarchy`
  - `engine.Logger`
  - `engine.Scene`
  - `engine.World`

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

`engine/services/datastructures`:
  - `engine/services/datastructures.GetIndex`
  - `engine/services/datastructures.SetReader`

`engine/services/ecs`:
  - `engine/services/ecs.ApplyEntity`
  - `engine/services/ecs.ApplyEntityEvent`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.Set`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`