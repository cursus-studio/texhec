# window
## Architecture
integrates window management

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             23              5            133
-------------------------------------------------------------------------------
SUM:                             3             23              5            133
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Logger`

`engine/modules/window`:
  - `engine/modules/window.Elem`
  - `engine/modules/window.MousePos`
  - `engine/modules/window.NewMousePos`
  - `engine/modules/window.Service`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/mix`
- `github.com/veandco/go-sdl2/sdl`