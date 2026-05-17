# hooks
## Architecture
initializes and handles git hooks

## Types
### type Service
Type: `cicd/modules/hooks.Service`

#### method Service Handle
Type: `func(command string) error`

#### method Service Setup
Type: `func() error`


## Dependencies
`cicd/modules/hooks`:
  - `cicd/modules/hooks.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`
  - `cicd/world.Docs`
  - `cicd/world.Git`

### Third Party
- `github.com/ogiusek/ioc/v2`