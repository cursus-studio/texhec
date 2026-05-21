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
## Types
### type Service
Type: `core/modules/fpslogger.Service`

#### method Service Register
Type: `func() error`


## Dependencies
`core/game`:
  - `core/game.GameWorld`

`core/modules/fpslogger`:
  - `core/modules/fpslogger.Service`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`

`engine/services/ecs`:
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`