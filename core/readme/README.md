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
[definitions](/core/modules/definitions)
[deploy](/core/modules/deploy)
[fpslogger](/core/modules/fpslogger)
[generation](/core/modules/generation)
[loading](/core/modules/loading)
[obstruction](/core/modules/obstruction)
[pathfind](/core/modules/pathfind)
[player](/core/modules/player)
[settings](/core/modules/settings)
[tile](/core/modules/tile)
[ui](/core/modules/ui)

## Challenges
Challenge of this module is to create something on incomplete foundation (`engine`)
without ending up in spagetti relations or broken code.
It needs to balance stability and features.

`grid` module which were made to deliver pre-mature features.
It stores whole map in a single slice where it should chunk the map.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              57            726            336           3945
GLSL                             3             29              2            114
Markdown                         6             15              0             65
-------------------------------------------------------------------------------
SUM:                            66            770            338           4124
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