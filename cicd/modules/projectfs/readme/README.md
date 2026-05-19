# projectfs
## Architecture
responsible for accessing project file structure and project specific directories

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

#### method Service ProjectModules
Type: `func(project string) []string`

#### method Service ProjectServices
Type: `func(project string) []string`

#### method Service Save
Type: `func(file string, content string) error`


## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             21              4            153
-------------------------------------------------------------------------------
SUM:                             3             21              4            153
-------------------------------------------------------------------------------

```
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