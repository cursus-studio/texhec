# smooth
## Architecture
catches changes on tick and applies them smoothly between ticks

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               4             37             13            230
-------------------------------------------------------------------------------
SUM:                             4             37             13            230
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/smooth.Service`

#### method Service Start
Type: `func() engine/modules/ecs.SystemRegister`

#### method Service Stop
Type: `func() engine/modules/ecs.SystemRegister`

### type ServiceT
Type: `engine/modules/smooth.ServiceT[StateComponent any]`

#### method ServiceT AddWaypoint
Type: `func(engine/modules/ecs.EntityID, StateComponent)`
AddWaypoint appends a state snapshot for the next tick interval.
Multiple waypoints within a single tick period are distributed evenly across the frame duration.

### type SmoothConstraint
Type: `engine/modules/smooth.SmoothConstraint[Component any]`

#### method SmoothConstraint Lerp
Type: `func(Component, float32) Component`

#### method SmoothConstraint Smooth
Type: `func()`
this method is a tag that component is smoothed
each lerpable component with this method will automatically be registered to be smoothed

### type AddWaypointEvent
Type: `engine/modules/smooth.AddWaypointEvent[StateComponent any]`

#### property AddWaypointEvent Entity
Type: `engine/modules/ecs.EntityID`

#### property AddWaypointEvent State
Type: `StateComponent`

## Functions
### func NewAddWaypointEvent
Type: `func[StateComponent any](entity engine/modules/ecs.EntityID, state StateComponent) engine/modules/smooth.AddWaypointEvent[StateComponent]`


## Dependencies
`engine`:
  - `engine.Delay`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Record`
  - `engine.World`

`engine/modules/delay`:
  - `engine/modules/delay.Delay`
  - `engine/modules/delay.NewDelayedEvent`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.TickEvent`

`engine/modules/record`:
  - `engine/modules/record.AddToConfig`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entities`
  - `engine/modules/record.Entity`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.StartBackwardsRecording`
  - `engine/modules/record.Stop`

`engine/modules/smooth`:
  - `engine/modules/smooth.AddWaypointEvent`
  - `engine/modules/smooth.Entity`
  - `engine/modules/smooth.Service`
  - `engine/modules/smooth.ServiceT`
  - `engine/modules/smooth.SmoothConstraint`
  - `engine/modules/smooth.State`

`engine/modules/transition`:
  - `engine/modules/transition.LerpConstraint`
  - `engine/modules/transition.NewTransition`
  - `engine/modules/transition.TransitionComponent`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`