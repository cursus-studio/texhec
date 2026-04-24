package definitionspkg

import (
	"core/modules/definitions"
	"core/modules/definitions/internal"
	"engine/modules/entityregistry"
	"engine/modules/transition"
	"engine/services/ecs"
	_ "image/png"
	"math"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) definitions.Service {
		return internal.NewService(c)
	})

	//
	//
	//

	// animations

	transitions := map[string]func(t transition.Progress) transition.Progress{
		"linear": func(t transition.Progress) transition.Progress {
			return t
		},
		"my easing": func(t transition.Progress) transition.Progress {
			const n1 = 7.5625
			const d1 = 2.75

			if t < 1/d1 { // First segment of the bounce (rising curve)
				return n1 * t * t
			} else if t < 2/d1 { // Second segment (peak of the first bounce)
				t -= 1.5 / d1
				return n1*t*t + 0.75
			} else if t < 2.5/d1 { // Third segment (peak of the second, smaller bounce)
				t -= 2.25 / d1
				return n1*t*t + 0.9375
			} else { // Final segment (settling)
				t -= 2.625 / d1
				return n1*t*t + 0.984375
			}
		},
		"ease out elastic": func(t transition.Progress) transition.Progress {
			const c1 float64 = 10
			const c2 float64 = .75
			const c3 float64 = (2 * math.Pi) / 3
			if t == 0 {
				return 0
			}
			if t == 1 {
				return 1
			}
			x := float64(t)
			x = math.Pow(2, -c1*x)*
				math.Sin((x*c1-c2)*c3) +
				1
			return transition.Progress(x)
		},
	}

	ioc.Wrap(b, func(c ioc.Dic, b entityregistry.Service) {
		b.Register("transition", func(entity ecs.EntityID, structTagValue string) {
			transitionService := ioc.Get[transition.Service](c)
			easing, ok := transitions[structTagValue]
			if !ok {
				easing = func(t transition.Progress) transition.Progress { return t }
			}
			transitionService.EasingFunction().Set(entity, transition.NewEasingFunction(easing))
		})
	})
})
