package internal

import (
	"core/game"
	"core/modules/obstruction"
	"core/modules/pathfind"
	"core/modules/tile"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/relation"
	"fmt"
	"slices"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

// regions are composed from fragments.
// this is intermediate f
type RegionFragment pathfind.Region

// this variable contains region index and is used for region connectivity
type Region = pathfind.Region

//

// key of which chunk is described
type ChunkObstructionComponent struct {
	ChunkCoords grid.ChunkCoordsComponent
	Obstruction obstruction.Obstruction
}

func NewChunkObstruction(coords grid.ChunkCoordsComponent, obstruction obstruction.Obstruction) ChunkObstructionComponent {
	return ChunkObstructionComponent{coords, obstruction}
}

//

type RegionData struct {
	Obstruction obstruction.Obstruction
	Fragments   datastructures.SparseSet[RegionFragment] // once upon no fragments contain region, region can be removed
}
type FragmentData struct {
	Region        Region
	ChunksCounter int // once upon no chunks contain fragment, fragment can be remved
}
type ChunkData struct {
	Fragments datastructures.SparseSet[RegionFragment]
}

func NewRegionData(obstruction obstruction.Obstruction) RegionData {
	return RegionData{obstruction, datastructures.NewSparseSet[RegionFragment]()}
}
func NewFragData(region pathfind.Region) FragmentData {
	return FragmentData{region, 0}
}
func NewChunkData(fragments datastructures.SparseSet[RegionFragment]) ChunkData {
	return ChunkData{fragments}
}

//

type regionsService struct {
	game.GameWorld   `inject:""`
	Relation         relation.Service[ChunkObstructionComponent] `inject:""`
	FragmentsChunks  grid.ServiceT[RegionFragment]               `inject:""`
	chunkObstruction ecs.ComponentArray[ChunkObstructionComponent]

	regions        datastructures.SparseArray[Region, RegionData]
	regionsCounter Region
	regionGaps     datastructures.SparseSet[Region]
	zeroRegion     Region

	fragments        datastructures.SparseArray[RegionFragment, FragmentData]
	fragmentsCounter RegionFragment
	fragmentsGaps    datastructures.SparseSet[RegionFragment]

	chunksData map[grid.ChunkCoordsComponent]ChunkData
}

func newRegionService(c ioc.Dic) *regionsService {
	s := ioc.GetServices[*regionsService](c)
	s.chunkObstruction = ecs.GetComponentArray[ChunkObstructionComponent](s.World())

	s.regions = datastructures.NewSparseArray[Region, RegionData]()
	s.regionsCounter = 1
	s.regionGaps = datastructures.NewSparseSet[Region]()

	s.fragments = datastructures.NewSparseArray[RegionFragment, FragmentData]()
	s.fragmentsCounter = 1
	s.fragmentsGaps = datastructures.NewSparseSet[RegionFragment]()

	s.chunksData = make(map[grid.ChunkCoordsComponent]ChunkData)

	s.Tile().Grid().Chunk().OnUpsert(s.OnChunkUpsert)
	events.Listen(s.EventsBuilder(), s.OnChunkUnload)
	return s
}

func (s *regionsService) newFragment() RegionFragment {
	fragment := s.fragmentsCounter
	s.fragmentsCounter++
	return fragment
}
func (s *regionsService) newRegion() Region {
	regions := s.regionGaps.GetIndices()
	if len(regions) == 0 {
		region := s.regionsCounter
		s.regionsCounter++
		return region
	}
	region := regions[0]
	s.regionGaps.Remove(region)
	return region
}
func (s *regionsService) mergeRegions(fragments ...RegionFragment) {
	if len(fragments) <= 1 {
		return
	}
	f1 := fragments[0]
	f1Data, ok := s.fragments.Get(f1)
	if !ok {
		return
	}
	r1Data, ok := s.regions.Get(f1Data.Region)
	if !ok {
		return
	}
	for _, f2 := range fragments[1:] {
		f2Data, ok := s.fragments.Get(f2)
		if !ok || f2Data.Region == f1Data.Region {
			continue
		}
		r2Data, ok := s.regions.Get(f2Data.Region)
		if !ok {
			continue
		}
		for _, fragment := range r2Data.Fragments.GetIndices() {
			r1Data.Fragments.Add(fragment)
		}
		for _, frag := range r2Data.Fragments.GetIndices() {
			s.fragments.Set(frag, NewFragData(f1Data.Region))
		}
		s.regions.Remove(f2Data.Region)
		s.regionGaps.Add(f2Data.Region)
	}
	s.regions.Set(f1Data.Region, r1Data)
}
func (s *regionsService) getOrCreate(key ChunkObstructionComponent) ecs.EntityID {
	if entity, ok := s.Relation.Get(key); ok {
		return entity
	}
	entity := s.World().NewEntity()
	s.chunkObstruction.Set(entity, key)
	return entity
}

func (s *regionsService) OnChunkUnload(event tile.UnloadChunkEvent) {
	data, ok := s.chunksData[event.Coords]
	if !ok {
		return
	}
	delete(s.chunksData, event.Coords)
	for _, fragment := range data.Fragments.GetIndices() {
		// modify fragment
		fragData, ok := s.fragments.Get(fragment)
		if !ok {
			continue
		}
		fragData.ChunksCounter--
		if fragData.ChunksCounter != 0 {
			s.fragments.Set(fragment, fragData)
			continue
		}
		s.fragments.Remove(fragment)
		s.fragmentsGaps.Add(fragment)

		// modify region if there is no fragment
		region := fragData.Region
		regData, ok := s.regions.Get(region)
		if !ok {
			continue
		}
		regData.Fragments.Remove(fragment)
		if len(regData.Fragments.GetIndices()) != 0 {
			continue
		}
		s.regions.Remove(region)
		s.regionGaps.Add(region)
	}
	for _, obstruction := range s.Obstruction().Obstructions().GetIndices() {
		entity, ok := s.Relation.Get(NewChunkObstruction(event.Coords, obstruction))
		if !ok {
			continue
		}
		s.World().RemoveEntity(entity)
	}
}

var neighborsCoords = []grid.Coords{
	grid.NewCoords(0, -1),
	grid.NewCoords(1, 0),
	grid.NewCoords(0, 1),
	grid.NewCoords(-1, 0),
}

func (s *regionsService) OnChunkUpsert(originalChunkEntity ecs.EntityID) {
	originalChunkCoords, ok := s.Grid().Coords().Get(originalChunkEntity)
	if !ok {
		s.Logger().Fatal(pathfind.ErrInvalidServiceOrder)
	}
	originalChunk, ok := s.Tile().Grid().Chunk().Get(originalChunkEntity)
	if !ok {
		s.Logger().Fatal(pathfind.ErrInvalidServiceOrder)
	}
	for _, obstruction := range s.Obstruction().Obstructions().GetIndices() {
		entityKey := NewChunkObstruction(originalChunkCoords, obstruction)
		entity := s.getOrCreate(entityKey)
		fragmentsChunk, ok := s.FragmentsChunks.Chunk().Get(entity)
		if !ok {
			fragmentsChunk = s.FragmentsChunks.NewChunk()
		}
		chunkFragments := datastructures.NewSparseSet[RegionFragment]()
		for i, tile := range originalChunk.GetTiles() {
			gridI := grid.Index(i)
			coords := s.Grid().IndexCoords(gridI)
			absoluteCoords := s.Grid().AbsoluteCoords(originalChunkCoords, coords)
			tileEntity, ok := s.Tile().GetTile(tile)
			if !ok {
				continue
			}
			tileObstruction, ok := s.Obstruction().Component().Get(tileEntity)
			if !ok || obstruction&tileObstruction.Obstruction != 0 {
				continue
			}
			neighborFragments := []RegionFragment{0}
			for _, neighborCoords := range neighborsCoords {
				absoluteCoords := grid.NewCoords(
					absoluteCoords.X+neighborCoords.X,
					absoluteCoords.Y+neighborCoords.Y)
				chunkCoords, coords := s.Grid().RelativeCoords(absoluteCoords)
				chunk := fragmentsChunk
				if originalChunkCoords != chunkCoords {
					comp := NewChunkObstruction(chunkCoords, obstruction)
					chunkEntity, ok := s.Relation.Get(comp)
					if !ok {
						continue
					}
					chunk, ok = s.FragmentsChunks.Chunk().Get(chunkEntity)
					if !ok {
						continue
					}
				}
				index, _ := s.Grid().CoordsIndex(coords)
				fragment := chunk.GetTile(index)
				if !slices.Contains(neighborFragments, fragment) {
					neighborFragments = append(neighborFragments, fragment)
				}
			}
			neighborFragments = neighborFragments[1:]

			s.mergeRegions(neighborFragments...)
			var fragment RegionFragment
			var fragmentData FragmentData
			var region Region
			var regionData RegionData
			if len(neighborFragments) == 0 { // 4 variables getters
				fragment = s.newFragment()
				region = s.newRegion()
				fragmentData = NewFragData(region)
				regionData = NewRegionData(obstruction)
			} else {
				fragment = neighborFragments[0]
				fragmentData, ok = s.fragments.Get(fragment)
				if !ok {
					s.Logger().Warn(fmt.Errorf("invalid internal state of pathfind module. Expected fragment data"))
					fragmentData = NewFragData(s.newRegion())
				}
				region = fragmentData.Region
				regionData, ok = s.regions.Get(region)
				if !ok {
					s.Logger().Warn(fmt.Errorf("invalid internal state of pathfind module. Expected region data"))
					regionData = NewRegionData(obstruction)
				}
			}
			chunkFragments.Add(fragment)
			regionData.Fragments.Add(fragment)

			s.fragments.Set(fragment, fragmentData)
			s.regions.Set(region, regionData)
			fragmentsChunk.SetTile(gridI, fragment)
		}
		for _, frag := range chunkFragments.GetIndices() {
			data, ok := s.fragments.Get(frag)
			if !ok {
				continue
			}
			data.ChunksCounter++
			s.fragments.Set(frag, data)
		}
		s.chunksData[originalChunkCoords] = NewChunkData(chunkFragments)
		s.FragmentsChunks.Chunk().Set(entity, fragmentsChunk)
	}
}

//

func (s *regionsService) RegionObstruction(region Region) (obstruction.Obstruction, bool) {
	regionData, ok := s.regions.Get(region)
	return regionData.Obstruction, ok
}
func (s *regionsService) CoordsRegion(coords grid.Coords, obstruction obstruction.Obstruction) (pathfind.Region, bool) {
	chunkCoords, relativeCoords := s.Grid().RelativeCoords(coords)
	chunkKey := NewChunkObstruction(chunkCoords, obstruction)
	coordsIndex, ok := s.Grid().CoordsIndex(relativeCoords)
	if !ok {
		return s.zeroRegion, false
	}

	chunkEntity, ok := s.Relation.Get(chunkKey)
	if !ok {
		return s.zeroRegion, false
	}
	chunk, ok := s.FragmentsChunks.Chunk().Get(chunkEntity)
	if !ok {
		return s.zeroRegion, false
	}
	fragment := chunk.GetTile(coordsIndex)
	fragmentData, ok := s.fragments.Get(fragment)
	return fragmentData.Region, ok
}
func (s *regionsService) EntityRegion(entity ecs.EntityID) (pathfind.Region, bool) {
	pos, ok := s.Tile().Pos().Get(entity)
	if !ok {
		return s.zeroRegion, false
	}
	coords, _ := pos.Aligned()
	obstruction, ok := s.Obstruction().Component().Get(entity)
	if !ok {
		return s.zeroRegion, false
	}
	return s.CoordsRegion(coords, obstruction.Obstruction)
}
func (s *regionsService) ShareRegion(entity ecs.EntityID, coords grid.Coords) bool {
	obstruction, ok := s.Obstruction().Component().Get(entity)
	if !ok {
		return false
	}
	entityRegion, ok := s.EntityRegion(entity)
	if !ok {
		return false
	}
	coordsRegion, ok := s.CoordsRegion(coords, obstruction.Obstruction)
	if !ok {
		return false
	}
	return entityRegion == coordsRegion
}
