package obstructionsystem

import (
	"core/modules/tile"
	"engine"
	"engine/modules/grid"
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
}

func NewSystem(c ioc.Dic) tile.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		s.config = record.NewConfig()
		s.dirtyEntities = ecs.NewDirtySet()

		record.AddToConfig[tile.PosComponent](s.config)
		record.AddToConfig[tile.ObstructionComponent](s.config)
		record.AddToConfig[tile.DeployedComponent](s.config)

		s.Tile.Pos().AddDirtySet(s.dirtyEntities)
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
			index, ok := obstructionGrid.GetIndex(grid.Coord(pos.X), grid.Coord(pos.Y))
			if !ok {
				s.Logger.Warn(errors.New("invalid position"))
				continue
			}
			obstruction, ok := s.Tile.Obstruction().Get(entity)
			if !ok {
				continue
			}
			res := obstructionGrid.GetTile(index) | obstruction.Obstruction
			obstructionGrid.SetTile(index, res)
		}
		goto cleanup
	}

	for _, entity := range recording.Entities.GetIndices() {
		if components, ok := recording.Entities.Get(entity); ok &&
			len(components) != 0 &&
			components[0] != nil &&
			components[1] != nil &&
			components[2] != nil {
			pos := components[0].(tile.PosComponent)
			obstruction := components[1].(tile.ObstructionComponent)
			index, ok := obstructionGrid.GetIndex(grid.Coord(pos.X), grid.Coord(pos.Y))
			if !ok {
				s.Logger.Warn(errors.New("invalid position"))
			} else {
				obstructionGrid.SetTile(index, obstructionGrid.GetTile(index)&^obstruction.Obstruction)
			}
		}
		if _, ok := s.Tile.Deployed().Get(entity); !ok {
			continue
		}
		pos, ok := s.Tile.Pos().Get(entity)
		if !ok {
			continue
		}
		obstruction, _ := s.Tile.Obstruction().Get(entity)
		index, ok := obstructionGrid.GetIndex(grid.Coord(pos.X), grid.Coord(pos.Y))
		if !ok {
			s.Logger.Warn(errors.New("invalid position"))
			continue
		}
		obstructionGrid.SetTile(index, obstructionGrid.GetTile(index)|obstruction.Obstruction)
	}

cleanup:
	s.recordingID = s.Record.Entity().StartBackwardsRecording(s.config)
	s.Tile.ObstructionGrid().Set(obstructionGridEntity, obstructionGrid)
}
