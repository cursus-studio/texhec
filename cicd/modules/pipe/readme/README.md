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
Go                               4             64            204            416
-------------------------------------------------------------------------------
SUM:                             4             64            204            416
-------------------------------------------------------------------------------
```
## Dependencies
`cicd/modules/git`:
  - `cicd/modules/git.DiffCompare`
  - `cicd/modules/git.DiffNotCommited`
  - `cicd/modules/git.Pending`
  - `cicd/modules/git.Release`
  - `cicd/modules/git.SetStatus`
  - `cicd/modules/git.Stage`
  - `cicd/modules/git.Version`

`cicd/modules/pipe`:
  - `cicd/modules/pipe.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`
  - `cicd/world.Docs`
  - `cicd/world.Git`
  - `cicd/world.ProjectFS`

### Third Party
- `github.com/ogiusek/ioc/v2`