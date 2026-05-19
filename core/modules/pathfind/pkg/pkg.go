package pathfindpkg

import (
	"core/game"
	"core/modules/pathfind"
	"core/modules/pathfind/internal"
	"engine/modules/entityregistry"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/services/ecs"
	"fmt"
	"strconv"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[pathfind.TargetComponent],
		typeregistrypkg.PkgT[pathfind.SpeedComponent],
		typeregistrypkg.PkgT[pathfind.StepComponent],

		typeregistrypkg.PkgT[pathfind.SelectEvent],
		typeregistrypkg.PkgT[pathfind.PreviewPathEvent],
		typeregistrypkg.PkgT[pathfind.FindPathEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) pathfind.Service {
		return internal.NewService(c)
	})

	ioc.Wrap(b, func(c ioc.Dic, b entityregistry.Service) {
		world := ioc.Get[game.GameWorld](c)
		b.Register("speed", func(entity ecs.EntityID, structTagValue string) {
			v, err := strconv.Atoi(structTagValue)
			if err != nil {
				world.Logger().Log(err)
				return
			}
			speed := pathfind.NewSpeed(v)
			if int(speed.InvSpeed) != v {
				world.Logger().Log(fmt.Errorf("speed has to be clamped between 0 and 255"))
				return
			}
			world.Pathfind().Speed().Set(entity, speed)
		})
	})
})
