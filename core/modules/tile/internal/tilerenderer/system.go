package tilerenderer

import (
	"core/game"
	"core/modules/tile"
	_ "embed"
	"engine/modules/assets"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/graphics"
	"engine/modules/grid"
	"engine/modules/render"
	"errors"
	"fmt"
	"image"
	"math"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"golang.org/x/exp/constraints"
)

//go:embed shader.vert
var vertSource string

//go:embed shader.geom
var geomSource string

//go:embed shader.frag
var fragSource string

type TileType struct {
	Texture image.Image
}

type Batch struct {
	buffer graphics.Buffer[int32]
}

func (b *Batch) Release() {
	b.buffer.Release()
}

//

type locations struct {
	Mvp  int32 `uniform:"mvp"`  // mat4
	Size int32 `uniform:"size"` // uint
	// widthInv and heightInv is 2/width and 2/height
	SizeInv int32 `uniform:"sizeInv"` // float
}

type system struct {
	game.GameWorld `inject:""`

	program   graphics.Program
	locations locations
	ids       datastructures.SparseArray[tile.ID, uint32]
	// lod shrinks
	lodTextureArrays []graphics.TextureArray
	texturesBuffer   graphics.Buffer[mgl32.Vec2] // [index, amount]
	vao              graphics.VAO

	tileTextures datastructures.SparseArray[uint32, mgl32.Vec2]
	textures     datastructures.SparseArray[uint32, image.Image]

	tilesDirtySet ecs.DirtySet
	gridDirtySet  ecs.DirtySet
	batches       datastructures.SparseArray[ecs.EntityID, Batch]
}

func NewSystem(c ioc.Dic) error {
	s := ioc.GetServices[*system](c)

	vert, err := s.Graphics().NewShader(vertSource, graphics.VertexShader)
	if err != nil {
		return err
	}
	defer vert.Release()

	geom, err := s.Graphics().NewShader(geomSource, graphics.GeomShader)
	if err != nil {
		return err
	}
	defer geom.Release()

	frag, err := s.Graphics().NewShader(fragSource, graphics.FragmentShader)
	if err != nil {
		return err
	}
	defer frag.Release()

	programID := gl.CreateProgram()
	gl.AttachShader(programID, vert.ID())
	gl.AttachShader(programID, geom.ID())
	gl.AttachShader(programID, frag.ID())

	p, err := s.Graphics().NewProgram(programID, nil)
	if err != nil {
		return err
	}

	locations, err := graphics.GetProgramLocations[locations](p)
	if err != nil {
		return err
	}

	s.program = p
	s.vao = s.Graphics().NewVAO(nil, nil)
	s.locations = locations
	s.ids = datastructures.NewSparseArray[tile.ID, uint32]()
	s.lodTextureArrays = []graphics.TextureArray{}

	s.texturesBuffer = graphics.NewBuffer[mgl32.Vec2](gl.SHADER_STORAGE_BUFFER, gl.DYNAMIC_DRAW, 1)

	s.tileTextures = datastructures.NewSparseArray[uint32, mgl32.Vec2]()
	s.textures = datastructures.NewSparseArray[uint32, image.Image]()

	s.tilesDirtySet = ecs.NewDirtySet()
	s.Tile().Component().AddDirtySet(s.tilesDirtySet)

	s.gridDirtySet = ecs.NewDirtySet()
	s.Tile().Grid().Chunk().AddDirtySet(s.gridDirtySet)

	s.batches = datastructures.NewSparseArray[ecs.EntityID, Batch]()

	events.Listen(s.EventsBuilder(), s.ListenRender)
	return nil
}

// padding neighbor

type Neighbor struct {
	ChunkCoords grid.Coords
}

func NewNeighbor[Number constraints.Integer](x, y Number) Neighbor {
	return Neighbor{grid.NewCoords(grid.NewCoord(x), grid.NewCoord(y))}
}

