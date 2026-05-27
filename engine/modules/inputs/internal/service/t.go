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

	textInput ecs.ComponentsArray[inputs.TextInputComponent]

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
	s := ioc.GetServices[*service](c)
	s.c = c
	s.textInput = ecs.GetComponentsArray[inputs.TextInputComponent](s.World())

	s.hovered = ecs.GetComponentsArray[inputs.HoveredComponent](s.World())
	s.dragged = ecs.GetComponentsArray[inputs.DraggedComponent](s.World())
	s.stacked = ecs.GetComponentsArray[inputs.StackedComponent](s.World())

	s.keepSelected = ecs.GetComponentsArray[inputs.KeepSelectedComponent](s.World())

	s.leftClick = ecs.GetComponentsArray[inputs.LeftClickComponent](s.World())
	s.doubleLeftClick = ecs.GetComponentsArray[inputs.DoubleLeftClickComponent](s.World())

	s.rightClick = ecs.GetComponentsArray[inputs.RightClickComponent](s.World())
	s.doubleRightClick = ecs.GetComponentsArray[inputs.DoubleRightClickComponent](s.World())

	s.mouseEnter = ecs.GetComponentsArray[inputs.MouseEnterComponent](s.World())
	s.mouseLeave = ecs.GetComponentsArray[inputs.MouseLeaveComponent](s.World())

	s.mouseHover = ecs.GetComponentsArray[inputs.HoverComponent](s.World())
	s.mouseDrag = ecs.GetComponentsArray[inputs.DragComponent](s.World())

	s.stack = ecs.GetComponentsArray[inputs.StackComponent](s.World())

	ecs.GetComponentsArray[inputs.StackComponent](s.World())

	stack := []inputs.Target{}
	s.stackData = &stack
	return s
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

func (s *service) TextInput() ecs.ComponentsArray[inputs.TextInputComponent] {
	return s.textInput
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
