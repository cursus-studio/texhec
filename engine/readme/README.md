# TEXHEC engine

## Architecture
This module defines everything what every game engine should have but with one caveat.
It places efficient data layout as a first class citizen.

It follows **DOD** (data oriented design) and stores all game objects using **ECS** (entity component system).

## Modules
- [assets](/engine/modules/assets/readme/README.md)
- [audio](/engine/modules/audio/readme/README.md)
- [batcher](/engine/modules/batcher/readme/README.md)
- [bitmasks](/engine/modules/bitmasks/readme/README.md)
- [camera](/engine/modules/camera/readme/README.md)
- [clock](/engine/modules/clock/readme/README.md)
- [codec](/engine/modules/codec/readme/README.md)
- [collider](/engine/modules/collider/readme/README.md)
- [connection](/engine/modules/connection/readme/README.md)
- [console](/engine/modules/console/readme/README.md)
- [datastructures](/engine/modules/datastructures/readme/README.md)
- [delay](/engine/modules/delay/readme/README.md)
- [drag](/engine/modules/drag/readme/README.md)
- [ecs](/engine/modules/ecs/readme/README.md)
- [entityregistry](/engine/modules/entityregistry/readme/README.md)
- [focus](/engine/modules/focus/readme/README.md)
- [graphics](/engine/modules/graphics/readme/README.md)
- [grid](/engine/modules/grid/readme/README.md)
- [groups](/engine/modules/groups/readme/README.md)
- [hierarchy](/engine/modules/hierarchy/readme/README.md)
- [inputs](/engine/modules/inputs/readme/README.md)
- [interactions](/engine/modules/interactions/readme/README.md)
- [layout](/engine/modules/layout/readme/README.md)
- [logger](/engine/modules/logger/readme/README.md)
- [loop](/engine/modules/loop/readme/README.md)
- [metadata](/engine/modules/metadata/readme/README.md)
- [netsync](/engine/modules/netsync/readme/README.md)
- [noise](/engine/modules/noise/readme/README.md)
- [prototype](/engine/modules/prototype/readme/README.md)
- [record](/engine/modules/record/readme/README.md)
- [relation](/engine/modules/relation/readme/README.md)
- [render](/engine/modules/render/readme/README.md)
- [scene](/engine/modules/scene/readme/README.md)
- [seed](/engine/modules/seed/readme/README.md)
- [smooth](/engine/modules/smooth/readme/README.md)
- [text](/engine/modules/text/readme/README.md)
- [transform](/engine/modules/transform/readme/README.md)
- [transition](/engine/modules/transition/readme/README.md)
- [typeregistry](/engine/modules/typeregistry/readme/README.md)
- [uuid](/engine/modules/uuid/readme/README.md)
- [warmup](/engine/modules/warmup/readme/README.md)
- [window](/engine/modules/window/readme/README.md)

## Challenges
Biggest challenge was creating framework while building on top of it.
Testing own framework edgecases while delivering is impossible without frequent refactors.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                             288           3131            964          14841
Markdown                        22             51              0            241
GLSL                             5             35              4             99
-------------------------------------------------------------------------------
SUM:                           315           3217            968          15181
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