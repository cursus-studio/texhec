# fpslogger
## Architecture
logs frames per second

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             15              2             56
-------------------------------------------------------------------------------
SUM:                             3             15              2             56
-------------------------------------------------------------------------------
```
## Dependencies
`core/game`:
  - `core/game.GameWorld`

`core/modules/fpslogger`:
  - `core/modules/fpslogger.Service`

`engine/modules/ecs`:
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`