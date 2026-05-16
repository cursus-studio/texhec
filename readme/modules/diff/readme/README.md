# diff
## Architecture
this module finds modified modules since last commit

## Types
### type Service
Type: `readme/modules/diff.Service`

#### method Service GetModifiedModules
Type: `func() ([]string, error)`


## Dependencies
`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSet`

`readme/modules/diff`:
  - `readme/modules/diff.Service`