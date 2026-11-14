// responsible for git integration
package git

type State string

const (
	Pending State = "pending"
	Success State = "success"
	Failure State = "failure"
	Error   State = "error"
)

// each method returns modules
type Service interface {
	DiffNotCommited() ([]string, error)
	DiffPrevCommit() ([]string, error)
	// passing empty string will compare to not commited
	DiffCompare(commitHash string) ([]string, error)

	// can take `patch`, `minor`, `major`, `$version` and outputs `$version`
	Version(arg string) (version string, err error)
	Release(version, title string, files []string) error

	Stage(...string) error
	SetStatus(status State, decs string) error
}
