# settings
## Architecture
renders settings GUI

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             30              8            106
-------------------------------------------------------------------------------
SUM:                             3             30              8            106
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/settings.Service`

#### method Service Register
Type: `func() error`

### type EnterSettingsEvent
Type: `core/modules/settings.EnterSettingsEvent`

### type EnterSettingsForParentEvent
Type: `core/modules/settings.EnterSettingsForParentEvent`

#### property EnterSettingsForParentEvent Parent
Type: `engine/services/ecs.EntityID`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.GameWorld`
  - `core/game.Ui`

`core/modules/definitions`:
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.Btn`
  - `core/modules/definitions.EffectChannel`
  - `core/modules/definitions.ExampleAudio`
  - `core/modules/definitions.Hud`
  - `core/modules/definitions.MenuID`

`core/modules/settings`:
  - `core/modules/settings.EnterSettingsEvent`
  - `core/modules/settings.EnterSettingsForParentEvent`
  - `core/modules/settings.Parent`
  - `core/modules/settings.Service`

`engine/modules/audio`:
  - `engine/modules/audio.NewPlayEvent`

`engine/modules/groups`:
  - `engine/modules/groups.Inherit`
  - `engine/modules/groups.InheritGroupsComponent`

`engine/modules/inputs`:
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

`engine/modules/loop`:
  - `engine/modules/loop.TickEvent`

`engine/modules/scene`:
  - `engine/modules/scene.NewChangeSceneEvent`

`engine/modules/text`:
  - `engine/modules/text.Align`
  - `engine/modules/text.Content`
  - `engine/modules/text.FontSize`
  - `engine/modules/text.NewAlign`
  - `engine/modules/text.NewFontSize`
  - `engine/modules/text.NewText`

`engine/modules/transform`:
  - `engine/modules/transform.NewParent`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.Parent`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeSizeX`
  - `engine/modules/transform.Size`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/ecs`:
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`