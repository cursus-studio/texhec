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

func RulePkg[RuleType feature.Rule](
	applies func(c ioc.Dic) func(RuleType, feature.AllFields) error,
	err error,
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		pkgs := []ioc.Pkg{
			typeregistrypkg.PkgT[feature.RuleRenderComponent[RuleType]],
		}
		for _, pkg := range pkgs {
			pkg(b)
		}
		ioc.Register(b, func(c ioc.Dic) feature.RuleService[RuleType] {
			return internal.NewRuleService(c, applies(c), err)
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
