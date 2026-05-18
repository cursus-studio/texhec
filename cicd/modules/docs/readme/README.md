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
- `readme/BENCH.md` is used to overwrite automatic benchmarks
- `readme/CHALLENGES.md` is used to spark discussions
- `readme/TODO.md` is great for contribution and notes

## Types
### type Service
Type: `cicd/modules/docs.Service`

#### method Service DiffModuleDocs
Type: `func(modulePath string) error`

#### method Service GenerateModuleDocs
Type: `func(modulePath string) error`
Generates module documentation in `$modulePath/readme/README.md`

## Variables
### var ErrMissingPackage
Type: `error`

### var ErrMissingPackageComments
Type: `error`

### var ErrInconsistentPackageComments
Type: `error`


## TODO
1. Check is there better readme extension than ".md".
Extension to preview:
- ".rst"
- ".adoc"

2. Re work module interface to take struct and output string

## Dependencies
`cicd/modules/docs`:
  - `cicd/modules/docs.ErrInconsistentPackageComments`
  - `cicd/modules/docs.ErrMissingPackage`
  - `cicd/modules/docs.ErrMissingPackageComments`
  - `cicd/modules/docs.Service`

`cicd/world`:
  - `cicd/world.CICDWorld`

### Third Party
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/tools/go/packages`