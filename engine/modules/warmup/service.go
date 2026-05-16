// runs all lazy listeners
// it should be used executing anything concurrently
package warmup

type Service interface {
	WarmUp()
}

type Event struct{}
