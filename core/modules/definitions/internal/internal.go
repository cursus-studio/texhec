package internal

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/deploy"
	"engine/modules/assets"
	"engine/modules/collider"
	"engine/modules/entityregistry"
	"engine/modules/graphics"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/services/ecs"
	"image"
	"image/color"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	World       game.GameWorld `inject:""`
	assets      *definitions.Assets
	hud         *definitions.Hud
	tiles       *definitions.Tiles
	objects     *definitions.Objects
	transitions *definitions.Transitions
}

func NewService(c ioc.Dic) definitions.Service {
	return ioc.GetServices[*service](c)
}

func (s *service) Load() {
	s.Assets()
	s.Hud()
	s.Tiles()
	s.Objects()
	s.Transitions()
}

func (s *service) Assets() definitions.Assets {
	if s.assets != nil {
		return *s.assets
	}
	def, err := entityregistry.GetRegistry[definitions.Assets](s.World.EntityRegistry())
	s.World.Logger().Warn(err)
	{
		img := image.NewRGBA(image.Rect(0, 0, 1, 1))
		white := color.RGBA{255, 255, 255, 255}
		img.Set(0, 0, white)
		asset, err := render.NewTextureAsset(img)
		s.World.Logger().Warn(err)
		def.Blank = s.World.World().NewEntity()
		s.World.Assets().Cache().Set(def.Blank, assets.NewCache(asset))
	}
	{
		vertices := []render.Vertex{
			{Pos: [3]float32{1, 1, 1}, TexturePos: [2]float32{1, 1}},
			{Pos: [3]float32{1, -1, 1}, TexturePos: [2]float32{1, 0}},
			{Pos: [3]float32{-1, -1, 1}, TexturePos: [2]float32{0, 0}},
			{Pos: [3]float32{-1, 1, 1}, TexturePos: [2]float32{0, 1}},
		}
		indices := []graphics.Index{
			0, 1, 2,
			0, 2, 3,
		}
		asset := render.NewMeshAsset(vertices, indices)
		def.SquareMesh = s.World.World().NewEntity()
		s.World.Assets().Cache().Set(def.SquareMesh, assets.NewCache(asset))
	}
	{
		asset := collider.NewColliderAsset(
			[]collider.AABB{collider.NewAABB(mgl32.Vec3{-1, -1}, mgl32.Vec3{1, 1})},
			[]collider.Range{collider.NewRange(collider.Leaf, 0, 2)},
			[]collider.Polygon{
				collider.NewPolygon(mgl32.Vec3{-1, -1, 0}, mgl32.Vec3{+1, -1, 0}, [3]float32{-1, +1, 0}),
				collider.NewPolygon(mgl32.Vec3{+1, +1, 0}, mgl32.Vec3{+1, -1, 0}, [3]float32{-1, +1, 0}),
			})
		def.SquareCollider = s.World.World().NewEntity()
		s.World.Assets().Cache().Set(def.SquareCollider, assets.NewCache(asset))
	}
	s.assets = &def
	return def
}
func (s *service) Hud() definitions.Hud {
	if s.hud != nil {
		return *s.hud
	}
	def, err := entityregistry.GetRegistry[definitions.Hud](s.World.EntityRegistry())
	s.World.Logger().Warn(err)
	{
		btnAsset, err := assets.GetAsset[render.TextureAsset](s.World.Assets(), def.Btn)
		if err != nil {
			s.World.Logger().Warn(err)
		}
		btnAspectRatio := btnAsset.AspectRatio()
		s.World.Groups().Inherit().Set(def.Btn, groups.InheritGroupsComponent{})
		s.World.Groups().Component().Set(def.Btn, groups.EmptyGroups())

		s.World.Transform().AspectRatio().Set(def.Btn, transform.NewAspectRatio(float32(btnAspectRatio.Dx()), float32(btnAspectRatio.Dy()), 0, transform.PrimaryAxisX))
		s.World.Transform().Parent().Set(def.Btn, transform.NewParent(transform.RelativePos|transform.RelativeSizeX))
		s.World.Transform().MaxSize().Set(def.Btn, transform.NewMaxSize(0, 50, 0))
		s.World.Transform().Size().Set(def.Btn, transform.NewSize(1, 50, 1))

		s.World.Render().Mesh().Set(def.Btn, render.NewMesh(s.World.Definitions().Assets().SquareMesh))
		s.World.Render().Texture().Set(def.Btn, render.NewTexture(def.Btn))

		s.World.Collider().Component().Set(def.Btn, collider.NewCollider(s.World.Definitions().Assets().SquareCollider))
		s.World.Inputs().KeepSelected().Set(def.Btn, inputs.KeepSelectedComponent{})

		s.World.Text().Align().Set(def.Btn, text.NewAlign(.5, .5))
		s.World.Text().FontSize().Set(def.Btn, text.NewFontSize(24))
	}
	{
		btnAsset, err := assets.GetAsset[render.TextureAsset](s.World.Assets(), def.Btn)
		if err != nil {
			s.World.Logger().Warn(err)
		}
		btnAspectRatio := btnAsset.AspectRatio()
		s.World.Groups().Inherit().Set(def.Text, groups.InheritGroupsComponent{})
		s.World.Groups().Component().Set(def.Text, groups.EmptyGroups())

		s.World.Transform().Size().Set(def.Text, transform.NewSize(150, 50, 1))
		s.World.Transform().AspectRatio().Set(def.Text, transform.NewAspectRatio(float32(btnAspectRatio.Dx()), float32(btnAspectRatio.Dy()), 0, transform.PrimaryAxisX))
		s.World.Transform().Parent().Set(def.Text, transform.NewParent(transform.RelativePos))

		s.World.Render().Mesh().Set(def.Text, render.NewMesh(s.World.Definitions().Assets().SquareMesh))
		s.World.Render().Texture().Set(def.Text, render.NewTexture(def.Btn))
		s.World.Render().Color().Set(def.Text, render.NewColor(mgl32.Vec4{0, 0, 0, 0}))

		s.World.Collider().Component().Set(def.Text, collider.NewCollider(s.World.Definitions().Assets().SquareCollider))
		s.World.Inputs().KeepSelected().Set(def.Text, inputs.KeepSelectedComponent{})

		s.World.Text().Align().Set(def.Text, text.NewAlign(.5, .5))
		s.World.Text().FontSize().Set(def.Text, text.NewFontSize(24))
	}
	s.hud = &def
	return def
}
func (s *service) Tiles() definitions.Tiles {
	if s.tiles != nil {
		return *s.tiles
	}
	def, err := entityregistry.GetRegistry[definitions.Tiles](s.World.EntityRegistry())
	s.World.Logger().Warn(err)
	s.tiles = &def
	return def
}
func (s *service) Objects() definitions.Objects {
	if s.objects != nil {
		return *s.objects
	}
	def, err := entityregistry.GetRegistry[definitions.Objects](s.World.EntityRegistry())
	s.World.Logger().Warn(err)

	s.World.Deploy().Component().Set(def.Tank, deploy.NewDeploy(def.Tank, def.Farm))
	s.World.Deploy().Component().Set(def.Farm, deploy.NewDeploy(def.Tank, def.Farm, def.HouseT1, def.HouseT2, def.HouseT3, def.HouseT4))
	s.objects = &def
	return def
}
func (s *service) Transitions() definitions.Transitions {
	if s.transitions != nil {
		return *s.transitions
	}
	def, err := entityregistry.GetRegistry[definitions.Transitions](s.World.EntityRegistry())
	s.World.Logger().Warn(err)
	transitionEntity := func(easing func(t transition.Progress) transition.Progress) ecs.EntityID {
		entity := s.World.World().NewEntity()
		s.World.Transition().EasingFunction().Set(entity, transition.NewEasingFunction(easing))
		return entity
	}
	def.Linear = transitionEntity(func(t transition.Progress) transition.Progress {
		return t
	})
	def.MyEasing = transitionEntity(func(t transition.Progress) transition.Progress {
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
	})
	def.EaseOutElastic = transitionEntity(func(t transition.Progress) transition.Progress {
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
	})
	s.transitions = &def
	return def
}
