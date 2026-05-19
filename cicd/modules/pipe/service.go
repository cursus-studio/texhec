// main entry point responsible for setuping pipe hooks, running pipe hooks, running cicd pipeline
//
// 1. code quality stages:
// - dependencies
// - compilation
// - gosec
// - golangci-lint
// - tests
// 2. pipeline quality stages:
// - trivy
// 3. docs quality stages:
// - generate or verify docs using [docs](/cicd/modules/docs)
// - lychee
package pipe

type Service interface {
	Setup() error

	Sync() error
	Fix() error
	Cloud(commitHash string) error
	Verify(commitHash string) error
}
