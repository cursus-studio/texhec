# loading
## Architecture
defines GUI to show up when batcher processes any task

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             19              1             88
-------------------------------------------------------------------------------
SUM:                             3             19              1             88
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/loading.Service`

#### method Service Register
Type: `func() error`


## Dependencies
`core/game`:
  - `core/game.GameWorld`
  - `core/game.Ui`

`core/modules/loading`:
  - `core/modules/loading.Service`

`core/modules/ui`:
  - `core/modules/ui.AnimatedBackground`
  - `core/modules/ui.AnimatedBackgroundComponent`

`engine/modules/camera`:
  - `engine/modules/camera.NewOrtho`
  - `engine/modules/camera.Ortho`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`

`engine/modules/text`:
  - `engine/modules/text.Align`
  - `engine/modules/text.Break`
  - `engine/modules/text.BreakNone`
  - `engine/modules/text.Content`
  - `engine/modules/text.FontSize`
  - `engine/modules/text.NewAlign`
  - `engine/modules/text.NewBreak`
  - `engine/modules/text.NewFontSize`
  - `engine/modules/text.NewText`

`engine/modules/transform`:
  - `engine/modules/transform.Inherit`
  - `engine/modules/transform.NewInherit`
  - `engine/modules/transform.NewParentPivotPoint`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.ParentPivotPoint`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.RelativePos`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`