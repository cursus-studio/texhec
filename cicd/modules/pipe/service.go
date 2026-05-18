// main entry point responsible for setuping pipe hooks, running pipe hooks, running cicd pipeline
package pipe

// check dependencies, does compile, gosec, trivy, golangci-lint
// verify docs generation, lychee
// tests
type Service interface {
	Setup() error

	Sync() error
	Fix() error
	Cloud(commitHash string) error
	Verify(commitHash string) error
}
