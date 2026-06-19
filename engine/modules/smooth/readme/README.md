# smooth
## Architecture
catches changes on tick and applies them smoothly between ticks

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               4             25              8            133
-------------------------------------------------------------------------------
SUM:                             4             25              8            133
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/smooth.Service`

#### method Service Start
Type: `func() engine/modules/ecs.SystemRegister`

#### method Service Stop
Type: `func() engine/modules/ecs.SystemRegister`

### type SmoothConstraint
Type: `engine/modules/smooth.SmoothConstraint[Component any]`

#### method SmoothConstraint Lerp
Type: `func(Component, float32) Component`

#### method SmoothConstraint Smooth
Type: `func()`
this method is a tag that component is smooothed


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Record`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
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
  - `engine/modules/smooth.Service`
  - `engine/modules/smooth.SmoothConstraint`

`engine/modules/transition`:
  - `engine/modules/transition.LerpConstraint`
  - `engine/modules/transition.NewTransition`
  - `engine/modules/transition.To`
  - `engine/modules/transition.TransitionComponent`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`