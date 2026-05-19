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

#### method Service SetStatus
Type: `func(status cicd/modules/git.State, decs string) error`

#### method Service Stage
Type: `func(...string) error`

### type State
Type: `cicd/modules/git.State`

## Variables
### var Pending
Type: `cicd/modules/git.State`

### var Success
Type: `cicd/modules/git.State`

### var Failure
Type: `cicd/modules/git.State`

### var Error
Type: `cicd/modules/git.State`


## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             23              3            122
-------------------------------------------------------------------------------
SUM:                             3             23              3            122
-------------------------------------------------------------------------------

```
## Dependencies
`cicd/modules/git`:
  - `cicd/modules/git.Service`
  - `cicd/modules/git.State`

`cicd/world`:
  - `cicd/world.CICDWorld`

### Third Party
- `github.com/google/go-github/v60/github`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/oauth2`