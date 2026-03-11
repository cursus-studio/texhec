package definitions

import "engine/modules/transition"

const (
	_ transition.EasingID = iota
	LinearEasingFunction
	MyEasingFunction
	EaseOutElastic
)
