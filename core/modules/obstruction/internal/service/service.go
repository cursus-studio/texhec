package service

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/modules/record"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld         `inject:""`
	ObstructionGridService grid.Service[obstruction.Obstruction] `inject:""`

	obstruction ecs.ComponentsArray[obstruction.Component]
	deployed    ecs.ComponentsArray[obstruction.DeployedComponent]

	// system
	config        record.Config
	recordingID   record.RecordingID
	dirtyEntities ecs.DirtySet

	posGetter         record.ComponentGetter[tile.PosComponent]
	sizeGetter        record.ComponentGetter[tile.SizeComponent]
	obstructionGetter record.ComponentGetter[obstruction.Component]
	deployedGetter    record.ComponentGetter[obstruction.DeployedComponent]
}

func NewService(c ioc.Dic) obstruction.Service {
	s := ioc.GetServices[*service](c)

	s.obstruction = ecs.GetComponentsArray[obstruction.Component](s.World())
	s.deployed = ecs.GetComponentsArray[obstruction.DeployedComponent](s.World())

	s.obstruction.SetEmpty(obstruction.NewObstruction(definitions.LowlandObstruction))

	//
	s.config = record.NewConfig()
	s.dirtyEntities = ecs.NewDirtySet()

	s.posGetter = record.AddToConfig[tile.PosComponent](s.config)
	s.sizeGetter = record.AddToConfig[tile.SizeComponent](s.config)
	s.obstructionGetter = record.AddToConfig[obstruction.Component](s.config)
	s.deployedGetter = record.AddToConfig[obstruction.DeployedComponent](s.config)

	return s
}

func (s *service) Grid() ecs.ComponentsArray[grid.SquareGridComponent[obstruction.Obstruction]] {
	return s.ObstructionGridService.Component()
}
func (s *service) Component() ecs.ComponentsArray[obstruction.Component]        { return s.obstruction }
func (s *service) Deployed() ecs.ComponentsArray[obstruction.DeployedComponent] { return s.deployed }

func (s *service) Collisions(aabb obstruction.AABB, obstruction obstruction.Obstruction) []grid.Coords {
	var collisions []grid.Coords
	obstructionGridEntity := s.Grid().GetEntities()[0]
	obstructed, ok := s.Grid().Get(obstructionGridEntity)
	if !ok {
		collisions = append(collisions, aabb.Tiles...)
		return collisions
	}
	for _, coords := range aabb.Tiles {
		index, ok := obstructed.GetIndex(coords.Coords())
		if !ok || obstruction&obstructed.GetTile(index) != 0 {
			collisions = append(collisions, coords)
			continue
		}
	}
	return collisions
}
