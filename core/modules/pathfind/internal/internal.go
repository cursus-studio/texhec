package internal

import (
	"core/modules/pathfind"
	"core/modules/tile"
	"engine"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.World `inject:"1"`
	Tile         tile.Service `inject:"1"`

	target ecs.ComponentsArray[pathfind.TargetComponent]
}

func NewService(c ioc.Dic) pathfind.Service {
	s := ioc.GetServices[*service](c)
	s.target = ecs.GetComponentsArray[pathfind.TargetComponent](s)
	return s
}

func (s *service) Target() ecs.ComponentsArray[pathfind.TargetComponent] { return s.target }

func (s *service) Select(event pathfind.SelectEvent)           { events.Emit(s.Events, event) }
func (s *service) PreviewPath(event pathfind.PreviewPathEvent) { events.Emit(s.Events, event) }
func (s *service) FindPath(event pathfind.FindPathEvent)       { events.Emit(s.Events, event) }
