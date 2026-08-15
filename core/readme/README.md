# TEXHEC CORE

## Architecture
Core module is responsible for implementing all `TEXHEC` game modules.

Implemented scenes:
- `menu`
- `game`
- `credits`
- `settings`

Map scale preview:
![Map scroll](/readme/map_scroll.gif)

## Modules
- [actions](/core/modules/actions/readme/README.md)
- [attack](/core/modules/attack/readme/README.md)
- [definitions](/core/modules/definitions/readme/README.md)
- [deploy](/core/modules/deploy/readme/README.md)
- [economy](/core/modules/economy/readme/README.md)
- [fpslogger](/core/modules/fpslogger/readme/README.md)
- [generation](/core/modules/generation/readme/README.md)
- [loading](/core/modules/loading/readme/README.md)
- [obstruction](/core/modules/obstruction/readme/README.md)
- [pathfind](/core/modules/pathfind/readme/README.md)
- [player](/core/modules/player/readme/README.md)
- [reach](/core/modules/reach/readme/README.md)
- [settings](/core/modules/settings/readme/README.md)
- [tile](/core/modules/tile/readme/README.md)
- [ui](/core/modules/ui/readme/README.md)

## Challenges
Challenge of this project is to create evolving foundation (`engine`) on which
(`core`) can rely while its still changing.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              77            919            272           5451
Markdown                        17             25              0            152
GLSL                             3             31              2            112
-------------------------------------------------------------------------------
SUM:                            97            975            274           5715
-------------------------------------------------------------------------------
```
## Dependencies
### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/go-gl/mathgl/mgl64`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/sdl`
- `golang.org/x/exp/constraints`