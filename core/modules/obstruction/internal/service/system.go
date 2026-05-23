package service

import (
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/services/ecs"
)

func (s *service) Register() error {
	s.Tile().Pos().AddDirtySet(s.dirtyEntities)
	s.Tile().Size().AddDirtySet(s.dirtyEntities)
	s.Obstruction().Component().AddDirtySet(s.dirtyEntities)
	s.Obstruction().Deployed().AddDirtySet(s.dirtyEntities)
	s.Obstruction().Grid().Chunk().BeforeGet(s.UpdateObstructionGridBeforeGet)
	return nil
}

func (s *service) UpdateObstructionGridBeforeGet() {
	if len(s.dirtyEntities.Get()) == 0 {
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
			data, ok := s.Obstruction().Grid().CoordsData(coords)
			if !ok {
				s.Logger().Log(tile.ErrInvalidPosition)
				continue
			}
			data.Component.SetTile(data.Index, data.Component.GetTile(data.Index)&^obstructionComp.Obstruction)
			s.Obstruction().Grid().Chunk().Set(data.Entity, data.Component)
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
		tilesData := make([]grid.CoordsData[obstruction.Obstruction], len(aabb.Tiles))
		for i, coords := range aabb.Tiles {
			data, ok := s.Obstruction().Grid().CoordsData(coords)
			if !ok {
				s.Logger().Log(tile.ErrInvalidPosition)
				continue
			}
			tilesData[i] = data
			if data.Component.GetTile(data.Index)&obstructionComp.Obstruction == 0 {
				continue
			}
			s.World().RemoveEntity(entity)
			s.Logger().Log(obstruction.ErrPositionIsOccupied)
			continue entityLoop
		}
		for i := range aabb.Tiles {
			data := tilesData[i]
			data.Component.SetTile(data.Index, data.Component.GetTile(data.Index)^obstructionComp.Obstruction)
			s.Obstruction().Grid().Chunk().Set(data.Entity, data.Component)
		}
	}

	s.recordingID = s.Record().Entity().StartBackwardsRecording(s.config)
}
