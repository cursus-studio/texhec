# TEXHEC
## What is TEXHEC ?
TEXHEC is a **HIGH-Performance** project where natural map size is **1.000.000** tiles with\
hundreds or thousands buildings and units **all** being **simulated** in real time on\
low end hardware Intel® Core™ i5-8350U × 8 Intel® and UHD Graphics 620 (KBL GT2).\
We use **DOD** and use our **own** **ECS** and **DI container** framework.

We focus on what works. What is scalable and what is performant.\
We achieved framework which reduced boilerplate and allowed us to write 40+ modules in less than 20k loc.

[More about **ECS** framework](/engine/services/ecs/readme/README.md)
[More about **IOC** framework](https://github.com/ogiusek/ioc)

## Architectural choices
### Why golang
Others would **discard golang** due to **garbage collector**.\
In reality garbage collector isn't an inconvenience because we follow **DOD** and\
we do not have enough pointers to be an inconvenience.

In reality using golang has benefits:
- its very performant (its compiled)
- its fast to write, understand and its very easy to use (necessary to deliver by a single developer)
- it lacks decades of building technical debt
- aligned philosophies (simplicity creates performance not other way around)

### Dependencies
Dependencies are only added if necessary.
- `sdl2`
- `opengl`
- `opengl math`
- `ghw` - to read gpu name for benchmarks
- `golang constraints`
- `golang images and text (used only to generate image per letter)`
- `google uuid`

Dependencies which are written by me:
- `ioc`
- `events`

### Module vs Service
Service is something separate from game engine which is basis for it.\
After creating **ECS** service i attempt to migrate everything to a module.\
Modules also have more struct rules and have dedicated file structure.\
Services are more detached from alone game engine and have less strict rules.

### Module structure
```
modules/
└─ `$module_name`/
    ├── internal/       # Defines implementation for `Service` and `System` (if exist in module)
    ├── pkg/            # This exposes `Package` function to register `Service` implementation.
    │                   # `pkg`, `internal` and `test` separation allows `modules`
    │                   # Decouples the interface definition from the construction logic to allow for flexible dependency wiring
    ├── test/           # Defines test
    ├── readme/         # Defines readme
    └── `$interface.go` # There is no strict file rule naming. This defines what module exposes
                        # Expects interface name `Service` so module name and service purpose were related
```
Everything in module file structure is optional and should be only added if used.

### Module readme schema
```md
# Module_name
## Architecture
How module is built and general flow of data and why this way in case of controversial choices.

## Benchmarks (optional)

## Usage examples
Code snippets of `Service` and of how to use it.
...
```

### Implicit dependency graph
#### How explicit dependency graph looks ?
```modules/mod3/internal.go
type service struct {
    Dep1 mod1.Service `inject:""`
    Dep2 mod2.Service `inject:""`
}
```
This explicitly states all used dependencies.

**Pros**:
- Easier to read module dependencies
- Faster
**Cons**:
- Additional maintenance cost and loc
- Fragments the engine

#### How implicit dependency graph looks ?
```engine.go
type EngineWorld struct {
    Dep1 ioc.Lazy[mod1.Service] `inject:""`
    Dep2 ioc.Lazy[mod2.Service] `inject:""`
    Dep3 ioc.Lazy[mod3.Service] `inject:""`
}
```

```modules/mod3/internal.go
type service struct {
    engine.EngineWorld `inject:""`
}
```

**Pros**:
- Easy wiring (wiring becomes a matter of 2 loc per module)
- Centralizes the app and treats it as a single object
**Cons**:
- Less performant (Additional boolean check within the lambda during access)

This makes additions to engine a piece of cake and reduces boilerplate and allows for circular dependencies.
While circular dependencies are a bad practice additional technical effort for each module
to explicitly ensure they won't occur isn't a way to avoid them.

#### Conclusion
The choice of implicit dependency management provides:
- Developer velocity
- Less code to maintain
For the price of:
- Performance (Additional if check during any service access)

Why performance cost high-performance project is negligible:
- Bool and pointer are so small that they'll always be loaded into memory so there won't be a cache miss (biggest performance cost)
- While the boolean check has a cost, it is negligible compared to calling the most expensive methods (these which need optimization) methods like pathfinding on 1M tiles map

This trade of is a no brainer for scenario where one developer builds whole game/simulation engine from scratch.
Developer velocity is everything in this scenario and this minor price is well worth it.
This project isn't in asm for a reason.

What, contrary to appearances, is not a price:
- ##### safety
All dependencies are resolved at startup so there won't be circular dependencies at runtime.
- ##### promoption of bad practices
Some could argue that this promotes circular dependencies but
maintaining this rule shouldn't be at a cost of boilerplate.
Code isn't about enforcing rules its about composition of functionalities.
Forbidding functionalities in code BECAUSE, with no technical reason is a bad practice.

### Modules
**Currently only cherry picked readmes are written**
Cherry picked readmes to show project complexity:
- [tile](/core/modules/tile/readme/README.md)
- [ecs](/engine/services/ecs/readme/README.md)
- [assets](/engine/modules/assets/readme/README.md)
- [hierarchy](/engine/modules/hierarchy/readme/README.md)
- [record](/engine/modules/record/readme/README.md)
- [transform](/engine/modules/transform/readme/README.md)

#### Core
Core is game source code.\
It defines game specific objects like `tiles`, `units`, map `generation` or `pathfinding`.\
Core modules:
- [definitions](/core/modules/definitions/readme/README.md)
- [deploy](/core/modules/deploy/readme/README.md)
- [fpslogger](/core/modules/fpslogger/readme/README.md)
- [generation](/core/modules/generation/readme/README.md)
- [loading](/core/modules/loading/readme/README.md)
- [pathfind](/core/modules/pathfind/readme/README.md)
- [player](/core/modules/player/readme/README.md)
- [settings](/core/modules/settings/readme/README.md)
- [tile](/core/modules/tile/readme/README.md)
- [ui](/core/modules/ui/readme/README.md)

#### Engine
Engine is reusable in other projects.\
It defines ecs framework and basic engine modules like `transform` or `hierarchy`.\
Engine modules:
- [assets](/engine/modules/assets/readme/README.md)
- [audio](/engine/modules/audio/readme/README.md)
- [batcher](/engine/modules/batcher/readme/README.md)
- [camera](/engine/modules/camera/readme/README.md)
- [codec](/engine/modules/codec/readme/README.md)
- [collider](/engine/modules/collider/readme/README.md)
- [connection](/engine/modules/connection/readme/README.md)
- [drag](/engine/modules/drag/readme/README.md)
- [entityregistry](/engine/modules/entityregistry/readme/README.md)
- [graphics](/engine/modules/graphics/readme/README.md)
- [grid](/engine/modules/grid/readme/README.md)
- [groups](/engine/modules/groups/readme/README.md)
- [hierarchy](/engine/modules/hierarchy/readme/README.md)
- [inputs](/engine/modules/inputs/readme/README.md)
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

Engine services:
- [clock](/engine/services/clock/readme/README.md)
- [console](/engine/services/console/readme/README.md)
- [datastructures](/engine/services/datastructures/readme/README.md)
- [ecs](/engine/services/ecs/readme/README.md)

### Technical challenges
Each and every module had unique challenges and they are described in these readmes.

Biggest challenge of the whole project was architecture.\
Finding file structure which allows for most logic with least friction between modules.\
Current approach reduces whole friction to a few interface files and often in a single `Service` interface.

## CI/CD
Current pipeline is hosted using **Jenkins** on **Raspberry PI**.
Current pipeline **CI** flow:
- `quality`: `golangci-lint`, dependency hygiene
- `correctness`: `unit tests`
- `security`: `gosec`, `trivy`

Using **GitHub Status API** integrates notifications into developer worflow.\
Current design leaves space to add in the future stages for **CD**.

[CI/CD directory](/ci-cd)

## Graphics

Example map generated in a matter of seconds and rendered in less than 6ms\
using 5 years old Intel® Core™ i5-8350U × 8 Intel® and UHD Graphics 620 (KBL GT2):
![Map scroll](map_scroll.gif)
![Whole map](whole_map.png)
![Bottom right map corner](bottom_right.png)

## How to run ?
### Install dependencies
Install packages for:
- opengl
- sdl2

ubuntu:
```
sudo apt install libsdl2-dev libsdl2-image-dev libsdl2-ttf-dev libsdl2-mixer-dev
sudo apt install mesa-common-dev libglew-dev libglu1-mesa-dev
```

arch:
```
sudo pacman -S sdl2 mesa libglew glue
sudo pacman -S sdl2_image sdl2_mixer sdl2_ttf
```

### Run
```
cd core
go run .
```

## Contribution
We are not currently seeking external contributions.\
However, we will review individual inquiries on a case-by-case basis.\
While we remain selective at this stage, we are open to discussion.\
Please note that this project is not yet open-source, as it is in the early stages of development.

## License
Copyright © 2026. All rights reserved.
Currently, this repository is public to allow for code review and demonstration of functionality for recruitment purposes.\
No part of this software may be reproduced, distributed, or transmitted in any form or by any means without prior written permission.
