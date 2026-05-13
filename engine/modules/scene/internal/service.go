package internal

import (
	"engine"
	"engine/modules/scene"
	"engine/services/ecs"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

// type Service interface {
// 	SetScene(id ID, loader func(sceneParent ecs.EntityID))
// }

type SceneComp struct{}

type service struct {
	engine.EngineWorld `inject:""`
	scenes             map[scene.ID]scene.Scene
	SceneArr           ecs.ComponentsArray[SceneComp]
}

func NewService(c ioc.Dic) scene.Service {
	s := ioc.GetServices[*service](c)
	s.scenes = make(map[scene.ID]scene.Scene)
	s.SceneArr = ecs.GetComponentsArray[SceneComp](s.World())
	entity := s.World().NewEntity()
	s.SceneArr.Set(entity, SceneComp{})

	events.Listen(s.EventsBuilder(), s.ChangeScene)
	return s
}

func (s *service) ChangeScene(event scene.ChangeSceneEvent) {
	for _, entity := range s.SceneArr.GetEntities() {
		s.World().RemoveEntity(entity)
	}
	sceneEntity := s.World().NewEntity()
	s.SceneArr.Set(sceneEntity, SceneComp{})

	scene, ok := s.scenes[event.ID]
	if !ok {
		s.Logger().Log(fmt.Errorf("scene with id %v doesn't exist", event.ID))
		return
	}
	scene(sceneEntity)
}

func (s *service) Scene() ecs.EntityID {
	return s.SceneArr.GetEntities()[0]
}

func (s *service) SetScene(id scene.ID, scene scene.Scene) {
	s.scenes[id] = scene
}
