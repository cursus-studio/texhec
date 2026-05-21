# TEXHEC
## About
**TEXHEC** is an experimental, high-performance, ECS-based simulation engine.

**The Goal:** Create a massive, immersive strategy game world that pushes hardware limits within a realistic timeframe.

## Projects
Pick readme for your niche:
- ### Devops [CI/CD](/cicd/readme/README.md)
Runs on both client (using `git hooks`) and server (using `jenkins`).
Before-commit documentation sections are automatically generated and added to main readmes.

- ### Low-Level [TEXHEC Engine](/engine/readme/README.md)
Defines data structures to store whole game state in continuous chunks and game engine foundations

- ### Game [TEXHEC Core](/core/readme/README.md)
Defines game and game objects. In [DOD](#dod) section you can see performance of **engine** and **core**.

- ### Dependencies [DI container](https://github.com/ogiusek/ioc)
Implements more efficient structures to store and access services while providing more sugary
abstractions reducing **LOC**.

- ### Events [Event bus](https://github.com/ogiusek/events)
**DI container** and **event bus** was made to have full control over this project.
In comparison to whole project scope these were like single features to implement so cost
of writing and maintaining them is well worth it.

## Principles
- ### Adapt. Engineers are meant to adapt to project requirements not the other way around.
- ### Legacy should be something that pushes us. One working feature tomorrow over two half baked today
- ### Pick lowest hanging fruit first. Cut corners to deliver without building technical debt

## Architecture
### DOD
Following **DOD** creates exceptionally performant software and avoids **GC**.
#### Zero allocation **ECS** framework
[ECS benchmark](/engine/services/ecs/readme/README.md#benchmarks)

Data oriented ECS framework implemented in Go completely avoids **GC**.

#### Performance showcase
[Tile benchmark](/core/modules/tile/readme/README.md#benchmarks)

A 36.000.000 tile map rendered in less than **8.6ms** using:
- CPU: `Intel(R) Core(TM) i7-14700KF`
- GPU: `NVIDIA GeForce RTX 4080 SUPER`

A 1.000.000 tile map is generated in seconds and rendered in less than **5.1ms** using:
- CPU: 5-year-old `Intel® Core™ i5-8350U × 8 Intel®`
- GPU: `UHD Graphics 620 (KBL GT2)`

![Map scroll](map_scroll.gif)

#### Architecture
In this project we follow **DOD** by using [ECS](/engine/services/ecs/readme/README.md)

### Determinism
Using **DOD** alongside **EDD** creates deterministic engine.
Using **TPS** (ticks per second) to perform modifications and **FPS** (frames per second)
to perform client GUI and input updates creates environment which is easy to debug and to send over network.

### Golang
Golang despite having **GC** is a perfect choice for a game engine.
By following [**DOD**](#dod) we leave no pointers behind which have to be heavily cleaned.
Pros:
- very performant (it's compiled)
- it's fast to write, understand and it's very easy to use
- it lacks decades of building technical debt
- aligned philosophies (simplicity creates performance not other way around)

### DI container
Why:
1. Services store data, and with such dependencies, we need to have an instance manager.
2. Service wrappers allow us to extend our services in a way that everything is contained.

### Implicit dependency injection graph
Implicit dependency injection graph in code is single struct with all services.
The point is to automate dependency management and to expose s single facade.
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
1. Forces separation of interface from implementation to avoid circular dependencies\
**Solution**: Own file structure separating interface from implementation.
2. Theoretically using lazy listeners decreases performance\
**Solution**: Uses bool under the hood stored alongside pointer therefore in practice hot path doesn't show up in profiler
3. Automatic dependencies spread in module.\
**Solution**: This is solved with automatic documentation listing all dependencies in one place
4. Testing requires whole world.\
**Solution**: Packages have default configurations.

These cons enforce good practices, while with these solutions, their impact is negligible.

### Documentation generation
Own documentation format allows for additional sections like:
- `Types`
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
**Module**: integral part of an engine with specific structure.
It is faster to write, use and has dedicated tooling (in [CI/CD](/cicd/readme/README.md)) but it is harder to separate.

**Service**: separate package which can be used in other projects without modifications.

Most of codebase leans towards modules for developer velocity and unification.

## Cherry picked readmes
**CI/CD**:
- [docs](/cicd/modules/docs/readme/README.md)
- [pipe](/cicd/modules/pipe/readme/README.md)
**Foundations**:
- [tile](/core/modules/tile/readme/README.md)
- [ecs](/engine/services/ecs/readme/README.md)
**Building Blocks**:
- [assets](/engine/modules/assets/readme/README.md)
- [hierarchy](/engine/modules/hierarchy/readme/README.md)
- [record](/engine/modules/record/readme/README.md)
- [transform](/engine/modules/transform/readme/README.md)

## How to run ?
### Clone repository
```
git clone https://github.com/cursus-studio/texhec.git
```

### Setup project
```
go run cicd setup
```

### Install dependencies
Install packages for:
- `opengl`
- `sdl2`
- `golang`
- `docker`

### Run
```
go -C core run .
```

## Contribution
We are not currently seeking external contributions.\
However, we will review individual inquiries on a case-by-case basis.\
While we remain selective at this stage, we are open to discussion.

## License
Copyright © 2026. All rights reserved.

Permission is hereby granted to download, compile, and run this software locally on your own machine, and to make local modifications solely for the purposes of personal evaluation, testing, and generating feedback.

**Strictly Prohibited:**
* **Public Distribution & Forking:** You may not distribute, publish, or transmit modified versions of this software (derivative works), nor host public forks intended to branch or continue the project independently.
* **Commercial Use:** You may not use this software, in whole or in part, for any commercial purposes or financial gain.
