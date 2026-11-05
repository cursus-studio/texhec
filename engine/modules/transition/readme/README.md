# transition
## Architecture
changes components from stateA to stateB across many frames

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             57             13            231
-------------------------------------------------------------------------------
SUM:                             6             57             13            231
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.EventsBuilder`
  - `engine.Transition`
  - `engine.World`

`engine/modules/codec/pkg`:
  - `engine/modules/codec/pkg.PkgT`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.Register`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.FrameEvent`

`engine/modules/prototype/pkg`:
  - `engine/modules/prototype/pkg.PkgT`

`engine/modules/transition`:
  - `engine/modules/transition.Component`
  - `engine/modules/transition.Duration`
  - `engine/modules/transition.EasingComponent`
  - `engine/modules/transition.EasingFunction`
  - `engine/modules/transition.EasingFunctionComponent`
  - `engine/modules/transition.Entity`
  - `engine/modules/transition.From`
  - `engine/modules/transition.ID`
  - `engine/modules/transition.Lerp`
  - `engine/modules/transition.LerpConstraint`
  - `engine/modules/transition.Progress`
  - `engine/modules/transition.Service`
  - `engine/modules/transition.To`
  - `engine/modules/transition.TransitionComponent`
  - `engine/modules/transition.TransitionEvent`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`