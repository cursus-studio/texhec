package obstructionsystem

import (
	"core/modules/tile"
	"engine"
	"engine/modules/record"
	"engine/services/ecs"
	"errors"

	"github.com/ogiusek/ioc/v2"
)

type system struct {
	engine.World `inject:"1"`
	Tile         tile.Service `inject:"1"`

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

		s.Tile.Pos().AddDirtySet(s.dirtyEntities)
		s.Tile.Size().AddDirtySet(s.dirtyEntities)
		s.Tile.Obstruction().AddDirtySet(s.dirtyEntities)
		s.Tile.Deployed().AddDirtySet(s.dirtyEntities)
		s.Tile.ObstructionGrid().BeforeGet(s.BeforeGet)
		return nil
	})
}

func (s *system) BeforeGet() {
	if len(s.dirtyEntities.Get()) == 0 {
		return
	}
	if len(s.Tile.ObstructionGrid().GetEntities()) == 0 {
		return
	}
	obstructionGridEntity := s.Tile.ObstructionGrid().GetEntities()[0]
	obstructionGrid, ok := s.Tile.ObstructionGrid().Get(obstructionGridEntity)
	if !ok {
		s.Logger.Warn(errors.New("didn't found obstruction grid"))
		return
	}

	recording, ok := s.Record.Entity().Stop(s.recordingID)
	if !ok {
		for _, entity := range s.Tile.Deployed().GetEntities() {
			pos, ok := s.Tile.Pos().Get(entity)
			if !ok {
				continue
			}
			size, _ := s.Tile.Size().Get(entity)
			aabb := tile.NewAABB(pos, size)
			for _, coords := range aabb.Tiles {
				index, ok := obstructionGrid.GetIndex(coords.Coords())
				if !ok {
					s.Logger.Warn(tile.ErrInvalidPosition)
					continue
				}
				obstruction, ok := s.Tile.Obstruction().Get(entity)
				if !ok {
					continue
				}
				res := obstructionGrid.GetTile(index) | obstruction.Obstruction
				obstructionGrid.SetTile(index, res)
			}
		}
		goto cleanup
	}

	// remove old positions
	for _, entity := range recording.Entities.GetIndices() {
		components, ok := recording.Entities.Get(entity)
		if !ok {
			continue
		}
		pos, ok := s.posGetter(components)
		if !ok {
			continue
		}
		size, _ := s.sizeGetter(components)
		aabb := tile.NewAABB(pos, size)
		for _, coords := range aabb.Tiles {
			obstruction, ok := s.obstructionGetter(components)
			if !ok {
				continue
			}
			if _, ok := s.deployedGetter(components); !ok {
				continue
			}
			index, ok := obstructionGrid.GetIndex(coords.Coords())
			if !ok {
				s.Logger.Warn(tile.ErrInvalidPosition)
				continue
			}
			obstructionGrid.SetTile(index, obstructionGrid.GetTile(index)&^obstruction.Obstruction)
		}
	}

	// add new positions
	for _, entity := range recording.Entities.GetIndices() {
		if _, ok := s.Tile.Deployed().Get(entity); !ok {
			continue
		}
		pos, ok := s.Tile.Pos().Get(entity)
		if !ok {
			continue
		}
		size, _ := s.Tile.Size().Get(entity)
		obstruction, _ := s.Tile.Obstruction().Get(entity)
		aabb := tile.NewAABB(pos, size)
		for _, coords := range aabb.Tiles {
			index, ok := obstructionGrid.GetIndex(coords.Coords())
			if !ok {
				s.Logger.Warn(tile.ErrInvalidPosition)
				continue
			}
			obstructionGrid.SetTile(index, obstructionGrid.GetTile(index)|obstruction.Obstruction)
		}
	}

cleanup:
	s.recordingID = s.Record.Entity().StartBackwardsRecording(s.config)
	s.Tile.ObstructionGrid().Set(obstructionGridEntity, obstructionGrid)
}
