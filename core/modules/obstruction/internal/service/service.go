package service

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/record"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld         `inject:""`
	ObstructionGridService grid.ServiceT[obstruction.Obstruction] `inject:""`

	obstruction ecs.ComponentArray[obstruction.Component]
	deployed    ecs.ComponentArray[obstruction.DeployedComponent]

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

	s.obstruction = ecs.GetComponentArray[obstruction.Component](s.World())
	s.deployed = ecs.GetComponentArray[obstruction.DeployedComponent](s.World())

	s.obstruction.SetEmpty(obstruction.NewObstruction(definitions.LowlandObstruction))

	s.Deployed().OnUpsert(s.OnDeployUpsert)

	//
	s.config = record.NewConfig()
	s.dirtyEntities = ecs.NewDirtySet()

	s.posGetter = record.AddToConfig[tile.PosComponent](s.config)
	s.sizeGetter = record.AddToConfig[tile.SizeComponent](s.config)
	s.obstructionGetter = record.AddToConfig[obstruction.Component](s.config)
	s.deployedGetter = record.AddToConfig[obstruction.DeployedComponent](s.config)

	return s
}

func (s *service) OnDeployUpsert(entity ecs.EntityID) {
	s.Inputs().Stack().Set(entity, inputs.StackComponent{})
}

func (s *service) Grid() grid.ServiceT[obstruction.Obstruction]                { return s.ObstructionGridService }
func (s *service) Component() ecs.ComponentArray[obstruction.Component]        { return s.obstruction }
func (s *service) Deployed() ecs.ComponentArray[obstruction.DeployedComponent] { return s.deployed }

func (s *service) Collisions(aabb obstruction.AABB, obstruction obstruction.Obstruction) []grid.Coords {
	var collisions []grid.Coords
	for _, coords := range aabb.Tiles {
		data, ok := s.Obstruction().Grid().CoordsData(coords)
		if !ok {
			collisions = append(collisions, coords)
			continue
		}
		if obstruction&data.Component.GetTile(data.Index) != 0 {
			collisions = append(collisions, coords)
			continue
		}
	}
	return collisions
}
