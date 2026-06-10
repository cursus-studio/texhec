// defines generic slice based data structure to store grids
// it implements unified chunk size
package grid

import (
	"engine/services/ecs"

	"golang.org/x/exp/constraints"
)

type TileConstraint interface {
	comparable
}

//

type ChunkSize uint8 // stores power of 2

// cannot use number bigger then 31 because this would overflow Coord (uint32)
func NewChunkSize(size uint8) ChunkSize { return ChunkSize(min(size, (2<<4)-1)) }
func (s ChunkSize) Val() Coord          { return 1 << s }

//

type Coord uint32
type Coords struct{ X, Y Coord }
type Index int64 // stores coords

// allows to create negative coords
func NewCoord[Num constraints.Integer](n Num) Coord            { return Coord(n) }
func NewCoords[Number constraints.Integer](x, y Number) Coords { return Coords{Coord(x), Coord(y)} }

func (c *Coords) Size() (Coord, Coord) { return c.X, c.Y }
func (c *Coords) Coords() (X, Y Coord) {
	return c.X, c.Y
}

//

type ChunkComponent[Tile TileConstraint] struct {
	slice []Tile
}
type ChunkCoordsComponent Coords

func NewChunk[Tile TileConstraint](s Coord) ChunkComponent[Tile] {
	return ChunkComponent[Tile]{slice: make([]Tile, s*s)}
}
func (c ChunkComponent[Tile]) GetTiles() []Tile {
	tiles := make([]Tile, len(c.slice))
	copy(tiles, c.slice)
	return tiles
}
func (c ChunkComponent[Tile]) GetTile(index Index) Tile       { return c.slice[index] }
func (c ChunkComponent[Tile]) SetTile(index Index, tile Tile) { c.slice[index] = tile }

func NewChunkCoords(x, y Coord) ChunkCoordsComponent {
	return ChunkCoordsComponent{X: x, Y: y}
}

//

type ClickEvent struct {
	Chunk  ecs.EntityID
	Coords Coords
}
type HoverEvent struct {
	Chunk  ecs.EntityID
	Coords Coords
}

func NewClickEvent(chunk ecs.EntityID, tile Coords) ClickEvent { return ClickEvent{chunk, tile} }
func NewHoverEvent(chunk ecs.EntityID, tile Coords) HoverEvent { return HoverEvent{chunk, tile} }

//

// stores coords chunk data
type CoordsData[Tile TileConstraint] struct {
	Entity    ecs.EntityID
	Component ChunkComponent[Tile]
	Index     Index
}

func NewCoordsData[Tile TileConstraint](
	entity ecs.EntityID,
	component ChunkComponent[Tile],
	index Index,
) CoordsData[Tile] {
	return CoordsData[Tile]{entity, component, index}
}

type Service interface {
	// arrays
	Coords() ecs.ComponentsArray[ChunkCoordsComponent]
	GetChunk(ChunkCoordsComponent) (ecs.EntityID, bool)

	// getters within chunk
	ChunkSize() Coord
	CoordsIndex(Coords) (Index, bool)
	IndexCoords(index Index) Coords
	GetLastIndex() Index

	// calculate chunk coords
	AbsoluteCoords(ChunkCoordsComponent, Coords) Coords
	RelativeCoords(Coords) (ChunkCoordsComponent, Coords)
}

type ServiceT[Tile TileConstraint] interface {
	// arrays
	Chunk() ecs.ComponentsArray[ChunkComponent[Tile]]

	// ctors
	NewChunk() ChunkComponent[Tile]

	// calculate chunk coords
	CoordsData(Coords) (CoordsData[Tile], bool)
}
