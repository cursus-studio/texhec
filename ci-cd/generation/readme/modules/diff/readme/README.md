# diff
## Architecture
lists modified modules

## Types
### type Service
Type: `readme/modules/diff.Service`
each method returns modules

#### method Service DiffCommited
Type: `func() ([]string, error)`
lists changed modules in previous commit

#### method Service DiffUncommited
Type: `func() ([]string, error)`
lists uncommited changed modules


## Dependencies
`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSet`

`readme/modules/diff`:
  - `readme/modules/diff.Service`