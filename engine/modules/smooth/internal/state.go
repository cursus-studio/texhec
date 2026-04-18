package internal

import (
	"engine"
	"engine/modules/record"
	"engine/modules/transition"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type Service[Component transition.LerpConstraint[Component]] struct {
	engine.EngineWorld `inject:""`
	recordingID        record.RecordingID
	config             record.Config

	componentArray ecs.ComponentsArray[Component]
	lerpArray      ecs.ComponentsArray[transition.TransitionComponent[Component]]
}

func NewService[Component transition.LerpConstraint[Component]](c ioc.Dic) *Service[Component] {
	config := record.NewConfig()
	record.AddToConfig[Component](config)

	s := ioc.GetServices[*Service[Component]](c)

	s.recordingID = 0
	s.config = config
	s.componentArray = ecs.GetComponentsArray[Component](s.World())
	s.lerpArray = ecs.GetComponentsArray[transition.TransitionComponent[Component]](s.World())
	return s
}

//

type system[Component transition.LerpConstraint[Component]] struct {
	engine.EngineWorld `inject:""`
	Service            *Service[Component] `inject:""`
}
