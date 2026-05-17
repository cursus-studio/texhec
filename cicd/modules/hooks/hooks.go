// initializes and handles git hooks
package hooks

type Service interface {
	Setup() error
	Handle(command string) error
}
