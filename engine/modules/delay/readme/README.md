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