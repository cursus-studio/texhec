package docspkg

import (
	"readme/modules/docs"
	"readme/modules/docs/internal"
)

func NewService() docs.Service {
	return internal.NewService()
}
