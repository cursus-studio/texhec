package interactionspkg

import (
	"engine/modules/interactions"
	"engine/modules/interactions/internal"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

func FeaturePkg[Event any](
	name interactions.Name,
	interactionStates []reflect.Type,
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) interactions.FeatureService[Event] {
			return internal.NewFeatureService[Event](c, name, interactionStates)
		})
		ioc.Wrap(b, func(c ioc.Dic, s internal.Service) {
			s.RegisterFeature(ioc.Get[interactions.FeatureService[Event]](c))
		})
	})
}

func InteractionPkg[State any](name interactions.Name) ioc.Pkg {
	return func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) interactions.InteractionService[State] {
			return internal.NewInteractionService[State](c, name)
		})
		ioc.Wrap(b, func(c ioc.Dic, s internal.Service) {
			s.RegisterInteraction(ioc.Get[interactions.InteractionService[State]](c))
		})
	}
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, internal.NewService)
	ioc.Register(b, func(c ioc.Dic) interactions.Service {
		return ioc.Get[internal.Service](c)
	})
})
