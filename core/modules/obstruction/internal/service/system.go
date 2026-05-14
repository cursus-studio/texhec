package service

import (
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/record"
	"engine/services/ecs"
	"errors"
)

func (s *service) Register() error {
	s.config = record.NewConfig()
	s.dirtyEntities = ecs.NewDirtySet()

	s.posGetter = record.AddToConfig[tile.PosComponent](s.config)
	s.sizeGetter = record.AddToConfig[tile.SizeComponent](s.config)
	s.obstructionGetter = record.AddToConfig[obstruction.ObstructionComponent](s.config)
	s.deployedGetter = record.AddToConfig[obstruction.DeployedComponent](s.config)

	s.Tile().Pos().AddDirtySet(s.dirtyEntities)
	s.Tile().Size().AddDirtySet(s.dirtyEntities)
	s.Obstruction().Component().AddDirtySet(s.dirtyEntities)
	s.Obstruction().Deployed().AddDirtySet(s.dirtyEntities)
	s.Obstruction().ObstructionGrid().BeforeGet(s.BeforeGet)
	return nil
}

func (s *service) BeforeGet() {
	if len(s.dirtyEntities.Get()) == 0 {
		return
	}
	if len(s.Obstruction().ObstructionGrid().GetEntities()) == 0 {
		return
	}
	obstructionGridEntity := s.Obstruction().ObstructionGrid().GetEntities()[0]
	obstructionGrid, ok := s.Obstruction().ObstructionGrid().Get(obstructionGridEntity)
	if !ok {
		s.Logger().Log(errors.New("didn't found obstruction grid"))
		return
	}

	var entities []ecs.EntityID
	recording, ok := s.Record().Entity().Stop(s.recordingID)
	if !ok {
		entities = s.Obstruction().Deployed().GetEntities()
		goto entityLoop
	} else {
		entities = recording.Entities.GetIndices()
	}

	// remove old positions
	for _, entity := range entities {
		components, ok := recording.Entities.Get(entity)
		if !ok {
			continue
		}

		if _, ok := s.deployedGetter(components); !ok {
			continue
		}
		pos, ok := s.posGetter(components)
		if !ok {
			continue
		}
		size, _ := s.sizeGetter(components)
		obstructionComp, _ := s.obstructionGetter(components)
		aabb := obstruction.NewAABB(pos, size)
		for _, coords := range aabb.Tiles {
			index, ok := obstructionGrid.GetIndex(coords.Coords())
			if !ok {
				s.Logger().Log(tile.ErrInvalidPosition)
				continue
			}
			obstructionGrid.SetTile(index, obstructionGrid.GetTile(index)&^obstructionComp.Obstruction)
		}
	}

	// add new positions
entityLoop:
	for _, entity := range entities {
		if _, ok := s.Obstruction().Deployed().Get(entity); !ok {
			continue
		}
		pos, ok := s.Tile().Pos().Get(entity)
		if !ok {
			continue
		}
		size, _ := s.Tile().Size().Get(entity)
		obstructionComp, _ := s.Obstruction().Component().Get(entity)
		aabb := obstruction.NewAABB(pos, size)
		for _, coords := range aabb.Tiles {
			index, ok := obstructionGrid.GetIndex(coords.Coords())
			if !ok {
				s.Logger().Log(tile.ErrInvalidPosition)
				continue
			}
			if obstructionGrid.GetTile(index)&obstructionComp.Obstruction == 0 {
				continue
			}
			s.World().RemoveEntity(entity)
			s.Logger().Log(tile.ErrPositionIsOccupied)
			continue entityLoop
		}
		for _, coords := range aabb.Tiles {
			// index, ok validation is performed in loop before
			index, _ := obstructionGrid.GetIndex(coords.Coords())
			obstructionGrid.SetTile(index, obstructionGrid.GetTile(index)^obstructionComp.Obstruction)
		}
	}

	s.recordingID = s.Record().Entity().StartBackwardsRecording(s.config)
	s.Obstruction().ObstructionGrid().Set(obstructionGridEntity, obstructionGrid)
}
