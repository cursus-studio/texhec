# docs
## Architecture
This module is responsible for automatically generating documentation from comments for
projects following this project file structure.

This module doesn't rely on generic tools for documentation it requires dedicated comments.
This module is meant to be used in CI-CD pipeline.

Legend:
+ marks required comments
- marks optional comments

Where and how to write comments:
+ in package comment define core know how of the module
- `readme/TITLE.md` is used to overwrite automatic title
- `readme/ARCHITECTURE.md` is used to overwrite automatic architecture
- `readme/BENCH.md` is used to overwrite automatic benchmarks
- `readme/CHALLENGES.md` is used to spark discussions
- `readme/TODO.md` is great for contribution and notes

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8            148             73            986
Markdown                         1              1              0              5
-------------------------------------------------------------------------------
SUM:                             9            149             73            991
-------------------------------------------------------------------------------
```
## TODO
1. Check is there better readme extension than ".md".
Extension to preview:
- ".rst"
- ".adoc"

2. Re work module interface to take struct and output string

## Types
### type Service
Type: `cicd/modules/docs.Service`

#### method Service DiffModule
Type: `func(modulePath string) error`

#### method Service DiffProject
Type: `func(projectPath string) error`

#### method Service DiffTODO
Type: `func() error`

#### method Service GenerateModule
Type: `func(modulePath string) error`
Generates module documentation in `$modulePath/readme/README.md`

#### method Service GenerateProject
Type: `func(projectPath string) error`

#### method Service GenerateTODO
Type: `func() error`
reads TODO.md in all modules
generates readme/TODO.md

### type Config
Type: `cicd/modules/docs.Config`

## Variables
### var ErrMissingPackage
Type: `error`

### var ErrMissingPackageComments
Type: `error`

### var ErrInconsistentPackageComments
Type: `error`


## Dependencies
`cicd/modules/docs`:
  - `cicd/modules/docs.ErrInconsistentPackageComments`
  - `cicd/modules/docs.ErrMissingPackage`
  - `cicd/modules/docs.ErrMissingPackageComments`
  - `cicd/modules/docs.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`
  - `cicd/world.ProjectFS`

### Third Party
- `github.com/go-git/go-git/v5/plumbing/format/gitignore`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/tools/go/packages`