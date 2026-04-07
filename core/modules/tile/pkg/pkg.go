package tilepkg

import (
	"bytes"
	"core/modules/definitions"
	"core/modules/tile"
	clicksystems "core/modules/tile/internal/clickSystems"
	"core/modules/tile/internal/tilerenderer"
	"core/modules/tile/internal/tileservice"
	"core/modules/tile/internal/tileui"
	"engine"
	"engine/modules/assets"
	"engine/modules/collider"
	gridpkg "engine/modules/grid/pkg"
	"engine/modules/groups"
	"engine/modules/inputs"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/registry"
	relationpkg "engine/modules/relation/pkg"
	"engine/modules/render"
	transitionpkg "engine/modules/transition/pkg"
	"engine/services/codec"
	"engine/services/ecs"
	gtexture "engine/services/graphics/texture"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct {
	pkgs []ioc.Pkg
}

func Package() ioc.Pkg {
	return pkg{
		[]ioc.Pkg{
			gridpkg.Package[tile.ID](tile.NewHoverEvent),
			gridpkg.Package[tile.Obstruction](nil),
			relationpkg.SpatialRelationPackage(
				func(w ecs.World) ecs.DirtySet {
					dirtySet := ecs.NewDirtySet()
					ecs.GetComponentsArray[tile.TypeComponent](w).AddDirtySet(dirtySet)
					return dirtySet
				},
				func(w ecs.World) func(entity ecs.EntityID) (tile.ID, bool) {
					componentArray := ecs.GetComponentsArray[tile.TypeComponent](w)
					return func(entity ecs.EntityID) (tile.ID, bool) {
						comp, ok := componentArray.Get(entity)
						return comp.ID, ok
					}
				},
				func(index tile.ID) uint32 { return uint32(index) },
			),
			tileservice.Package(),
			tilerenderer.Package(),
			prototypepkg.PackageT[tile.TypeComponent](),
			prototypepkg.PackageT[tile.PosComponent](),
			prototypepkg.PackageT[tile.SizeComponent](),
			prototypepkg.PackageT[tile.RotComponent](),
			prototypepkg.PackageT[tile.LayerComponent](),
		},
	}
}

func (pkg pkg) Register(b ioc.Builder) {
	for _, pkg := range pkg.pkgs {
		pkg.Register(b)
	}

	ioc.WrapService(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// events
			Register(tile.HoverEvent{})
	})

	for _, pkg := range []ioc.Pkg{
		transitionpkg.PackageT[tile.PosComponent](),
		transitionpkg.PackageT[tile.SizeComponent](),
		transitionpkg.PackageT[tile.RotComponent](),
	} {
		pkg.Register(b)
	}

	ioc.RegisterSingleton(b, func(c ioc.Dic) tile.System {
		systems := []tile.System{
			tileui.NewSystem(c),
			clicksystems.NewSystems(c),
		}
		return ecs.NewSystemRegister(func() error {
			for _, system := range systems {
				if err := system.Register(); err != nil {
					return err
				}
			}
			return nil
		})
	})

	ioc.WrapService(b, func(c ioc.Dic, registry registry.Service) {
		var counter tile.ID
		registry.Register("tile", func(entity ecs.EntityID, structTagValue string) {
			counter++
			tileService := ioc.Get[tile.Service](c)
			tileService.TileType().Set(entity, tile.NewTileType(counter))
			tileService.Obstruction().Set(entity, tile.NewObstruction(definitions.WaterObstruction))
		})
	})

	ioc.WrapService(b, func(c ioc.Dic, b assets.Service) {
		b.Register("biom", func(path assets.PathComponent) (assets.Asset, error) {
			images := [6][]image.Image{}
			directory, _ := strings.CutSuffix(path.Path, ".biom")

			for i := range 6 {
				tileDir := fmt.Sprintf("%v/%v", directory, i+1)
				files, err := os.ReadDir(tileDir)
				if err != nil {
					return nil, err
				}
				if len(files) == 0 {
					return nil, fmt.Errorf("there is no tile variant for %v tile", i)
				}

				for _, file := range files {
					filePath := fmt.Sprintf("%v/%v", tileDir, file.Name())
					source, err := os.ReadFile(filePath)
					if err != nil {
						return nil, err
					}
					imgFile := bytes.NewBuffer(source)
					img, _, err := image.Decode(imgFile)
					if err != nil {
						return nil, err
					}
					img = gtexture.NewImage(img).FlipV().Image()
					images[i] = append(images[i], img)
				}
			}

			return tile.NewBiomAsset(images)
		})
	})

	ioc.WrapService(b, func(c ioc.Dic, b registry.Service) {
		type World struct {
			engine.World `inject:"1"`
			Tile         tile.Service        `inject:"1"`
			Definitions  definitions.BuiltIn `inject:"1"`
		}
		b.Register("unit", func(entity ecs.EntityID, structTagValue string) {
			world := ioc.GetServices[World](c)
			world.Tile.Layer().Set(entity, tile.NewLayer(definitions.UnitLayer))
			world.Tile.Obstruction().Set(entity, tile.NewObstruction(definitions.LowlandsObstruction))

			world.Render.Mesh().Set(entity, render.NewMesh(world.Definitions.SquareMesh))
			world.Render.Texture().Set(entity, render.NewTexture(entity))
			world.Groups.Component().Set(entity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())

			world.Collider.Component().Set(entity, collider.NewCollider(world.Definitions.SquareCollider))
			world.Inputs.LeftClick().Set(entity, inputs.NewLeftClick(tile.NewClickEntityEvent()))
			world.Inputs.Stack().Set(entity, inputs.StackComponent{})
		})
		b.Register("construct", func(entity ecs.EntityID, structTagValue string) {
			world := ioc.GetServices[World](c)
			world.Tile.Layer().Set(entity, tile.NewLayer(definitions.ConstructLayer))
			world.Tile.Obstruction().Set(entity, tile.NewObstruction(definitions.LowlandsObstruction))

			world.Render.Mesh().Set(entity, render.NewMesh(world.Definitions.SquareMesh))
			world.Render.Texture().Set(entity, render.NewTexture(entity))
			world.Groups.Component().Set(entity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())

			world.Collider.Component().Set(entity, collider.NewCollider(world.Definitions.SquareCollider))
			world.Inputs.LeftClick().Set(entity, inputs.NewLeftClick(tile.NewClickEntityEvent()))
			world.Inputs.Stack().Set(entity, inputs.StackComponent{})
		})
	})
}
