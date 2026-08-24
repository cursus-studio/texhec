package interactionspkg

import (
	"engine/modules/interactions"
	"engine/modules/interactions/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) internal.Service {
		return internal.NewService(c)
	})
	ioc.Register(b, func(c ioc.Dic) interactions.Service {
		s := ioc.Get[internal.Service](c)
		s.Init()
		return s
	})
})

//

func FeaturePkg[Feature interactions.Feature](
	relations ...internal.RawRelation,
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) internal.FeatureService[Feature] {
			return internal.NewFeatureService[Feature](c, relations)
		})
		ioc.Wrap(b, func(c ioc.Dic, s internal.Service) {
			s.RegisterFeature(ioc.Get[internal.FeatureService[Feature]](c))
		})
	})
}

//

func StepPkg[StepT interactions.Step[State], State any](
	rule func(c ioc.Dic) func(state State) error,
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) internal.StepService[StepT, State] {
			return internal.NewStepService[StepT](c, rule(c))
		})
		ioc.Wrap(b, func(c ioc.Dic, s internal.Service) {
			s.RegisterStep(ioc.Get[internal.StepService[StepT, State]](c))
		})
	})
}

//

func InteractionPkg[State any]() ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) internal.InteractionService[State] {
			return internal.NewInteractionService[State](c)
		})
		ioc.Register(b, func(c ioc.Dic) interactions.InteractionService[State] {
			return ioc.Get[internal.InteractionService[State]](c)
		})
		ioc.Wrap(b, func(c ioc.Dic, s internal.Service) {
			s.RegisterInteraction(ioc.Get[internal.InteractionService[State]](c))
		})
	})
}
