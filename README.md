# TEXHEC
## What is TEXHEC ?

### Answer for interested players

TEXHEC is a game during development.

In TEXHEC we can win by:
- technological win
- total enemies annihilation

Point of the game is to allow complex tactics (all moves allowed) in seemingly endless world.

Where everything can be done manually or can be automated.
Everything is under player control where he has to find order in endless chaos and has to conquer it before others do so.

### Technical answer

TEXHEC is a vision **RTS** (real time strategy) game which currently is being brought to life.\
On current stage of TEXHEC i wrote **dozens** of **unique** **modules** from every solved other problems.\
Its completely built from scratch with **less** **than** a **dozen** of **dependencies** not controled by me.

TEXHEC is a **HIGH-Performance** project where natural map size is **1.000.000*** tiles with hundreds or thousands buildings and units **all** being **simulated** in real time.\
We use **DOD** and use our **own** **ECS** framework.

[More about this **ECS** framework](/engine/services/ecs/readme/README.md)

#### Why golang
Others would **discard golang** due to **garbage collector**.\
In reality garbage collector isn't an inconvenience because we follow **DOD** and\
we do not have enough pointers to be an inconvenience.

In reality using golang has benefits:
- its very performant (its compiled)
- its fast to write, understand and its very easy to use (necessary to deliver by a single developer)
- it lacks decades of building technical debt
- aligned philosophies (simplicity creates performance not other way around)

#### Dependencies
- `sdl2`
- `opengl`
- `opengl math`
- `golang constraints`
- `golang images and text (used only to generate image per letter)`
- `thread safe hash map`
- `google uuid`

Dependencies which are written by me:
- `ioc`
- `events`

#### Module vs Service
Service is something separate from game engine which is basis for it.\
After creating **ECS** service i attempt to migrate everything to a module.\
Modules also have more struct rules and have dedicated file structure.\
Services are more detached from alone game engine and have less strict rules.

#### Module structure
```
modules/
└─ `$module_name`/
    ├── internal/       # Defines implementation for `Service` and `System` (if exist in module)
    ├── pkg/            # This exposes `Package` function to register `Service` implementation.
    │                   # `pkg`, `internal` and `tests` separation allows `modules`
    │                   # Decouples the interface definition from the construction logic to allow for flexible dependency wiring
    ├── tests/          # Defines tests
    ├── readme/         # Defines readme
    └── `$interface.go` # There is no strict file rule naming. This defines what module exposes
                        # Expects interface name `Service` so module name and service purpose were related
```
Everything in module file structure is optional and should be only added if used.

#### Module readme schema
```md
# Module_name
## Architecture
How module is built and general flow of data and why this way in case of controversial choices.

## Benchmarks (optional)

## Usage examples
Code snippets of `Service` and of how to use it.
...

## Dependencies
- [module](/engine/modules/module_name/readme/README.md)
```

#### Engine
Engine is the core which can be re-used in other projects.\
It defines ecs framework and basic engine modules like `transform` or `hierarchy`

**Currently only cherry picked readmes are written**

Cherry picked readmes to show project complexity:
- [ecs](/engine/services/ecs/readme/README.md)
- [assets](/engine/modules/assets/readme/README.md)
- [hierarchy](/engine/modules/hierarchy/readme/README.md)
- [record](/engine/modules/record/readme/README.md)
- [transform](/engine/modules/transform/readme/README.md)

Engine modules:
- [assets](/engine/modules/assets/readme/README.md)
- [audio (placeholder)](/engine/modules/audio/readme/README.md)
- [batcher (placeholder)](/engine/modules/batcher/readme/README.md)
- [camera (placeholder)](/engine/modules/camera/readme/README.md)
- [collider (placeholder)](/engine/modules/collider/readme/README.md)
- [connection (placeholder)](/engine/modules/connection/readme/README.md)
- [drag (placeholder)](/engine/modules/drag/readme/README.md)
- [grid (placeholder)](/engine/modules/grid/readme/README.md)
- [groups (placeholder)](/engine/modules/groups/readme/README.md)
- [hierarchy](/engine/modules/hierarchy/readme/README.md)
- [inputs (placeholder)](/engine/modules/inputs/readme/README.md)
- [layout (placeholder)](/engine/modules/layout/readme/README.md)
- [netsync (placeholder)](/engine/modules/netsync/readme/README.md)
- [noise (placeholder)](/engine/modules/noise/readme/README.md)
- [record](/engine/modules/record/readme/README.md)
- [registry (placeholder)](/engine/modules/registry/readme/README.md)
- [relation (placeholder)](/engine/modules/relation/readme/README.md)
- [render (placeholder)](/engine/modules/render/readme/README.md)
- [scene (placeholder)](/engine/modules/scene/readme/README.md)
- [seed (placeholder)](/engine/modules/seed/readme/README.md)
- [smooth (placeholder)](/engine/modules/smooth/readme/README.md)
- [text (placeholder)](/engine/modules/text/readme/README.md)
- [transform](/engine/modules/transform/readme/README.md)
- [transition (placeholder)](/engine/modules/transition/readme/README.md)
- [uuid (placeholder)](/engine/modules/uuid/readme/README.md)

Engine services:
- [clock (placeholder)](/engine/services/clock/readme/README.md)
- [codec (placeholder)](/engine/services/codec/readme/README.md)
- [console (placeholder)](/engine/services/console/readme/README.md)
- [datastructures (placeholder)](/engine/services/datastructures/readme/README.md)
- [ecs](/engine/services/ecs/readme/README.md)
- [frames (placeholder)](/engine/services/frames/readme/README.md)
- [graphics (placeholder)](/engine/services/graphics/readme/README.md)
- [httperrors (placeholder)](/engine/services/httperrors/readme/README.md)
- [logger (placeholder)](/engine/services/logger/readme/README.md)
- [media (placeholder)](/engine/services/media/readme/README.md)
- [runtime (placeholder)](/engine/services/runtime/readme/README.md)

#### Technical challenges
Each and every module had unique challenges and they are described in these readmes.

Biggest challenge of the whole project was architecture.\
Finding file structure which allows for most logic with least friction between modules.\
Current approach reduces whole friction to a few interface files and often in a single `Service` interface.

### Graphics

Example map generated in a matter of seconds and rendered in less than 6ms\
using 5 years old Intel® Core™ i5-8350U × 8 Intel® UHD Graphics 620 (KBL GT2):
![Whole map](/readme/whole_map.png)
![Bottom right map corner](/readme/bottom_right.png)

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
Currently we do not expect any contributions and each will be handled individually.\
But currently we're rather sceptical and assertive but open to discussion.

## License
Copyright © 2026. All rights reserved.
Currently, this repository is public to allow for code review and demonstration of functionality for recruitment purposes.\
No part of this software may be reproduced, distributed, or transmitted in any form or by any means without prior written permission.
