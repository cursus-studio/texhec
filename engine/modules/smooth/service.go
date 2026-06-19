// catches changes on tick and applies them smoothly between ticks
package smooth

import (
	"engine/modules/ecs"
	"engine/modules/transition"
)

type Service interface {
	Start() ecs.SystemRegister
	Stop() ecs.SystemRegister
}

type SmoothConstraint[Component any] interface {
	transition.LerpConstraint[Component]
	// this method is a tag that component is smooothed
	Smooth()
}