// returns required (padding indices and grid indices) on grid with padding
// these can be taken to paste values from neighbor chunks to current chunks
// (padding is of size 1 around whole grid so new grid size is n+2)
func (n *Neighbor) GetIndices(size int) (gridIndices, paddingIndices []int) {
	s := size + 2
	dx, dy := n.ChunkCoords.X, n.ChunkCoords.Y
	xLen, yLen := size, size
	if dx != 0 {
		xLen = 1
	}
	if dy != 0 {
		yLen = 1
	}
	gridIndices = make([]int, xLen*yLen)
	paddingIndices = make([]int, xLen*yLen)

	idx := 0
	for i := range yLen {
		var gY, pY int
		switch dy {
		case grid.NewCoord(-1):
			gY, pY = size, 0
		case 0:
			gY, pY = i+1, i+1
		case 1:
			gY, pY = 1, size+1
		}

		for j := range xLen {
			var gX, pX int
			switch dx {
			case grid.NewCoord(-1):
				gX, pX = size, 0
			case 0:
				gX, pX = j+1, j+1
			case 1:
				gX, pX = 1, size+1
			}

			gridIndices[idx] = gY*s + gX
			paddingIndices[idx] = pY*s + pX
			idx++
		}
	}
	return gridIndices, paddingIndices
}

var Neighbors = []Neighbor{
	NewNeighbor(-1, -1),
	NewNeighbor(0, -1),
	NewNeighbor(1, -1),
	NewNeighbor(-1, 0),
	NewNeighbor(1, 0),
	NewNeighbor(-1, 1),
	NewNeighbor(0, 1),
	NewNeighbor(1, 1),
}

func getLastID(s *system) int {
	size := s.Grid().ChunkSize() + 2
	return int(size * size)
}
func getIndexOnPadding[Num constraints.Integer](s *system, i Num) Num {
	size := Num(s.Grid().ChunkSize())
	return i + size + 3 + (i/(size))*2
}
func (s *system) applyPadding(
	chunkCoords grid.ChunkCoordsComponent,
	batch Batch,
) {
	// get neighbor to copy from and neighbor to copy to
	for _, neighbor := range Neighbors {
		neighborCoords := grid.NewChunkCoords(
			chunkCoords.X+neighbor.ChunkCoords.X,
			chunkCoords.Y+neighbor.ChunkCoords.Y,
		)
		neighborEntity, ok := s.Grid().GetChunk(neighborCoords)
		if !ok {
			continue
		}
		neighborBatch, ok := s.batches.Get(neighborEntity)
		if !ok {
			continue
		}
		// apply padding on current chunk
		gridIndices, paddingIndices := neighbor.GetIndices(int(s.Grid().ChunkSize()))
		for i, gridIndex := range gridIndices {
			paddingIndex := paddingIndices[i]
			val := neighborBatch.buffer.Get()[gridIndex]
			batch.buffer.Set(paddingIndex, val)
		}
	}
	for _, neighbor := range Neighbors {
		neighborCoords := grid.NewChunkCoords(
			chunkCoords.X-neighbor.ChunkCoords.X,
			chunkCoords.Y-neighbor.ChunkCoords.Y,
		)
		neighborEntity, ok := s.Grid().GetChunk(neighborCoords)
		if !ok {
			continue
		}
		neighborBatch, ok := s.batches.Get(neighborEntity)
		if !ok {
			continue
		}
		// apply padding on neighbor chunk
		gridIndices, paddingIndices := neighbor.GetIndices(int(s.Grid().ChunkSize()))
		for i, gridIndex := range gridIndices {
			paddingIndex := paddingIndices[i]
			val := batch.buffer.Get()[gridIndex]
			neighborBatch.buffer.Set(paddingIndex, val)
		}
		neighborBatch.buffer.Flush()
	}
}

