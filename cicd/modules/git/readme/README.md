# git
## Architecture
responsible for git integration

## Types
### type Service
Type: `cicd/modules/git.Service`
each method returns modules

#### method Service DiffCommited
Type: `func() ([]string, error)`
lists changed modules in previous commit

#### method Service DiffUncommited
Type: `func() ([]string, error)`
lists uncommited changed modules

#### method Service Stage
Type: `func(...string) error`


## Dependencies
`cicd/modules/git`:
  - `cicd/modules/git.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSet`

### Third Party
- `github.com/ogiusek/ioc/v2`