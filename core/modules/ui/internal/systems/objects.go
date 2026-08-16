package systems

import (
	"core/game"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/render"
	"engine/modules/transform"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type objectsSystem struct {
	game.GameWorld `inject:""`
}

func NewObjectsSystem(c ioc.Dic) ecs.SystemRegister {
	s := ioc.GetServices[*objectsSystem](c)
	return s
}

func (s *objectsSystem) Register() error {
	events.Listen(s.EventsBuilder(), s.OnFrame)
	return nil
}

func (s *objectsSystem) AddHealthBar(entity ecs.EntityID) {
	health, ok := s.Attack().Health().Get(entity)
	if !ok {
		return
	}
	fullHealth, ok := s.Attack().FullHealth(entity)
	if !ok || health == fullHealth {
		return
	}
	healthBarSize := float32(health.Health) / float32(fullHealth.Health)
	blankBar := s.World().NewEntity()
	s.Hierarchy().SetParent(blankBar, entity)
	s.Transform().Inherit().Set(blankBar, transform.NewInherit(transform.RelativePos|transform.RelativeSizeX))
	s.Transform().PivotPoint().Set(blankBar, transform.NewPivotPoint(.5, 0, .5))
	s.Transform().ParentPivotPoint().Set(blankBar, transform.NewParentPivotPoint(.5, 0, .5))
	s.Transform().Pos().Set(blankBar, transform.NewPos(0, 0, 1))
	s.Transform().Size().Set(blankBar, transform.NewSize(1, 10, 1))

	s.Render().Mesh().Set(blankBar, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(blankBar, render.NewTexture(s.Definitions().Assets().Blank))
	s.Render().Color().Set(blankBar, render.NewColor(mgl32.Vec4{0, 0, 0, 1}))

	healthBar := s.World().NewEntity()
	s.Hierarchy().SetParent(healthBar, entity)
	s.Transform().Inherit().Set(healthBar, transform.NewInherit(transform.RelativePos|transform.RelativeSizeX))
	s.Transform().PivotPoint().Set(healthBar, transform.NewPivotPoint(.5, 0, .5))
	s.Transform().ParentPivotPoint().Set(healthBar, transform.NewParentPivotPoint(.5, 0, .5))
	s.Transform().Pos().Set(healthBar, transform.NewPos(0, 0, 2))
	s.Transform().Size().Set(healthBar, transform.NewSize(healthBarSize, 10, 1))

	s.Render().Mesh().Set(healthBar, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(healthBar, render.NewTexture(s.Definitions().Assets().Blank))
	s.Render().Color().Set(healthBar, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
}

// resets objects preview objects
func (s *objectsSystem) OnFrame(loop.FrameEvent) {
	for _, entity := range s.Tile().Pos().GetEntities() {
		for _, child := range s.Hierarchy().Children(entity).GetIndices() {
			s.World().RemoveEntity(child)
		}
		s.AddHealthBar(entity)
	}
}
