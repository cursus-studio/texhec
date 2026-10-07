package featurepkg

import (
	"engine/modules/feature"
	"engine/modules/feature/internal"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

func FieldPkg[FieldType any](b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[feature.FieldRenderComponent[FieldType]],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, internal.NewFieldService[FieldType])
}

func NoInputRulePkg[RuleType feature.NoInputRule[AppliedTo], AppliedTo any](
	applies func(c ioc.Dic) func(AppliedTo) error,
	err error,
) ioc.Pkg {
	return RulePkg[RuleType](func(c ioc.Dic) func(AppliedTo, AppliedTo) error {
		applies := applies(c)
		return func(at1, at2 AppliedTo) error {
			return applies(at2)
		}
	}, err)
}
func RulePkg[RuleType feature.Rule[DataIn, AppliedTo], DataIn any, AppliedTo any](
	applies func(c ioc.Dic) func(DataIn, AppliedTo) error,
	err error,
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		pkgs := []ioc.Pkg{
			typeregistrypkg.PkgT[feature.RuleRenderComponent[RuleType, DataIn, AppliedTo]],
		}
		for _, pkg := range pkgs {
			pkg(b)
		}
		ioc.Register(b, func(c ioc.Dic) feature.RuleService[RuleType, DataIn, AppliedTo] {
			return internal.NewRuleService[RuleType](c, applies(c), err)
		})
	})
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[feature.RenderComponent],
		typeregistrypkg.PkgT[feature.InteractionsComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, internal.NewConfig)
	ioc.Register(b, internal.NewService)
	ioc.Register(b, func(c ioc.Dic) feature.Service {
		return ioc.Get[internal.Service](c)
	})
})
