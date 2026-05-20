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
Type: `func(Key) (engine/services/ecs.EntityID, bool)`


## Dependencies
`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/warmup`:
  - `engine/modules/warmup.Event`

`engine/services/datastructures`:
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.Remove`
  - `engine/services/datastructures.Set`
  - `engine/services/datastructures.SparseArray`

`engine/services/ecs`:
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.World`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`