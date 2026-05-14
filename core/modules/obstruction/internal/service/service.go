package service

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/modules/record"
	"engine/modules/relation"
	"engine/modules/transform"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld         `inject:""`
	ObstructionGridService grid.Service[obstruction.Obstruction] `inject:""`
	TileTypeRelation       relation.Service[tile.ID]             `inject:""`

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

	return s
}

func (s *service) Grid() ecs.ComponentsArray[grid.SquareGridComponent[obstruction.Obstruction]] {
	return s.ObstructionGridService.Component()
}
func (s *service) GetTileType(id tile.ID) (ecs.EntityID, bool) {
	return s.TileTypeRelation.Get(id)
}

func (s *service) Component() ecs.ComponentsArray[obstruction.Component] {
	return s.obstruction
}
func (s *service) Deployed() ecs.ComponentsArray[obstruction.DeployedComponent] { return s.deployed }

func (s *service) GetPos(coords grid.Coords) transform.PosComponent {
	size := s.GetTileSize().Size
	return transform.NewPos(
		size.X()*(float32(coords.X)+.5),
		size.Y()*(float32(coords.Y)+.5),
		size.Z(),
	)
}
func (s *service) GetTileSize() transform.SizeComponent {
	return transform.NewSize(100, 100, 1)
}

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
