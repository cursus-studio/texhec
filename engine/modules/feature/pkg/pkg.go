package featurepkg

import (
	"engine/modules/feature"
	"engine/modules/feature/internal"

	"github.com/ogiusek/ioc/v2"
)

func FieldPkg[FieldType any](b ioc.Builder) {
	ioc.Register(b, internal.NewFieldService[FieldType])
}

func RulePkg[RuleType feature.Rule[DataIn, AppliedTo], DataIn any, AppliedTo any](
	applies func(c ioc.Dic) func(DataIn, AppliedTo) error,
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) feature.RuleService[RuleType, DataIn, AppliedTo] {
			return internal.NewRuleService[RuleType](c, applies(c))
		})
	})
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, internal.NewConfig)
	ioc.Register(b, internal.NewService)
	ioc.Register(b, func(c ioc.Dic) feature.Service {
		return ioc.Get[internal.Service](c)
	})
})
