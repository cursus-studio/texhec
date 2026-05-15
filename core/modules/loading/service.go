package loading

import "engine/services/ecs"

type Service interface {
	ecs.SystemRegister
}
