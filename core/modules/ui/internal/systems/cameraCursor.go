package systems

import (
	gamescenes "core/scenes"
	"engine/modules/groups"
	"engine/modules/loop"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/services/ecs"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type CursorComponent struct{}

type cursorSystem struct {
	gamescenes.GameWorld `inject:""`

	CursorComponent ecs.ComponentsArray[CursorComponent]
}

func NewCursorSystem(c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*cursorSystem](c)

		s.CursorComponent = ecs.GetComponentsArray[CursorComponent](s.World())

		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (s *cursorSystem) Listen(loop.FrameEvent) {
	mousePos := s.Window().GetMousePos()

	cameras := s.Ui().CursorCamera().GetEntities()
	if len(cameras) > 1 {
		s.Logger().Warn(fmt.Errorf("expected at most one cursor camera component"))
	}
	if len(cameras) != 1 {
		for _, cursor := range s.CursorComponent.GetEntities() {
			s.World().RemoveEntity(cursor)
		}
		_, _ = sdl.ShowCursor(sdl.ENABLE)
		return
	}
	camera := cameras[0]
	// hide cursor
	var cursor ecs.EntityID
	if entities := s.CursorComponent.GetEntities(); len(entities) == 1 {
		cursor = entities[0]
	} else if len(entities) == 0 {
		for _, cursor := range s.CursorComponent.GetEntities() {
			s.World().RemoveEntity(cursor)
		}
		cursor = s.World().NewEntity()
		s.CursorComponent.Set(cursor, CursorComponent{})
	}

	ray := s.Camera().ShootRay(camera, mousePos)
	pos := transform.NewPos(ray.Pos.Add(ray.Direction).Elem())

	s.Hierarchy().SetParent(cursor, camera)
	s.Transform().Parent().Set(cursor, transform.NewParent(transform.Absolute))
	s.Transform().Pos().Set(cursor, pos)
	s.Render().Mesh().Set(cursor, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(cursor, render.NewTexture(s.Definitions().Hud().Cursor))
	s.Groups().Inherit().Set(cursor, groups.InheritGroupsComponent{})
	s.Transform().Size().Set(cursor, transform.NewSize(50, 50, 1))
	_, _ = sdl.ShowCursor(sdl.DISABLE)
}
