# git
## Architecture
responsible for git integration

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             23              3            123
-------------------------------------------------------------------------------
SUM:                             3             23              3            123
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