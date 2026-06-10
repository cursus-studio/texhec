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
	"github.com/veandco/go-sdl2/sdl"
)

type menuComponent struct {
	Visible bool
}
type childrenComponent struct{}

type service struct {
	game.GameWorld `inject:""`
	systems        []ecs.SystemRegister

	animationDuration time.Duration

	uiCameraArray           ecs.ComponentsArray[ui.UiCameraComponent]
	animatedBackgroundArray ecs.ComponentsArray[ui.AnimatedBackgroundComponent]
	cursorCameraArray       ecs.ComponentsArray[ui.CursorCameraComponent]

	objects ui.SelectionGroup[ui.ObjectComponent]
	actions ui.SelectionGroup[ui.ActionComponent]

	menuArray            ecs.ComponentsArray[menuComponent]
	childrenWrapperArray ecs.ComponentsArray[childrenComponent]
}

func NewService(
	c ioc.Dic,
	systems []ecs.SystemRegister,
	animationDuration time.Duration,
) ui.Service {
	s := ioc.GetServices[*service](c)
	s.systems = systems
	s.animationDuration = animationDuration

	s.uiCameraArray = ecs.GetComponentsArray[ui.UiCameraComponent](s.World())
	s.animatedBackgroundArray = ecs.GetComponentsArray[ui.AnimatedBackgroundComponent](s.World())
	s.cursorCameraArray = ecs.GetComponentsArray[ui.CursorCameraComponent](s.World())

	s.objects = NewSelectionGroup[ui.ObjectComponent](c, s)
	s.actions = NewSelectionGroup[ui.ActionComponent](c, s)

	s.menuArray = ecs.GetComponentsArray[menuComponent](s.World())
	s.childrenWrapperArray = ecs.GetComponentsArray[childrenComponent](s.World())

	s.systems = append(s.systems, ecs.NewSystemRegister(func() error {
		events.Listen(s.EventsBuilder(), func(e sdl.MouseButtonEvent) {
			if e.Button != sdl.BUTTON_RIGHT || e.State != sdl.RELEASED {
				return
			}
			events.Emit(s.Events(), ui.NewUnselect[ui.ObjectComponent]())
		})

		events.Listen(s.EventsBuilder(), func(ui.UnselectEvent[ui.ObjectComponent]) {
			events.Emit(s.Events(), ui.NewUnselect[ui.ActionComponent]())
			s.World().RemoveEntity(s.Interactions().FeatureEntity())
			s.HideMenu()
		})
		return nil
	}))

	s.EnsureExists()

	return s
}

func (s *service) Register() error {
	errs := ecs.RegisterSystems(s.systems...)
	if len(errs) != 0 {
		return errs[0]
	}
	return nil
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

func (s *service) Objects() ui.SelectionGroup[ui.ObjectComponent] { return s.objects }
func (s *service) Actions() ui.SelectionGroup[ui.ActionComponent] { return s.actions }

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
