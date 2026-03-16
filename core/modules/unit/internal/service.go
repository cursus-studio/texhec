package internal

import (
	"core/modules/definitions"
	"core/modules/tile"
	"core/modules/ui"
	"core/modules/unit"
	"engine"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.World `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	GameAssets   definitions.Definitions `inject:"1"`
	Ui           ui.Service              `inject:"1"`

	units     ecs.ComponentsArray[unit.UnitComponent]
	coords    ecs.ComponentsArray[unit.CoordsComponent]
	rotations ecs.ComponentsArray[unit.RotationComponent]
}

func NewService(c ioc.Dic) unit.Service {
	s := ioc.GetServices[*service](c)

	s.units = ecs.GetComponentsArray[unit.UnitComponent](s.World)
	s.coords = ecs.GetComponentsArray[unit.CoordsComponent](s.World)
	s.rotations = ecs.GetComponentsArray[unit.RotationComponent](s.World)

	return s
}

func (s *service) Unit() ecs.ComponentsArray[unit.UnitComponent] {
	return s.units
}
func (s *service) Coords() ecs.ComponentsArray[unit.CoordsComponent] {
	return s.coords
}
func (s *service) Rotation() ecs.ComponentsArray[unit.RotationComponent] {
	return s.rotations
}
