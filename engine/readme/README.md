# TEXHEC engine

## Architecture
This module defines everything what every game engine should have but with one caveat.
It places efficient data layout as a first class citizen.

It follows **DOD** (data oriented design) and stores all game objects using **ECS** (entity component system).

## Modules
- [assets](/engine/modules/assets)
- [audio](/engine/modules/audio)
- [batcher](/engine/modules/batcher)
- [camera](/engine/modules/camera)
- [codec](/engine/modules/codec)
- [collider](/engine/modules/collider)
- [connection](/engine/modules/connection)
- [drag](/engine/modules/drag)
- [entityregistry](/engine/modules/entityregistry)
- [graphics](/engine/modules/graphics)
- [grid](/engine/modules/grid)
- [groups](/engine/modules/groups)
- [hierarchy](/engine/modules/hierarchy)
- [inputs](/engine/modules/inputs)
- [layout](/engine/modules/layout)
- [logger](/engine/modules/logger)
- [loop](/engine/modules/loop)
- [metadata](/engine/modules/metadata)
- [netsync](/engine/modules/netsync)
- [noise](/engine/modules/noise)
- [prototype](/engine/modules/prototype)
- [record](/engine/modules/record)
- [relation](/engine/modules/relation)
- [render](/engine/modules/render)
- [scene](/engine/modules/scene)
- [seed](/engine/modules/seed)
- [smooth](/engine/modules/smooth)
- [text](/engine/modules/text)
- [transform](/engine/modules/transform)
- [transition](/engine/modules/transition)
- [typeregistry](/engine/modules/typeregistry)
- [uuid](/engine/modules/uuid)
- [warmup](/engine/modules/warmup)
- [window](/engine/modules/window)
## Services
- [bitmasks](/engine/services/bitmasks)
- [clock](/engine/services/clock)
- [console](/engine/services/console)
- [datastructures](/engine/services/datastructures)
- [ecs](/engine/services/ecs)

## Challenges
Biggest challenge was creating framework while building on top of it.
Testing own framework edgecases while delivering is impossible without frequent refactors.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                             264           2900            766          13681
GLSL                             5             35              4             99
Markdown                         5              4              0             50
-------------------------------------------------------------------------------
SUM:                           274           2939            770          13830
-------------------------------------------------------------------------------

```
## Dependencies
### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/go-gl/mathgl/mgl64`
- `github.com/google/uuid`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/mix`
- `github.com/veandco/go-sdl2/sdl`
- `golang.org/x/exp/constraints`
- `golang.org/x/image/draw`
- `golang.org/x/image/font`
- `golang.org/x/image/font/opentype`
- `golang.org/x/image/math/fixed`