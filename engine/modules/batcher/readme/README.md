# batcher
## Architecture
this module allows us to write tasks and to progress them across frames without stuterring

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             49              5            221
-------------------------------------------------------------------------------
SUM:                             6             49              5            221
-------------------------------------------------------------------------------

```
## Types
### type Service
Type: `engine/modules/batcher.Service`

#### method Service NewTask
Type: `func() engine/modules/batcher.TaskFactory`

#### method Service Progress
Type: `func() float32`
progress of first task in queue
when there is no tasks in queue -1 is returned

#### method Service Queue
Type: `func(engine/modules/batcher.Task)`

#### method Service Register
Type: `func() error`

### type TaskFactory
Type: `engine/modules/batcher.TaskFactory`

#### method TaskFactory AddConcurrentBatch
Type: `func(engine/modules/batcher.Batch) engine/modules/batcher.TaskFactory`

#### method TaskFactory AddOrderedBatch
Type: `func(engine/modules/batcher.Batch) engine/modules/batcher.TaskFactory`

#### method TaskFactory Build
Type: `func() engine/modules/batcher.Task`

### type Task
Type: `engine/modules/batcher.Task`

#### method Task Perform
Type: `func()`

#### method Task Progress
Type: `func() float32`
progress of first task in queue
when there is no tasks in queue -1 is returned

#### method Task Step
Type: `func()`

### type Batch
Type: `engine/modules/batcher.Batch`

#### property Batch Steps
Type: `int`

#### property Batch Handler
Type: `func(int)`

## Functions
### func NewBatch
Type: `func(steps int, handler func(int)) engine/modules/batcher.Batch`


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

`engine/modules/loop`:
  - `engine/modules/loop.FrameBudgetLeft`
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.Stats`

`engine/services/ecs`:
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`