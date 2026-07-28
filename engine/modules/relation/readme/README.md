# relation
## Architecture
this is a generic package which is used to access entities by component (id) in O(1) time

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8             49              1            266
-------------------------------------------------------------------------------
SUM:                             8             49              1            266
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/relation.Service[Key any]`

#### method Service Get
Type: `func(Key) (engine/modules/ecs.EntityID, bool)`


## Dependencies
`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.SparseArray`

`engine/modules/ecs`:
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.World`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/warmup`:
  - `engine/modules/warmup.Event`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`