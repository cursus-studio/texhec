// lists modified modules
package diff

// each method returns modules
type Service interface {
	// lists uncommited changed modules
	DiffUncommited() ([]string, error)

	// lists changed modules in previous commit
	DiffCommited() ([]string, error)
}
