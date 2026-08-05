# pipe
## Architecture
main entry point responsible for setuping pipe hooks, running pipe hooks, running cicd pipeline

1. code quality stages:
- dependencies
- compilation
- gosec
- golangci-lint
- tests
2. pipeline quality stages:
- trivy
3. docs quality stages:
- generate or verify docs using [docs](/cicd/modules/docs)
- lychee

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               4             59             38            544
-------------------------------------------------------------------------------
SUM:                             4             59             38            544
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `cicd/modules/pipe.Service`

#### method Service Cloud
Type: `func(commitHash string) error`

#### method Service Fix
Type: `func() error`

#### method Service Setup
Type: `func() error`

#### method Service Sync
Type: `func() error`

#### method Service Verify
Type: `func(commitHash string) error`


## Dependencies
`cicd/modules/git`:
  - `cicd/modules/git.DiffCompare`
  - `cicd/modules/git.DiffNotCommited`
  - `cicd/modules/git.Pending`
  - `cicd/modules/git.SetStatus`
  - `cicd/modules/git.Stage`

`cicd/modules/pipe`:
  - `cicd/modules/pipe.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`
  - `cicd/world.Docs`
  - `cicd/world.Git`
  - `cicd/world.ProjectFS`

### Third Party
- `github.com/ogiusek/ioc/v2`