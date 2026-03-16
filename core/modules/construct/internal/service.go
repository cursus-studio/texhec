package internal

import (
	"core/modules/construct"
	"core/modules/definitions"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.World `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	GameAssets   definitions.Definitions `inject:"1"`
	Ui           ui.Service              `inject:"1"`

	constructs      ecs.ComponentsArray[construct.ConstructComponent]
	constructCoords ecs.ComponentsArray[construct.CoordsComponent]
}

func NewService(c ioc.Dic) construct.Service {
	s := ioc.GetServices[*service](c)

	s.constructs = ecs.GetComponentsArray[construct.ConstructComponent](s)
	s.constructCoords = ecs.GetComponentsArray[construct.CoordsComponent](s)

	return s
}

func (s *service) Construct() ecs.ComponentsArray[construct.ConstructComponent] {
	return s.constructs
}
func (s *service) Coords() ecs.ComponentsArray[construct.CoordsComponent] {
	return s.constructCoords
}
