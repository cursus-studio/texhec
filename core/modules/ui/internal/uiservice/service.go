package uiservice

import (
	"core/game"
	"core/modules/ui"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/services/ecs"
	"time"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type menuComponent struct {
	Visible bool
}
type childrenComponent struct{}

type service struct {
	game.GameWorld `inject:""`

	animationDuration time.Duration

	uiCameraArray           ecs.ComponentsArray[ui.UiCameraComponent]
	animatedBackgroundArray ecs.ComponentsArray[ui.AnimatedBackgroundComponent]
	cursorCameraArray       ecs.ComponentsArray[ui.CursorCameraComponent]

	objects ui.SelectionGroup
	actions ui.SelectionGroup

	menuArray            ecs.ComponentsArray[menuComponent]
	childrenWrapperArray ecs.ComponentsArray[childrenComponent]
}

func NewService(
	c ioc.Dic,
	animationDuration time.Duration,
) *service {
	s := ioc.GetServices[*service](c)
	s.animationDuration = animationDuration

	s.uiCameraArray = ecs.GetComponentsArray[ui.UiCameraComponent](s.World())
	s.animatedBackgroundArray = ecs.GetComponentsArray[ui.AnimatedBackgroundComponent](s.World())
	s.cursorCameraArray = ecs.GetComponentsArray[ui.CursorCameraComponent](s.World())

	type ObjectSelectionComponent struct{}
	s.objects = NewSelectionGroup[ObjectSelectionComponent](c)
	type ActionSelectionComponent struct{}
	s.actions = NewSelectionGroup[ActionSelectionComponent](c)

	s.menuArray = ecs.GetComponentsArray[menuComponent](s.World())
	s.childrenWrapperArray = ecs.GetComponentsArray[childrenComponent](s.World())

	s.Objects().OnUnselect(s.Actions().Unselect)
	s.Objects().OnUnselect(s.HideMenu)

	s.EnsureExists()

	return s
}

func (s *service) ResetChildWrapper() {
	s.EnsureExists()

	for _, childWrapper := range s.childrenWrapperArray.GetEntities() {
		for _, child := range s.Hierarchy().Children(childWrapper).GetIndices() {
			s.World().RemoveEntity(child)
		}
	}
}

func (s *service) UiCamera() ecs.ComponentsArray[ui.UiCameraComponent] { return s.uiCameraArray }
func (s *service) AnimatedBackground() ecs.ComponentsArray[ui.AnimatedBackgroundComponent] {
	return s.animatedBackgroundArray
}
func (s *service) CursorCamera() ecs.ComponentsArray[ui.CursorCameraComponent] {
	return s.cursorCameraArray
}

func (s *service) Objects() ui.SelectionGroup { return s.objects }
func (s *service) Actions() ui.SelectionGroup { return s.actions }

func (s *service) ShowMenu() []ecs.EntityID {
	s.ResetChildWrapper()

	for _, menu := range s.menuArray.GetEntities() {
		if component, _ := s.menuArray.Get(menu); !component.Visible {
			s.menuArray.Set(menu, menuComponent{true})
			events.Emit(s.Events(), transition.NewTransitionEvent(
				menu,
				transform.NewPivotPoint(0, 1, .5),
				transform.NewPivotPoint(1, 1, .5),
				s.animationDuration,
			))
		}
	}
	return s.childrenWrapperArray.GetEntities()
}

func (s *service) HideMenu() {
	s.EnsureExists()

	for _, menu := range s.menuArray.GetEntities() {
		if component, _ := s.menuArray.Get(menu); component.Visible {
			s.menuArray.Set(menu, menuComponent{false})
			events.Emit(s.Events(), transition.NewTransitionEvent(
				menu,
				transform.NewPivotPoint(1, 1, .5),
				transform.NewPivotPoint(0, 1, .5),
				s.animationDuration,
			))
		}
	}
}
