// finds modified modules since last commit
package diff

type Service interface {
	GetModifiedModules() ([]string, error)
}
