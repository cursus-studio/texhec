# projectfs
## Architecture
saves results in files

## Types
### type Service
Type: `cicd/modules/projectfs.Service`

#### method Service AllModules
Type: `func() ([]string, error)`

#### method Service AllProjects
Type: `func() ([]string, error)`

#### method Service FilesModules
Type: `func(files []string) []string`

#### method Service FilesProjects
Type: `func(files []string) []string`

#### method Service Save
Type: `func(file string, content string) error`


## Dependencies
`cicd/modules/projectfs`:
  - `cicd/modules/projectfs.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSet`

### Third Party
- `github.com/ogiusek/ioc/v2`