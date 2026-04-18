package obstructionsystem

import (
	"core/modules/tile"
	gamescenes "core/scenes"
	"engine/modules/record"
	"engine/services/ecs"
	"errors"

	"github.com/ogiusek/ioc/v2"
)

type system struct {
	gamescenes.GameWorld `inject:""`

	config        record.Config
	recordingID   record.RecordingID
	dirtyEntities ecs.DirtySet

	posGetter         record.ComponentGetter[tile.PosComponent]
	sizeGetter        record.ComponentGetter[tile.SizeComponent]
	obstructionGetter record.ComponentGetter[tile.ObstructionComponent]
	deployedGetter    record.ComponentGetter[tile.DeployedComponent]
}

func NewSystem(c ioc.Dic) tile.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		s.config = record.NewConfig()
		s.dirtyEntities = ecs.NewDirtySet()

		s.posGetter = record.AddToConfig[tile.PosComponent](s.config)
		s.sizeGetter = record.AddToConfig[tile.SizeComponent](s.config)
		s.obstructionGetter = record.AddToConfig[tile.ObstructionComponent](s.config)
		s.deployedGetter = record.AddToConfig[tile.DeployedComponent](s.config)

		s.Tile().Pos().AddDirtySet(s.dirtyEntities)
		s.Tile().Size().AddDirtySet(s.dirtyEntities)
		s.Tile().Obstruction().AddDirtySet(s.dirtyEntities)
		s.Tile().Deployed().AddDirtySet(s.dirtyEntities)
		s.Tile().ObstructionGrid().BeforeGet(s.BeforeGet)
		return nil
	})
}

func (s *system) BeforeGet() {
	if len(s.dirtyEntities.Get()) == 0 {
		return
	}
	if len(s.Tile().ObstructionGrid().GetEntities()) == 0 {
		return
	}
	obstructionGridEntity := s.Tile().ObstructionGrid().GetEntities()[0]
	obstructionGrid, ok := s.Tile().ObstructionGrid().Get(obstructionGridEntity)
	if !ok {
		s.Logger().Warn(errors.New("didn't found obstruction grid"))
		return
	}

	var entities []ecs.EntityID
	recording, ok := s.Record().Entity().Stop(s.recordingID)
	if !ok {
		entities = s.Tile().Deployed().GetEntities()
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
		obstruction, _ := s.obstructionGetter(components)
		aabb := tile.NewAABB(pos, size)
		for _, coords := range aabb.Tiles {
			index, ok := obstructionGrid.GetIndex(coords.Coords())
			if !ok {
				s.Logger().Warn(tile.ErrInvalidPosition)
				continue
			}
			obstructionGrid.SetTile(index, obstructionGrid.GetTile(index)&^obstruction.Obstruction)
		}
	}

	// add new positions
entityLoop:
	for _, entity := range entities {
		if _, ok := s.Tile().Deployed().Get(entity); !ok {
			continue
		}
		pos, ok := s.Tile().Pos().Get(entity)
		if !ok {
			continue
		}
		size, _ := s.Tile().Size().Get(entity)
		obstruction, _ := s.Tile().Obstruction().Get(entity)
		aabb := tile.NewAABB(pos, size)
		for _, coords := range aabb.Tiles {
			index, ok := obstructionGrid.GetIndex(coords.Coords())
			if !ok {
				s.Logger().Warn(tile.ErrInvalidPosition)
				continue
			}
			if obstructionGrid.GetTile(index)&obstruction.Obstruction == 0 {
				continue
			}
			s.World().RemoveEntity(entity)
			s.Logger().Warn(tile.ErrPositionIsOccupied)
			continue entityLoop
		}
		for _, coords := range aabb.Tiles {
			// index, ok validation is performed in loop before
			index, _ := obstructionGrid.GetIndex(coords.Coords())
			obstructionGrid.SetTile(index, obstructionGrid.GetTile(index)^obstruction.Obstruction)
		}
	}

	s.recordingID = s.Record().Entity().StartBackwardsRecording(s.config)
	s.Tile().ObstructionGrid().Set(obstructionGridEntity, obstructionGrid)
}
