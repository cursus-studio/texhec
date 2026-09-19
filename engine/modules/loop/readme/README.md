# loop
## Architecture
defines game loop

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             51             17            157
-------------------------------------------------------------------------------
SUM:                             3             51             17            157
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/loop.Service`

#### method Service Configure
Type: `func(engine/modules/loop.ConfigureEvent)`

#### method Service EmitOnTick
Type: `func(event any)`

#### method Service LastTickUnixNano
Type: `func() int64`

#### method Service Run
Type: `func(initialConfiguration engine/modules/loop.ConfigureEvent)`
Starts the game loop if it isn't started.
Waits until game loop stops.

#### method Service Stats
Type: `func() engine/modules/loop.Stats`

#### method Service Stop
Type: `func()`

#### method Service SyncToUnixNano
Type: `func(int64)`
also re-ticks

### type Stats
Type: `engine/modules/loop.Stats`

#### method Stats FrameBudget
Type: `func() time.Duration`

#### method Stats FrameBudgetLeft
Type: `func() time.Duration`

### type StopEvent
Type: `engine/modules/loop.StopEvent`
stops game loop
can listen to it to clean up

### type ConfigureEvent
Type: `engine/modules/loop.ConfigureEvent`
changes game loop configuration

#### property ConfigureEvent FPS
Type: `int`

#### property ConfigureEvent TPS
Type: `int`

### type TickEvent
Type: `engine/modules/loop.TickEvent`
tick has fixed delta to ensure determinism
tick is triggered before frame as many times as many ticks passed between frames

#### property TickEvent Delta
Type: `time.Duration`

### type FrameEvent
Type: `engine/modules/loop.FrameEvent`

#### property FrameEvent Delta
Type: `time.Duration`

### type EmitOnTickComponent
Type: `engine/modules/loop.EmitOnTickComponent`

#### property EmitOnTickComponent Event
Type: `any`

### type EmitOnTickEvent
Type: `engine/modules/loop.EmitOnTickEvent`

#### property EmitOnTickEvent Event
Type: `any`

## Functions
### func NewStopEvent
Type: `func() engine/modules/loop.StopEvent`

### func NewConfigureEvent
Type: `func(fps int, tps int) engine/modules/loop.ConfigureEvent`

### func NewEmitOnTickEvent
Type: `func(event any) engine/modules/loop.EmitOnTickEvent`


## Dependencies
`engine`:
  - `engine.Clock`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/loop`:
  - `engine/modules/loop.ConfigureEvent`
  - `engine/modules/loop.EmitOnTickComponent`
  - `engine/modules/loop.EmitOnTickEvent`
  - `engine/modules/loop.Event`
  - `engine/modules/loop.FPS`
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.Service`
  - `engine/modules/loop.Stats`
  - `engine/modules/loop.StopEvent`
  - `engine/modules/loop.TPS`
  - `engine/modules/loop.TickEvent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`