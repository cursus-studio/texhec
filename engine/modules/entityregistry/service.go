// `entityregistry` allows us to define entities and components using struct tags
// example:
// ```go
//
//	type OurEntities struct {
//	  OurEntity ecs.EntityID `registered_component:"its_value"`
//	}
//
// ```
package entityregistry

import (
	"engine/services/ecs"
	"fmt"
)

var (
	ErrExpectedPointerToAStruct error = fmt.Errorf("expected pointer to a struct")
)

type Service interface {
	Register(structTagKey string, handler func(entity ecs.EntityID, structTagValue string))

	// can return ErrExpectedPointerToAStruct
	Populate(any) error
}

func GetRegistry[Registry any](s Service) (Registry, error) {
	var r Registry
	err := s.Populate(&r)
	return r, err
}
