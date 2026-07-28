// defines loging abstraction which can be easily expanded to log information in GUI or sent to developer
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

	// wrapper around Log(errors.Join(ErrInfo, err))
	Info(error)
	// wrapper around Log(err)
	Warn(error)
	// wrapper around Log(errors.Join(ErrFatal, err))
	Fatal(error)
}
