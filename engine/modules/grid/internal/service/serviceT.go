package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/grid"

	"github.com/ogiusek/ioc/v2"
)

type serviceT[Tile grid.TileConstraint] struct {
	engine.EngineWorld `inject:""`
	chunk              ecs.ComponentArray[grid.ChunkComponent[Tile]]
}

func NewServiceT[Tile grid.TileConstraint](c ioc.Dic) grid.ServiceT[Tile] {
	s := ioc.GetServices[*serviceT[Tile]](c)
	s.chunk = ecs.GetComponentArray[grid.ChunkComponent[Tile]](s.World())
	return s
}

// arrays
func (s *serviceT[Tile]) Chunk() ecs.ComponentArray[grid.ChunkComponent[Tile]] { return s.chunk }

// ctors
func (s *serviceT[Tile]) NewChunk() grid.ChunkComponent[Tile] {
	return grid.NewChunk[Tile](s.Grid().ChunkSize())
}

// getters within chunk
func (s *serviceT[Tile]) ChunkSize() grid.Coord { return s.Grid().ChunkSize() }
func (s *serviceT[Tile]) CoordsIndex(coords grid.Coords) (grid.Index, bool) {
	size := s.Grid().ChunkSize()
	if coords.X >= size || coords.Y >= size {
		return 0, false
	}
	return grid.Index(coords.X) + grid.Index(coords.Y)*grid.Index(size), true
}
func (s *serviceT[Tile]) IndexCoords(index grid.Index) grid.Coords {
	size := grid.Index(s.Grid().ChunkSize())
	return grid.NewCoords(
		// #nosec G115
		grid.Coord(index%size),
		// #nosec G115
		grid.Coord(index/size),
	)
}
func (s *serviceT[Tile]) GetLastIndex() grid.Index {
	chunkSize := s.Grid().ChunkSize()
	return grid.Index(chunkSize * chunkSize)
}

// calculate chunk coords
func (s *serviceT[Tile]) CoordsData(coords grid.Coords) (grid.CoordsData[Tile], bool) {
	chunkCoords, tileCoords := s.Grid().RelativeCoords(coords)
	chunkEntity, ok := s.Grid().GetChunk(chunkCoords)
	if !ok {
		return grid.CoordsData[Tile]{}, false
	}
	chunk, ok := s.chunk.Get(chunkEntity)
	if !ok {
		return grid.CoordsData[Tile]{}, false
	}
	index, ok := s.CoordsIndex(tileCoords)
	if !ok {
		return grid.CoordsData[Tile]{}, false
	}
	data := grid.NewCoordsData(chunkEntity, chunk, index)
	return data, true
}
