package registry

import (
	"engine/services/ecs"
	"fmt"
)

var (
	ErrExpectedPointerToAStruct error = fmt.Errorf("expected pointer to a struct")
	ErrNotFoundHandlerForAField error = fmt.Errorf("not found handler for a field")
)

type Service interface {
	Register(structTagKey string, handler func(structTagValue string) ecs.EntityID)

	// can return ErrExpectedPointerToAStruct
	Populate(any) error
}
