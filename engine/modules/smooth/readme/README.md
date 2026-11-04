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