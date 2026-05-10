package smooth

import (
	"engine/modules/transition"
	"engine/services/ecs"
)

type StartSystem ecs.SystemRegister
type StopSystem ecs.SystemRegister

type Service any

type SmoothConstraint[Component any] interface {
	transition.LerpConstraint[Component]
	// this method is a tag that component is smooothed
	Smooth()
}
