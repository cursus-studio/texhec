# projectfs
## Architecture
responsible for accessing project file structure and project specific directories

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               5             40             10            296
-------------------------------------------------------------------------------
SUM:                             5             40             10            296
-------------------------------------------------------------------------------
```
## Dependencies
`cicd/modules/projectfs`:
  - `cicd/modules/projectfs.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSet`

### Third Party
- `github.com/ogiusek/ioc/v2`