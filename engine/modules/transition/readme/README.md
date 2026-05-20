# transition
## Architecture
changes components from stateA to stateB across many frames

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             67             15            265
-------------------------------------------------------------------------------
SUM:                             6             67             15            265
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/transition.Service`

#### method Service Easing
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/transition.EasingComponent]`

#### method Service EasingFunction
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/transition.EasingFunctionComponent]`

#### method Service Register
Type: `func() error`

### type LerpConstraint
Type: `engine/modules/transition.LerpConstraint[Component any]`

#### method LerpConstraint Lerp
Type: `func(Component, float32) Component`

### type Progress
Type: `engine/modules/transition.Progress`

### type TransitionComponent
Type: `engine/modules/transition.TransitionComponent[Component any]`
type TransitionComponent[Component LerpConstraint[Component]] struct {

#### property TransitionComponent From
Type: `Component`

#### property TransitionComponent To
Type: `Component`

#### property TransitionComponent Progress
Type: `time.Duration`

#### property TransitionComponent Duration
Type: `time.Duration`

### type TransitionEvent
Type: `engine/modules/transition.TransitionEvent[Component any]`
saves transition component
type TransitionEvent[Component LerpConstraint[Component]] struct {

#### property TransitionEvent Entity
Type: `engine/services/ecs.EntityID`

#### property TransitionEvent Component
Type: `engine/modules/transition.TransitionComponent[Component]`

### type DelayedEvent
Type: `engine/modules/transition.DelayedEvent`

#### property DelayedEvent Event
Type: `any`

#### property DelayedEvent Duration
Type: `time.Duration`

### type EasingComponent
Type: `engine/modules/transition.EasingComponent`

#### property EasingComponent ID
Type: `engine/services/ecs.EntityID`

### type EasingFunctionComponent
Type: `engine/modules/transition.EasingFunctionComponent`

#### property EasingFunctionComponent EasingFunction
Type: `func(t engine/modules/transition.Progress) engine/modules/transition.Progress`

## Functions
### func Lerp
Type: `func[Number, T golang.org/x/exp/constraints.Float](a Number, b Number, t T) Number`

### func NewTransition
Type: `func[Component any](from Component, to Component, duration time.Duration) engine/modules/transition.TransitionComponent[Component]`
func NewTransition[Component LerpConstraint[Component]](

### func NewTransitionEvent
Type: `func[Component engine/modules/transition.LerpConstraint[Component]](entity engine/services/ecs.EntityID, from Component, to Component, duration time.Duration) engine/modules/transition.TransitionEvent[Component]`

### func NewDelayedEvent
Type: `func(event any, duration time.Duration) engine/modules/transition.DelayedEvent`

### func NewEasing
Type: `func(id engine/services/ecs.EntityID) engine/modules/transition.EasingComponent`

### func NewEasingFunction
Type: `func(easingFunction func(t engine/modules/transition.Progress) engine/modules/transition.Progress) engine/modules/transition.EasingFunctionComponent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Transition`
  - `engine.World`

`engine/modules/codec/pkg`:
  - `engine/modules/codec/pkg.PkgT`

`engine/modules/loop`:
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.FrameEvent`

`engine/modules/prototype/pkg`:
  - `engine/modules/prototype/pkg.PkgT`

`engine/modules/transition`:
  - `engine/modules/transition.Component`
  - `engine/modules/transition.DelayedEvent`
  - `engine/modules/transition.Duration`
  - `engine/modules/transition.EasingComponent`
  - `engine/modules/transition.EasingFunction`
  - `engine/modules/transition.EasingFunctionComponent`
  - `engine/modules/transition.Entity`
  - `engine/modules/transition.Event`
  - `engine/modules/transition.From`
  - `engine/modules/transition.ID`
  - `engine/modules/transition.Lerp`
  - `engine/modules/transition.LerpConstraint`
  - `engine/modules/transition.Progress`
  - `engine/modules/transition.Service`
  - `engine/modules/transition.To`
  - `engine/modules/transition.TransitionComponent`
  - `engine/modules/transition.TransitionEvent`

`engine/services/ecs`:
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.Register`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`