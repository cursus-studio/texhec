# diff
## Architecture
lists modified modules

## Types
### type Service
Type: `cicd/modules/diff.Service`
each method returns modules

#### method Service DiffCommited
Type: `func() ([]string, error)`
lists changed modules in previous commit

#### method Service DiffUncommited
Type: `func() ([]string, error)`
lists uncommited changed modules


## Dependencies
`cicd/modules/diff`:
  - `cicd/modules/diff.Service`

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSet`