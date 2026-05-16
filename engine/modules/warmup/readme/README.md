# warmup
## Architecture
runs all lazy listeners
it should be used executing anything concurrently

## Types
### type Service
Type: `engine/modules/warmup.Service`

#### method Service WarmUp
Type: `func()`

### type Event
Type: `engine/modules/warmup.Event`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`

`engine/modules/warmup`:
  - `engine/modules/warmup.Event`
  - `engine/modules/warmup.Service`

`engine/services/ecs`:
  - `engine/services/ecs.WarmUp`
  - `engine/services/ecs.World`

### Third Party
`github.com/ogiusek/events`
`github.com/ogiusek/ioc/v2`