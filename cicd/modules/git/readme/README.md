# git
## Architecture
responsible for git integration

## Types
### type Service
Type: `cicd/modules/git.Service`
each method returns modules

#### method Service DiffCompare
Type: `func(commitHash string) ([]string, error)`
passing empty string will compare to not commited

#### method Service DiffNotCommited
Type: `func() ([]string, error)`

#### method Service DiffPrevCommit
Type: `func() ([]string, error)`

#### method Service Stage
Type: `func(...string) error`


## Dependencies
`cicd/modules/git`:
  - `cicd/modules/git.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`

### Third Party
- `github.com/ogiusek/ioc/v2`