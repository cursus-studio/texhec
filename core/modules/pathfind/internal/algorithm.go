package internal

import (
	"container/heap"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/services/datastructures"
	"math"
)

// algorithm is ai generated

// Item represents a node in the priority queue
type Item struct {
	index    grid.Index
	coords   grid.Coords
	priority int // fScore
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// heuristic calculates Manhattan distance for 4-directional movement
func heuristic(a, b grid.Coords) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func (s *service) findPath(
	from, to grid.Coords,
	size tile.SizeComponent,
	obstruction tile.ObstructionComponent,
) (path []tile.PosComponent, ok bool) {
	obstructionGridEntity := s.Tile().ObstructionGrid().GetEntities()[0]
	obstructed, ok := s.Tile().ObstructionGrid().Get(obstructionGridEntity)
	if !ok {
		return nil, false
	}

	fromIndex, _ := obstructed.GetIndex(from.Coords())
	toIndex, _ := obstructed.GetIndex(to.Coords())
	if obstruction.Obstruction&obstructed.GetTile(toIndex) != 0 {
		return nil, false
	}

	// 2. Initialize A* data structures
	// gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := datastructures.NewSparseArray[grid.Index, int]()
	gScore.Set(fromIndex, 0)

	// fScore[n] = gScore[n] + h(n). fScore[n] represents our current best guess.
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, &Item{index: fromIndex, coords: from, priority: heuristic(from, to)})

	// cameFrom tracks the path for reconstruction
	cameFrom := datastructures.NewSparseArray[grid.Index, grid.Coords]()
	parentIndex := datastructures.NewSparseArray[grid.Index, grid.Index]()

	// Directions for 4 nearest neighbors
	dirs := []struct{ x, y grid.Coord }{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for openSet.Len() > 0 {
		// Get node with lowest fScore
		current := heap.Pop(openSet).(*Item)

		// Goal reached?
		if current.index == toIndex {
			path := reconstructPath(cameFrom, parentIndex, toIndex)
			path = append(path, tile.NewPos(to.Coords()))
			return path, true
		}

		// Check neighbors
		for _, d := range dirs {
			neighborCoords := grid.Coords{
				X: current.coords.X + d.x,
				Y: current.coords.Y + d.y,
			}

			neighborIndex, valid := obstructed.GetIndex(neighborCoords.Coords())
			if !valid {
				continue
			}

			step := tile.NewStep(neighborCoords.X, neighborCoords.Y)
			if !s.Tile().CanStep(current.coords, size, obstruction, step) {
				continue
			}

			tentativeGScore, _ := gScore.Get(current.index)
			tentativeGScore++

			if score, exists := gScore.Get(neighborIndex); !exists || tentativeGScore < score {
				cameFrom.Set(neighborIndex, current.coords)
				parentIndex.Set(neighborIndex, current.index)
				gScore.Set(neighborIndex, tentativeGScore)
				fScore := tentativeGScore + heuristic(neighborCoords, to)

				heap.Push(openSet, &Item{
					index:    neighborIndex,
					coords:   neighborCoords,
					priority: fScore,
				})
			}
		}
	}

	return nil, false
}

func reconstructPath(
	cameFrom datastructures.SparseArray[grid.Index, grid.Coords],
	parentIndex datastructures.SparseArray[grid.Index, grid.Index],
	current grid.Index,
) []tile.PosComponent {
	var path []tile.PosComponent

	for {
		coords, ok := cameFrom.Get(current)
		if !ok {
			break
		}
		pos := tile.NewPos(coords.X, coords.Y)
		path = append([]tile.PosComponent{pos}, path...)
		current, _ = parentIndex.Get(current)
	}

	path = path[1:]

	return path
}

// func (s *service) findPath(
// 	from, to grid.Coords,
// 	size tile.SizeComponent,
// 	obstruction tile.ObstructionComponent,
// ) (path []tile.PosComponent, ok bool) {
// 	obstructionGridEntity := s.Tile.ObstructionGrid().GetEntities()[0]
// 	obstructed, ok := s.Tile.ObstructionGrid().Get(obstructionGridEntity)
// 	if !ok {
// 		s.Logger.Warn(tile.ErrPositionIsOccupied)
// 		return nil, false
// 	}
// 	fromIndex, ok := obstructed.GetIndex(from.Coords())
// 	if !ok {
// 		s.Logger.Warn(tile.ErrInvalidPosition)
// 		return nil, false
// 	}
// 	toIndex, ok := obstructed.GetIndex(to.Coords())
// 	if !ok {
// 		s.Logger.Warn(tile.ErrInvalidPosition)
// 		return nil, false
// 	}
//
// 	// in this comments is everything you need to know
// 	// neighbours := []tile.PosComponent{
// 	// 	tile.NewPos(0, 1),
// 	// 	tile.NewPos(1, 0),
// 	// 	tile.NewPos(-1, 0),
// 	// 	tile.NewPos(0, -1),
// 	// }
// 	// step := tile.NewStep(grid.Coord(from.X+neighbours[0].X), grid.Coord(from.Y+neighbours[1].Y))
// 	// canStep := s.Tile.CanStep(from, size, obstruction, step)
//
// 	type SearchedTile struct {
// 		Distance grid.Coord
// 		Parent   grid.Index
// 		Coords   grid.Coords
// 	}
// 	visited := datastructures.NewSparseArray[grid.Index, SearchedTile]()
// 	queued := []grid.Index{fromIndex}
//
// 	for len(queued) != 0 && queued[0] != toIndex {
// 		visiting := queued[0]
// 		queued = queued[1:]
//
// 		// can step
// 		t := SearchedTile{
// 			Distance: 0,
// 			// Parent: ,
// 			// Coords: ,
// 		}
// 		visited.Set(visiting)
//
// 		s.Tile.CanStep()
// 	}
//
// 	return nil, false
// }
