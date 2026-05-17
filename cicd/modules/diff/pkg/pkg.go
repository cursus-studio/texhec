package diffpkg

import (
	"cicd/modules/diff"
	"cicd/modules/diff/internal"
)

func NewService() diff.Service {
	return internal.NewService()
}
