# CI-CD

## Architecture
This project is responsible for code quality and commit history.

### Hosting
Hosted on Rasperry Pi and uses prometheus and grafana for monitoring.

### Commands
CI/CD tool exposes 5 commands:
- #### `go run cicd setup`
It initializes git hooks.

- #### `go run cicd sync`
It generates readmes in whole codebase.
It is used when readme generation algorithm is changes.

- #### `go run cicd fix`
It takes uncommited files and runs all tests on them and updates readmes and stages them.

- #### `go run cicd verify $compared-commit-hash`
`$compared-commit-hash` - By default only not commited changed are tested. In CI-CD pipeline previous successful commit is used here.

It runs all quality tests only on specific changed files.

- #### `go run cicd cloud $compared-commit-hash`
`$compared-commit-hash` - By default only not commited changed are tested. In CI-CD pipeline previous successful commit is used here.

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
- [docs](/cicd/modules/docs/readme/README.md)
- [git](/cicd/modules/git/readme/README.md)
- [pipe](/cicd/modules/pipe/readme/README.md)
- [projectfs](/cicd/modules/projectfs/readme/README.md)

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              19            225            107           1643
Groovy                           1              5              1             76
Markdown                         4             10              0             33
Dockerfile                       1              7              4             32
-------------------------------------------------------------------------------
SUM:                            25            247            112           1784
-------------------------------------------------------------------------------
```
## TODO
Ensure script wraps itself in `Dockerfile` if it runs on local machine

## Dependencies
### Third Party
- `github.com/go-git/go-git/v5/plumbing/format/gitignore`
- `github.com/google/go-github/v60/github`
- `github.com/ogiusek/ioc/v2`
- `github.com/spf13/cobra`
- `golang.org/x/oauth2`
- `golang.org/x/tools/go/packages`