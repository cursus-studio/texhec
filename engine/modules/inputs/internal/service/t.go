package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/inputs"
	"engine/modules/inputs/internal"
	"engine/modules/inputs/internal/mouse"
	"engine/modules/inputs/internal/systems"
	"engine/modules/loop"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type service struct {
	engine.EngineWorld `inject:""`
	c                  ioc.Dic

	textInput ecs.ComponentArray[inputs.TextInputComponent]

	hovered ecs.ComponentArray[inputs.HoveredComponent]
	dragged ecs.ComponentArray[inputs.DraggedComponent]
	stacked ecs.ComponentArray[inputs.StackedComponent]

	keepSelected ecs.ComponentArray[inputs.KeepSelectedComponent]

	leftClick       ecs.ComponentArray[inputs.LeftClickComponent]
	doubleLeftClick ecs.ComponentArray[inputs.DoubleLeftClickComponent]

	rightClick       ecs.ComponentArray[inputs.RightClickComponent]
	doubleRightClick ecs.ComponentArray[inputs.DoubleRightClickComponent]

	mouseEnter ecs.ComponentArray[inputs.MouseEnterComponent]
	mouseLeave ecs.ComponentArray[inputs.MouseLeaveComponent]

	mouseHover ecs.ComponentArray[inputs.HoverComponent]
	mouseDrag  ecs.ComponentArray[inputs.DragComponent]

	stack ecs.ComponentArray[inputs.StackComponent]

	stackData *[]inputs.Target
}

func NewService(c ioc.Dic) inputs.Service {
	s := ioc.GetServices[*service](c)
	s.c = c
	s.textInput = ecs.GetComponentArray[inputs.TextInputComponent](s.World())

	s.hovered = ecs.GetComponentArray[inputs.HoveredComponent](s.World())
	s.dragged = ecs.GetComponentArray[inputs.DraggedComponent](s.World())
	s.stacked = ecs.GetComponentArray[inputs.StackedComponent](s.World())

	s.keepSelected = ecs.GetComponentArray[inputs.KeepSelectedComponent](s.World())

	s.leftClick = ecs.GetComponentArray[inputs.LeftClickComponent](s.World())
	s.doubleLeftClick = ecs.GetComponentArray[inputs.DoubleLeftClickComponent](s.World())

	s.rightClick = ecs.GetComponentArray[inputs.RightClickComponent](s.World())
	s.doubleRightClick = ecs.GetComponentArray[inputs.DoubleRightClickComponent](s.World())

	s.mouseEnter = ecs.GetComponentArray[inputs.MouseEnterComponent](s.World())
	s.mouseLeave = ecs.GetComponentArray[inputs.MouseLeaveComponent](s.World())

	s.mouseHover = ecs.GetComponentArray[inputs.HoverComponent](s.World())
	s.mouseDrag = ecs.GetComponentArray[inputs.DragComponent](s.World())

	s.stack = ecs.GetComponentArray[inputs.StackComponent](s.World())

	ecs.GetComponentArray[inputs.StackComponent](s.World())

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

func (s *service) TextInput() ecs.ComponentArray[inputs.TextInputComponent] {
	return s.textInput
}

func (s *service) Hovered() ecs.ComponentArray[inputs.HoveredComponent] { return s.hovered }
func (s *service) Dragged() ecs.ComponentArray[inputs.DraggedComponent] { return s.dragged }
func (s *service) Stacked() ecs.ComponentArray[inputs.StackedComponent] { return s.stacked }

func (s *service) KeepSelected() ecs.ComponentArray[inputs.KeepSelectedComponent] {
	return s.keepSelected
}

func (s *service) LeftClick() ecs.ComponentArray[inputs.LeftClickComponent] { return s.leftClick }
func (s *service) DoubleLeftClick() ecs.ComponentArray[inputs.DoubleLeftClickComponent] {
	return s.doubleLeftClick
}

func (s *service) RightClick() ecs.ComponentArray[inputs.RightClickComponent] { return s.rightClick }
func (s *service) DoubleRightClick() ecs.ComponentArray[inputs.DoubleRightClickComponent] {
	return s.doubleRightClick
}

func (s *service) MouseEnter() ecs.ComponentArray[inputs.MouseEnterComponent] { return s.mouseEnter }
func (s *service) MouseLeave() ecs.ComponentArray[inputs.MouseLeaveComponent] { return s.mouseLeave }

func (s *service) Hover() ecs.ComponentArray[inputs.HoverComponent] { return s.mouseHover }
func (s *service) Drag() ecs.ComponentArray[inputs.DragComponent]   { return s.mouseDrag }

func (s *service) Stack() ecs.ComponentArray[inputs.StackComponent] { return s.stack }

func (s *service) StackedData() []inputs.Target {
	stackCopy := make([]inputs.Target, len(*s.stackData))
	copy(stackCopy, *s.stackData)
	return stackCopy
}
