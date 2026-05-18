# pipe
## Architecture
pipeline:
-

## Types
### type Service
Type: `cicd/modules/pipe.Service`
check dependencies, does compile, gosec, trivy, golangci-lint
verify docs generation, lychee
tests

#### method Service Fix
Type: `func() error`

#### method Service Setup
Type: `func() error`

#### method Service Sync
Type: `func() error`

#### method Service Verify
Type: `func(commitHash string) error`


## Dependencies
`cicd/modules/pipe`:
  - `cicd/modules/pipe.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`
  - `cicd/world.Docs`
  - `cicd/world.Git`
  - `cicd/world.ProjectFS`

### Third Party
- `github.com/ogiusek/ioc/v2`