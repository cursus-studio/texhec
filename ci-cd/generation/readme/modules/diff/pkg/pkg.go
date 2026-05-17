package diffpkg

import (
	"readme/modules/diff"
	"readme/modules/diff/internal"
)

func NewService() diff.Service {
	return internal.NewService()
}
