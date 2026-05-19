# TEXHEC
## About
TEXHEC is experimental High-Performance Simulation **ECS** based Engine.
Target is to create biggest immersive simulated strategy game world.

## Projects
- ### [CI/CD](/cicd/readme/README.md)
- ### [TEXHEC Core](/core/readme/README.md)
- ### [TEXHEC Engine](/engine/readme/README.md)
- ### [DI container](https://github.com/ogiusek/ioc)
- ### [Event bus](https://github.com/ogiusek/events)

## Principles
- ### Software is for users. It is an engineers job to make users lifes easier
- ### Legacy should be something that pushes us. Dozens of features tommorow over one today
- ### Focus on what matters most now instead of optimizing working code
- ### Handle all edge cases upfront. Do not leave errors for later

## Architecture
### Golang
Golang despite having **GC** is a perfect choice for a game engine.
By following [**DOD**](#dod) we leave no pointers behind which have to be heavily cleaned.
Pros:
- very performant (its compiled)
- its fast to write, understand and its very easy to use
- it lacks decades of building technical debt
- aligned philosophies (simplicity creates performance not other way around)

### DI container
Why:
1. Services store data and with such dependencies we need to have instance manager.
2. Service wrappers allow us to extend our services in a way that everything is contained.

### Implicit dependency injection graph
Implicit dependency injection graph in code is single struct with all services.
The point is to automate dependency management and to expose single facade.
```go
type EngineWorld struct {
	World         ioc.Lazy[ecs.World]      `inject:""`
	EventsBuilder ioc.Lazy[events.Builder] `inject:""`
	Events        ioc.Lazy[events.Events]  `inject:""`

	Assets         ioc.Lazy[assets.Service]         `inject:""`
	Audio          ioc.Lazy[audio.Service]          `inject:""`
    // ...
}
```

#### Pros
- DX. Unified structure with all services in project
- Loc reduction. Less imports and repeating service properties
- Less maintenance on adding or removing unused services injected

#### Cons
1. Forces separation of interface from implementation to avoid circular dependencies
Solution: Own file structure separating interface from implementation.
2. Theoreticaly using lazy listeners decreases performance
Solution: Uses bool under the hood stored along side pointer therefor in practice hot path doesn't show up in profiler
3. Automatic dependencies spread in module.
Solution: This is solved with automatic documentation listing all dependencies in one place
4. Testing requires whole world.
Solution: Packages have default configurations.

Despite more cons than pros these cons enforce good practices while with these solutions these cons are negligible.

### DOD
Following **DOD** creates exceptionally performant software and to avoids **GC**.
#### Proof
[Tile benchmark](/core/modules/tile/readme/README.md#benchmarks)
Map generated in seconds and rendered in less than **6ms** using
5 years old Intel® Core™ i5-8350U × 8 Intel® and UHD Graphics 620 (KBL GT2):
![Map scroll](map_scroll.gif)
![Whole map](whole_map.png)
![Bottom right map corner](bottom_right.png)

#### Architecture
In this project we follow **DOD** by using [ECS](/engine/services/ecs/readme/README.md)

### Determinism
Using **DOD** along side **EDD** creates deterministic engine.
Using **TPS** (ticks per second) to perform modifications and **FPS** (frames per second)
to perform client GUI and input updates creates environment which is easy to debug and to send over network.

### CI/CD
[CI/CD](/cicd/readme/README.md) runs on both client and server.
Before commit documentations are automatically generated.

### Documentation generation
Own documentation format allows for additional sections like:
- `Benchmarks`
- `Lines of code`
- `Dependencies`

### Module structure
```
modules/
└─ `$module_name`/
    ├── internal/          # Defines implementation for `Service` and `System` (if exist in module)
    │
    ├── pkg/               # This exposes `Package` function to register `Service` implementation.
    │                      # `pkg`, `internal` and `test` separation allows `modules`
    │                      # Decouples the interface definition from the construction logic to allow for flexible dependency wiring
    │
    ├── test/              # Defines test
    │                      # Benchmarks here are automatically used in generated readme
    │
    ├── readme/            # Defines readme
    │   ├─ TITLE.md        # Overwrites `package name` as automatic readme header
    │   ├─ ARCHITECTURE.md # Overwrites package comments as automatic readme architecture section
    │   ├─ BENCH.md        # Overwrites automatic benchmarks
    │   ├─ CHALLENGES.md   # Challenges section in generated readme
    │   ├─ TODO.md         # TODO  section in generated readme
    │   └─ README.md       # generated readme
    │
    └── `service.go`       # There is no strict file rule naming. This defines what module exposes
                           # Expects interface name `Service` so module name and service purpose were related
```

### Module vs Service
Service is something separate from game engine which is basis for it.\
After creating **ECS** service i attempt to migrate everything to a module.\
Modules also have more struct rules and have dedicated file structure.\
Services are more detached from alone game engine and have less strict rules.

## Cherry picked readmes
- [docs](/cicd/modules/docs/readme/README.md)
- [pipe](/cicd/modules/pipe/readme/README.md)

- [tile](/core/modules/tile/readme/README.md)
- [ecs](/engine/services/ecs/readme/README.md)

- [assets](/engine/modules/assets/readme/README.md)
- [hierarchy](/engine/modules/hierarchy/readme/README.md)
- [record](/engine/modules/record/readme/README.md)
- [transform](/engine/modules/transform/readme/README.md)

## How to run ?
### Clone repository
```
git clone github.com/cursus-studio/texhec
```

### Setup project
```
go run cicd setup
```

### Install dependencies
Install packages for:
- `opengl`
- `sdl2`

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
go -C core run .
```

## Contribution
We are not currently seeking external contributions.\
However, we will review individual inquiries on a case-by-case basis.\
While we remain selective at this stage, we are open to discussion.\
Please note that this project is not yet open-source, as it is in to early stages of development.

## License
Copyright © 2026. All rights reserved.
Currently, this repository is public to allow for code review and demonstration of functionality for recruitment purposes.\
No part of this software may be reproduced, distributed, or transmitted in any form or by any means without prior written permission.
