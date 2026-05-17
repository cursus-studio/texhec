package docspkg

import (
	"cicd/modules/docs"
	"cicd/modules/docs/internal"
)

func NewService() docs.Service {
	return internal.NewService()
}
