package tilepkg

import (
	"bytes"
	"core/modules/definitions"
	"core/modules/tile"
	clicksystem "core/modules/tile/internal/clickSystem"
	obstructionsystem "core/modules/tile/internal/obstructionSystem"
	"core/modules/tile/internal/tilerenderer"
	"core/modules/tile/internal/tileservice"
	"core/modules/tile/internal/tilesystem"
	"engine"
	"engine/modules/assets"
	"engine/modules/collider"
	gridpkg "engine/modules/grid/pkg"
	"engine/modules/groups"
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
	"strconv"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct {
}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg pkg) Register(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
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
		tilerenderer.Package(),
		prototypepkg.PackageT[tile.TypeComponent](),
		prototypepkg.PackageT[tile.PosComponent](),
		prototypepkg.PackageT[tile.SizeComponent](),
		prototypepkg.PackageT[tile.RotComponent](),
		prototypepkg.PackageT[tile.LayerComponent](),

		prototypepkg.PackageT[tile.ObstructionComponent](),

		prototypepkg.PackageT[tile.SpeedComponent](),

		transitionpkg.PackageT[tile.PosComponent](),
		transitionpkg.PackageT[tile.RotComponent](),
	} {
		pkg.Register(b)
	}

	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// events
			Register(tile.HoverEvent{})
	})

	ioc.Register(b, func(c ioc.Dic) tile.Service {
		return tileservice.NewService(c)
	})

	ioc.Register(b, func(c ioc.Dic) tile.System {
		systems := []tile.System{
			tilesystem.NewSystem(c),
			clicksystem.NewSystem(c),
			obstructionsystem.NewSystem(c),
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

	ioc.Wrap(b, func(c ioc.Dic, b assets.Service) {
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

	ioc.Wrap(b, func(c ioc.Dic, b registry.Service) {
		type World struct {
			engine.World `inject:"1"`
			Tile         ioc.Lazy[tile.Service] `inject:"1"`
			Definitions  definitions.Assets     `inject:"1"`
		}
		world := ioc.GetServices[World](c)
		var counter tile.ID
		b.Register("object", func(entity ecs.EntityID, structTagValue string) {
			var layer tile.Coord
			switch structTagValue {
			case "construct":
				layer = definitions.ConstructLayer
			case "unit":
				layer = definitions.UnitLayer
			default:
				return
			}
			world.Tile().Rot().Set(entity, tile.NewRot(0))
			world.Tile().Layer().Set(entity, tile.NewLayer(layer))
			world.Render.Mesh().Set(entity, render.NewMesh(world.Definitions.SquareMesh))
			world.Render.Texture().Set(entity, render.NewTexture(entity))
			world.Groups.Component().Set(entity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())

			world.Collider.Component().Set(entity, collider.NewCollider(world.Definitions.SquareCollider))
		})
		b.Register("tile", func(entity ecs.EntityID, structTagValue string) {
			counter++
			world.Tile().TileType().Set(entity, tile.NewTileType(counter))
		})
		b.Register("obstruction", func(entity ecs.EntityID, structTagValue string) {
			var obstruction tile.Obstruction
			if strings.Contains(structTagValue, "water") {
				obstruction |= definitions.WaterObstruction
			}
			if strings.Contains(structTagValue, "lowland") {
				obstruction |= definitions.LowlandObstruction
			}
			if strings.Contains(structTagValue, "air") {
				obstruction |= definitions.AirspaceObstruction
			}
			world.Tile().Obstruction().Set(entity, tile.NewObstruction(obstruction))
		})
		b.Register("size", func(entity ecs.EntityID, structTagValue string) {
			errInvalidFormat := fmt.Errorf("size should be in format \"1x1\" where first number is width and second is height")
			xy := strings.Split(structTagValue, "x")
			if len(xy) != 2 {
				world.Logger.Warn(errInvalidFormat)
				return
			}
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				world.Logger.Warn(errInvalidFormat)
				return
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				world.Logger.Warn(errInvalidFormat)
				return
			}
			world.Tile().Size().Set(entity, tile.NewSize(x, y))
		})
		b.Register("speed", func(entity ecs.EntityID, structTagValue string) {
			v, err := strconv.Atoi(structTagValue)
			if err != nil {
				world.Logger.Warn(err)
				return
			}
			speed := tile.NewSpeed(v)
			if int(speed.InvSpeed) != v {
				world.Logger.Warn(fmt.Errorf("speed has to be clamped between 0 and 255"))
				return
			}
			world.Tile().Speed().Set(entity, speed)
		})
	})
}
