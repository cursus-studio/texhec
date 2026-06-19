package transitionimpl

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/transition"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

// type sysT[Component transition.LerpConstraint[Component]] struct {
type sysT[Component any] struct {
	engine.EngineWorld `inject:""`

	dirtySet ecs.DirtySet

	transitionArray ecs.ComponentArray[transition.TransitionComponent[Component]]
	easingArray     ecs.ComponentArray[transition.EasingComponent]
	componentArray  ecs.ComponentArray[Component]
}

// func NewSysT[Component transition.LerpConstraint[Component]](c ioc.Dic) transition.System {
func NewSysT[Component any](c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*sysT[Component]](c)

		s.dirtySet = ecs.NewDirtySet()
		s.transitionArray = ecs.GetComponentArray[transition.TransitionComponent[Component]](s.World())
		s.easingArray = ecs.GetComponentArray[transition.EasingComponent](s.World())
		s.componentArray = ecs.GetComponentArray[Component](s.World())

		events.Listen(s.EventsBuilder(), s.ListenTransition)

		s.transitionArray.AddDirtySet(s.dirtySet)
		events.Listen(s.EventsBuilder(), s.ListenFrame)

		return nil
	})
}

func (s *sysT[Component]) ListenTransition(event transition.TransitionEvent[Component]) {
	s.transitionArray.Set(event.Entity, event.Component)
}

func (s *sysT[Component]) ListenFrame(event loop.FrameEvent) {
	ei := s.dirtySet.Get()

	for _, entity := range ei {
		transitionComponent, ok := s.transitionArray.Get(entity)
		if !ok {
			continue
		}

		transitionComponent.Progress = min(
			transitionComponent.Duration,
			transitionComponent.Progress+event.Delta,
		)
		progress := transition.Progress(transitionComponent.Progress) / transition.Progress(transitionComponent.Duration)

		easingComponent, ok := s.easingArray.Get(entity)
		if ok {
			if fn, ok := s.Transition().EasingFunction().Get(easingComponent.ID); ok {
				progress = fn.EasingFunction(progress)
			}
		}

		component := any(transitionComponent.From).(transition.LerpConstraint[Component]).
			Lerp(transitionComponent.To, float32(progress))

		s.transitionArray.Set(entity, transitionComponent)
		s.componentArray.Set(entity, component)
	}
}
