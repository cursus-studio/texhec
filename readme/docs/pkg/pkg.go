package docspkg

import (
	"readme/docs"
	"readme/docs/internal"
)

func NewService() docs.Service {
	return internal.NewService()
}
