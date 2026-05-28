package tilepkg

import (
	"bytes"
	"core/game"
	"core/modules/definitions"
	"core/modules/tile"
	clicksystem "core/modules/tile/internal/clickSystem"
	"core/modules/tile/internal/tilerenderer"
	"core/modules/tile/internal/tileservice"
	"core/modules/tile/internal/tilesystem"
	"engine/modules/assets"
	"engine/modules/collider"
	"engine/modules/entityregistry"
	"engine/modules/graphics"
	gridpkg "engine/modules/grid/pkg"
	relationpkg "engine/modules/relation/pkg"
	"engine/modules/render"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/services/ecs"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/ioc/v2"
)

type config struct {
	tileSize float32
}

func NewConfig() Config {
	return &config{
		tileSize: 100,
	}
}

func (c *config) GetTileSize() float32     { return c.tileSize }
func (c *config) SetTileSize(size float32) { c.tileSize = size }

type Config interface {
	GetTileSize() float32
	SetTileSize(float32)
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		gridpkg.PkgT[tile.ID](),
		relationpkg.SpatialRelationPkg(
			func(w ecs.World) ecs.DirtySet {
				dirtySet := ecs.NewDirtySet()
				ecs.GetComponentsArray[tile.Component](w).AddDirtySet(dirtySet)
				return dirtySet
			},
			func(w ecs.World) func(entity ecs.EntityID) (tile.ID, bool) {
				componentArray := ecs.GetComponentsArray[tile.Component](w)
				return func(entity ecs.EntityID) (tile.ID, bool) {
					comp, ok := componentArray.Get(entity)
					return comp.ID, ok
				}
			},
			func(index tile.ID) uint32 { return uint32(index) },
		),
		typeregistrypkg.PkgT[tile.HoverEvent],

		typeregistrypkg.PkgT[tile.Component],
		typeregistrypkg.PkgT[tile.PosComponent],
		typeregistrypkg.PkgT[tile.SizeComponent],
		typeregistrypkg.PkgT[tile.RotComponent],
		typeregistrypkg.PkgT[tile.LayerComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) Config {
		return NewConfig()
	})
	ioc.Wrap(b, func(c ioc.Dic, config gridpkg.ConfigT[tile.ID]) {
		config.SetHoverEvent(tile.NewHoverEvent)
	})

	ioc.Register(b, func(c ioc.Dic) graphics.VBOFactory[tile.ID] {
		return func() graphics.VBOSetter[tile.ID] {
			vbo := graphics.NewVBO[tile.ID](func() {
				var i uint32 = 0

				gl.VertexAttribIPointerWithOffset(i, 1, gl.UNSIGNED_BYTE,
					int32(unsafe.Sizeof(tile.ID(0))), uintptr(0))
				gl.EnableVertexAttribArray(i)
			})
			return vbo
		}
	})

	ioc.Register(b, func(c ioc.Dic) tile.Service {
		return tileservice.NewService(c,
			ecs.NewSystemRegister(func() error {
				errs := ecs.RegisterSystems(
					tilesystem.NewSystem(c),
					clicksystem.NewSystem(c),
				)
				if len(errs) != 0 {
					return errs[0]
				}
				return nil
			}),
			ecs.NewSystemRegister(func() error {
				// tilerenderer
				// TODO
				// currently doesn't support animated tiles
				// always renderes first frame if something is animated

				return tilerenderer.NewSystem(c)
			}),
			ioc.Get[Config](c).GetTileSize(),
		)
	})

	ioc.Wrap(b, func(c ioc.Dic, b assets.Service) {
		world := ioc.Get[game.GameWorld](c)
		b.Register("biome", func(path assets.PathComponent) (assets.Asset, error) {
			images := [6][]image.Image{}
			directory, _ := strings.CutSuffix(path.Path, ".biome")

			for i := range 6 {
				tileDir := fmt.Sprintf("%v/%v", directory, i+1)
				files, err := os.ReadDir(tileDir)
				if err != nil {
					return nil, err
				}
				if len(files) == 0 {
					return nil, fmt.Errorf("there is no tile variant for %v tile", i)
				}
				root, err := os.OpenRoot(tileDir)
				if err != nil {
					return nil, err
				}
				defer func() {
					_ = root.Close()
				}()

				for _, file := range files {
					source, err := root.ReadFile(file.Name())
					if err != nil {
						return nil, err
					}
					imgFile := bytes.NewBuffer(source)
					img, _, err := image.Decode(imgFile)
					if err != nil {
						return nil, err
					}
					img = world.Graphics().NewImage(img).FlipV().Image()
					images[i] = append(images[i], img)
				}
			}

			return world.Tile().NewBiomeAsset(images)
		})
	})

	ioc.Wrap(b, func(c ioc.Dic, b entityregistry.Service) {
		world := ioc.Get[game.GameWorld](c)
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
			world.Render().Mesh().Set(entity, render.NewMesh(world.Definitions().Assets().SquareMesh))
			world.Render().Texture().Set(entity, render.NewTexture(entity))
			world.Groups().InheritGroups(entity)

			world.Collider().Component().Set(entity, collider.NewCollider(world.Definitions().Assets().SquareCollider))
		})
		b.Register("tile", func(entity ecs.EntityID, structTagValue string) {
			counter++
			world.Tile().Component().Set(entity, tile.NewTile(counter))
		})
		b.Register("size", func(entity ecs.EntityID, structTagValue string) {
			errInvalidFormat := fmt.Errorf("size should be in format \"1x1\" where first number is width and second is height")
			xy := strings.Split(structTagValue, "x")
			if len(xy) != 2 {
				world.Logger().Log(errInvalidFormat)
				return
			}
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				world.Logger().Log(errInvalidFormat)
				return
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				world.Logger().Log(errInvalidFormat)
				return
			}
			world.Tile().Size().Set(entity, tile.NewSize(x, y))
		})
	})
})
