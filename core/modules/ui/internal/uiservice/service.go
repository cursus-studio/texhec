package uiservice

import (
	"core/game"
	"core/modules/ui"
	"engine/modules/ecs"
	"engine/modules/transform"
	"engine/modules/transition"
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

	uiCameraArray           ecs.ComponentArray[ui.UiCameraComponent]
	animatedBackgroundArray ecs.ComponentArray[ui.AnimatedBackgroundComponent]
	cursorCameraArray       ecs.ComponentArray[ui.CursorCameraComponent]

	menuArray            ecs.ComponentArray[menuComponent]
	childrenWrapperArray ecs.ComponentArray[childrenComponent]
}

func NewService(
	c ioc.Dic,
	systems []ecs.SystemRegister,
	animationDuration time.Duration,
) ui.Service {
	s := ioc.GetServices[*service](c)
	s.systems = systems
	s.animationDuration = animationDuration

	s.uiCameraArray = ecs.GetComponentArray[ui.UiCameraComponent](s.World())
	s.animatedBackgroundArray = ecs.GetComponentArray[ui.AnimatedBackgroundComponent](s.World())
	s.cursorCameraArray = ecs.GetComponentArray[ui.CursorCameraComponent](s.World())

	s.menuArray = ecs.GetComponentArray[menuComponent](s.World())
	s.childrenWrapperArray = ecs.GetComponentArray[childrenComponent](s.World())

	s.systems = append(s.systems, ecs.NewSystemRegister(func() error {
		events.Listen(s.EventsBuilder(), func(e sdl.MouseButtonEvent) {
			if e.Button != sdl.BUTTON_RIGHT || e.State != sdl.RELEASED {
				return
			}
			removeEntityEvent := ecs.NewRemoveEntityEvent(s.Interactions().FeatureEntity())
			events.Emit(s.Events(), removeEntityEvent)
		})

		s.Interactions().Instance().OnRemove(func(ecs.EntityID) {
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

func (s *service) UiCamera() ecs.ComponentArray[ui.UiCameraComponent] { return s.uiCameraArray }
func (s *service) AnimatedBackground() ecs.ComponentArray[ui.AnimatedBackgroundComponent] {
	return s.animatedBackgroundArray
}
func (s *service) CursorCamera() ecs.ComponentArray[ui.CursorCameraComponent] {
	return s.cursorCameraArray
}

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
