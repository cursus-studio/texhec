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
## Types
### type Service
Type: `engine/modules/window.Service`

#### method Service Ctx
Type: `func() github.com/veandco/go-sdl2/sdl.GLContext`

#### method Service GetClipboard
Type: `func() (string, bool)`

#### method Service GetMousePos
Type: `func() engine/modules/window.MousePos`

#### method Service NormalizeMousePos
Type: `func(engine/modules/window.MousePos) github.com/go-gl/mathgl/mgl32.Vec2`

#### method Service SetClipboard
Type: `func(string)`

#### method Service Window
Type: `func() *github.com/veandco/go-sdl2/sdl.Window`

### type MousePos
Type: `engine/modules/window.MousePos`

#### property MousePos X
Type: `int32`

#### property MousePos Y
Type: `int32`

#### method MousePos Elem
Type: `func() (x int32, y int32)`

## Functions
### func NewMousePos
Type: `func(x int32, y int32) engine/modules/window.MousePos`


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