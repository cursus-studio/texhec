# batcher
## Architecture
this module allows us to write tasks and to progress them across frames without stuterring

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             49              5            226
-------------------------------------------------------------------------------
SUM:                             6             49              5            226
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.EventsBuilder`
  - `engine.Loop`
  - `engine.WarmUp`

`engine/modules/batcher`:
  - `engine/modules/batcher.AddOrderedBatch`
  - `engine/modules/batcher.Batch`
  - `engine/modules/batcher.Handler`
  - `engine/modules/batcher.NewBatch`
  - `engine/modules/batcher.Progress`
  - `engine/modules/batcher.Service`
  - `engine/modules/batcher.Step`
  - `engine/modules/batcher.Steps`
  - `engine/modules/batcher.Task`
  - `engine/modules/batcher.TaskFactory`

`engine/modules/ecs`:
  - `engine/modules/ecs.SystemRegister`

`engine/modules/loop`:
  - `engine/modules/loop.FrameBudgetLeft`
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.Stats`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`