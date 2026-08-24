package pathfindpkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/pathfind"
	"core/modules/pathfind/internal"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	gridpkg "engine/modules/grid/pkg"
	interactionspkg "engine/modules/interactions/pkg"
	relationpkg "engine/modules/relation/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/ogiusek/ioc/v2"
)

type FindPathFeature struct {
	Entity actions.FriendlyMobileEntityStep
	Coords actions.CoordsStep
}

func (f FindPathFeature) Event() any {
	return pathfind.NewFindPathEvent(f.Entity.State().Entity, f.Coords.State().Coords)
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[pathfind.TargetComponent],
		typeregistrypkg.PkgT[pathfind.SpeedComponent],
		typeregistrypkg.PkgT[pathfind.StepComponent],

		typeregistrypkg.PkgT[pathfind.FindPathEvent],

		interactionspkg.FeaturePkg[FindPathFeature](
			interactionspkg.NewCopyRelation[actions.CoordsCursorComponent](
				unsafe.Offsetof(FindPathFeature{}.Entity), unsafe.Offsetof(FindPathFeature{}.Coords)),
			interactionspkg.NewCopyRelation[actions.RegionAnchorComponent](
				unsafe.Offsetof(FindPathFeature{}.Entity), unsafe.Offsetof(FindPathFeature{}.Coords)),
		),
		relationpkg.MapRelationPkg(
			func(w ecs.World) ecs.DirtySet {
				set := ecs.NewDirtySet()
				arr := ecs.GetComponentArray[internal.ChunkObstructionComponent](w)
				arr.AddDirtySet(set)
				return set
			},
			func(w ecs.World) func(entity ecs.EntityID) (indexType internal.ChunkObstructionComponent, ok bool) {
				return ecs.GetComponentArray[internal.ChunkObstructionComponent](w).Get
			},
		),
		gridpkg.PkgT[internal.RegionFragment],
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
