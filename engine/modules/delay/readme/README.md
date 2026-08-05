# delay
## Architecture
package responsible for delaying events

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             21              3             85
-------------------------------------------------------------------------------
SUM:                             3             21              3             85
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/delay.Service`

#### method Service Delay
Type: `func(engine/modules/delay.DelayedEvent)`

#### method Service Register
Type: `func() error`

### type ApplyDelay
Type: `engine/modules/delay.ApplyDelay`

#### method ApplyDelay ApplyDelay
Type: `func(time.Duration) any`

### type DelayedEvent
Type: `engine/modules/delay.DelayedEvent`
emits event after duration on frame

#### property DelayedEvent Event
Type: `engine/modules/delay.ApplyDelay`

#### property DelayedEvent Duration
Type: `time.Duration`

## Functions
### func NewDelayedEvent
Type: `func(event any, duration time.Duration) engine/modules/delay.DelayedEvent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`

`engine/modules/delay`:
  - `engine/modules/delay.ApplyDelay`
  - `engine/modules/delay.DelayedEvent`
  - `engine/modules/delay.Duration`
  - `engine/modules/delay.Event`
  - `engine/modules/delay.Service`

`engine/modules/ecs`:
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.FrameEvent`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`