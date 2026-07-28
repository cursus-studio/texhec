package internal

import (
	"container/heap"
	"core/modules/obstruction"
	"core/modules/pathfind"
	"core/modules/tile"
	"engine/modules/grid"
	"math"
)

// algorithm is ai generated

// Item represents a node in the priority queue
type Item struct {
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
	obstruction obstruction.Component,
) (path []tile.PosComponent, ok bool) {
	toData, ok := s.Obstruction().Grid().CoordsData(to)
	if !ok {
		return nil, false
	}
	if obstruction.Obstruction&toData.Component.GetTile(toData.Index) != 0 {
		return nil, false
	}

	// gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := map[grid.Coords]int{}
	gScore[from] = 0

	// fScore[n] = gScore[n] + h(n). fScore[n] represents our current best guess.
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, &Item{coords: from, priority: heuristic(from, to)})

	// cameFrom tracks the path for reconstruction
	cameFrom := map[grid.Coords]grid.Coords{}
	parentIndex := map[grid.Coords]grid.Coords{}

	// Directions for 4 nearest neighbors
	dirs := []struct{ x, y grid.Coord }{
		{0, 1}, {1, 0}, {0, grid.NewCoord(-1)}, {grid.NewCoord(-1), 0},
	}

	for openSet.Len() > 0 {
		// Get node with lowest fScore
		current := heap.Pop(openSet).(*Item)

		// Goal reached?
		if current.coords == to {
			path := reconstructPath(cameFrom, parentIndex, to)
			path = append(path, tile.NewPos(to.Coords()))
			return path, true
		}

		// Check neighbors
		for _, d := range dirs {
			neighborCoords := grid.Coords{
				X: current.coords.X + d.x,
				Y: current.coords.Y + d.y,
			}
			step := pathfind.NewStep(neighborCoords.X, neighborCoords.Y)
			if !s.Pathfind().CanStep(current.coords, size, obstruction, step) {
				continue
			}

			tentativeGScore := gScore[current.coords]
			tentativeGScore++

			if score, exists := gScore[neighborCoords]; !exists || tentativeGScore < score {
				cameFrom[neighborCoords] = current.coords
				parentIndex[neighborCoords] = current.coords
				gScore[neighborCoords] = tentativeGScore
				fScore := tentativeGScore + heuristic(neighborCoords, to)

				heap.Push(openSet, &Item{
					coords:   neighborCoords,
					priority: fScore,
				})
			}
		}
	}

	return nil, false
}

func reconstructPath(
	cameFrom map[grid.Coords]grid.Coords,
	parentIndex map[grid.Coords]grid.Coords,
	current grid.Coords,
) []tile.PosComponent {
	var path []tile.PosComponent

	for {
		coords, ok := cameFrom[current]
		if !ok {
			break
		}
		pos := tile.NewPos(coords.X, coords.Y)
		path = append([]tile.PosComponent{pos}, path...)
		current = parentIndex[current]
	}

	path = path[1:]

	return path
}
