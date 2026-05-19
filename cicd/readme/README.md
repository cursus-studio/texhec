# CI-CD

## Architecture
This project is responsible for code quality and commit history.

CI/CD tool exposes 5 commands:
- ### `go run cicd setup`
It initializes git hooks.

- ### `go run cicd sync`
It generates readmes in whole codebase.
It is used when readme generation algorithm is changes.

- ### `go run cicd fix`
It takes uncommited files and runs all tests on them and updates readmes and stages them.

- ### `go run cicd verify $compared-commit-hash`
`$compared-commit-hash` - By deafult only not commited changed are tested. In CI-CD pipeline previous successful commit is used here.

It runs all quality tests only on specific changed files.

- ### `go run cicd cloud $compared-commit-hash`
`$compared-commit-hash` - By deafult only not commited changed are tested. In CI-CD pipeline previous successful commit is used here.

It is effectively the save as `go run cicd verify` but also uses github status api.
To send changes using github status api it needs ENV variables:
- `TOKEN`: github access token
- `OWNER`: repository owner (e.g. `cursus-studio`)
- `REPO`: repository name (e.g. `texhec`)
- `GIT_COMMIT`: git hash which status should be updated
- `BUILD_URL`: url where this command runs
- `CONTEXT`: git status key

To see implementation details of documentation or pipeline go to its module readme.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              18            211            131           1528
Groovy                           1              5              1             74
Dockerfile                       1              7              3             29
Markdown                         3              9              0             29
-------------------------------------------------------------------------------
SUM:                            23            232            135           1660
-------------------------------------------------------------------------------

```
## Dependencies
`cicd/modules/docs`:
  - `cicd/modules/docs.DiffModuleDocs`
  - `cicd/modules/docs.DiffProjectDocs`
  - `cicd/modules/docs.ErrInconsistentPackageComments`
  - `cicd/modules/docs.ErrMissingPackage`
  - `cicd/modules/docs.ErrMissingPackageComments`
  - `cicd/modules/docs.GenerateModuleDocs`
  - `cicd/modules/docs.GenerateProjectDocs`
  - `cicd/modules/docs.Service`

`cicd/modules/docs/pkg`:
  - `cicd/modules/docs/pkg.Pkg`

`cicd/modules/git`:
  - `cicd/modules/git.DiffCompare`
  - `cicd/modules/git.DiffNotCommited`
  - `cicd/modules/git.Pending`
  - `cicd/modules/git.Service`
  - `cicd/modules/git.SetStatus`
  - `cicd/modules/git.Stage`
  - `cicd/modules/git.State`

`cicd/modules/git/pkg`:
  - `cicd/modules/git/pkg.Pkg`

`cicd/modules/pipe`:
  - `cicd/modules/pipe.Cloud`
  - `cicd/modules/pipe.Fix`
  - `cicd/modules/pipe.Service`
  - `cicd/modules/pipe.Setup`
  - `cicd/modules/pipe.Sync`
  - `cicd/modules/pipe.Verify`

`cicd/modules/pipe/pkg`:
  - `cicd/modules/pipe/pkg.Pkg`

`cicd/modules/projectfs`:
  - `cicd/modules/projectfs.AllModules`
  - `cicd/modules/projectfs.AllProjects`
  - `cicd/modules/projectfs.FilesModules`
  - `cicd/modules/projectfs.FilesProjects`
  - `cicd/modules/projectfs.ProjectModules`
  - `cicd/modules/projectfs.ProjectServices`
  - `cicd/modules/projectfs.Service`

`cicd/modules/projectfs/pkg`:
  - `cicd/modules/projectfs/pkg.Pkg`

`cicd/pkg`:
  - `cicd/pkg.Pkg`

`cicd/world`:
  - `cicd/world.CICDWorld`
  - `cicd/world.Docs`
  - `cicd/world.Git`
  - `cicd/world.Pipe`
  - `cicd/world.ProjectFS`

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSet`

### Third Party
- `github.com/google/go-github/v60/github`
- `github.com/ogiusek/ioc/v2`
- `github.com/spf13/cobra`
- `golang.org/x/oauth2`
- `golang.org/x/tools/go/packages`