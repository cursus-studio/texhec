package economypkg

import (
	"core/game"
	"core/modules/economy"
	"core/modules/economy/internal"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"errors"
	"fmt"
	"strconv"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[economy.WalletComponent],
		typeregistrypkg.PkgT[economy.FactoryComponent],
		typeregistrypkg.PkgT[economy.CostComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) economy.Service {
		return internal.NewService(c)
	})

	ioc.Wrap(b, func(c ioc.Dic, b entityregistry.Service) {
		world := ioc.Get[game.GameWorld](c)
		b.Register("factory", func(entity ecs.EntityID, structTagValue string) {
			val, err := strconv.Atoi(structTagValue)
			if err != nil {
				world.Logger().Warn(errors.Join(
					fmt.Errorf("couldn't set for entity \"%v\" factory", entity),
					err,
				))
				return
			}
			money := economy.Money(val)
			world.Economy().Factory().Set(entity, economy.NewFactory(money))
		})
		b.Register("cost", func(entity ecs.EntityID, structTagValue string) {
			val, err := strconv.Atoi(structTagValue)
			if err != nil {
				world.Logger().Warn(errors.Join(
					fmt.Errorf("couldn't set for entity \"%v\" cost", entity),
					err,
				))
				return
			}
			money := economy.Money(val)
			world.Economy().Cost().Set(entity, economy.NewCost(money))
		})
	})
})
