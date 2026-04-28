package logger

import "errors"

var ( // Built in severity tags
	ErrInfo error = errors.New("info")
	// warn is default
	ErrFatal error = errors.New("fatal")
)

func IsWarning(meta error) bool {
	return !errors.Is(meta, ErrInfo) && !errors.Is(meta, ErrFatal)
}

type Service interface {
	// error is composed from:
	// - multiple meta tags which can contain audience or severity (optional)
	// - error message
	Log(error)
}
