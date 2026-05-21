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
- [definitions](/core/modules/definitions/readme/README.md)
- [deploy](/core/modules/deploy/readme/README.md)
- [fpslogger](/core/modules/fpslogger/readme/README.md)
- [generation](/core/modules/generation/readme/README.md)
- [loading](/core/modules/loading/readme/README.md)
- [obstruction](/core/modules/obstruction/readme/README.md)
- [pathfind](/core/modules/pathfind/readme/README.md)
- [player](/core/modules/player/readme/README.md)
- [settings](/core/modules/settings/readme/README.md)
- [tile](/core/modules/tile/readme/README.md)
- [ui](/core/modules/ui/readme/README.md)

## Challenges
Challenge of this module is to create something on incomplete foundation (`engine`)
without ending up in spagetti relations or broken code.
It needs to balance stability and features.

`grid` module which were made to deliver pre-mature features.
It stores whole map in a single slice where it should chunk the map.

## Lines of code
```

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