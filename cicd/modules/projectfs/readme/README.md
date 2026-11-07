# projectfs
## Architecture
responsible for accessing project file structure and project specific directories

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             22              4            168
-------------------------------------------------------------------------------
SUM:                             3             22              4            168
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