// This module is responsible for automatically generating documentation from comments for
// projects following this project file structure.
//
// This module doesn't rely on generic tools for documentation it requires dedicated comments.
// This module is meant to be used in CI-CD pipeline.
//
// Legend:
// + marks required comments
// - marks optional comments
//
// Where and how to write comments:
// + in package comment define core know how of the module
// - `readme/TITLE.md` is used to overwrite automatic title
// - `readme/ARCHITECTURE.md` is used to overwrite automatic architecture
// - `readme/BENCH.md` is used to overwrite automatic benchmarks
// - `readme/CHALLENGES.md` is used to spark discussions
// - `readme/TODO.md` is great for contribution and notes
package docs

import "errors"

var (
	ErrMissingPackage              = errors.New("missing `package name` to compose title")
	ErrMissingPackageComments      = errors.New("missing package comments")
	ErrInconsistentPackageComments = errors.New("inconsistent package comments")
)

type Config struct {
}

type Service interface {
	// Generates module documentation in `$modulePath/readme/README.md`
	GenerateModule(modulePath string) error
	DiffModule(modulePath string) error

	GenerateProject(projectPath string) error
	DiffProject(projectPath string) error

	// reads TODO.md in all modules
	// generates readme/TODO.md
	GenerateTODO() error
	DiffTODO() error
}
