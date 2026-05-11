package service

import (
	"engine"
	"engine/modules/inputs"
	"engine/modules/inputs/internal"
	"engine/modules/inputs/internal/mouse"
	"engine/modules/inputs/internal/systems"
	"engine/modules/loop"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type service struct {
	engine.EngineWorld `inject:""`
	c                  ioc.Dic

	hovered ecs.ComponentsArray[inputs.HoveredComponent]
	dragged ecs.ComponentsArray[inputs.DraggedComponent]
	stacked ecs.ComponentsArray[inputs.StackedComponent]

	keepSelected ecs.ComponentsArray[inputs.KeepSelectedComponent]

	leftClick       ecs.ComponentsArray[inputs.LeftClickComponent]
	doubleLeftClick ecs.ComponentsArray[inputs.DoubleLeftClickComponent]

	rightClick       ecs.ComponentsArray[inputs.RightClickComponent]
	doubleRightClick ecs.ComponentsArray[inputs.DoubleRightClickComponent]

	mouseEnter ecs.ComponentsArray[inputs.MouseEnterComponent]
	mouseLeave ecs.ComponentsArray[inputs.MouseLeaveComponent]

	mouseHover ecs.ComponentsArray[inputs.HoverComponent]
	mouseDrag  ecs.ComponentsArray[inputs.DragComponent]

	stack ecs.ComponentsArray[inputs.StackComponent]

	stackData *[]inputs.Target
}

func NewService(c ioc.Dic) inputs.Service {
	t := ioc.GetServices[*service](c)
	t.c = c
	t.hovered = ecs.GetComponentsArray[inputs.HoveredComponent](t.World())
	t.dragged = ecs.GetComponentsArray[inputs.DraggedComponent](t.World())
	t.stacked = ecs.GetComponentsArray[inputs.StackedComponent](t.World())

	t.keepSelected = ecs.GetComponentsArray[inputs.KeepSelectedComponent](t.World())

	t.leftClick = ecs.GetComponentsArray[inputs.LeftClickComponent](t.World())
	t.doubleLeftClick = ecs.GetComponentsArray[inputs.DoubleLeftClickComponent](t.World())

	t.rightClick = ecs.GetComponentsArray[inputs.RightClickComponent](t.World())
	t.doubleRightClick = ecs.GetComponentsArray[inputs.DoubleRightClickComponent](t.World())

	t.mouseEnter = ecs.GetComponentsArray[inputs.MouseEnterComponent](t.World())
	t.mouseLeave = ecs.GetComponentsArray[inputs.MouseLeaveComponent](t.World())

	t.mouseHover = ecs.GetComponentsArray[inputs.HoverComponent](t.World())
	t.mouseDrag = ecs.GetComponentsArray[inputs.DragComponent](t.World())

	t.stack = ecs.GetComponentsArray[inputs.StackComponent](t.World())

	ecs.GetComponentsArray[inputs.StackComponent](t.World())

	stack := []inputs.Target{}
	t.stackData = &stack
	return t
}

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), func(e internal.RayChangedTargetEvent) {
		*s.stackData = e.Targets
	})
	events.Listen(s.EventsBuilder(), func(loop.FrameEvent) {
		events.Emit(s.Events(), mouse.NewShootRayEvent())
	})
	events.Listen(s.EventsBuilder(), func(sdl.QuitEvent) {
		events.Emit(s.Events(), loop.NewStopEvent())
	})
	ecs.RegisterSystems(
		systems.NewInputsSystem(s.c),
		mouse.NewCameraRaySystem(s.c),
		mouse.NewHoverSystem(s.c),
		mouse.NewHoverEventsSystem(s.c),
		mouse.NewClickSystem(s.c),
	)
	return nil
}

func (s *service) Hovered() ecs.ComponentsArray[inputs.HoveredComponent] { return s.hovered }
func (s *service) Dragged() ecs.ComponentsArray[inputs.DraggedComponent] { return s.dragged }
func (s *service) Stacked() ecs.ComponentsArray[inputs.StackedComponent] { return s.stacked }

func (s *service) KeepSelected() ecs.ComponentsArray[inputs.KeepSelectedComponent] {
	return s.keepSelected
}

func (s *service) LeftClick() ecs.ComponentsArray[inputs.LeftClickComponent] { return s.leftClick }
func (s *service) DoubleLeftClick() ecs.ComponentsArray[inputs.DoubleLeftClickComponent] {
	return s.doubleLeftClick
}

func (s *service) RightClick() ecs.ComponentsArray[inputs.RightClickComponent] { return s.rightClick }
func (s *service) DoubleRightClick() ecs.ComponentsArray[inputs.DoubleRightClickComponent] {
	return s.doubleRightClick
}

func (s *service) MouseEnter() ecs.ComponentsArray[inputs.MouseEnterComponent] { return s.mouseEnter }
func (s *service) MouseLeave() ecs.ComponentsArray[inputs.MouseLeaveComponent] { return s.mouseLeave }

func (s *service) Hover() ecs.ComponentsArray[inputs.HoverComponent] { return s.mouseHover }
func (s *service) Drag() ecs.ComponentsArray[inputs.DragComponent]   { return s.mouseDrag }

func (s *service) Stack() ecs.ComponentsArray[inputs.StackComponent] { return s.stack }

func (s *service) StackedData() []inputs.Target {
	stackCopy := make([]inputs.Target, len(*s.stackData))
	copy(stackCopy, *s.stackData)
	return stackCopy
}
