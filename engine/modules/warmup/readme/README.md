# warmup
## Architecture
runs all lazy listeners
it should be used executing anything concurrently

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             10              2             40
-------------------------------------------------------------------------------
SUM:                             3             10              2             40
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`

`engine/modules/ecs`:
  - `engine/modules/ecs.World`

`engine/modules/warmup`:
  - `engine/modules/warmup.Event`
  - `engine/modules/warmup.Service`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`