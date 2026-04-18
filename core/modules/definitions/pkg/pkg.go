package definitionspkg

import (
	"core/modules/definitions"
	"core/modules/deploy"
	gamescenes "core/scenes"
	"engine/modules/assets"
	"engine/modules/collider"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/registry"
	"engine/modules/render"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/services/ecs"
	"engine/services/graphics/vao/ebo"
	"image"
	"image/color"
	_ "image/png"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	// register specific files
	ioc.Wrap(b, func(c ioc.Dic, b assets.Service) {
		b.Register("blank texture", func(_ assets.PathComponent) (assets.Asset, error) {
			img := image.NewRGBA(image.Rect(0, 0, 1, 1))
			white := color.RGBA{255, 255, 255, 255}
			img.Set(0, 0, white)
			asset, err := render.NewTextureAsset(img)
			return asset, err
		})
		b.Register("square mesh", func(_ assets.PathComponent) (assets.Asset, error) {
			vertices := []render.Vertex{
				{Pos: [3]float32{1, 1, 1}, TexturePos: [2]float32{1, 1}},
				{Pos: [3]float32{1, -1, 1}, TexturePos: [2]float32{1, 0}},
				{Pos: [3]float32{-1, -1, 1}, TexturePos: [2]float32{0, 0}},
				{Pos: [3]float32{-1, 1, 1}, TexturePos: [2]float32{0, 1}},
			}

			indices := []ebo.Index{
				0, 1, 2,
				0, 2, 3,
			}
			asset := render.NewMeshAsset(vertices, indices)
			return asset, nil
		})

		b.Register("square collider", func(_ assets.PathComponent) (assets.Asset, error) {
			asset := collider.NewColliderAsset(
				[]collider.AABB{collider.NewAABB(mgl32.Vec3{-1, -1}, mgl32.Vec3{1, 1})},
				[]collider.Range{collider.NewRange(collider.Leaf, 0, 2)},
				[]collider.Polygon{
					collider.NewPolygon(mgl32.Vec3{-1, -1, 0}, mgl32.Vec3{+1, -1, 0}, [3]float32{-1, +1, 0}),
					collider.NewPolygon(mgl32.Vec3{+1, +1, 0}, mgl32.Vec3{+1, -1, 0}, [3]float32{-1, +1, 0}),
				})
			return asset, nil
		})
	})

	ioc.Register(b, func(c ioc.Dic) definitions.Tiles {
		world := ioc.GetServices[gamescenes.GameWorld](c)
		def, err := registry.GetRegistry[definitions.Tiles](world.Registry())
		world.Logger().Warn(err)
		return def
	})
	ioc.Register(b, func(c ioc.Dic) definitions.Objects {
		world := ioc.GetServices[gamescenes.GameWorld](c)
		def, err := registry.GetRegistry[definitions.Objects](world.Registry())
		world.Logger().Warn(err)

		{
			world.Deploy().Component().Set(def.Tank, deploy.NewDeploy(def.Tank, def.Farm))
			world.Deploy().Link().Set(def.Tank, deploy.NewLink(def.Tank))
		}
		{
			world.Deploy().Component().Set(def.Farm, deploy.NewDeploy(def.Tank, def.Farm))
			world.Deploy().Link().Set(def.Farm, deploy.NewLink(def.Farm))
		}
		return def
	})
	ioc.Register(b, func(c ioc.Dic) definitions.Hud {
		world := ioc.GetServices[gamescenes.GameWorld](c)
		def, err := registry.GetRegistry[definitions.Hud](world.Registry())
		world.Logger().Warn(err)

		{
			btnAsset, err := assets.GetAsset[render.TextureAsset](world.Assets(), def.Btn)
			if err != nil {
				world.Logger().Warn(err)
			}
			btnAspectRatio := btnAsset.AspectRatio()
			world.Groups().Inherit().Set(def.Btn, groups.InheritGroupsComponent{})
			world.Groups().Component().Set(def.Btn, groups.EmptyGroups())

			world.Transform().AspectRatio().Set(def.Btn, transform.NewAspectRatio(float32(btnAspectRatio.Dx()), float32(btnAspectRatio.Dy()), 0, transform.PrimaryAxisX))
			world.Transform().Parent().Set(def.Btn, transform.NewParent(transform.RelativePos|transform.RelativeSizeX))
			world.Transform().MaxSize().Set(def.Btn, transform.NewMaxSize(0, 50, 0))
			world.Transform().Size().Set(def.Btn, transform.NewSize(1, 50, 1))

			world.Render().Mesh().Set(def.Btn, render.NewMesh(world.Definitions().SquareMesh))
			world.Render().Texture().Set(def.Btn, render.NewTexture(def.Btn))

			world.Collider().Component().Set(def.Btn, collider.NewCollider(world.Definitions().SquareCollider))
			world.Inputs().KeepSelected().Set(def.Btn, inputs.KeepSelectedComponent{})

			world.Text().Align().Set(def.Btn, text.TextAlignComponent{Vertical: .5, Horizontal: .5})
			world.Text().FontSize().Set(def.Btn, text.FontSizeComponent{FontSize: 24})
		}
		{
			btnAsset, err := assets.GetAsset[render.TextureAsset](world.Assets(), def.Btn)
			if err != nil {
				world.Logger().Warn(err)
			}
			btnAspectRatio := btnAsset.AspectRatio()
			world.Groups().Inherit().Set(def.Text, groups.InheritGroupsComponent{})
			world.Groups().Component().Set(def.Text, groups.EmptyGroups())

			world.Transform().Size().Set(def.Text, transform.NewSize(150, 50, 1))
			world.Transform().AspectRatio().Set(def.Text, transform.NewAspectRatio(float32(btnAspectRatio.Dx()), float32(btnAspectRatio.Dy()), 0, transform.PrimaryAxisX))
			world.Transform().Parent().Set(def.Text, transform.NewParent(transform.RelativePos))

			world.Render().Mesh().Set(def.Text, render.NewMesh(world.Definitions().SquareMesh))
			world.Render().Texture().Set(def.Text, render.NewTexture(def.Btn))
			world.Render().Color().Set(def.Text, render.NewColor(mgl32.Vec4{0, 0, 0, 0}))

			world.Collider().Component().Set(def.Text, collider.NewCollider(world.Definitions().SquareCollider))
			world.Inputs().KeepSelected().Set(def.Text, inputs.KeepSelectedComponent{})

			world.Text().Align().Set(def.Text, text.TextAlignComponent{Vertical: .5, Horizontal: .5})
			world.Text().FontSize().Set(def.Text, text.FontSizeComponent{FontSize: 24})
		}
		return def
	})
	ioc.Register(b, func(c ioc.Dic) definitions.Transitions {
		world := ioc.GetServices[gamescenes.GameWorld](c)
		def, err := registry.GetRegistry[definitions.Transitions](world.Registry())
		world.Logger().Warn(err)
		return def
	})

	ioc.Register(b, func(c ioc.Dic) definitions.Definitions {
		world := ioc.GetServices[gamescenes.GameWorld](c)
		def, err := registry.GetRegistry[definitions.Definitions](world.Registry())
		world.Logger().Warn(err)
		world.Logger().Warn(c.InjectServices(&def))
		return def
	})

	//
	//
	//

	// animations

	transitions := map[string]func(t transition.Progress) transition.Progress{
		"linear": func(t transition.Progress) transition.Progress {
			return t
		},
		"my easing": func(t transition.Progress) transition.Progress {
			const n1 = 7.5625
			const d1 = 2.75

			if t < 1/d1 { // First segment of the bounce (rising curve)
				return n1 * t * t
			} else if t < 2/d1 { // Second segment (peak of the first bounce)
				t -= 1.5 / d1
				return n1*t*t + 0.75
			} else if t < 2.5/d1 { // Third segment (peak of the second, smaller bounce)
				t -= 2.25 / d1
				return n1*t*t + 0.9375
			} else { // Final segment (settling)
				t -= 2.625 / d1
				return n1*t*t + 0.984375
			}
		},
		"ease out elastic": func(t transition.Progress) transition.Progress {
			const c1 float64 = 10
			const c2 float64 = .75
			const c3 float64 = (2 * math.Pi) / 3
			if t == 0 {
				return 0
			}
			if t == 1 {
				return 1
			}
			x := float64(t)
			x = math.Pow(2, -c1*x)*
				math.Sin((x*c1-c2)*c3) +
				1
			return transition.Progress(x)
		},
	}

	ioc.Wrap(b, func(c ioc.Dic, b registry.Service) {
		b.Register("transition", func(entity ecs.EntityID, structTagValue string) {
			transitionService := ioc.Get[transition.Service](c)
			easing, ok := transitions[structTagValue]
			if !ok {
				easing = func(t transition.Progress) transition.Progress { return t }
			}
			transitionService.EasingFunction().Set(entity, transition.NewEasingFunction(easing))
		})
	})
})
