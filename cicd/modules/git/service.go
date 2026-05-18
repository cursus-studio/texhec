// responsible for git integration
package git

// each method returns modules
type Service interface {
	DiffNotCommited() ([]string, error)
	DiffPrevCommit() ([]string, error)
	// passing empty string will compare to not commited
	DiffCompare(commitHash string) ([]string, error)

	Stage(...string) error
}
