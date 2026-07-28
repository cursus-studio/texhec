// defines GUI to show up when batcher processes any task
package loading

import "engine/modules/ecs"

type Service interface {
	ecs.SystemRegister
}
