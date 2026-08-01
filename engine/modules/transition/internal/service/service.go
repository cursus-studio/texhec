package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/transition"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	easing             ecs.ComponentArray[transition.EasingComponent]
	easingFunction     ecs.ComponentArray[transition.EasingFunctionComponent]
	register           ecs.SystemRegister
}

func NewService(c ioc.Dic, register ecs.SystemRegister) transition.Service {
	s := ioc.GetServices[*service](c)
	s.easing = ecs.GetComponentArray[transition.EasingComponent](s.World())
	s.easingFunction = ecs.GetComponentArray[transition.EasingFunctionComponent](s.World())
	s.register = register

	return s
}

func (s *service) Register() error {
	return s.register.Register()
}
func (s *service) Easing() ecs.ComponentArray[transition.EasingComponent] {
	return s.easing
}
func (s *service) EasingFunction() ecs.ComponentArray[transition.EasingFunctionComponent] {
	return s.easingFunction
}
