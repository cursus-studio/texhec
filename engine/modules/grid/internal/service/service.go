package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/relation"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld                          `inject:""`
	relation.Service[grid.ChunkCoordsComponent] `inject:""`
	chunkSize                                   grid.ChunkSize
	chunkSizeTileMask                           grid.Coord
	coords                                      ecs.ComponentArray[grid.ChunkCoordsComponent]
}

func NewService(c ioc.Dic, chunkSize grid.ChunkSize) grid.Service {
	s := ioc.GetServices[*service](c)
	s.chunkSize = chunkSize
	s.chunkSizeTileMask = chunkSize.Val() - 1
	s.coords = ecs.GetComponentArray[grid.ChunkCoordsComponent](s.World())

	s.Coords().OnUpsert(s.OnUpsert)
	events.Listen(s.EventsBuilder(), s.OnHover)
	events.Listen(s.EventsBuilder(), s.OnClick)
	return s
}

// arrays
func (s *service) Coords() ecs.ComponentArray[grid.ChunkCoordsComponent] {
	return s.coords
}

func (s *service) GetChunk(coords grid.ChunkCoordsComponent) (ecs.EntityID, bool) {
	return s.Get(coords)
}

// getters within chunk
func (s *service) ChunkSize() grid.Coord { return s.chunkSize.Val() }
func (s *service) CoordsIndex(coords grid.Coords) (grid.Index, bool) {
	size := s.chunkSize.Val()
	if coords.X >= size || coords.Y >= size {
		return 0, false
	}
	return grid.Index(coords.X) + grid.Index(coords.Y)*grid.Index(size), true
}
func (s *service) IndexCoords(index grid.Index) grid.Coords {
	return grid.NewCoords(
		// #nosec G115
		grid.Coord(index&grid.Index(s.chunkSizeTileMask)), // &  -> index % size
		// #nosec G115
		grid.Coord(index>>s.chunkSize), // >> -> index / size
	)
}
func (s *service) GetLastIndex() grid.Index {
	chunkSize := s.chunkSize.Val()
	return grid.Index(chunkSize * chunkSize)
}

// calculate chunk coords
func (s *service) AbsoluteCoords(chunkCoords grid.ChunkCoordsComponent, coords grid.Coords) grid.Coords {
	coords.X += s.chunkSize.Val() * chunkCoords.X
	coords.Y += s.chunkSize.Val() * chunkCoords.Y
	return coords
}
func (s *service) RelativeCoords(coords grid.Coords) (grid.ChunkCoordsComponent, grid.Coords) {
	chunkCoords := grid.NewChunkCoords(
		coords.X>>s.chunkSize, // >> -> /
		coords.Y>>s.chunkSize, // >> -> /
	)
	relCoords := grid.NewCoords(
		coords.X&s.chunkSizeTileMask, // & -> %
		coords.Y&s.chunkSizeTileMask, // & -> %
	)
	return chunkCoords, relCoords
}
