// defines GUI to show up when batcher processes any task
package loading

import "engine/services/ecs"

type Service interface {
	ecs.SystemRegister
}
