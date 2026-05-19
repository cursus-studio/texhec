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

## Modules
[docs](/cicd/modules/docs)
[git](/cicd/modules/git)
[pipe](/cicd/modules/pipe)
[projectfs](/cicd/modules/projectfs)

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              18            211            131           1528
Groovy                           1              5              1             74
Dockerfile                       1              8              3             30
Markdown                         3              9              0             29
-------------------------------------------------------------------------------
SUM:                            23            233            135           1661
-------------------------------------------------------------------------------

```
## Dependencies
### Third Party
- `github.com/google/go-github/v60/github`
- `github.com/ogiusek/ioc/v2`
- `github.com/spf13/cobra`
- `golang.org/x/oauth2`
- `golang.org/x/tools/go/packages`