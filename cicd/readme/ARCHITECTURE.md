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
