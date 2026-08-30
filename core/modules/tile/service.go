package tile

import (
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/modules/uuid"
	"errors"
	"image"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/exp/constraints"
)

var (
	// error logged when grid.GetIndex returns !ok
	ErrInvalidPosition                  error = errors.New("tile:position not found on the grid")
	ErrInvalidStep                      error = errors.New("tile:invalid step")
	ErrPositionAndSpeedIsRequiredToStep error = errors.New("tile:to step you need to have speed and position")
	ErrBlueprintIsMissingUUID           error = errors.New("blueprint is missing UUID")
)

type ID uint8

//

type Component struct {
	ID ID
}

func NewTile(id ID) Component {
	return Component{id}
}

//

type Coord float32
type PosComponent struct {
	X, Y Coord
}

func NewPos[Number constraints.Integer | constraints.Float](x, y Number) PosComponent {
	return PosComponent{Coord(x), Coord(y)}
}
func (PosComponent) Smooth() {}
func (c1 PosComponent) Lerp(c2 PosComponent, mix32 float32) PosComponent {
	return PosComponent{
		transition.Lerp(c1.X, c2.X, mix32),
		transition.Lerp(c1.Y, c2.Y, mix32),
	}
}
func abs[Number constraints.Float | constraints.Integer](n Number) Number { return max(-n, n) }
func (p *PosComponent) Aligned() (coords grid.Coords, isAligned bool) {
	const epsilon Coord = 1e-3
	x, y := grid.Coord(p.X+.5-epsilon), grid.Coord(p.Y+.5-epsilon)
	return grid.NewCoords(x, y), abs(Coord(x)-p.X) < epsilon && abs(Coord(y)-p.Y) < epsilon
}

//

type LayerComponent struct {
	Z Coord
}

func NewLayer[Number constraints.Integer | constraints.Float](z Number) LayerComponent {
	return LayerComponent{Coord(z)}
}

//

type NameComponent struct{ Name string }
type LinkComponent struct{ UUID uuid.Component }       // link to object
type CachedLinkComponent struct{ Entity ecs.EntityID } // cached link to object

func NewName(name string) NameComponent                     { return NameComponent{name} }
func NewLink(uuid uuid.Component) LinkComponent             { return LinkComponent{uuid} }
func NewCachedLink(entity ecs.EntityID) CachedLinkComponent { return CachedLinkComponent{entity} }

//

type SizeComponent struct {
	X, Y grid.Coord
}

func NewSize[Number constraints.Integer](x, y Number) SizeComponent {
	return SizeComponent{grid.Coord(x), grid.Coord(y)}
}
func (c *SizeComponent) Size() (grid.Coord, grid.Coord) {
	return c.X, c.Y
}

//

type RotComponent struct {
	Radians float32
}

func NewRot(radians float32) RotComponent {
	return RotComponent{radians}
}
func (RotComponent) Smooth() {}
func (c1 RotComponent) Lerp(c2 RotComponent, mix32 float32) RotComponent {
	const Tau = 2 * math.Pi
	c2.Radians = c1.Radians + float32(math.Remainder(float64(c2.Radians-c1.Radians), Tau))
	return RotComponent{transition.Lerp(c1.Radians, c2.Radians, mix32)}
}
func (e *RotComponent) Quat() mgl32.Quat {
	return mgl32.QuatRotate(e.Radians, mgl32.Vec3{0, 0, 1})
}

//

type BiomeAsset interface {
	Images() [15][]image.Image
	Res() image.Rectangle
	AspectRatio() image.Rectangle
	Release()
}

//

type MissingChunkEvent struct{ Coords grid.ChunkCoordsComponent }
type UnloadChunkEvent struct{ Coords grid.ChunkCoordsComponent }

func NewMissingChunkEvent(coords grid.ChunkCoordsComponent) MissingChunkEvent {
	return MissingChunkEvent{coords}
}
func NewUnloadChunkEvent(coords grid.ChunkCoordsComponent) UnloadChunkEvent {
	return UnloadChunkEvent{coords}
}

//

type ApplyCoordsEvent interface {
	ApplyCoords(grid.Coords) any
}

//

type ClickEntityEvent struct{ Entity ecs.EntityID }
type ClickBlueprintEvent struct{ Entity ecs.EntityID }

func NewClickEntityEvent() ClickEntityEvent { return ClickEntityEvent{} }
func NewClickBlueprintEvent(deployed ecs.EntityID) ClickBlueprintEvent {
	return ClickBlueprintEvent{deployed}
}
func (e ClickEntityEvent) ApplyEntity(entity ecs.EntityID) any {
	e.Entity = entity
	return e
}

//

type Service interface {
	ecs.SystemRegister
	Renderer() ecs.SystemRegister

	Component() ecs.ComponentArray[Component]
	Grid() grid.ServiceT[ID]
	GetTile(ID) (ecs.EntityID, bool)

	// is often used to check is element a prototype or a placed object
	Pos() ecs.ComponentArray[PosComponent]
	Size() ecs.ComponentArray[SizeComponent]
	Rot() ecs.ComponentArray[RotComponent]
	Layer() ecs.ComponentArray[LayerComponent]

	Name() ecs.ComponentArray[NameComponent]
	Link() ecs.ComponentArray[LinkComponent]
	CachedLink() ecs.ComponentArray[CachedLinkComponent]

	GetLink(object ecs.EntityID) (source ecs.EntityID, ok bool)

	// src images should be:
	// - 1111
	// - 1110
	// - 1010
	// - 1001
	// - 0001
	NewBiomeAsset(srcImages [6][]image.Image) (BiomeAsset, error)

	GetPos(coords grid.Coords) transform.PosComponent
}