func (s *system) ListenRender(render render.RenderEvent) {
	{ // reload bioms. it reloads definitions, buffers, texture arrays
		dirtyTiles := s.tilesDirtySet.Get()
		for _, entity := range dirtyTiles {
			tileComp, ok := s.Tile().Component().Get(entity)
			if !ok {
				continue
			}

			idsCount := s.ids.Size()
			if idsCount < 0 || idsCount > math.MaxUint32 {
				s.Logger().Log(fmt.Errorf("cannot create tile type. run out of ids"))
				continue
			}

			id := uint32(idsCount)
			s.ids.Set(tileComp.ID, id)
			texture, err := assets.GetAsset[tile.BiomeAsset](s.Assets(), entity)
			if err != nil {
				s.Logger().Log(err)
				continue
			}

			rangeBase := id*15 + 1
			for i, images := range texture.Images() {
				size := s.textures.Size()
				tileRange := mgl32.Vec2{float32(size), float32(len(images))}
				s.tileTextures.Set(rangeBase+uint32(i), tileRange)

				imageBase := size
				for i, img := range images {
					index := imageBase + i
					if index < 0 || index > math.MaxUint32 {
						s.Logger().Log(fmt.Errorf("run out of ids for images"))
						break
					}
					s.textures.Set(uint32(index), img)
				}
			}
		}
		if len(dirtyTiles) != 0 {
			highLodTextureArray, err := s.Graphics().TextureArray().New(s.textures)
			if err != nil {
				s.Logger().Log(err)
				return
			}

			lowLodTextures := datastructures.NewSparseArray[uint32, image.Image]()
			for _, texture := range s.textures.GetIndices() {
				img, _ := s.textures.Get(texture)
				img = s.Graphics().NewImage(img).Scale(2, 2).Opaque().Image()
				lowLodTextures.Set(texture, img)
			}
			lowLodTextureArray, err := s.Graphics().TextureArray().New(lowLodTextures)
			if err != nil {
				s.Logger().Log(err)
				return
			}

			dirtySet := ecs.NewDirtySet()
			s.Tile().Grid().Chunk().AddDirtySet(dirtySet)

			for _, t := range s.lodTextureArrays {
				t.Release()
			}

			s.lodTextureArrays = []graphics.TextureArray{
				highLodTextureArray,
				lowLodTextureArray,
			}

			for _, id := range s.tileTextures.GetIndices() {
				value, _ := s.tileTextures.Get(id)
				s.texturesBuffer.Set(int(id), value)
			}
			s.texturesBuffer.Flush()
		}
		if len(s.lodTextureArrays) == 0 {
			return
		}
	}

	// reload per grid buffers
	for _, entity := range s.gridDirtySet.Get() {
		batch, batchOk := s.batches.Get(entity)
		grid, compOk := s.Tile().Grid().Chunk().Get(entity)
		chunkCoords, _ := s.Grid().Coords().Get(entity)

		if !batchOk && !compOk {
			continue
		}
		if batchOk && !compOk {
			batch.Release()
			s.batches.Remove(entity)
			continue
		}
		if !batchOk && compOk {
			batch = Batch{
				graphics.NewBuffer[int32](gl.SHADER_STORAGE_BUFFER, gl.DYNAMIC_DRAW, 0),
			}
			// set size buffer
			batch.buffer.Set(getLastID(s), 0)
			s.batches.Set(entity, batch)
		}

		for i, tile := range grid.GetTiles() {
			// there is a conflict
			// we use definitionID to define tile and textures used
			// but tile values are tile.Type diffrentiate it
			id, ok := s.ids.Get(tile)
			if !ok {
				continue
			}
			// #nosec G115
			batch.buffer.Set(getIndexOnPadding(s, i), int32(id))
		}
		// apply paddings
		s.applyPadding(chunkCoords, batch)
		batch.buffer.Flush()
	}

	// render
	s.texturesBuffer.Bind()

	s.program.Bind()
	s.vao.Bind()
	var lod int
	if ortho, ok := s.Camera().Ortho().Get(render.Camera); ok && ortho.Zoom < .25 {
		lod = 1
	}
	s.lodTextureArrays[lod].Bind()

	cameraGroups, _ := s.Groups().Component().Get(render.Camera)
	cameraMatrix := s.Camera().Mat4(render.Camera)

	size := s.Grid().ChunkSize()
	gl.Uniform1ui(s.locations.Size, uint32(size))
	gl.Uniform1f(s.locations.SizeInv, 2/float32(size))

	for _, entity := range s.batches.GetIndices() {
		batch, ok := s.batches.Get(entity)
		if !ok {
			continue
		}
		if groups, _ := s.Groups().Component().Get(entity); !cameraGroups.SharesAnyGroup(groups) {
			continue
		}
		batch.buffer.Bind()

		mvp := cameraMatrix.Mul4(s.Transform().Mat4(entity))
		gl.UniformMatrix4fv(s.locations.Mvp, 1, false, &mvp[0])

		verticesCount := (size + 1) * (size + 1)
		if verticesCount > math.MaxInt32 {
			s.Logger().Warn(errors.New("tiles have to many vertices"))
			verticesCount = math.MaxInt32
		}
		gl.DrawArrays(gl.POINTS, 0, int32(verticesCount))
	}
}
